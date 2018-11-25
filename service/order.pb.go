// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/freightio/backend/service/order.proto

package backend

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// contact person name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tel  string `protobuf:"bytes,3,opt,name=tel,proto3" json:"tel,omitempty"`
	// freight type
	Type                 string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	From                 string            `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Tos                  []string          `protobuf:"bytes,6,rep,name=tos" json:"tos,omitempty"`
	Fee                  float32           `protobuf:"fixed32,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,8,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created              uint64            `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_433a343decaa5e82, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *Order) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Order) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Order) GetTos() []string {
	if m != nil {
		return m.Tos
	}
	return nil
}

func (m *Order) GetFee() float32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Order) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Order) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type OrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_433a343decaa5e82, []int{1}
}
func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (dst *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(dst, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OrderList struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderList) Reset()         { *m = OrderList{} }
func (m *OrderList) String() string { return proto.CompactTextString(m) }
func (*OrderList) ProtoMessage()    {}
func (*OrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_433a343decaa5e82, []int{2}
}
func (m *OrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderList.Unmarshal(m, b)
}
func (m *OrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderList.Marshal(b, m, deterministic)
}
func (dst *OrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderList.Merge(dst, src)
}
func (m *OrderList) XXX_Size() int {
	return xxx_messageInfo_OrderList.Size(m)
}
func (m *OrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderList proto.InternalMessageInfo

func (m *OrderList) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "backend.Order")
	proto.RegisterMapType((map[string]string)(nil), "backend.Order.AnnotationsEntry")
	proto.RegisterType((*OrderRequest)(nil), "backend.OrderRequest")
	proto.RegisterType((*OrderList)(nil), "backend.OrderList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrdersClient interface {
	Add(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	Get(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error)
	Update(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	List(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderList, error)
	Delete(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type ordersClient struct {
	cc *grpc.ClientConn
}

func NewOrdersClient(cc *grpc.ClientConn) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) Add(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/backend.Orders/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) Get(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/backend.Orders/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) Update(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/backend.Orders/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) List(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderList, error) {
	out := new(OrderList)
	err := c.cc.Invoke(ctx, "/backend.Orders/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersClient) Delete(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/backend.Orders/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
type OrdersServer interface {
	Add(context.Context, *Order) (*Order, error)
	Get(context.Context, *OrderRequest) (*Order, error)
	Update(context.Context, *Order) (*Order, error)
	List(context.Context, *Order) (*OrderList, error)
	Delete(context.Context, *OrderRequest) (*Order, error)
}

func RegisterOrdersServer(s *grpc.Server, srv OrdersServer) {
	s.RegisterService(&_Orders_serviceDesc, srv)
}

func _Orders_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Orders/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Add(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Orders/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Get(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Orders/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Update(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Orders/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).List(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orders_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Orders/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).Delete(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Orders_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Orders_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Orders_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Orders_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Orders_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/freightio/backend/service/order.proto",
}

func init() {
	proto.RegisterFile("github.com/freightio/backend/service/order.proto", fileDescriptor_order_433a343decaa5e82)
}

var fileDescriptor_order_433a343decaa5e82 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6a, 0xf2, 0x40,
	0x14, 0x35, 0x3f, 0xc6, 0xcf, 0xeb, 0x87, 0xc8, 0xf0, 0x7d, 0x30, 0xb8, 0x68, 0x43, 0x16, 0x6d,
	0x28, 0x25, 0x69, 0x75, 0x53, 0xba, 0x28, 0x08, 0x2d, 0xdd, 0x14, 0x0a, 0x81, 0x3e, 0x40, 0xcc,
	0x5c, 0x75, 0x50, 0x33, 0x76, 0x66, 0x14, 0x7c, 0xa1, 0x3e, 0x5e, 0x9f, 0xa1, 0xcc, 0x24, 0xd2,
	0x1a, 0xba, 0x70, 0x77, 0xee, 0xb9, 0xe7, 0x0c, 0xe7, 0x5c, 0x06, 0x6e, 0xe6, 0x5c, 0x2f, 0xb6,
	0xd3, 0xa4, 0x10, 0xeb, 0x74, 0x26, 0x91, 0xcf, 0x17, 0x9a, 0x8b, 0x74, 0x9a, 0x17, 0x4b, 0x2c,
	0x59, 0xaa, 0x50, 0xee, 0x78, 0x81, 0xa9, 0x90, 0x0c, 0x65, 0xb2, 0x91, 0x42, 0x0b, 0xd2, 0xa9,
	0x97, 0xd1, 0x87, 0x0b, 0xed, 0x57, 0xb3, 0x20, 0x7d, 0x70, 0x39, 0xa3, 0x4e, 0xe8, 0xc4, 0xdd,
	0xcc, 0xe5, 0x8c, 0x10, 0xf0, 0xcb, 0x7c, 0x8d, 0xd4, 0xb5, 0x8c, 0xc5, 0x64, 0x00, 0x9e, 0xc6,
	0x15, 0xf5, 0x2c, 0x65, 0xa0, 0x51, 0xe9, 0xfd, 0x06, 0xa9, 0x5f, 0xa9, 0x0c, 0x36, 0xdc, 0x4c,
	0x8a, 0x35, 0x6d, 0x57, 0x9c, 0xc1, 0xd6, 0x29, 0x14, 0x0d, 0x42, 0xcf, 0x3a, 0x85, 0x32, 0xcc,
	0x0c, 0x91, 0x76, 0x42, 0x27, 0x76, 0x33, 0x03, 0xc9, 0x04, 0x7a, 0x79, 0x59, 0x0a, 0x9d, 0x6b,
	0x2e, 0x4a, 0x45, 0xff, 0x84, 0x5e, 0xdc, 0x1b, 0x9d, 0x27, 0x75, 0xd4, 0xc4, 0xc6, 0x4c, 0x26,
	0xdf, 0x8a, 0xa7, 0x52, 0xcb, 0x7d, 0xf6, 0xd3, 0x43, 0x28, 0x74, 0x0a, 0x89, 0xb9, 0x46, 0x46,
	0xbb, 0xa1, 0x13, 0xfb, 0xd9, 0x61, 0x1c, 0x3e, 0xc0, 0xa0, 0x69, 0x35, 0x11, 0x96, 0xb8, 0xaf,
	0x3b, 0x1b, 0x48, 0xfe, 0x41, 0x7b, 0x97, 0xaf, 0xb6, 0x87, 0xd6, 0xd5, 0x70, 0xef, 0xde, 0x39,
	0xd1, 0x19, 0xfc, 0xb5, 0x01, 0x32, 0x7c, 0xdf, 0xa2, 0xd2, 0xcd, 0x73, 0x45, 0x63, 0xe8, 0xda,
	0xfd, 0x0b, 0x57, 0x9a, 0x5c, 0x40, 0x60, 0xaf, 0xad, 0xa8, 0x63, 0x4b, 0xf4, 0x8f, 0x4b, 0x64,
	0xf5, 0x76, 0xf4, 0xe9, 0x40, 0x60, 0x19, 0x45, 0x2e, 0xc1, 0x9b, 0x30, 0x46, 0x1a, 0xca, 0x61,
	0x63, 0x8e, 0x5a, 0x24, 0x01, 0xef, 0x19, 0x35, 0xf9, 0xdf, 0x78, 0xb2, 0x8a, 0xf5, 0x8b, 0xfe,
	0x0a, 0x82, 0xb7, 0x0d, 0xcb, 0x35, 0x9e, 0xf0, 0xf6, 0x35, 0xf8, 0x36, 0x7f, 0x53, 0x49, 0x8e,
	0x67, 0xa3, 0x89, 0x5a, 0xe4, 0x16, 0x82, 0x47, 0x5c, 0xa1, 0xc6, 0x93, 0xc3, 0x4c, 0x03, 0xfb,
	0xfd, 0xc6, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xcb, 0xd1, 0x74, 0xb2, 0x02, 0x00, 0x00,
}