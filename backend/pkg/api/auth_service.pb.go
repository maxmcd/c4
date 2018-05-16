// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	auth_service.proto

It has these top-level messages:
	AuthenticateRequest
	AuthenticateResponse
	AuthenticateEmailRequest
	AuthenticateEmailResponse
	RegisterEmailRequest
	RegisterEmailResponse
	SmsCodeCheckRequest
	SmsCodeCheckResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type AuthenticateRequest struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthenticateRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type AuthenticateResponse struct {
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthenticateResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type AuthenticateEmailRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthenticateEmailRequest) Reset()                    { *m = AuthenticateEmailRequest{} }
func (m *AuthenticateEmailRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateEmailRequest) ProtoMessage()               {}
func (*AuthenticateEmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthenticateEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthenticateEmailRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateEmailResponse struct {
	Token  string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *AuthenticateEmailResponse) Reset()                    { *m = AuthenticateEmailResponse{} }
func (m *AuthenticateEmailResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticateEmailResponse) ProtoMessage()               {}
func (*AuthenticateEmailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthenticateEmailResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthenticateEmailResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type RegisterEmailRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *RegisterEmailRequest) Reset()                    { *m = RegisterEmailRequest{} }
func (m *RegisterEmailRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterEmailRequest) ProtoMessage()               {}
func (*RegisterEmailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterEmailRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterEmailResponse struct {
	Token  string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *RegisterEmailResponse) Reset()                    { *m = RegisterEmailResponse{} }
func (m *RegisterEmailResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterEmailResponse) ProtoMessage()               {}
func (*RegisterEmailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegisterEmailResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegisterEmailResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SmsCodeCheckRequest struct {
	Phone     string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Code      string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *SmsCodeCheckRequest) Reset()                    { *m = SmsCodeCheckRequest{} }
func (m *SmsCodeCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*SmsCodeCheckRequest) ProtoMessage()               {}
func (*SmsCodeCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SmsCodeCheckRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *SmsCodeCheckRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *SmsCodeCheckRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type SmsCodeCheckResponse struct {
	Token    string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	UserId   uint64 `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *SmsCodeCheckResponse) Reset()                    { *m = SmsCodeCheckResponse{} }
func (m *SmsCodeCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*SmsCodeCheckResponse) ProtoMessage()               {}
func (*SmsCodeCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SmsCodeCheckResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SmsCodeCheckResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SmsCodeCheckResponse) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthenticateRequest)(nil), "api.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "api.AuthenticateResponse")
	proto.RegisterType((*AuthenticateEmailRequest)(nil), "api.AuthenticateEmailRequest")
	proto.RegisterType((*AuthenticateEmailResponse)(nil), "api.AuthenticateEmailResponse")
	proto.RegisterType((*RegisterEmailRequest)(nil), "api.RegisterEmailRequest")
	proto.RegisterType((*RegisterEmailResponse)(nil), "api.RegisterEmailResponse")
	proto.RegisterType((*SmsCodeCheckRequest)(nil), "api.SmsCodeCheckRequest")
	proto.RegisterType((*SmsCodeCheckResponse)(nil), "api.SmsCodeCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	AuthenticateEmail(ctx context.Context, in *AuthenticateEmailRequest, opts ...grpc.CallOption) (*AuthenticateEmailResponse, error)
	RegisterEmail(ctx context.Context, in *RegisterEmailRequest, opts ...grpc.CallOption) (*RegisterEmailResponse, error)
	SmsCodeCheck(ctx context.Context, in *SmsCodeCheckRequest, opts ...grpc.CallOption) (*SmsCodeCheckResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthenticateEmail(ctx context.Context, in *AuthenticateEmailRequest, opts ...grpc.CallOption) (*AuthenticateEmailResponse, error) {
	out := new(AuthenticateEmailResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/AuthenticateEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RegisterEmail(ctx context.Context, in *RegisterEmailRequest, opts ...grpc.CallOption) (*RegisterEmailResponse, error) {
	out := new(RegisterEmailResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/RegisterEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SmsCodeCheck(ctx context.Context, in *SmsCodeCheckRequest, opts ...grpc.CallOption) (*SmsCodeCheckResponse, error) {
	out := new(SmsCodeCheckResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/SmsCodeCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	AuthenticateEmail(context.Context, *AuthenticateEmailRequest) (*AuthenticateEmailResponse, error)
	RegisterEmail(context.Context, *RegisterEmailRequest) (*RegisterEmailResponse, error)
	SmsCodeCheck(context.Context, *SmsCodeCheckRequest) (*SmsCodeCheckResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthenticateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthenticateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/AuthenticateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthenticateEmail(ctx, req.(*AuthenticateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RegisterEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/RegisterEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterEmail(ctx, req.(*RegisterEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SmsCodeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsCodeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SmsCodeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/SmsCodeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SmsCodeCheck(ctx, req.(*SmsCodeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
		{
			MethodName: "AuthenticateEmail",
			Handler:    _AuthService_AuthenticateEmail_Handler,
		},
		{
			MethodName: "RegisterEmail",
			Handler:    _AuthService_RegisterEmail_Handler,
		},
		{
			MethodName: "SmsCodeCheck",
			Handler:    _AuthService_SmsCodeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}

func init() { proto.RegisterFile("auth_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xd5, 0xee, 0xb6, 0xa5, 0x1d, 0x8a, 0x04, 0x6e, 0x0a, 0x69, 0x4a, 0xf9, 0xb0, 0x38, 0xa0,
	0x45, 0xdd, 0x20, 0x10, 0x97, 0xde, 0x50, 0x05, 0xa2, 0x88, 0xd3, 0xf6, 0x4e, 0x31, 0xc9, 0x68,
	0x63, 0xb5, 0xb1, 0x83, 0xed, 0x94, 0x3b, 0x7f, 0x81, 0x9f, 0xc5, 0x91, 0xbf, 0xc0, 0x0f, 0x41,
	0x76, 0x1c, 0xe4, 0xcd, 0x9a, 0x3d, 0xc0, 0x6d, 0xdf, 0x8e, 0x67, 0xde, 0x9b, 0x79, 0x4f, 0x01,
	0xc2, 0x5a, 0x53, 0x5d, 0x68, 0x54, 0xd7, 0xbc, 0xc0, 0x59, 0xa3, 0xa4, 0x91, 0x64, 0xc2, 0x1a,
	0x9e, 0xdd, 0x5f, 0x48, 0xb9, 0xb8, 0xc2, 0x9c, 0x35, 0x3c, 0x67, 0x42, 0x48, 0xc3, 0x0c, 0x97,
	0x42, 0x77, 0x4f, 0xe8, 0x33, 0xd8, 0x7b, 0xdd, 0x9a, 0x0a, 0x85, 0xe1, 0x05, 0x33, 0x38, 0xc7,
	0x2f, 0x2d, 0x6a, 0x43, 0x12, 0xd8, 0x6c, 0x2a, 0x29, 0x30, 0x1d, 0x3d, 0x1a, 0x3d, 0xdd, 0x99,
	0x77, 0x80, 0xbe, 0x82, 0x64, 0xf9, 0xb1, 0x6e, 0xa4, 0xd0, 0x48, 0x8e, 0x00, 0x54, 0xd7, 0x78,
	0xc1, 0x4b, 0xdf, 0xb2, 0xe3, 0xff, 0x39, 0x2b, 0xe9, 0x07, 0x48, 0xc3, 0xb6, 0x37, 0x35, 0xe3,
	0x57, 0x01, 0x11, 0x5a, 0xdc, 0x13, 0x39, 0x40, 0x32, 0xd8, 0x6e, 0x98, 0xd6, 0x5f, 0xa5, 0x2a,
	0xd3, 0xb1, 0x2b, 0xfc, 0xc1, 0xf4, 0x3d, 0x1c, 0x44, 0xa6, 0x79, 0x25, 0x09, 0x6c, 0x1a, 0x79,
	0x89, 0xa2, 0x1f, 0xe7, 0x00, 0xb9, 0x07, 0x37, 0x5a, 0x8d, 0xca, 0x8a, 0xb3, 0xd3, 0x36, 0xe6,
	0x5b, 0x16, 0x9e, 0x95, 0xf4, 0x1d, 0x24, 0x73, 0x5c, 0x70, 0x6d, 0x50, 0xfd, 0xa7, 0xaa, 0xb7,
	0xb0, 0x3f, 0x98, 0xf4, 0x6f, 0x8a, 0x3e, 0xc2, 0xde, 0x79, 0xad, 0x4f, 0x65, 0x89, 0xa7, 0x15,
	0x16, 0x97, 0x6b, 0xfd, 0x18, 0xdc, 0x7d, 0x3c, 0xb8, 0x3b, 0x21, 0xb0, 0x51, 0xc8, 0x12, 0xd3,
	0x89, 0x2b, 0xb8, 0xdf, 0x94, 0x41, 0xb2, 0x3c, 0x7f, 0xad, 0xcc, 0x0c, 0xb6, 0xad, 0x2e, 0xc1,
	0x6a, 0xec, 0x37, 0xee, 0x71, 0xb8, 0xc2, 0x24, 0x5c, 0xe1, 0xc5, 0x8f, 0x09, 0xdc, 0xb4, 0x0e,
	0x9d, 0x77, 0x59, 0x24, 0x9f, 0x60, 0x37, 0x34, 0x8c, 0xa4, 0x33, 0xd6, 0xf0, 0x59, 0x24, 0x75,
	0xd9, 0x41, 0xa4, 0xd2, 0xe9, 0xa3, 0x87, 0xdf, 0x7e, 0xfe, 0xfa, 0x3e, 0xde, 0xa7, 0xb7, 0xf3,
	0xeb, 0xe7, 0x39, 0x0b, 0x5e, 0x9c, 0x8c, 0xa6, 0xa4, 0x85, 0x3b, 0x2b, 0x91, 0x20, 0x47, 0x2b,
	0xc3, 0x42, 0x8b, 0xb3, 0x07, 0x7f, 0x2b, 0x7b, 0xc2, 0xc7, 0x8e, 0xf0, 0x90, 0xde, 0x1d, 0x12,
	0xe6, 0x2e, 0x0c, 0x96, 0x56, 0xc1, 0xad, 0x25, 0xcf, 0x49, 0xa7, 0x3f, 0x96, 0xa8, 0x2c, 0x8b,
	0x95, 0x3c, 0xd5, 0xd4, 0x51, 0x3d, 0xa1, 0x0f, 0xe3, 0x54, 0xb9, 0xf2, 0x5d, 0x96, 0x53, 0xc2,
	0x6e, 0xe8, 0x9f, 0x3f, 0x66, 0x24, 0x32, 0xfe, 0x98, 0x31, 0xb3, 0xd7, 0x10, 0xea, 0x5a, 0x1f,
	0xdb, 0x9c, 0x1c, 0x17, 0xb6, 0xe1, 0x64, 0x34, 0xfd, 0xbc, 0xe5, 0xbe, 0x13, 0x2f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x7d, 0x30, 0x81, 0x60, 0x04, 0x00, 0x00,
}