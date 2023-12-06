// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: ethereum_action_service.proto

package protogen

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

type AddressWatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddressWatcher) Reset() {
	*x = AddressWatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_action_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressWatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressWatcher) ProtoMessage() {}

func (x *AddressWatcher) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_action_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressWatcher.ProtoReflect.Descriptor instead.
func (*AddressWatcher) Descriptor() ([]byte, []int) {
	return file_ethereum_action_service_proto_rawDescGZIP(), []int{0}
}

type EventWatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventWatcher) Reset() {
	*x = EventWatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_action_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventWatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventWatcher) ProtoMessage() {}

func (x *EventWatcher) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_action_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventWatcher.ProtoReflect.Descriptor instead.
func (*EventWatcher) Descriptor() ([]byte, []int) {
	return file_ethereum_action_service_proto_rawDescGZIP(), []int{1}
}

type AddressWatcher_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResponseType Format_Kind `protobuf:"varint,2,opt,name=responseType,proto3,enum=pmodel.Format_Kind" json:"responseType,omitempty"`
	Address      string      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressWatcher_Request) Reset() {
	*x = AddressWatcher_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_action_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressWatcher_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressWatcher_Request) ProtoMessage() {}

func (x *AddressWatcher_Request) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_action_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressWatcher_Request.ProtoReflect.Descriptor instead.
func (*AddressWatcher_Request) Descriptor() ([]byte, []int) {
	return file_ethereum_action_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AddressWatcher_Request) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddressWatcher_Request) GetResponseType() Format_Kind {
	if x != nil {
		return x.ResponseType
	}
	return Format_UndefinedKind
}

func (x *AddressWatcher_Request) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type EventWatcher_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResponseType Format_Kind `protobuf:"varint,2,opt,name=responseType,proto3,enum=pmodel.Format_Kind" json:"responseType,omitempty"`
	Event        string      `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventWatcher_Request) Reset() {
	*x = EventWatcher_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethereum_action_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventWatcher_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventWatcher_Request) ProtoMessage() {}

func (x *EventWatcher_Request) ProtoReflect() protoreflect.Message {
	mi := &file_ethereum_action_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventWatcher_Request.ProtoReflect.Descriptor instead.
func (*EventWatcher_Request) Descriptor() ([]byte, []int) {
	return file_ethereum_action_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *EventWatcher_Request) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventWatcher_Request) GetResponseType() Format_Kind {
	if x != nil {
		return x.ResponseType
	}
	return Format_UndefinedKind
}

func (x *EventWatcher_Request) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

var File_ethereum_action_service_proto protoreflect.FileDescriptor

var file_ethereum_action_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x1a,
	0x6c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x78, 0x0a,
	0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x68, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x9c, 0x01, 0x0a, 0x15, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x43, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethereum_action_service_proto_rawDescOnce sync.Once
	file_ethereum_action_service_proto_rawDescData = file_ethereum_action_service_proto_rawDesc
)

func file_ethereum_action_service_proto_rawDescGZIP() []byte {
	file_ethereum_action_service_proto_rawDescOnce.Do(func() {
		file_ethereum_action_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethereum_action_service_proto_rawDescData)
	})
	return file_ethereum_action_service_proto_rawDescData
}

var file_ethereum_action_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ethereum_action_service_proto_goTypes = []interface{}{
	(*AddressWatcher)(nil),         // 0: AddressWatcher
	(*EventWatcher)(nil),           // 1: EventWatcher
	(*AddressWatcher_Request)(nil), // 2: AddressWatcher.Request
	(*EventWatcher_Request)(nil),   // 3: EventWatcher.Request
	(Format_Kind)(0),               // 4: pmodel.Format.Kind
	(*Empty)(nil),                  // 5: pmodel.Empty
}
var file_ethereum_action_service_proto_depIdxs = []int32{
	4, // 0: AddressWatcher.Request.responseType:type_name -> pmodel.Format.Kind
	4, // 1: EventWatcher.Request.responseType:type_name -> pmodel.Format.Kind
	2, // 2: EthereumServiceAction.RegisterAddresseWatcher:input_type -> AddressWatcher.Request
	3, // 3: EthereumServiceAction.RegisterEventWatcher:input_type -> EventWatcher.Request
	5, // 4: EthereumServiceAction.RegisterAddresseWatcher:output_type -> pmodel.Empty
	5, // 5: EthereumServiceAction.RegisterEventWatcher:output_type -> pmodel.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ethereum_action_service_proto_init() }
func file_ethereum_action_service_proto_init() {
	if File_ethereum_action_service_proto != nil {
		return
	}
	file_pmodel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ethereum_action_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressWatcher); i {
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
		file_ethereum_action_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventWatcher); i {
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
		file_ethereum_action_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressWatcher_Request); i {
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
		file_ethereum_action_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventWatcher_Request); i {
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
			RawDescriptor: file_ethereum_action_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ethereum_action_service_proto_goTypes,
		DependencyIndexes: file_ethereum_action_service_proto_depIdxs,
		MessageInfos:      file_ethereum_action_service_proto_msgTypes,
	}.Build()
	File_ethereum_action_service_proto = out.File
	file_ethereum_action_service_proto_rawDesc = nil
	file_ethereum_action_service_proto_goTypes = nil
	file_ethereum_action_service_proto_depIdxs = nil
}