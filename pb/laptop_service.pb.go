// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateHelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateHelloRequest) Reset()         { *m = CreateHelloRequest{} }
func (m *CreateHelloRequest) String() string { return proto.CompactTextString(m) }
func (*CreateHelloRequest) ProtoMessage()    {}
func (*CreateHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{2}
}

func (m *CreateHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHelloRequest.Unmarshal(m, b)
}
func (m *CreateHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHelloRequest.Marshal(b, m, deterministic)
}
func (m *CreateHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHelloRequest.Merge(m, src)
}
func (m *CreateHelloRequest) XXX_Size() int {
	return xxx_messageInfo_CreateHelloRequest.Size(m)
}
func (m *CreateHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHelloRequest proto.InternalMessageInfo

func (m *CreateHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateHelloResponse struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateHelloResponse) Reset()         { *m = CreateHelloResponse{} }
func (m *CreateHelloResponse) String() string { return proto.CompactTextString(m) }
func (*CreateHelloResponse) ProtoMessage()    {}
func (*CreateHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{3}
}

func (m *CreateHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHelloResponse.Unmarshal(m, b)
}
func (m *CreateHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHelloResponse.Marshal(b, m, deterministic)
}
func (m *CreateHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHelloResponse.Merge(m, src)
}
func (m *CreateHelloResponse) XXX_Size() int {
	return xxx_messageInfo_CreateHelloResponse.Size(m)
}
func (m *CreateHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHelloResponse proto.InternalMessageInfo

func (m *CreateHelloResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

type CreateFindRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFindRequest) Reset()         { *m = CreateFindRequest{} }
func (m *CreateFindRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFindRequest) ProtoMessage()    {}
func (*CreateFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{4}
}

func (m *CreateFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFindRequest.Unmarshal(m, b)
}
func (m *CreateFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFindRequest.Marshal(b, m, deterministic)
}
func (m *CreateFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFindRequest.Merge(m, src)
}
func (m *CreateFindRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFindRequest.Size(m)
}
func (m *CreateFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFindRequest proto.InternalMessageInfo

func (m *CreateFindRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateFindResponse struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFindResponse) Reset()         { *m = CreateFindResponse{} }
func (m *CreateFindResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFindResponse) ProtoMessage()    {}
func (*CreateFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{5}
}

func (m *CreateFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFindResponse.Unmarshal(m, b)
}
func (m *CreateFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFindResponse.Marshal(b, m, deterministic)
}
func (m *CreateFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFindResponse.Merge(m, src)
}
func (m *CreateFindResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFindResponse.Size(m)
}
func (m *CreateFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFindResponse proto.InternalMessageInfo

func (m *CreateFindResponse) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "lemonade.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "lemonade.CreateLaptopResponse")
	proto.RegisterType((*CreateHelloRequest)(nil), "lemonade.CreateHelloRequest")
	proto.RegisterType((*CreateHelloResponse)(nil), "lemonade.CreateHelloResponse")
	proto.RegisterType((*CreateFindRequest)(nil), "lemonade.CreateFindRequest")
	proto.RegisterType((*CreateFindResponse)(nil), "lemonade.CreateFindResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x48, 0x8b, 0x8e, 0x55, 0x74, 0xcc, 0x41, 0x62, 0x2b, 0xb2, 0x82, 0x04, 0x84,
	0x1c, 0xea, 0x5d, 0xc1, 0x42, 0xf1, 0x20, 0x08, 0xf5, 0xe6, 0x45, 0x52, 0x33, 0x87, 0xc0, 0x36,
	0xbb, 0x66, 0x57, 0xc1, 0x1f, 0xed, 0x7f, 0x10, 0x77, 0x76, 0x35, 0x86, 0x16, 0xbc, 0x25, 0xf3,
	0xbe, 0x79, 0x33, 0xf3, 0x58, 0x48, 0x65, 0xa9, 0xad, 0xd2, 0xcf, 0x86, 0xda, 0xf7, 0xfa, 0x85,
	0x0a, 0xdd, 0x2a, 0xab, 0x70, 0x5b, 0xd2, 0x4a, 0x35, 0x65, 0x45, 0xd9, 0x88, 0x75, 0xae, 0x8b,
	0x1b, 0x38, 0x9a, 0xb5, 0x54, 0x5a, 0xba, 0x77, 0xd5, 0x05, 0xbd, 0xbe, 0x91, 0xb1, 0x98, 0xc3,
	0x90, 0xb1, 0xe3, 0xe8, 0x2c, 0xca, 0x77, 0xa7, 0x07, 0x45, 0xe8, 0x2f, 0x3c, 0xe8, 0x75, 0x71,
	0x01, 0xe9, 0x5f, 0x03, 0xa3, 0x55, 0x63, 0x08, 0xf7, 0x21, 0xae, 0x2b, 0xd7, 0xbd, 0xb3, 0x88,
	0xeb, 0x4a, 0xe4, 0x80, 0xcc, 0xdd, 0x91, 0x94, 0x2a, 0xcc, 0x41, 0x48, 0x9a, 0x72, 0x45, 0x9e,
	0x73, 0xdf, 0xe2, 0x32, 0xac, 0xe4, 0x49, 0x6f, 0x98, 0xc2, 0xa0, 0x25, 0x2d, 0x3f, 0x3c, 0xcb,
	0x3f, 0xe2, 0x1c, 0x0e, 0x19, 0x9e, 0xd7, 0x4d, 0x15, 0x5c, 0xfb, 0xb3, 0xaf, 0xc3, 0x6c, 0x86,
	0xbc, 0xe1, 0xbf, 0x6f, 0x9c, 0x7e, 0x46, 0xb0, 0xc7, 0xa5, 0x47, 0x0e, 0x15, 0x1f, 0x60, 0xd4,
	0xbd, 0x1a, 0x27, 0xbf, 0xbd, 0x6b, 0xe2, 0xcc, 0x4e, 0x37, 0xc9, 0xbc, 0x8a, 0xd8, 0xc2, 0x39,
	0x0c, 0xdc, 0xb9, 0x38, 0xee, 0xa3, 0xdd, 0xbc, 0xb2, 0xc9, 0x06, 0xf5, 0xc7, 0x67, 0x06, 0xc9,
	0xf7, 0x91, 0x78, 0xd2, 0x07, 0x3b, 0xf9, 0x64, 0xe3, 0xf5, 0x62, 0x30, 0xb9, 0x4d, 0x9e, 0x62,
	0xbd, 0x5c, 0x0e, 0xdd, 0x0b, 0xb9, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x14, 0x42, 0xd4, 0x77,
	0x51, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	Hello(ctx context.Context, in *CreateHelloRequest, opts ...grpc.CallOption) (*CreateHelloResponse, error)
	Find(ctx context.Context, in *CreateFindRequest, opts ...grpc.CallOption) (*CreateFindResponse, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/lemonade.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) Hello(ctx context.Context, in *CreateHelloRequest, opts ...grpc.CallOption) (*CreateHelloResponse, error) {
	out := new(CreateHelloResponse)
	err := c.cc.Invoke(ctx, "/lemonade.LaptopService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) Find(ctx context.Context, in *CreateFindRequest, opts ...grpc.CallOption) (*CreateFindResponse, error) {
	out := new(CreateFindResponse)
	err := c.cc.Invoke(ctx, "/lemonade.LaptopService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	Hello(context.Context, *CreateHelloRequest) (*CreateHelloResponse, error)
	Find(context.Context, *CreateFindRequest) (*CreateFindResponse, error)
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) Hello(ctx context.Context, req *CreateHelloRequest) (*CreateHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedLaptopServiceServer) Find(ctx context.Context, req *CreateFindRequest) (*CreateFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade.LaptopService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).Hello(ctx, req.(*CreateHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemonade.LaptopService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).Find(ctx, req.(*CreateFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lemonade.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _LaptopService_Hello_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _LaptopService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laptop_service.proto",
}
