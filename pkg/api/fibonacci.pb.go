// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/fibonacci.proto

package api

import (
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

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_fibonacci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_fibonacci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_grpc_fibonacci_proto_rawDescGZIP(), []int{0}
}

func (x *FibonacciRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *FibonacciRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []int64 `protobuf:"varint,1,rep,packed,name=result,proto3" json:"result,omitempty"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_fibonacci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_fibonacci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_grpc_fibonacci_proto_rawDescGZIP(), []int{1}
}

func (x *FibonacciResponse) GetResult() []int64 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_grpc_fibonacci_proto protoreflect.FileDescriptor

var file_grpc_fibonacci_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2e, 0x0a, 0x10,
	0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x2b, 0x0a, 0x11,
	0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x53, 0x0a, 0x09, 0x46, 0x69, 0x62,
	0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x12, 0x46, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e,
	0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x6b,
	0x69, 0x74, 0x61, 0x39, 0x39, 0x31, 0x30, 0x30, 0x2f, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_fibonacci_proto_rawDescOnce sync.Once
	file_grpc_fibonacci_proto_rawDescData = file_grpc_fibonacci_proto_rawDesc
)

func file_grpc_fibonacci_proto_rawDescGZIP() []byte {
	file_grpc_fibonacci_proto_rawDescOnce.Do(func() {
		file_grpc_fibonacci_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_fibonacci_proto_rawDescData)
	})
	return file_grpc_fibonacci_proto_rawDescData
}

var file_grpc_fibonacci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_fibonacci_proto_goTypes = []interface{}{
	(*FibonacciRequest)(nil),  // 0: api.FibonacciRequest
	(*FibonacciResponse)(nil), // 1: api.FibonacciResponse
}
var file_grpc_fibonacci_proto_depIdxs = []int32{
	0, // 0: api.Fibonacci.FibonacciSequence:input_type -> api.FibonacciRequest
	1, // 1: api.Fibonacci.FibonacciSequence:output_type -> api.FibonacciResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_fibonacci_proto_init() }
func file_grpc_fibonacci_proto_init() {
	if File_grpc_fibonacci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_fibonacci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciRequest); i {
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
		file_grpc_fibonacci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciResponse); i {
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
			RawDescriptor: file_grpc_fibonacci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_fibonacci_proto_goTypes,
		DependencyIndexes: file_grpc_fibonacci_proto_depIdxs,
		MessageInfos:      file_grpc_fibonacci_proto_msgTypes,
	}.Build()
	File_grpc_fibonacci_proto = out.File
	file_grpc_fibonacci_proto_rawDesc = nil
	file_grpc_fibonacci_proto_goTypes = nil
	file_grpc_fibonacci_proto_depIdxs = nil
}
