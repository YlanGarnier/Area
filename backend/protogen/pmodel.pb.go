// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pmodel.proto

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

type Format_Kind int32

const (
	Format_UndefinedKind           Format_Kind = 0
	Format_MessageAndDestinaryKind Format_Kind = 1
	Format_GHIncidentReportKind    Format_Kind = 2
	Format_OnlyTitleKind           Format_Kind = 3
	Format_ManyFilesKind           Format_Kind = 4
	Format_TagsKind                Format_Kind = 5
	Format_BranchKind              Format_Kind = 6
)

// Enum value maps for Format_Kind.
var (
	Format_Kind_name = map[int32]string{
		0: "UndefinedKind",
		1: "MessageAndDestinaryKind",
		2: "GHIncidentReportKind",
		3: "OnlyTitleKind",
		4: "ManyFilesKind",
		5: "TagsKind",
		6: "BranchKind",
	}
	Format_Kind_value = map[string]int32{
		"UndefinedKind":           0,
		"MessageAndDestinaryKind": 1,
		"GHIncidentReportKind":    2,
		"OnlyTitleKind":           3,
		"ManyFilesKind":           4,
		"TagsKind":                5,
		"BranchKind":              6,
	}
)

func (x Format_Kind) Enum() *Format_Kind {
	p := new(Format_Kind)
	*p = x
	return p
}

func (x Format_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_pmodel_proto_enumTypes[0].Descriptor()
}

func (Format_Kind) Type() protoreflect.EnumType {
	return &file_pmodel_proto_enumTypes[0]
}

func (x Format_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format_Kind.Descriptor instead.
func (Format_Kind) EnumDescriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 0}
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Base) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1}
}

// empty
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{2}
}

type Format_MessageAndDestinary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destinary string `protobuf:"bytes,1,opt,name=destinary,proto3" json:"destinary,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Format_MessageAndDestinary) Reset() {
	*x = Format_MessageAndDestinary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_MessageAndDestinary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_MessageAndDestinary) ProtoMessage() {}

func (x *Format_MessageAndDestinary) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_MessageAndDestinary.ProtoReflect.Descriptor instead.
func (*Format_MessageAndDestinary) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Format_MessageAndDestinary) GetDestinary() string {
	if x != nil {
		return x.Destinary
	}
	return ""
}

func (x *Format_MessageAndDestinary) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Format_GHIncidentReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base    *Base  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Format_GHIncidentReport) Reset() {
	*x = Format_GHIncidentReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_GHIncidentReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_GHIncidentReport) ProtoMessage() {}

func (x *Format_GHIncidentReport) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_GHIncidentReport.ProtoReflect.Descriptor instead.
func (*Format_GHIncidentReport) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Format_GHIncidentReport) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Format_GHIncidentReport) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Format_GHIncidentReport) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Format_OnlyTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *Base  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Format_OnlyTitle) Reset() {
	*x = Format_OnlyTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_OnlyTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_OnlyTitle) ProtoMessage() {}

func (x *Format_OnlyTitle) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_OnlyTitle.ProtoReflect.Descriptor instead.
func (*Format_OnlyTitle) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Format_OnlyTitle) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Format_OnlyTitle) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Format_ManyFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *Base                    `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Files []*Format_ManyFiles_File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Format_ManyFiles) Reset() {
	*x = Format_ManyFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_ManyFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_ManyFiles) ProtoMessage() {}

func (x *Format_ManyFiles) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_ManyFiles.ProtoReflect.Descriptor instead.
func (*Format_ManyFiles) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Format_ManyFiles) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Format_ManyFiles) GetFiles() []*Format_ManyFiles_File {
	if x != nil {
		return x.Files
	}
	return nil
}

