// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserMessage
	UserID
	UserDeleted
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type UserRequest struct {
	Categ string `protobuf:"bytes,1,opt,name=categ" json:"categ,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto1.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetCateg() string {
	if m != nil {
		return m.Categ
	}
	return ""
}

func (m *UserRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UserMessage struct {
	User []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *UserMessage) Reset()                    { *m = UserMessage{} }
func (m *UserMessage) String() string            { return proto1.CompactTextString(m) }
func (*UserMessage) ProtoMessage()               {}
func (*UserMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserMessage) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

type UserID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *UserID) Reset()                    { *m = UserID{} }
func (m *UserID) String() string            { return proto1.CompactTextString(m) }
func (*UserID) ProtoMessage()               {}
func (*UserID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserDeleted struct {
	Deleted bool   `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *UserDeleted) Reset()                    { *m = UserDeleted{} }
func (m *UserDeleted) String() string            { return proto1.CompactTextString(m) }
func (*UserDeleted) ProtoMessage()               {}
func (*UserDeleted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserDeleted) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *UserDeleted) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto1.RegisterType((*UserRequest)(nil), "proto.UserRequest")
	proto1.RegisterType((*UserMessage)(nil), "proto.UserMessage")
	proto1.RegisterType((*UserID)(nil), "proto.UserID")
	proto1.RegisterType((*UserDeleted)(nil), "proto.UserDeleted")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserSvc service

type UserSvcClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserMessage, error)
	CreateUser(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error)
	DeleteUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserDeleted, error)
}

type userSvcClient struct {
	cc *grpc.ClientConn
}

func NewUserSvcClient(cc *grpc.ClientConn) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := grpc.Invoke(ctx, "/proto.UserSvc/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) CreateUser(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := grpc.Invoke(ctx, "/proto.UserSvc/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) DeleteUser(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserDeleted, error) {
	out := new(UserDeleted)
	err := grpc.Invoke(ctx, "/proto.UserSvc/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserSvc service

type UserSvcServer interface {
	GetUser(context.Context, *UserRequest) (*UserMessage, error)
	CreateUser(context.Context, *UserMessage) (*UserMessage, error)
	DeleteUser(context.Context, *UserID) (*UserDeleted, error)
}

func RegisterUserSvcServer(s *grpc.Server, srv UserSvcServer) {
	s.RegisterService(&_UserSvc_serviceDesc, srv)
}

func _UserSvc_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSvc/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSvc/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).CreateUser(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserSvc/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).DeleteUser(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserSvc_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserSvc_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserSvc_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto1.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xcb, 0x4a, 0xc6, 0x30,
	0x10, 0x85, 0xdb, 0x60, 0x5b, 0x1d, 0x2f, 0x8b, 0xc1, 0x45, 0x70, 0xa5, 0x59, 0xb9, 0x2a, 0x68,
	0x41, 0x71, 0x6d, 0x41, 0xba, 0x70, 0x13, 0xf1, 0x01, 0x62, 0x33, 0x94, 0x42, 0xa1, 0x9a, 0xa4,
	0x7d, 0x27, 0xdf, 0x52, 0x72, 0x29, 0x54, 0x7e, 0xfe, 0xd5, 0xcc, 0x39, 0x9c, 0x6f, 0x98, 0x03,
	0xb0, 0x58, 0x32, 0xf5, 0xb7, 0x99, 0xdd, 0x8c, 0x45, 0x18, 0xe2, 0x05, 0xce, 0x3f, 0x2d, 0x19,
	0x49, 0x3f, 0x0b, 0x59, 0x87, 0xd7, 0x50, 0xf4, 0xca, 0xd1, 0xc0, 0xf3, 0xdb, 0xfc, 0xfe, 0x4c,
	0x46, 0xe1, 0xdd, 0x55, 0x4d, 0x0b, 0x71, 0x16, 0xdd, 0x20, 0xc4, 0x5d, 0x44, 0xdf, 0xc9, 0x5a,
	0x35, 0x10, 0x22, 0x9c, 0xf8, 0xf3, 0x81, 0xbc, 0x90, 0x61, 0x17, 0x1c, 0x4a, 0x1f, 0xe9, 0x5a,
	0xbc, 0x02, 0x36, 0xea, 0x74, 0x95, 0x8d, 0x5a, 0x3c, 0x47, 0xb8, 0xa5, 0x89, 0x1c, 0x69, 0xe4,
	0x50, 0xe9, 0xb8, 0x86, 0xcc, 0xa9, 0xdc, 0x64, 0x02, 0xd9, 0x06, 0x3e, 0xfe, 0xe6, 0x50, 0x79,
	0xf2, 0x63, 0xed, 0xb1, 0x81, 0xea, 0x8d, 0x9c, 0x57, 0x88, 0xb1, 0x56, 0xbd, 0x2b, 0x73, 0xb3,
	0xf7, 0xd2, 0x97, 0x22, 0xc3, 0x27, 0x80, 0x57, 0x43, 0xca, 0xd1, 0x01, 0x97, 0x32, 0x47, 0xb8,
	0x07, 0x80, 0xf8, 0x6d, 0xe0, 0x2e, 0x77, 0x99, 0xae, 0xfd, 0x87, 0xa4, 0x4e, 0x22, 0xfb, 0x2a,
	0x83, 0xd9, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x0c, 0x7a, 0x64, 0x78, 0x01, 0x00, 0x00,
}
