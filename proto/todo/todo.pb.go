// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	// Types that are valid to be assigned to Response:
	//	*LoginResponse_Token
	//	*LoginResponse_Error
	Response             isLoginResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

type isLoginResponse_Response interface {
	isLoginResponse_Response()
}

type LoginResponse_Token struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3,oneof"`
}

type LoginResponse_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*LoginResponse_Token) isLoginResponse_Response() {}

func (*LoginResponse_Error) isLoginResponse_Response() {}

func (m *LoginResponse) GetResponse() isLoginResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *LoginResponse) GetToken() string {
	if x, ok := m.GetResponse().(*LoginResponse_Token); ok {
		return x.Token
	}
	return ""
}

func (m *LoginResponse) GetError() *Error {
	if x, ok := m.GetResponse().(*LoginResponse_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoginResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoginResponse_Token)(nil),
		(*LoginResponse_Error)(nil),
	}
}

type Error struct {
	ErrorMessage         string   `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	ErrorCode            int32    `protobuf:"varint,2,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Error) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() {
	proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639)
}

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xad, 0xb0, 0xa5, 0x9d, 0xb6, 0x1e, 0xf6, 0x20, 0xa5, 0x88, 0xc8, 0x1e, 0xc4, 0xd3,
	0x1e, 0x2a, 0xbe, 0x80, 0xa2, 0x28, 0xe8, 0x65, 0xe3, 0x0b, 0x44, 0x77, 0x08, 0x41, 0xdc, 0x89,
	0x33, 0x1b, 0x7d, 0x7d, 0xc9, 0x24, 0xd1, 0xe6, 0xf8, 0xfd, 0x1f, 0x7c, 0xec, 0x2c, 0x40, 0xa6,
	0x48, 0xbe, 0x61, 0xca, 0xe4, 0x1e, 0x60, 0xfd, 0x4c, 0x55, 0x9d, 0x02, 0x7e, 0xb5, 0x28, 0xd9,
	0xee, 0x60, 0xd1, 0x0a, 0x72, 0x2a, 0x3f, 0x71, 0x3b, 0xbb, 0x98, 0x5d, 0x2d, 0xc3, 0x1f, 0x77,
	0xae, 0x29, 0x45, 0x7e, 0x88, 0xe3, 0xf6, 0xb8, 0x77, 0x23, 0xbb, 0x02, 0x36, 0x43, 0x47, 0x1a,
	0x4a, 0x82, 0xf6, 0x14, 0x4c, 0xa6, 0x0f, 0x4c, 0x7d, 0xe5, 0xf1, 0x28, 0xf4, 0x68, 0xcf, 0xc1,
	0x20, 0x33, 0xb1, 0x16, 0x56, 0xfb, 0xb9, 0xbf, 0xef, 0xa8, 0xf3, 0x3a, 0xdf, 0x02, 0x2c, 0x78,
	0x68, 0xb8, 0x27, 0x30, 0x6a, 0xad, 0x83, 0xb5, 0xda, 0x17, 0x14, 0x29, 0xab, 0xf1, 0x65, 0x93,
	0xcd, 0x9e, 0xc1, 0x52, 0xf9, 0x8e, 0x22, 0x6a, 0xdc, 0x84, 0xff, 0x61, 0x7f, 0x03, 0xab, 0x57,
	0x8a, 0x54, 0x20, 0x7f, 0xd7, 0xef, 0x68, 0x2f, 0xc1, 0xe8, 0x73, 0xed, 0xc6, 0x1f, 0x9e, 0xbf,
	0x3b, 0xf1, 0x93, 0x2b, 0xde, 0xe6, 0xfa, 0x4b, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0xc6, 0x35, 0xa4, 0x33, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/TodoService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TodoService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}