type Format_Tags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base    `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Format_Tags) Reset() {
	*x = Format_Tags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_Tags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_Tags) ProtoMessage() {}

func (x *Format_Tags) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_Tags.ProtoReflect.Descriptor instead.
func (*Format_Tags) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 4}
}

func (x *Format_Tags) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Format_Tags) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Format_Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *Format_Branch) Reset() {
	*x = Format_Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_Branch) ProtoMessage() {}

func (x *Format_Branch) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_Branch.ProtoReflect.Descriptor instead.
func (*Format_Branch) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 5}
}

func (x *Format_Branch) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type Format_ManyFiles_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Format_ManyFiles_File) Reset() {
	*x = Format_ManyFiles_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmodel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format_ManyFiles_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format_ManyFiles_File) ProtoMessage() {}

func (x *Format_ManyFiles_File) ProtoReflect() protoreflect.Message {
	mi := &file_pmodel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format_ManyFiles_File.ProtoReflect.Descriptor instead.
func (*Format_ManyFiles_File) Descriptor() ([]byte, []int) {
	return file_pmodel_proto_rawDescGZIP(), []int{1, 3, 0}
}

func (x *Format_ManyFiles_File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Format_ManyFiles_File) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_pmodel_proto protoreflect.FileDescriptor

var file_pmodel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x34, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x9e, 0x05, 0x0a,
	0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a, 0x4d, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x64, 0x0a, 0x10, 0x47, 0x48, 0x49, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x43, 0x0a, 0x09,
	0x4f, 0x6e, 0x6c, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x1a, 0x98, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x34, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x3c, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x2a, 0x0a, 0x06, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x4b, 0x69, 0x6e, 0x64,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x64,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x47, 0x48, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c,
	0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4b, 0x69, 0x6e, 0x64, 0x10, 0x06, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pmodel_proto_rawDescOnce sync.Once
	file_pmodel_proto_rawDescData = file_pmodel_proto_rawDesc
)

func file_pmodel_proto_rawDescGZIP() []byte {
	file_pmodel_proto_rawDescOnce.Do(func() {
		file_pmodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_pmodel_proto_rawDescData)
	})
	return file_pmodel_proto_rawDescData
}

var file_pmodel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pmodel_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pmodel_proto_goTypes = []interface{}{
	(Format_Kind)(0),                   // 0: pmodel.Format.Kind
	(*Base)(nil),                       // 1: pmodel.Base
	(*Format)(nil),                     // 2: pmodel.Format
	(*Empty)(nil),                      // 3: pmodel.Empty
	(*Format_MessageAndDestinary)(nil), // 4: pmodel.Format.MessageAndDestinary
	(*Format_GHIncidentReport)(nil),    // 5: pmodel.Format.GHIncidentReport
	(*Format_OnlyTitle)(nil),           // 6: pmodel.Format.OnlyTitle
	(*Format_ManyFiles)(nil),           // 7: pmodel.Format.ManyFiles
	(*Format_Tags)(nil),                // 8: pmodel.Format.Tags
	(*Format_Branch)(nil),              // 9: pmodel.Format.Branch
	(*Format_ManyFiles_File)(nil),      // 10: pmodel.Format.ManyFiles.File
}
var file_pmodel_proto_depIdxs = []int32{
	1,  // 0: pmodel.Format.GHIncidentReport.base:type_name -> pmodel.Base
	1,  // 1: pmodel.Format.OnlyTitle.base:type_name -> pmodel.Base
	1,  // 2: pmodel.Format.ManyFiles.base:type_name -> pmodel.Base
	10, // 3: pmodel.Format.ManyFiles.files:type_name -> pmodel.Format.ManyFiles.File
	1,  // 4: pmodel.Format.Tags.base:type_name -> pmodel.Base
	1,  // 5: pmodel.Format.Branch.base:type_name -> pmodel.Base
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pmodel_proto_init() }
func file_pmodel_proto_init() {
	if File_pmodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pmodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_pmodel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
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
		file_pmodel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pmodel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_MessageAndDestinary); i {
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
		file_pmodel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_GHIncidentReport); i {
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
		file_pmodel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_OnlyTitle); i {
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
		file_pmodel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_ManyFiles); i {
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
		file_pmodel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_Tags); i {
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
		file_pmodel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_Branch); i {
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
		file_pmodel_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format_ManyFiles_File); i {
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
			RawDescriptor: file_pmodel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pmodel_proto_goTypes,
		DependencyIndexes: file_pmodel_proto_depIdxs,
		EnumInfos:         file_pmodel_proto_enumTypes,
		MessageInfos:      file_pmodel_proto_msgTypes,
	}.Build()
	File_pmodel_proto = out.File
	file_pmodel_proto_rawDesc = nil
	file_pmodel_proto_goTypes = nil
	file_pmodel_proto_depIdxs = nil
}
