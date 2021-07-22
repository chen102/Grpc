// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ordermanagement.proto

package errors

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items       []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"` //这个字段在消息中可以重复出现任意次，包括0次
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination string   `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordermanagement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_ordermanagement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_ordermanagement_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

var File_ordermanagement_proto protoreflect.FileDescriptor

var file_ordermanagement_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb5, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x67, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x33,
	0x0a, 0x08, 0x61, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x01, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ordermanagement_proto_rawDescOnce sync.Once
	file_ordermanagement_proto_rawDescData = file_ordermanagement_proto_rawDesc
)

func file_ordermanagement_proto_rawDescGZIP() []byte {
	file_ordermanagement_proto_rawDescOnce.Do(func() {
		file_ordermanagement_proto_rawDescData = protoimpl.X.CompressGZIP(file_ordermanagement_proto_rawDescData)
	})
	return file_ordermanagement_proto_rawDescData
}

var file_ordermanagement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ordermanagement_proto_goTypes = []interface{}{
	(*Order)(nil),                  // 0: pb.Order
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_ordermanagement_proto_depIdxs = []int32{
	1, // 0: pb.OrderManagement.getOrder:input_type -> google.protobuf.StringValue
	0, // 1: pb.OrderManagement.addOrder:input_type -> pb.Order
	1, // 2: pb.OrderManagement.searchOrder:input_type -> google.protobuf.StringValue
	0, // 3: pb.OrderManagement.getOrder:output_type -> pb.Order
	1, // 4: pb.OrderManagement.addOrder:output_type -> google.protobuf.StringValue
	0, // 5: pb.OrderManagement.searchOrder:output_type -> pb.Order
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ordermanagement_proto_init() }
func file_ordermanagement_proto_init() {
	if File_ordermanagement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ordermanagement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ordermanagement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ordermanagement_proto_goTypes,
		DependencyIndexes: file_ordermanagement_proto_depIdxs,
		MessageInfos:      file_ordermanagement_proto_msgTypes,
	}.Build()
	File_ordermanagement_proto = out.File
	file_ordermanagement_proto_rawDesc = nil
	file_ordermanagement_proto_goTypes = nil
	file_ordermanagement_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error)
	AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	SearchOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrderClient, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/pb.OrderManagement/getOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) AddOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/pb.OrderManagement/addOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) SearchOrder(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[0], "/pb.OrderManagement/searchOrder", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementSearchOrderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_SearchOrderClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementSearchOrderClient struct {
	grpc.ClientStream
}

func (x *orderManagementSearchOrderClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
type OrderManagementServer interface {
	GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error)
	AddOrder(context.Context, *Order) (*wrapperspb.StringValue, error)
	SearchOrder(*wrapperspb.StringValue, OrderManagement_SearchOrderServer) error
}

// UnimplementedOrderManagementServer can be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (*UnimplementedOrderManagementServer) GetOrder(context.Context, *wrapperspb.StringValue) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (*UnimplementedOrderManagementServer) AddOrder(context.Context, *Order) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (*UnimplementedOrderManagementServer) SearchOrder(*wrapperspb.StringValue, OrderManagement_SearchOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrder not implemented")
}

func RegisterOrderManagementServer(s *grpc.Server, srv OrderManagementServer) {
	s.RegisterService(&_OrderManagement_serviceDesc, srv)
}

func _OrderManagement_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderManagement/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrder(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderManagement/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).AddOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_SearchOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrder(m, &orderManagementSearchOrderServer{stream})
}

type OrderManagement_SearchOrderServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementSearchOrderServer struct {
	grpc.ServerStream
}

func (x *orderManagementSearchOrderServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

var _OrderManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrder",
			Handler:    _OrderManagement_GetOrder_Handler,
		},
		{
			MethodName: "addOrder",
			Handler:    _OrderManagement_AddOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrder",
			Handler:       _OrderManagement_SearchOrder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ordermanagement.proto",
}
