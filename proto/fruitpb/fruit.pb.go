// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: fruit.proto

package fruitpb

import (
	filterccpb "github.com/marcos-wz/capstone/proto/filterccpb"
	filterpb "github.com/marcos-wz/capstone/proto/filterpb"
	loaderpb "github.com/marcos-wz/capstone/proto/loaderpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_fruit_proto protoreflect.FileDescriptor

var file_fruit_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x72, 0x75, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63,
	0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63, 0x6f, 0x73, 0x2d, 0x77, 0x7a, 0x2f, 0x63, 0x61,
	0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x72, 0x63, 0x6f, 0x73, 0x2d, 0x77, 0x7a, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x70, 0x62,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63, 0x6f, 0x73,
	0x2d, 0x77, 0x7a, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x63, 0x63, 0x70, 0x62, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd1, 0x01, 0x0a,
	0x0c, 0x46, 0x72, 0x75, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06,
	0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x43, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x43, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x72, 0x63, 0x6f, 0x73, 0x2d, 0x77, 0x7a, 0x2f, 0x63, 0x61, 0x70, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x75, 0x69, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_fruit_proto_goTypes = []interface{}{
	(*filterpb.FilterRequest)(nil),      // 0: capstone.FilterRequest
	(*loaderpb.LoaderRequest)(nil),      // 1: capstone.LoaderRequest
	(*filterccpb.FilterCCRequest)(nil),  // 2: capstone.FilterCCRequest
	(*filterpb.FilterResponse)(nil),     // 3: capstone.FilterResponse
	(*loaderpb.LoaderResponse)(nil),     // 4: capstone.LoaderResponse
	(*filterccpb.FilterCCResponse)(nil), // 5: capstone.FilterCCResponse
}
var file_fruit_proto_depIdxs = []int32{
	0, // 0: capstone.FruitService.Filter:input_type -> capstone.FilterRequest
	1, // 1: capstone.FruitService.Loader:input_type -> capstone.LoaderRequest
	2, // 2: capstone.FruitService.FilterCC:input_type -> capstone.FilterCCRequest
	3, // 3: capstone.FruitService.Filter:output_type -> capstone.FilterResponse
	4, // 4: capstone.FruitService.Loader:output_type -> capstone.LoaderResponse
	5, // 5: capstone.FruitService.FilterCC:output_type -> capstone.FilterCCResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fruit_proto_init() }
func file_fruit_proto_init() {
	if File_fruit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fruit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fruit_proto_goTypes,
		DependencyIndexes: file_fruit_proto_depIdxs,
	}.Build()
	File_fruit_proto = out.File
	file_fruit_proto_rawDesc = nil
	file_fruit_proto_goTypes = nil
	file_fruit_proto_depIdxs = nil
}