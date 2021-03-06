// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

package light_sequel

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{1}
}

func (m *AuthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReply.Unmarshal(m, b)
}
func (m *AuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReply.Marshal(b, m, deterministic)
}
func (m *AuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReply.Merge(m, src)
}
func (m *AuthReply) XXX_Size() int {
	return xxx_messageInfo_AuthReply.Size(m)
}
func (m *AuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReply proto.InternalMessageInfo

func (m *AuthReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthReply) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SrvRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrvRequest) Reset()         { *m = SrvRequest{} }
func (m *SrvRequest) String() string { return proto.CompactTextString(m) }
func (*SrvRequest) ProtoMessage()    {}
func (*SrvRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{2}
}

func (m *SrvRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvRequest.Unmarshal(m, b)
}
func (m *SrvRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvRequest.Marshal(b, m, deterministic)
}
func (m *SrvRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvRequest.Merge(m, src)
}
func (m *SrvRequest) XXX_Size() int {
	return xxx_messageInfo_SrvRequest.Size(m)
}
func (m *SrvRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SrvRequest proto.InternalMessageInfo

func (m *SrvRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SrvReply struct {
	Ip                   []string `protobuf:"bytes,1,rep,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrvReply) Reset()         { *m = SrvReply{} }
func (m *SrvReply) String() string { return proto.CompactTextString(m) }
func (*SrvReply) ProtoMessage()    {}
func (*SrvReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{3}
}

func (m *SrvReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvReply.Unmarshal(m, b)
}
func (m *SrvReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvReply.Marshal(b, m, deterministic)
}
func (m *SrvReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvReply.Merge(m, src)
}
func (m *SrvReply) XXX_Size() int {
	return xxx_messageInfo_SrvReply.Size(m)
}
func (m *SrvReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvReply.DiscardUnknown(m)
}

var xxx_messageInfo_SrvReply proto.InternalMessageInfo

func (m *SrvReply) GetIp() []string {
	if m != nil {
		return m.Ip
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "no_sequel.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "no_sequel.AuthReply")
	proto.RegisterType((*SrvRequest)(nil), "no_sequel.SrvRequest")
	proto.RegisterType((*SrvReply)(nil), "no_sequel.SrvReply")
}

func init() {
	proto.RegisterFile("main.proto", fileDescriptor_7ed94b0a22d11796)
}

var fileDescriptor_7ed94b0a22d11796 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x14, 0xc4, 0xdd, 0xad, 0x2b, 0xbb, 0x4f, 0x50, 0x78, 0x56, 0x59, 0x72, 0x2a, 0x39, 0xf5, 0xb4,
	0x60, 0x3d, 0x08, 0x5e, 0xc4, 0x83, 0x7f, 0x0e, 0x9e, 0xd2, 0x0f, 0x20, 0x2b, 0x7d, 0xd4, 0xe0,
	0x36, 0x89, 0x49, 0xb6, 0xb2, 0xf8, 0xe5, 0x25, 0x89, 0xad, 0x45, 0xbd, 0x78, 0x9c, 0x37, 0xf9,
	0xcd, 0x30, 0x04, 0x60, 0xd5, 0x4a, 0xd5, 0x18, 0xab, 0xbd, 0xc6, 0x4a, 0xe9, 0x27, 0x47, 0x6f,
	0x3d, 0x75, 0xfc, 0x16, 0x0e, 0x6f, 0x7a, 0xff, 0x22, 0x82, 0x72, 0x1e, 0x19, 0x94, 0xbd, 0x23,
	0xab, 0xda, 0x15, 0xd5, 0xd9, 0x24, 0x9b, 0x56, 0x62, 0xab, 0x83, 0x67, 0x5a, 0xe7, 0xde, 0xb5,
	0x5d, 0xd4, 0x79, 0xf2, 0x36, 0x9a, 0x9f, 0x43, 0x95, 0x62, 0x4c, 0x37, 0xe0, 0x18, 0x0a, 0xaf,
	0x5f, 0x49, 0x7d, 0x25, 0x24, 0x81, 0x47, 0x90, 0xcb, 0x04, 0x16, 0x22, 0x97, 0x0b, 0xce, 0x01,
	0xe6, 0x76, 0xbd, 0x29, 0xfe, 0x93, 0xe1, 0x0c, 0xca, 0xf8, 0x26, 0xa4, 0x06, 0xde, 0xd4, 0xd9,
	0x64, 0x34, 0xad, 0x44, 0x2e, 0xcd, 0xec, 0x03, 0xf6, 0x43, 0x25, 0x5e, 0x42, 0xf1, 0xa8, 0x97,
	0x52, 0xe1, 0x59, 0xb3, 0x9d, 0xd5, 0xec, 0x6c, 0x62, 0xe3, 0x5f, 0x77, 0xd3, 0x0d, 0x7c, 0x0f,
	0xaf, 0xa0, 0x14, 0xb4, 0x94, 0xce, 0x93, 0xfd, 0x2f, 0x3b, 0xbb, 0x83, 0xd1, 0xdc, 0xae, 0xf1,
	0x1a, 0x8e, 0xef, 0xc9, 0xc7, 0xfa, 0x07, 0xe9, 0xbc, 0xb6, 0x03, 0x9e, 0xee, 0x10, 0xdf, 0xfb,
	0xd8, 0xc9, 0xcf, 0x73, 0xcc, 0x79, 0x3e, 0x88, 0x1f, 0x72, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff,
	0xd2, 0x7d, 0xee, 0x63, 0x9e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	Register(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/no_sequel.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Register(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/no_sequel.Auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Login(context.Context, *AuthRequest) (*AuthReply, error)
	Register(context.Context, *AuthRequest) (*AuthReply, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Login(ctx context.Context, req *AuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServer) Register(ctx context.Context, req *AuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/no_sequel.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/no_sequel.Auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "no_sequel.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

// SrvClient is the client API for Srv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SrvClient interface {
	GetLoginHistory(ctx context.Context, in *SrvRequest, opts ...grpc.CallOption) (*SrvReply, error)
}

type srvClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvClient(cc grpc.ClientConnInterface) SrvClient {
	return &srvClient{cc}
}

func (c *srvClient) GetLoginHistory(ctx context.Context, in *SrvRequest, opts ...grpc.CallOption) (*SrvReply, error) {
	out := new(SrvReply)
	err := c.cc.Invoke(ctx, "/no_sequel.Srv/GetLoginHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvServer is the server API for Srv service.
type SrvServer interface {
	GetLoginHistory(context.Context, *SrvRequest) (*SrvReply, error)
}

// UnimplementedSrvServer can be embedded to have forward compatible implementations.
type UnimplementedSrvServer struct {
}

func (*UnimplementedSrvServer) GetLoginHistory(ctx context.Context, req *SrvRequest) (*SrvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginHistory not implemented")
}

func RegisterSrvServer(s *grpc.Server, srv SrvServer) {
	s.RegisterService(&_Srv_serviceDesc, srv)
}

func _Srv_GetLoginHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SrvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvServer).GetLoginHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/no_sequel.Srv/GetLoginHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvServer).GetLoginHistory(ctx, req.(*SrvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Srv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "no_sequel.Srv",
	HandlerType: (*SrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLoginHistory",
			Handler:    _Srv_GetLoginHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
