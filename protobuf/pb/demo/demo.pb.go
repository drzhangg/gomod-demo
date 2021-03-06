// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package demo

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

type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

type UserReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserReq) Reset()         { *m = UserReq{} }
func (m *UserReq) String() string { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()    {}
func (*UserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *UserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReq.Unmarshal(m, b)
}
func (m *UserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReq.Marshal(b, m, deterministic)
}
func (m *UserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReq.Merge(m, src)
}
func (m *UserReq) XXX_Size() int {
	return xxx_messageInfo_UserReq.Size(m)
}
func (m *UserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserReq proto.InternalMessageInfo

func (m *UserReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserRsp struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  string   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRsp) Reset()         { *m = UserRsp{} }
func (m *UserRsp) String() string { return proto.CompactTextString(m) }
func (*UserRsp) ProtoMessage()    {}
func (*UserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *UserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRsp.Unmarshal(m, b)
}
func (m *UserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRsp.Marshal(b, m, deterministic)
}
func (m *UserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRsp.Merge(m, src)
}
func (m *UserRsp) XXX_Size() int {
	return xxx_messageInfo_UserRsp.Size(m)
}
func (m *UserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRsp proto.InternalMessageInfo

func (m *UserRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRsp) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

type Phone struct {
	Type                 PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PhoneType" json:"type,omitempty"`
	Number               string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phones               []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

type ContactBook struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{4}
}

func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (m *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(m, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*UserReq)(nil), "pb.UserReq")
	proto.RegisterType((*UserRsp)(nil), "pb.UserRsp")
	proto.RegisterType((*Phone)(nil), "pb.Phone")
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*ContactBook)(nil), "pb.ContactBook")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xff, 0x66, 0x82, 0x25, 0xcc, 0x41, 0xa2, 0x17, 0xdb, 0x45, 0xa1, 0x08, 0xa6,
	0x90, 0x7e, 0x83, 0x88, 0x28, 0x88, 0xa4, 0x2c, 0x8a, 0xe0, 0x2d, 0x31, 0x83, 0x16, 0x49, 0x76,
	0xcd, 0xc6, 0x43, 0xbf, 0xbd, 0xec, 0xba, 0xc6, 0xde, 0xde, 0xce, 0x6f, 0xe6, 0xed, 0xe3, 0x01,
	0xb4, 0xd4, 0x89, 0x5c, 0x0e, 0x62, 0x14, 0xe8, 0xca, 0x86, 0x9d, 0x42, 0xf4, 0xac, 0x68, 0xe0,
	0xf4, 0x85, 0x73, 0x70, 0x77, 0x6d, 0xe6, 0x2c, 0x9c, 0x55, 0xcc, 0xdd, 0x5d, 0xcb, 0xd6, 0x16,
	0x29, 0x89, 0x08, 0x7e, 0x5f, 0x77, 0x64, 0xa1, 0xd1, 0x98, 0x82, 0x57, 0xbf, 0x53, 0xe6, 0x9a,
	0x91, 0x96, 0xac, 0x84, 0x60, 0xfb, 0x21, 0x7a, 0xc2, 0x25, 0xf8, 0xe3, 0x5e, 0xfe, 0xae, 0xcf,
	0x8b, 0xe3, 0x5c, 0x36, 0xb9, 0x01, 0x4f, 0x7b, 0x49, 0xdc, 0x20, 0x3c, 0x81, 0xb0, 0xff, 0xee,
	0x1a, 0x1a, 0xac, 0x81, 0x7d, 0xb1, 0x0a, 0xc2, 0x2d, 0x0d, 0x4a, 0xf4, 0x07, 0x71, 0x02, 0x1d,
	0x67, 0xca, 0xe0, 0x1e, 0x64, 0x58, 0x42, 0x28, 0xb5, 0xb1, 0xca, 0xbc, 0x85, 0xb7, 0x4a, 0x8a,
	0x78, 0xfa, 0x8a, 0x5b, 0xc0, 0x36, 0x90, 0xdc, 0x88, 0x7e, 0xac, 0xdf, 0xc6, 0x52, 0x88, 0x4f,
	0xbc, 0x80, 0x48, 0x1a, 0x7f, 0x95, 0x39, 0xe6, 0x04, 0xcc, 0x89, 0x19, 0xf1, 0x3f, 0x74, 0x75,
	0x0e, 0xf1, 0x14, 0x18, 0x8f, 0xc0, 0xbf, 0xaf, 0x1e, 0x6f, 0xd3, 0x99, 0x56, 0x2f, 0x15, 0x7f,
	0x48, 0x9d, 0xe2, 0x1a, 0x7c, 0x5d, 0x24, 0x5e, 0x42, 0x74, 0x47, 0xa3, 0xae, 0x09, 0x13, 0x6d,
	0x64, 0xbb, 0x3c, 0xfb, 0x7f, 0x28, 0xc9, 0x66, 0x65, 0xf4, 0x1a, 0xac, 0xf5, 0x7e, 0x13, 0x9a,
	0xe6, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x29, 0xb2, 0x9a, 0x87, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoClient is the server API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoClient interface {
	GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRsp, error)
}

type demoClient struct {
	cc *grpc.ClientConn
}

func NewDemoClient(cc *grpc.ClientConn) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) GetUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserRsp, error) {
	out := new(UserRsp)
	err := c.cc.Invoke(ctx, "/pb.demo/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo service.
type DemoServer interface {
	GetUser(context.Context, *UserReq) (*UserRsp, error)
}

// UnimplementedDemoServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServer struct {
}

func (*UnimplementedDemoServer) GetUser(ctx context.Context, req *UserReq) (*UserRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterDemoServer(s *grpc.Server, srv DemoServer) {
	s.RegisterService(&_Demo_serviceDesc, srv)
}

func _Demo_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.demo/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).GetUser(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Demo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Demo_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
