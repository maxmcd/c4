package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/maxmcd/c4/backend/pkg/api"
	"github.com/maxmcd/c4/backend/pkg/config"
	"github.com/maxmcd/c4/backend/pkg/migrations"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	migrations.RunMigrations()

	// Set "logtostderr" flag so that glog messages go stderr and not /tmp.
	flag.Set("logtostderr", "true")
	flag.Parse()

	runApiServer()
}

func runApiServer() {
	// Start proxy
	go runProxy()

	s, err := NewAPIServer()
	if err != nil {
		glog.Fatal(err)
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(streamInterceptor),
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	api.RegisterAuthServiceServer(grpcServer, s)
	api.RegisterGameServiceServer(grpcServer, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",
		config.API.Port))
	if err != nil {
		glog.Fatal(err)
	}

	glog.Infoln("API listening on port:", config.API.Port)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Fatal(err)
	}
}

func runProxy() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	apiURL := fmt.Sprintf("localhost:%d", config.API.Port)

	if err := api.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, apiURL, opts); err != nil {
		glog.Fatalf("Error registering auth proxy server: %s", err.Error())
	}
	if err := api.RegisterGameServiceHandlerFromEndpoint(ctx, mux, apiURL, opts); err != nil {
		glog.Fatalf("Error registering auth proxy server: %s", err.Error())
	}

	glog.Infoln("API proxy is listening on port:", config.API.ProxyPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.API.ProxyPort), allowCORS(wsproxy.WebsocketProxy(mux))); err != nil {
		glog.Fatalf("Error starting proxy server: %s", err.Error())
	}
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// TODO: Fix this for prod deployment
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PATCH", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

type streamWithModifiedContext struct {
	ctx context.Context
	grpc.ServerStream
}

func (s *streamWithModifiedContext) Context() context.Context {
	return s.ctx
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()
	var err error
	if !publicEndpoint(info.FullMethod) {
		ctx, err = authorize(stream.Context())
		if err != nil {
			glog.Error(err)
			return err
		}
	}
	stream = &streamWithModifiedContext{
		ctx,
		stream,
	}
	return handler(srv, stream)
}

func publicEndpoint(fullMethod string) bool {
	suffixes := []string{
		"Authenticate",
		"SmsCodeCheck",
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(fullMethod, suffix) {
			return true
		}
	}
	return false
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error
	if !publicEndpoint(info.FullMethod) {
		// Authorize every request except for the auth endpoint
		if ctx, err = authorize(ctx); err != nil {
			return nil, err
		}
	}
	return handler(ctx, req)
}

func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Missing metadata")
	}

	bearerToken, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Missing authorization token")
	}
	if len(bearerToken) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Expected only 1 auth token, received %d", len(bearerToken))
	}

	bearerToken[0] = strings.TrimPrefix(bearerToken[0], "Bearer ")
	token, err := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {
		// We are expecting HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	// TODO: Ensure JWT claims are truthy (may be done by the jwt lib?)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid claims map")
	}
	if !token.Valid {
		return nil, status.Error(codes.InvalidArgument, "Invalid JWT token")
	}
	userId, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return nil, errors.New("Error converting user string user_id")
	}
	return context.WithValue(ctx, "userId", userId), nil
}
