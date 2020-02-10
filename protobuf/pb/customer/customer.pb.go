// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

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

type Tp int32

const (
	Tp_Tp_UnKnown  Tp = 0
	Tp_Tp_NotStart Tp = 1
	Tp_Tp_Process  Tp = 2
	Tp_Tp_End      Tp = 3
)

var Tp_name = map[int32]string{
	0: "Tp_UnKnown",
	1: "Tp_NotStart",
	2: "Tp_Process",
	3: "Tp_End",
}

var Tp_value = map[string]int32{
	"Tp_UnKnown":  0,
	"Tp_NotStart": 1,
	"Tp_Process":  2,
	"Tp_End":      3,
}

func (x Tp) String() string {
	return proto.EnumName(Tp_name, int32(x))
}

func (Tp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Name struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Activity struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tp                   Tp       `protobuf:"varint,2,opt,name=tp,proto3,enum=customer.Tp" json:"tp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Activity) GetTp() Tp {
	if m != nil {
		return m.Tp
	}
	return Tp_Tp_UnKnown
}

func init() {
	proto.RegisterEnum("customer.Tp", Tp_name, Tp_value)
	proto.RegisterType((*Id)(nil), "customer.Id")
	proto.RegisterType((*Name)(nil), "customer.Name")
	proto.RegisterType((*User)(nil), "customer.User")
	proto.RegisterType((*Activity)(nil), "customer.Activity")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0x4e, 0xae, 0xb1, 0xa6, 0x2f, 0x12, 0xc3, 0xc3, 0xa1, 0x04, 0x87, 0x72, 0x53, 0x75, 0x88,
	0xd0, 0xae, 0x2e, 0x55, 0xa4, 0x14, 0xa1, 0x48, 0x4c, 0x17, 0x97, 0x10, 0x93, 0x1b, 0x4e, 0xc8,
	0xdd, 0x71, 0x79, 0x2a, 0xfe, 0xf7, 0x92, 0xa3, 0x69, 0x3b, 0x74, 0x7b, 0xdf, 0xaf, 0xbb, 0x8f,
	0x0f, 0xe2, 0xfa, 0xbb, 0x23, 0xdd, 0x0a, 0x9b, 0x19, 0xab, 0x49, 0x63, 0x38, 0x60, 0x7e, 0x03,
	0x6c, 0xd3, 0x60, 0x0c, 0x4c, 0x36, 0x53, 0x7f, 0xe6, 0xcf, 0x2f, 0x72, 0x26, 0x1b, 0x9e, 0x42,
	0xb0, 0xad, 0x5a, 0x81, 0x08, 0x81, 0xaa, 0x5a, 0xe1, 0x94, 0x49, 0xee, 0x6e, 0x9e, 0x41, 0xb0,
	0xeb, 0x84, 0x3d, 0xa7, 0xf5, 0x1c, 0xc9, 0x56, 0x4c, 0xd9, 0xcc, 0x9f, 0x8f, 0x72, 0x77, 0xf3,
	0x47, 0x08, 0x57, 0x35, 0xc9, 0x1f, 0x49, 0x7f, 0x67, 0x33, 0xb7, 0xc0, 0xc8, 0xb8, 0x44, 0xbc,
	0xb8, 0xca, 0x0e, 0x45, 0x0b, 0x93, 0x33, 0x32, 0xf7, 0x2b, 0x60, 0x85, 0xc1, 0x18, 0xa0, 0x30,
	0xe5, 0x4e, 0xbd, 0x2a, 0xfd, 0xab, 0x12, 0x0f, 0xaf, 0x21, 0x2a, 0x4c, 0xb9, 0xd5, 0xf4, 0x4e,
	0x95, 0xa5, 0xc4, 0xdf, 0x1b, 0xde, 0xac, 0xae, 0x45, 0xd7, 0x25, 0x0c, 0x01, 0xc6, 0x85, 0x29,
	0x5f, 0x54, 0x93, 0x8c, 0x16, 0x5f, 0x10, 0x3e, 0xef, 0x5f, 0xc5, 0x3b, 0xb8, 0x5c, 0x0b, 0x72,
	0xfd, 0x4f, 0xfe, 0xda, 0x34, 0x69, 0x7c, 0x44, 0xbd, 0xca, 0x3d, 0x5c, 0x42, 0xb4, 0x16, 0x74,
	0xa8, 0x7e, 0x62, 0xe8, 0xa7, 0x49, 0xf1, 0x88, 0x07, 0x0f, 0xf7, 0x9e, 0xa2, 0x8f, 0xc9, 0xc3,
	0xc0, 0x7f, 0x8e, 0xdd, 0xd8, 0xcb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x4e, 0x52, 0x0d,
	0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error)
	GetActivity(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Activity, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/customer.Customer/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetActivity(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/customer.Customer/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	GetUser(context.Context, *Id) (*User, error)
	GetActivity(context.Context, *Name) (*Activity, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) GetUser(ctx context.Context, req *Id) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedCustomerServer) GetActivity(ctx context.Context, req *Name) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetActivity(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Customer_GetUser_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _Customer_GetActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}