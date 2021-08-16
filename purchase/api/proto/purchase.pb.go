// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: purchase/api/proto/purchase.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PurchaseResponse_Error int32

const (
	PurchaseResponse_None       PurchaseResponse_Error = 0
	PurchaseResponse_Info       PurchaseResponse_Error = 1
	PurchaseResponse_Credit     PurchaseResponse_Error = 2
	PurchaseResponse_Permission PurchaseResponse_Error = 3
	PurchaseResponse_Unknown    PurchaseResponse_Error = 4
)

// Enum value maps for PurchaseResponse_Error.
var (
	PurchaseResponse_Error_name = map[int32]string{
		0: "None",
		1: "Info",
		2: "Credit",
		3: "Permission",
		4: "Unknown",
	}
	PurchaseResponse_Error_value = map[string]int32{
		"None":       0,
		"Info":       1,
		"Credit":     2,
		"Permission": 3,
		"Unknown":    4,
	}
)

func (x PurchaseResponse_Error) Enum() *PurchaseResponse_Error {
	p := new(PurchaseResponse_Error)
	*p = x
	return p
}

func (x PurchaseResponse_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PurchaseResponse_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_purchase_api_proto_purchase_proto_enumTypes[0].Descriptor()
}

func (PurchaseResponse_Error) Type() protoreflect.EnumType {
	return &file_purchase_api_proto_purchase_proto_enumTypes[0]
}

func (x PurchaseResponse_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PurchaseResponse_Error.Descriptor instead.
func (PurchaseResponse_Error) EnumDescriptor() ([]byte, []int) {
	return file_purchase_api_proto_purchase_proto_rawDescGZIP(), []int{1, 0}
}

type PayByCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPhoneNo string `protobuf:"bytes,1,opt,name=user_phone_no,json=userPhoneNo,proto3" json:"user_phone_no,omitempty"`
	CartNumber  string `protobuf:"bytes,2,opt,name=cart_number,json=cartNumber,proto3" json:"cart_number,omitempty"`
	Amount      int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PayByCartRequest) Reset() {
	*x = PayByCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_api_proto_purchase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayByCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayByCartRequest) ProtoMessage() {}

func (x *PayByCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_api_proto_purchase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayByCartRequest.ProtoReflect.Descriptor instead.
func (*PayByCartRequest) Descriptor() ([]byte, []int) {
	return file_purchase_api_proto_purchase_proto_rawDescGZIP(), []int{0}
}

func (x *PayByCartRequest) GetUserPhoneNo() string {
	if x != nil {
		return x.UserPhoneNo
	}
	return ""
}

func (x *PayByCartRequest) GetCartNumber() string {
	if x != nil {
		return x.CartNumber
	}
	return ""
}

func (x *PayByCartRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPhoneNo string                 `protobuf:"bytes,1,opt,name=user_phone_no,json=userPhoneNo,proto3" json:"user_phone_no,omitempty"`
	Error       PurchaseResponse_Error `protobuf:"varint,2,opt,name=error,proto3,enum=proto.PurchaseResponse_Error" json:"error,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_api_proto_purchase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_api_proto_purchase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_purchase_api_proto_purchase_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseResponse) GetUserPhoneNo() string {
	if x != nil {
		return x.UserPhoneNo
	}
	return ""
}

func (x *PurchaseResponse) GetError() PurchaseResponse_Error {
	if x != nil {
		return x.Error
	}
	return PurchaseResponse_None
}

var File_purchase_api_proto_purchase_proto protoreflect.FileDescriptor

var file_purchase_api_proto_purchase_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x10, 0x50, 0x61,
	0x79, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x10,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x6f, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x04, 0x32,
	0x49, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x50,
	0x61, 0x79, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x79, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchase_api_proto_purchase_proto_rawDescOnce sync.Once
	file_purchase_api_proto_purchase_proto_rawDescData = file_purchase_api_proto_purchase_proto_rawDesc
)

func file_purchase_api_proto_purchase_proto_rawDescGZIP() []byte {
	file_purchase_api_proto_purchase_proto_rawDescOnce.Do(func() {
		file_purchase_api_proto_purchase_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchase_api_proto_purchase_proto_rawDescData)
	})
	return file_purchase_api_proto_purchase_proto_rawDescData
}

var file_purchase_api_proto_purchase_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_purchase_api_proto_purchase_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_purchase_api_proto_purchase_proto_goTypes = []interface{}{
	(PurchaseResponse_Error)(0), // 0: proto.PurchaseResponse.Error
	(*PayByCartRequest)(nil),    // 1: proto.PayByCartRequest
	(*PurchaseResponse)(nil),    // 2: proto.PurchaseResponse
}
var file_purchase_api_proto_purchase_proto_depIdxs = []int32{
	0, // 0: proto.PurchaseResponse.error:type_name -> proto.PurchaseResponse.Error
	1, // 1: proto.Purchase.PayByCart:input_type -> proto.PayByCartRequest
	2, // 2: proto.Purchase.PayByCart:output_type -> proto.PurchaseResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_purchase_api_proto_purchase_proto_init() }
func file_purchase_api_proto_purchase_proto_init() {
	if File_purchase_api_proto_purchase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_purchase_api_proto_purchase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayByCartRequest); i {
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
		file_purchase_api_proto_purchase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseResponse); i {
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
			RawDescriptor: file_purchase_api_proto_purchase_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_purchase_api_proto_purchase_proto_goTypes,
		DependencyIndexes: file_purchase_api_proto_purchase_proto_depIdxs,
		EnumInfos:         file_purchase_api_proto_purchase_proto_enumTypes,
		MessageInfos:      file_purchase_api_proto_purchase_proto_msgTypes,
	}.Build()
	File_purchase_api_proto_purchase_proto = out.File
	file_purchase_api_proto_purchase_proto_rawDesc = nil
	file_purchase_api_proto_purchase_proto_goTypes = nil
	file_purchase_api_proto_purchase_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PurchaseClient is the client API for Purchase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PurchaseClient interface {
	PayByCart(ctx context.Context, in *PayByCartRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type purchaseClient struct {
	cc grpc.ClientConnInterface
}

func NewPurchaseClient(cc grpc.ClientConnInterface) PurchaseClient {
	return &purchaseClient{cc}
}

func (c *purchaseClient) PayByCart(ctx context.Context, in *PayByCartRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/proto.Purchase/PayByCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PurchaseServer is the server API for Purchase service.
type PurchaseServer interface {
	PayByCart(context.Context, *PayByCartRequest) (*PurchaseResponse, error)
}

// UnimplementedPurchaseServer can be embedded to have forward compatible implementations.
type UnimplementedPurchaseServer struct {
}

func (*UnimplementedPurchaseServer) PayByCart(context.Context, *PayByCartRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayByCart not implemented")
}

func RegisterPurchaseServer(s *grpc.Server, srv PurchaseServer) {
	s.RegisterService(&_Purchase_serviceDesc, srv)
}

func _Purchase_PayByCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayByCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).PayByCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Purchase/PayByCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).PayByCart(ctx, req.(*PayByCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Purchase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Purchase",
	HandlerType: (*PurchaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayByCart",
			Handler:    _Purchase_PayByCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "purchase/api/proto/purchase.proto",
}