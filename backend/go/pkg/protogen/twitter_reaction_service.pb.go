// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: twitter_reaction_service.proto

package protogen

import (
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

var File_twitter_reaction_service_proto protoreflect.FileDescriptor

var file_twitter_reaction_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a,
	0x02, 0x0a, 0x16, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0d,
	0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x14, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4f, 0x6e, 0x6c, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x0d, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x1c, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x50, 0x6f, 0x6c, 0x6c,
	0x12, 0x18, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x4f, 0x6e, 0x6c, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x0d, 0x2e, 0x70, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_twitter_reaction_service_proto_goTypes = []interface{}{
	(*Format_NoParam)(nil),   // 0: pmodel.Format.NoParam
	(*Format_OnlyTitle)(nil), // 1: pmodel.Format.OnlyTitle
	(*Empty)(nil),            // 2: pmodel.Empty
}
var file_twitter_reaction_service_proto_depIdxs = []int32{
	0, // 0: TwitterServiceReaction.postTweet:input_type -> pmodel.Format.NoParam
	1, // 1: TwitterServiceReaction.postTweetWithContent:input_type -> pmodel.Format.OnlyTitle
	0, // 2: TwitterServiceReaction.postTweetWithPoll:input_type -> pmodel.Format.NoParam
	1, // 3: TwitterServiceReaction.postTweetWithContentWithPoll:input_type -> pmodel.Format.OnlyTitle
	2, // 4: TwitterServiceReaction.postTweet:output_type -> pmodel.Empty
	2, // 5: TwitterServiceReaction.postTweetWithContent:output_type -> pmodel.Empty
	2, // 6: TwitterServiceReaction.postTweetWithPoll:output_type -> pmodel.Empty
	2, // 7: TwitterServiceReaction.postTweetWithContentWithPoll:output_type -> pmodel.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_twitter_reaction_service_proto_init() }
func file_twitter_reaction_service_proto_init() {
	if File_twitter_reaction_service_proto != nil {
		return
	}
	file_pmodel_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_twitter_reaction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_twitter_reaction_service_proto_goTypes,
		DependencyIndexes: file_twitter_reaction_service_proto_depIdxs,
	}.Build()
	File_twitter_reaction_service_proto = out.File
	file_twitter_reaction_service_proto_rawDesc = nil
	file_twitter_reaction_service_proto_goTypes = nil
	file_twitter_reaction_service_proto_depIdxs = nil
}
