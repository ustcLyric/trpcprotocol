// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: user.proto

package trpcprotocol

import (
	reflect "reflect"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x62, 0x65, 0x61, 0x72, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x93, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x62, 0x65,
	0x61, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x62, 0x65, 0x61, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x62, 0x65, 0x61, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x62, 0x65, 0x61, 0x72, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x22, 0x00, 0x12, 0x54, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x31, 0x12, 0x23, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x62, 0x65, 0x61, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x23, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x62, 0x65,
	0x61, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x74, 0x63, 0x4c, 0x79, 0x72, 0x69, 0x63,
	0x2f, 0x74, 0x72, 0x70, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_user_proto_goTypes = []any{
	(*GetUserInfoReq)(nil), // 0: trpc.polarbear.user.GetUserInfoReq
}
var file_user_proto_depIdxs = []int32{
	0, // 0: trpc.polarbear.user.UserService.GetUserInfo:input_type -> trpc.polarbear.user.GetUserInfoReq
	0, // 1: trpc.polarbear.user.UserService.Hello:input_type -> trpc.polarbear.user.GetUserInfoReq
	0, // 2: trpc.polarbear.user.UserService.Hello1:input_type -> trpc.polarbear.user.GetUserInfoReq
	0, // 3: trpc.polarbear.user.UserService.GetUserInfo:output_type -> trpc.polarbear.user.GetUserInfoReq
	0, // 4: trpc.polarbear.user.UserService.Hello:output_type -> trpc.polarbear.user.GetUserInfoReq
	0, // 5: trpc.polarbear.user.UserService.Hello1:output_type -> trpc.polarbear.user.GetUserInfoReq
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
