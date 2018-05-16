package main

import (
	"context"

	"github.com/golang/glog"
	"github.com/maxmcd/c4/backend/pkg/api"
	"google.golang.org/grpc"
)

type jwt struct {
	token string
}

func (j jwt) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": j.token,
	}, nil
}

func (j *jwt) RequireTransportSecurity() bool {
	return false
}

func main() {
	creds := &jwt{
		token: "Bearer notarealjwttoken",
	}
	conn, err := grpc.Dial("localhost:8083", grpc.WithInsecure(), grpc.WithPerRPCCredentials(creds))
	if err != nil {
		glog.Fatal(err)
	}

	c := api.NewStreamServiceClient(conn)

	resp, err := c.CreateStream(context.Background(), &api.CreateStreamRequest{
		Id: "imarealid",
	})
	if err != nil {
		glog.Fatal(err)
	}
	glog.Infoln(resp)
}
