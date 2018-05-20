// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_service.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JoinPoolRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinPoolRequest) Reset()         { *m = JoinPoolRequest{} }
func (m *JoinPoolRequest) String() string { return proto.CompactTextString(m) }
func (*JoinPoolRequest) ProtoMessage()    {}
func (*JoinPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{0}
}
func (m *JoinPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinPoolRequest.Unmarshal(m, b)
}
func (m *JoinPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinPoolRequest.Marshal(b, m, deterministic)
}
func (dst *JoinPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinPoolRequest.Merge(dst, src)
}
func (m *JoinPoolRequest) XXX_Size() int {
	return xxx_messageInfo_JoinPoolRequest.Size(m)
}
func (m *JoinPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinPoolRequest proto.InternalMessageInfo

type JoinPoolResponse struct {
	Game                 *Game    `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinPoolResponse) Reset()         { *m = JoinPoolResponse{} }
func (m *JoinPoolResponse) String() string { return proto.CompactTextString(m) }
func (*JoinPoolResponse) ProtoMessage()    {}
func (*JoinPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{1}
}
func (m *JoinPoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinPoolResponse.Unmarshal(m, b)
}
func (m *JoinPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinPoolResponse.Marshal(b, m, deterministic)
}
func (dst *JoinPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinPoolResponse.Merge(dst, src)
}
func (m *JoinPoolResponse) XXX_Size() int {
	return xxx_messageInfo_JoinPoolResponse.Size(m)
}
func (m *JoinPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinPoolResponse proto.InternalMessageInfo

func (m *JoinPoolResponse) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type SendMoveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Column               int32    `protobuf:"varint,3,opt,name=column" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMoveRequest) Reset()         { *m = SendMoveRequest{} }
func (m *SendMoveRequest) String() string { return proto.CompactTextString(m) }
func (*SendMoveRequest) ProtoMessage()    {}
func (*SendMoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{2}
}
func (m *SendMoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMoveRequest.Unmarshal(m, b)
}
func (m *SendMoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMoveRequest.Marshal(b, m, deterministic)
}
func (dst *SendMoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMoveRequest.Merge(dst, src)
}
func (m *SendMoveRequest) XXX_Size() int {
	return xxx_messageInfo_SendMoveRequest.Size(m)
}
func (m *SendMoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMoveRequest proto.InternalMessageInfo

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

func (m *SendMoveRequest) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

type SendMoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMoveResponse) Reset()         { *m = SendMoveResponse{} }
func (m *SendMoveResponse) String() string { return proto.CompactTextString(m) }
func (*SendMoveResponse) ProtoMessage()    {}
func (*SendMoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{3}
}
func (m *SendMoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMoveResponse.Unmarshal(m, b)
}
func (m *SendMoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMoveResponse.Marshal(b, m, deterministic)
}
func (dst *SendMoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMoveResponse.Merge(dst, src)
}
func (m *SendMoveResponse) XXX_Size() int {
	return xxx_messageInfo_SendMoveResponse.Size(m)
}
func (m *SendMoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMoveResponse proto.InternalMessageInfo

type ReceiveMoveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMoveRequest) Reset()         { *m = ReceiveMoveRequest{} }
func (m *ReceiveMoveRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveMoveRequest) ProtoMessage()    {}
func (*ReceiveMoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{4}
}
func (m *ReceiveMoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMoveRequest.Unmarshal(m, b)
}
func (m *ReceiveMoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMoveRequest.Marshal(b, m, deterministic)
}
func (dst *ReceiveMoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMoveRequest.Merge(dst, src)
}
func (m *ReceiveMoveRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveMoveRequest.Size(m)
}
func (m *ReceiveMoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMoveRequest proto.InternalMessageInfo

func (m *ReceiveMoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReceiveMoveResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Command              string         `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Board                string         `protobuf:"bytes,3,opt,name=board" json:"board,omitempty"`
	RedTimeRemaining     uint64         `protobuf:"varint,4,opt,name=red_time_remaining,json=redTimeRemaining" json:"red_time_remaining,omitempty"`
	BlackTimeRemaining   uint64         `protobuf:"varint,5,opt,name=black_time_remaining,json=blackTimeRemaining" json:"black_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReceiveMoveResponse) Reset()         { *m = ReceiveMoveResponse{} }
func (m *ReceiveMoveResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiveMoveResponse) ProtoMessage()    {}
func (*ReceiveMoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{5}
}
func (m *ReceiveMoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMoveResponse.Unmarshal(m, b)
}
func (m *ReceiveMoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMoveResponse.Marshal(b, m, deterministic)
}
func (dst *ReceiveMoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMoveResponse.Merge(dst, src)
}
func (m *ReceiveMoveResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiveMoveResponse.Size(m)
}
func (m *ReceiveMoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMoveResponse proto.InternalMessageInfo

func (m *ReceiveMoveResponse) GetStatus() *status.Status {
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

func (m *ReceiveMoveResponse) GetRedTimeRemaining() uint64 {
	if m != nil {
		return m.RedTimeRemaining
	}
	return 0
}

func (m *ReceiveMoveResponse) GetBlackTimeRemaining() uint64 {
	if m != nil {
		return m.BlackTimeRemaining
	}
	return 0
}

type Game struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RedUser              *User    `protobuf:"bytes,2,opt,name=red_user,json=redUser" json:"red_user,omitempty"`
	BlackUser            *User    `protobuf:"bytes,3,opt,name=black_user,json=blackUser" json:"black_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{6}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (dst *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(dst, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Game) GetRedUser() *User {
	if m != nil {
		return m.RedUser
	}
	return nil
}

func (m *Game) GetBlackUser() *User {
	if m != nil {
		return m.BlackUser
	}
	return nil
}

type User struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_service_17958d6cc7ea3098, []int{7}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinPoolRequest)(nil), "api.JoinPoolRequest")
	proto.RegisterType((*JoinPoolResponse)(nil), "api.JoinPoolResponse")
	proto.RegisterType((*SendMoveRequest)(nil), "api.SendMoveRequest")
	proto.RegisterType((*SendMoveResponse)(nil), "api.SendMoveResponse")
	proto.RegisterType((*ReceiveMoveRequest)(nil), "api.ReceiveMoveRequest")
	proto.RegisterType((*ReceiveMoveResponse)(nil), "api.ReceiveMoveResponse")
	proto.RegisterType((*Game)(nil), "api.Game")
	proto.RegisterType((*User)(nil), "api.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	err := c.cc.Invoke(ctx, "/api.GameService/JoinPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) SendMove(ctx context.Context, in *SendMoveRequest, opts ...grpc.CallOption) (*SendMoveResponse, error) {
	out := new(SendMoveResponse)
	err := c.cc.Invoke(ctx, "/api.GameService/SendMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ReceiveMove(ctx context.Context, in *ReceiveMoveRequest, opts ...grpc.CallOption) (GameService_ReceiveMoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GameService_serviceDesc.Streams[0], "/api.GameService/ReceiveMove", opts...)
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

// GameServiceServer is the server API for GameService service.
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

func init() { proto.RegisterFile("game_service.proto", fileDescriptor_game_service_17958d6cc7ea3098) }

var fileDescriptor_game_service_17958d6cc7ea3098 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xba, 0x69, 0xb7, 0x9d, 0x4a, 0x6c, 0x19, 0x0a, 0x8d, 0xa2, 0x5d, 0xa9, 0xb2, 0xf6,
	0x50, 0xad, 0xd8, 0xa4, 0x94, 0x1b, 0x3f, 0x80, 0x84, 0x84, 0x84, 0x52, 0xe0, 0xc2, 0xa1, 0x72,
	0xe3, 0xd9, 0xca, 0xd0, 0xd8, 0xc1, 0x49, 0x7b, 0x41, 0x5c, 0xf8, 0x05, 0xfe, 0x82, 0xbf, 0x41,
	0xfc, 0x02, 0x1f, 0x82, 0xec, 0xb8, 0x5b, 0xda, 0x4a, 0x7b, 0xcb, 0xcc, 0x7b, 0x7e, 0x6f, 0x3c,
	0x7e, 0x01, 0x5c, 0xf1, 0x82, 0x16, 0x15, 0x99, 0xad, 0xcc, 0x29, 0x29, 0x8d, 0xae, 0x35, 0x9e,
	0xf1, 0x52, 0xc6, 0x97, 0x2b, 0xad, 0x57, 0x6b, 0x4a, 0x79, 0x29, 0x53, 0xae, 0x94, 0xae, 0x79,
	0x2d, 0xb5, 0xaa, 0x1a, 0x4a, 0x3c, 0xf2, 0xa8, 0x29, 0xf3, 0xb4, 0xaa, 0x79, 0xbd, 0xf1, 0x00,
	0x7b, 0x0c, 0x17, 0x6f, 0xb4, 0x54, 0xef, 0xb4, 0x5e, 0x67, 0xf4, 0x75, 0x43, 0x55, 0xcd, 0x5e,
	0xc0, 0x60, 0xdf, 0xaa, 0x4a, 0xad, 0x2a, 0xc2, 0x2b, 0x08, 0xad, 0x71, 0x14, 0x8c, 0x83, 0x49,
	0x7f, 0xd6, 0x4b, 0x78, 0x29, 0x93, 0xd7, 0xbc, 0xa0, 0xcc, 0xb5, 0xd9, 0x1c, 0x2e, 0xe6, 0xa4,
	0xc4, 0x5b, 0xbd, 0x25, 0xaf, 0x82, 0x8f, 0xa0, 0x25, 0x85, 0xe3, 0xf7, 0xb2, 0x96, 0x14, 0x18,
	0xc1, 0x79, 0xae, 0x8b, 0x82, 0x2b, 0x11, 0xb5, 0x5c, 0x73, 0x57, 0xe2, 0x33, 0xe8, 0xe4, 0x7a,
	0xbd, 0x29, 0x54, 0x74, 0x36, 0x0e, 0x26, 0xed, 0xcc, 0x57, 0x0c, 0x61, 0xb0, 0x17, 0x6d, 0xe6,
	0x60, 0xd7, 0x80, 0x19, 0xe5, 0x24, 0xb7, 0xf4, 0x80, 0x17, 0xfb, 0x1d, 0xc0, 0x93, 0x03, 0x9a,
	0xbf, 0xc5, 0x0d, 0x74, 0x9a, 0xcb, 0xfb, 0x7b, 0x60, 0xd2, 0xac, 0x25, 0x31, 0x65, 0x9e, 0xcc,
	0x1d, 0x92, 0x79, 0xc6, 0x03, 0xf3, 0x0e, 0xa1, 0xbd, 0xd4, 0xdc, 0x08, 0x37, 0x6e, 0x2f, 0x6b,
	0x0a, 0x7c, 0x0e, 0x68, 0x48, 0x2c, 0x6a, 0x59, 0xd0, 0xc2, 0x50, 0xc1, 0xa5, 0x92, 0x6a, 0x15,
	0x85, 0xe3, 0x60, 0x12, 0x66, 0x03, 0x43, 0xe2, 0xbd, 0x2c, 0x28, 0xdb, 0xf5, 0x71, 0x0a, 0xc3,
	0xe5, 0x9a, 0xe7, 0x5f, 0x8e, 0xf9, 0x6d, 0xc7, 0x47, 0x87, 0x1d, 0x9c, 0x60, 0x77, 0x10, 0xda,
	0x85, 0x9f, 0xec, 0xf5, 0x1a, 0xba, 0xd6, 0x77, 0x53, 0x91, 0x71, 0x83, 0xee, 0x5e, 0xe7, 0x43,
	0x45, 0x26, 0x3b, 0x37, 0x24, 0xec, 0x07, 0x4e, 0x00, 0x1a, 0x3f, 0xc7, 0x3b, 0x3b, 0xe6, 0xf5,
	0x1c, 0x68, 0x3f, 0xd9, 0x0c, 0x42, 0x77, 0x62, 0xef, 0x13, 0x3a, 0x9f, 0x18, 0xba, 0xf6, 0xac,
	0xb2, 0x29, 0x68, 0x16, 0x72, 0x5f, 0xcf, 0x7e, 0xb5, 0xa0, 0x6f, 0x87, 0x9b, 0x37, 0xb1, 0xc4,
	0x8f, 0xd0, 0xdd, 0x25, 0x08, 0x87, 0xce, 0xe5, 0x28, 0x63, 0xf1, 0xd3, 0xa3, 0xae, 0x7f, 0xde,
	0xab, 0x1f, 0x7f, 0xfe, 0xfe, 0x6c, 0x8d, 0x18, 0xa6, 0xdb, 0x69, 0x6a, 0x93, 0x95, 0x7e, 0xd6,
	0x52, 0xdd, 0x96, 0x5a, 0xaf, 0x5f, 0x05, 0x37, 0xf8, 0x09, 0xba, 0xbb, 0x44, 0x78, 0xdd, 0xa3,
	0xd4, 0x79, 0xdd, 0x93, 0xd8, 0x30, 0xa7, 0x7b, 0xc9, 0x46, 0xf7, 0xba, 0xdf, 0xa4, 0xf8, 0x9e,
	0x56, 0xa4, 0xc4, 0x6d, 0xa1, 0xb7, 0x64, 0xc5, 0xef, 0xa0, 0xff, 0x5f, 0x66, 0x70, 0xe4, 0x94,
	0x4e, 0xc3, 0x16, 0x47, 0xa7, 0xc0, 0xa1, 0x0b, 0xc6, 0x87, 0x2e, 0xa6, 0xa1, 0x3a, 0xa3, 0x69,
	0xb0, 0xec, 0xb8, 0x1f, 0xef, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x1c, 0x53, 0x62,
	0xca, 0x03, 0x00, 0x00,
}
