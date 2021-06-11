// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: Mysql/Options/Annotations.proto

package pboptions

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var file_Mysql_Options_Annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*MysqlFile)(nil),
		Field:         5002,
		Name:          "Mysql.Options.File",
		Tag:           "bytes,5002,opt,name=File",
		Filename:      "Mysql/Options/Annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MysqlMethod)(nil),
		Field:         5002,
		Name:          "Mysql.Options.Method",
		Tag:           "bytes,5002,opt,name=Method",
		Filename:      "Mysql/Options/Annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MysqlMessage)(nil),
		Field:         5002,
		Name:          "Mysql.Options.Message",
		Tag:           "bytes,5002,opt,name=Message",
		Filename:      "Mysql/Options/Annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*MysqlService)(nil),
		Field:         5002,
		Name:          "Mysql.Options.Service",
		Tag:           "bytes,5002,opt,name=Service",
		Filename:      "Mysql/Options/Annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*MysqlField)(nil),
		Field:         5002,
		Name:          "Mysql.Options.Field",
		Tag:           "bytes,5002,opt,name=Field",
		Filename:      "Mysql/Options/Annotations.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional Mysql.Options.MysqlFile File = 5002;
	E_File = &file_Mysql_Options_Annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional Mysql.Options.MysqlMethod Method = 5002;
	E_Method = &file_Mysql_Options_Annotations_proto_extTypes[1]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional Mysql.Options.MysqlMessage Message = 5002;
	E_Message = &file_Mysql_Options_Annotations_proto_extTypes[2]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional Mysql.Options.MysqlService Service = 5002;
	E_Service = &file_Mysql_Options_Annotations_proto_extTypes[3]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional Mysql.Options.MysqlField Field = 5002;
	E_Field = &file_Mysql_Options_Annotations_proto_extTypes[4]
)

var File_Mysql_Options_Annotations_proto protoreflect.FileDescriptor

var file_Mysql_Options_Annotations_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x4b, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x27, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x3a, 0x53, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x27, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x57, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x8a, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a,
	0x57, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x27, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x4f, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x8a, 0x27, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x52, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0xa3, 0x01, 0x0a, 0x1b, 0x68, 0x66,
	0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x50, 0x42, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x4a, 0x64, 0x65, 0x76, 0x2e, 0x70, 0x68, 0x6f, 0x65,
	0x6e, 0x69, 0x78, 0x2d, 0x65, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6f, 0x6e, 0x61,
	0x6e, 0x6b, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x66, 0x64, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x70,
	0x62, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x70, 0x62, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x09, 0x50, 0x42, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0xaa, 0x02, 0x1b, 0x48, 0x66, 0x64, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_Mysql_Options_Annotations_proto_goTypes = []interface{}{
	(*descriptorpb.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptorpb.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 4: google.protobuf.FieldOptions
	(*MysqlFile)(nil),                   // 5: Mysql.Options.MysqlFile
	(*MysqlMethod)(nil),                 // 6: Mysql.Options.MysqlMethod
	(*MysqlMessage)(nil),                // 7: Mysql.Options.MysqlMessage
	(*MysqlService)(nil),                // 8: Mysql.Options.MysqlService
	(*MysqlField)(nil),                  // 9: Mysql.Options.MysqlField
}
var file_Mysql_Options_Annotations_proto_depIdxs = []int32{
	0,  // 0: Mysql.Options.File:extendee -> google.protobuf.FileOptions
	1,  // 1: Mysql.Options.Method:extendee -> google.protobuf.MethodOptions
	2,  // 2: Mysql.Options.Message:extendee -> google.protobuf.MessageOptions
	3,  // 3: Mysql.Options.Service:extendee -> google.protobuf.ServiceOptions
	4,  // 4: Mysql.Options.Field:extendee -> google.protobuf.FieldOptions
	5,  // 5: Mysql.Options.File:type_name -> Mysql.Options.MysqlFile
	6,  // 6: Mysql.Options.Method:type_name -> Mysql.Options.MysqlMethod
	7,  // 7: Mysql.Options.Message:type_name -> Mysql.Options.MysqlMessage
	8,  // 8: Mysql.Options.Service:type_name -> Mysql.Options.MysqlService
	9,  // 9: Mysql.Options.Field:type_name -> Mysql.Options.MysqlField
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	5,  // [5:10] is the sub-list for extension type_name
	0,  // [0:5] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_Mysql_Options_Annotations_proto_init() }
func file_Mysql_Options_Annotations_proto_init() {
	if File_Mysql_Options_Annotations_proto != nil {
		return
	}
	file_Mysql_Options_MysqlOptions_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Mysql_Options_Annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_Mysql_Options_Annotations_proto_goTypes,
		DependencyIndexes: file_Mysql_Options_Annotations_proto_depIdxs,
		ExtensionInfos:    file_Mysql_Options_Annotations_proto_extTypes,
	}.Build()
	File_Mysql_Options_Annotations_proto = out.File
	file_Mysql_Options_Annotations_proto_rawDesc = nil
	file_Mysql_Options_Annotations_proto_goTypes = nil
	file_Mysql_Options_Annotations_proto_depIdxs = nil
}
