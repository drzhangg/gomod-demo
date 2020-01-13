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
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x3d, 0x6b, 0xc3, 0x30,
	0x14, 0xb4, 0x15, 0x37, 0x75, 0x5f, 0x8a, 0x6b, 0x1e, 0x1d, 0x82, 0xe9, 0x10, 0x34, 0xa5, 0x1d,
	0x3c, 0x24, 0x6b, 0x97, 0x50, 0x4a, 0x08, 0x85, 0x50, 0x5c, 0x65, 0x36, 0xae, 0xa5, 0x41, 0x05,
	0x4b, 0x42, 0x7e, 0x6d, 0xe9, 0xbf, 0x2f, 0x16, 0xf9, 0x1a, 0xb2, 0xdd, 0xe9, 0xee, 0xf4, 0x8e,
	0x83, 0xac, 0xfd, 0xee, 0xc9, 0x76, 0xca, 0x97, 0xce, 0x5b, 0xb2, 0x98, 0x1e, 0x38, 0xbf, 0x07,
	0xb6, 0x91, 0x98, 0x01, 0xd3, 0x72, 0x1a, 0xcf, 0xe2, 0xf9, 0x55, 0xc5, 0xb4, 0xe4, 0x05, 0x24,
	0xdb, 0xa6, 0x53, 0x88, 0x90, 0x98, 0xa6, 0x53, 0x41, 0xb9, 0xa9, 0x02, 0xe6, 0x25, 0x24, 0xbb,
	0x5e, 0xf9, 0x4b, 0xda, 0xf0, 0x46, 0xba, 0x53, 0x53, 0x36, 0x8b, 0xe7, 0xa3, 0x2a, 0x60, 0xfe,
	0x0c, 0xe9, 0xaa, 0x25, 0xfd, 0xa3, 0xe9, 0xef, 0x62, 0xe6, 0x01, 0x18, 0xb9, 0x90, 0xc8, 0x16,
	0xb7, 0xe5, 0xb1, 0xa8, 0x70, 0x15, 0x23, 0xf7, 0xb4, 0x02, 0x26, 0x1c, 0x66, 0x00, 0xc2, 0xd5,
	0x3b, 0xf3, 0x66, 0xec, 0xaf, 0xc9, 0x23, 0xbc, 0x83, 0x89, 0x70, 0xf5, 0xd6, 0xd2, 0x07, 0x35,
	0x9e, 0xf2, 0x78, 0x6f, 0x78, 0xf7, 0xb6, 0x55, 0x7d, 0x9f, 0x33, 0x04, 0x18, 0x0b, 0x57, 0xbf,
	0x1a, 0x99, 0x8f, 0x16, 0x5f, 0x90, 0xbe, 0xec, 0x7f, 0xc5, 0x47, 0xb8, 0x5e, 0x2b, 0x0a, 0xfd,
	0xcf, 0x6e, 0x6d, 0x64, 0x91, 0x9d, 0xd8, 0xa0, 0xf2, 0x08, 0x97, 0x30, 0x59, 0x2b, 0x3a, 0x56,
	0x3f, 0x33, 0x0c, 0xd3, 0x14, 0x78, 0xe2, 0x07, 0x0f, 0x8f, 0x3e, 0xc7, 0x61, 0xdf, 0xe5, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb5, 0x2e, 0x51, 0x71, 0x01, 0x00, 0x00,
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