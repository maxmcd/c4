// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_service.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JoinPoolRequest struct {
}

func (m *JoinPoolRequest) Reset()                    { *m = JoinPoolRequest{} }
func (m *JoinPoolRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinPoolRequest) ProtoMessage()               {}
func (*JoinPoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type JoinPoolResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *JoinPoolResponse) Reset()                    { *m = JoinPoolResponse{} }
func (m *JoinPoolResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinPoolResponse) ProtoMessage()               {}
func (*JoinPoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *JoinPoolResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SendMoveRequest struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Board   string `protobuf:"bytes,3,opt,name=board" json:"board,omitempty"`
}

func (m *SendMoveRequest) Reset()                    { *m = SendMoveRequest{} }
func (m *SendMoveRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMoveRequest) ProtoMessage()               {}
func (*SendMoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SendMoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendMoveRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *SendMoveRequest) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

type SendMoveResponse struct {
}

func (m *SendMoveResponse) Reset()                    { *m = SendMoveResponse{} }
func (m *SendMoveResponse) String() string            { return proto.CompactTextString(m) }
func (*SendMoveResponse) ProtoMessage()               {}
func (*SendMoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type ReceiveMoveRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReceiveMoveRequest) Reset()                    { *m = ReceiveMoveRequest{} }
func (m *ReceiveMoveRequest) String() string            { return proto.CompactTextString(m) }
func (*ReceiveMoveRequest) ProtoMessage()               {}
func (*ReceiveMoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ReceiveMoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReceiveMoveResponse struct {
	Status  *google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Command string             `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Board   string             `protobuf:"bytes,3,opt,name=board" json:"board,omitempty"`
}

func (m *ReceiveMoveResponse) Reset()                    { *m = ReceiveMoveResponse{} }
func (m *ReceiveMoveResponse) String() string            { return proto.CompactTextString(m) }
func (*ReceiveMoveResponse) ProtoMessage()               {}
func (*ReceiveMoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ReceiveMoveResponse) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReceiveMoveResponse) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ReceiveMoveResponse) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

type Game struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RedUserId   uint64 `protobuf:"varint,2,opt,name=red_user_id,json=redUserId" json:"red_user_id,omitempty"`
	BlackUserId uint64 `protobuf:"varint,3,opt,name=black_user_id,json=blackUserId" json:"black_user_id,omitempty"`
}

func (m *Game) Reset()                    { *m = Game{} }
func (m *Game) String() string            { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()               {}
func (*Game) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Game) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Game) GetRedUserId() uint64 {
	if m != nil {
		return m.RedUserId
	}
	return 0
}

func (m *Game) GetBlackUserId() uint64 {
	if m != nil {
		return m.BlackUserId
	}
	return 0
}

func init() {
	proto.RegisterType((*JoinPoolRequest)(nil), "api.JoinPoolRequest")
	proto.RegisterType((*JoinPoolResponse)(nil), "api.JoinPoolResponse")
	proto.RegisterType((*SendMoveRequest)(nil), "api.SendMoveRequest")
	proto.RegisterType((*SendMoveResponse)(nil), "api.SendMoveResponse")
	proto.RegisterType((*ReceiveMoveRequest)(nil), "api.ReceiveMoveRequest")
	proto.RegisterType((*ReceiveMoveResponse)(nil), "api.ReceiveMoveResponse")
	proto.RegisterType((*Game)(nil), "api.Game")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GameService service

type GameServiceClient interface {
	JoinPool(ctx context.Context, in *JoinPoolRequest, opts ...grpc.CallOption) (*JoinPoolResponse, error)
	SendMove(ctx context.Context, in *SendMoveRequest, opts ...grpc.CallOption) (*SendMoveResponse, error)
	ReceiveMove(ctx context.Context, in *ReceiveMoveRequest, opts ...grpc.CallOption) (GameService_ReceiveMoveClient, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) JoinPool(ctx context.Context, in *JoinPoolRequest, opts ...grpc.CallOption) (*JoinPoolResponse, error) {
	out := new(JoinPoolResponse)
	err := grpc.Invoke(ctx, "/api.GameService/JoinPool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) SendMove(ctx context.Context, in *SendMoveRequest, opts ...grpc.CallOption) (*SendMoveResponse, error) {
	out := new(SendMoveResponse)
	err := grpc.Invoke(ctx, "/api.GameService/SendMove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ReceiveMove(ctx context.Context, in *ReceiveMoveRequest, opts ...grpc.CallOption) (GameService_ReceiveMoveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GameService_serviceDesc.Streams[0], c.cc, "/api.GameService/ReceiveMove", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceReceiveMoveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_ReceiveMoveClient interface {
	Recv() (*ReceiveMoveResponse, error)
	grpc.ClientStream
}

type gameServiceReceiveMoveClient struct {
	grpc.ClientStream
}

func (x *gameServiceReceiveMoveClient) Recv() (*ReceiveMoveResponse, error) {
	m := new(ReceiveMoveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GameService service

type GameServiceServer interface {
	JoinPool(context.Context, *JoinPoolRequest) (*JoinPoolResponse, error)
	SendMove(context.Context, *SendMoveRequest) (*SendMoveResponse, error)
	ReceiveMove(*ReceiveMoveRequest, GameService_ReceiveMoveServer) error
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_JoinPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).JoinPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GameService/JoinPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).JoinPool(ctx, req.(*JoinPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_SendMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).SendMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GameService/SendMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).SendMove(ctx, req.(*SendMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ReceiveMove_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveMoveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).ReceiveMove(m, &gameServiceReceiveMoveServer{stream})
}

type GameService_ReceiveMoveServer interface {
	Send(*ReceiveMoveResponse) error
	grpc.ServerStream
}

type gameServiceReceiveMoveServer struct {
	grpc.ServerStream
}

func (x *gameServiceReceiveMoveServer) Send(m *ReceiveMoveResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPool",
			Handler:    _GameService_JoinPool_Handler,
		},
		{
			MethodName: "SendMove",
			Handler:    _GameService_SendMove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMove",
			Handler:       _GameService_ReceiveMove_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "game_service.proto",
}

func init() { proto.RegisterFile("game_service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xd9, 0x6c, 0xad, 0xed, 0x09, 0xda, 0x7a, 0xac, 0x6c, 0x08, 0x55, 0x64, 0xf0, 0x42,
	0x0a, 0x4d, 0x4a, 0xbd, 0xf3, 0x05, 0x44, 0x41, 0xd0, 0x2c, 0x7a, 0xa1, 0x17, 0xcb, 0x6c, 0xe6,
	0xb8, 0x8c, 0x6e, 0xe6, 0xa4, 0x33, 0xd9, 0xdc, 0x88, 0x37, 0xbe, 0x82, 0x6f, 0xe1, 0xeb, 0xf8,
	0x0a, 0x3e, 0x88, 0xec, 0xcc, 0x6c, 0xeb, 0x6e, 0x40, 0xf0, 0x72, 0xfe, 0xf3, 0xe7, 0xfb, 0x27,
	0xe7, 0x1f, 0xc0, 0x85, 0x6c, 0x68, 0xe6, 0xc8, 0xf6, 0xba, 0xa6, 0xa2, 0xb5, 0xdc, 0x31, 0x8e,
	0x65, 0xab, 0xf3, 0xd3, 0x05, 0xf3, 0x62, 0x49, 0xa5, 0x6c, 0x75, 0x29, 0x8d, 0xe1, 0x4e, 0x76,
	0x9a, 0x8d, 0x0b, 0x96, 0x7c, 0x12, 0xa7, 0xb6, 0xad, 0x4b, 0xd7, 0xc9, 0x6e, 0x15, 0x07, 0xe2,
	0x1e, 0x1c, 0xbd, 0x62, 0x6d, 0xde, 0x30, 0x2f, 0x2b, 0xba, 0x5a, 0x91, 0xeb, 0x84, 0x80, 0xe3,
	0x1b, 0xc9, 0xb5, 0x6c, 0x1c, 0xe1, 0x5d, 0x48, 0xb4, 0xca, 0x46, 0x8f, 0x47, 0x4f, 0x0f, 0xab,
	0x44, 0x2b, 0xf1, 0x16, 0x8e, 0xa6, 0x64, 0xd4, 0x6b, 0xee, 0x29, 0x7e, 0xb6, 0x6b, 0xc1, 0x0c,
	0x6e, 0xd7, 0xdc, 0x34, 0xd2, 0xa8, 0x2c, 0xf1, 0xe2, 0xe6, 0x88, 0x27, 0x70, 0x6b, 0xce, 0xd2,
	0xaa, 0x6c, 0xec, 0xf5, 0x70, 0x10, 0x08, 0xc7, 0x37, 0xc8, 0x10, 0x2b, 0x9e, 0x00, 0x56, 0x54,
	0x93, 0xee, 0xe9, 0x1f, 0x49, 0xe2, 0x0a, 0xee, 0x6f, 0xb9, 0xe2, 0x9d, 0xcf, 0x60, 0x3f, 0xfc,
	0xaa, 0xb7, 0xa6, 0x97, 0x58, 0x84, 0x25, 0x14, 0xb6, 0xad, 0x8b, 0xa9, 0x9f, 0x54, 0xd1, 0xf1,
	0xdf, 0x97, 0xfd, 0x00, 0x7b, 0x2f, 0x64, 0x33, 0xd8, 0x0b, 0x3e, 0x82, 0xd4, 0x92, 0x9a, 0xad,
	0x1c, 0xd9, 0x99, 0x0e, 0xac, 0xbd, 0xea, 0xd0, 0x92, 0x7a, 0xe7, 0xc8, 0xbe, 0x54, 0x28, 0xe0,
	0xce, 0x7c, 0x29, 0xeb, 0x2f, 0xd7, 0x8e, 0xb1, 0x77, 0xa4, 0x5e, 0x0c, 0x9e, 0xcb, 0x9f, 0x09,
	0xa4, 0x6b, 0xf8, 0x34, 0x94, 0x8c, 0xef, 0xe1, 0x60, 0xd3, 0x07, 0x9e, 0x14, 0xb2, 0xd5, 0xc5,
	0x4e, 0x63, 0xf9, 0x83, 0x1d, 0x35, 0x6e, 0xef, 0xe1, 0xf7, 0x5f, 0xbf, 0x7f, 0x24, 0x13, 0x81,
	0x65, 0x7f, 0x51, 0xae, 0xdf, 0x4d, 0xf9, 0x99, 0xb5, 0x39, 0x6f, 0x99, 0x97, 0xcf, 0x47, 0x67,
	0xf8, 0x11, 0x0e, 0x36, 0x0b, 0x8f, 0xdc, 0x9d, 0x4a, 0x23, 0x77, 0xd0, 0x8a, 0xf0, 0xdc, 0x53,
	0x31, 0xb9, 0xe6, 0x7e, 0xd5, 0xea, 0x5b, 0xe9, 0xc8, 0xa8, 0xf3, 0x86, 0x7b, 0x5a, 0xc3, 0x3f,
	0x41, 0xfa, 0x57, 0x27, 0x38, 0xf1, 0xa4, 0x61, 0x97, 0x79, 0x36, 0x1c, 0x6c, 0xa7, 0x60, 0xbe,
	0x9d, 0x62, 0x83, 0xd5, 0x07, 0x5d, 0x8c, 0xe6, 0xfb, 0xfe, 0x19, 0x3f, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x84, 0xd7, 0xa2, 0x3a, 0x18, 0x03, 0x00, 0x00,
}