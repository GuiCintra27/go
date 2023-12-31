// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_message.proto

package pb

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

type UserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7965fb6944d08275, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "user.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
}

func init() {
	proto.RegisterFile("user_message.proto", fileDescriptor_7965fb6944d08275)
}

var fileDescriptor_7965fb6944d08275 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0xcf, 0x82, 0x40,
	0x0c, 0x86, 0x3f, 0xf8, 0xd0, 0x68, 0x75, 0xb1, 0x71, 0x20, 0x4c, 0x86, 0x38, 0x38, 0x31, 0x68,
	0xfc, 0x03, 0xea, 0x2f, 0x80, 0xb8, 0xb8, 0x98, 0x43, 0x1b, 0x43, 0x14, 0xee, 0x6c, 0x0f, 0xfd,
	0xfb, 0x86, 0x3b, 0x63, 0xd8, 0xfa, 0x3e, 0x4d, 0x9e, 0xbe, 0x05, 0x6c, 0x85, 0xf8, 0x5c, 0x93,
	0x88, 0xba, 0x51, 0x66, 0x58, 0x5b, 0x8d, 0x51, 0xc7, 0xd2, 0x02, 0x26, 0x47, 0x21, 0xce, 0xe9,
	0xd9, 0x92, 0x58, 0x44, 0x88, 0x1a, 0x55, 0x53, 0x1c, 0x2c, 0x82, 0xd5, 0x38, 0x77, 0x33, 0xce,
	0x61, 0x40, 0xb5, 0xaa, 0x1e, 0x71, 0xe8, 0xa0, 0x0f, 0x98, 0xc0, 0xc8, 0x28, 0x91, 0xb7, 0xe6,
	0x6b, 0xfc, 0xef, 0x16, 0xbf, 0x9c, 0x2e, 0x61, 0xea, 0xa5, 0x62, 0x74, 0x23, 0xce, 0x60, 0xf5,
	0x9d, 0x9a, 0xaf, 0xd6, 0x87, 0xf5, 0xc1, 0x9f, 0x2e, 0x88, 0x5f, 0xd5, 0x85, 0x70, 0x0b, 0xb0,
	0x67, 0x52, 0x96, 0x3a, 0x88, 0xb3, 0xac, 0xab, 0x97, 0xf5, 0xba, 0x25, 0xd8, 0x47, 0xde, 0x9c,
	0xfe, 0xed, 0xa2, 0x53, 0x68, 0xca, 0x72, 0xe8, 0x7e, 0xda, 0x7c, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0f, 0xf3, 0xcf, 0x2d, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_message.proto",
}
