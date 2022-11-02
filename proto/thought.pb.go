//nolint

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: thought.proto

package proto

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

type Thought struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId                string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	CreationTime             int64  `protobuf:"varint,2,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ServerMergedMediaPath    string `protobuf:"bytes,11,opt,name=serverMergedMediaPath,proto3" json:"serverMergedMediaPath,omitempty"`
	ServerMergedMediaSize    int64  `protobuf:"varint,12,opt,name=serverMergedMediaSize,proto3" json:"serverMergedMediaSize,omitempty"`
	ServerExportedMediaPath  string `protobuf:"bytes,13,opt,name=serverExportedMediaPath,proto3" json:"serverExportedMediaPath,omitempty"`
	ServerExportedMediaSize  int64  `protobuf:"varint,14,opt,name=serverExportedMediaSize,proto3" json:"serverExportedMediaSize,omitempty"`
	ServerThumbnailImagePath string `protobuf:"bytes,15,opt,name=serverThumbnailImagePath,proto3" json:"serverThumbnailImagePath,omitempty"`
	ServerThumbnailSize      int64  `protobuf:"varint,16,opt,name=serverThumbnailSize,proto3" json:"serverThumbnailSize,omitempty"`
	ServerHDMediaPath        string `protobuf:"bytes,17,opt,name=serverHDMediaPath,proto3" json:"serverHDMediaPath,omitempty"`
	ServerHDMediaSize        int64  `protobuf:"varint,18,opt,name=serverHDMediaSize,proto3" json:"serverHDMediaSize,omitempty"`
	ServerSDMediaPath        string `protobuf:"bytes,19,opt,name=serverSDMediaPath,proto3" json:"serverSDMediaPath,omitempty"`
	ServerSDMediaSize        int64  `protobuf:"varint,20,opt,name=serverSDMediaSize,proto3" json:"serverSDMediaSize,omitempty"`
}

func (x *Thought) Reset() {
	*x = Thought{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thought_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thought) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thought) ProtoMessage() {}

func (x *Thought) ProtoReflect() protoreflect.Message {
	mi := &file_thought_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thought.ProtoReflect.Descriptor instead.
func (*Thought) Descriptor() ([]byte, []int) {
	return file_thought_proto_rawDescGZIP(), []int{0}
}

func (x *Thought) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Thought) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *Thought) GetServerMergedMediaPath() string {
	if x != nil {
		return x.ServerMergedMediaPath
	}
	return ""
}

func (x *Thought) GetServerMergedMediaSize() int64 {
	if x != nil {
		return x.ServerMergedMediaSize
	}
	return 0
}

func (x *Thought) GetServerExportedMediaPath() string {
	if x != nil {
		return x.ServerExportedMediaPath
	}
	return ""
}

func (x *Thought) GetServerExportedMediaSize() int64 {
	if x != nil {
		return x.ServerExportedMediaSize
	}
	return 0
}

func (x *Thought) GetServerThumbnailImagePath() string {
	if x != nil {
		return x.ServerThumbnailImagePath
	}
	return ""
}

func (x *Thought) GetServerThumbnailSize() int64 {
	if x != nil {
		return x.ServerThumbnailSize
	}
	return 0
}

func (x *Thought) GetServerHDMediaPath() string {
	if x != nil {
		return x.ServerHDMediaPath
	}
	return ""
}

func (x *Thought) GetServerHDMediaSize() int64 {
	if x != nil {
		return x.ServerHDMediaSize
	}
	return 0
}

func (x *Thought) GetServerSDMediaPath() string {
	if x != nil {
		return x.ServerSDMediaPath
	}
	return ""
}

func (x *Thought) GetServerSDMediaSize() int64 {
	if x != nil {
		return x.ServerSDMediaSize
	}
	return 0
}

var File_thought_proto protoreflect.FileDescriptor

var file_thought_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x04, 0x0a, 0x07, 0x54, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x30, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x44, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x44, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x48, 0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c,
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x44, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6f, 0x67, 0x77, 0x61, 0x79, 0x6c,
	0x61, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thought_proto_rawDescOnce sync.Once
	file_thought_proto_rawDescData = file_thought_proto_rawDesc
)

func file_thought_proto_rawDescGZIP() []byte {
	file_thought_proto_rawDescOnce.Do(func() {
		file_thought_proto_rawDescData = protoimpl.X.CompressGZIP(file_thought_proto_rawDescData)
	})
	return file_thought_proto_rawDescData
}

var file_thought_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_thought_proto_goTypes = []interface{}{
	(*Thought)(nil), // 0: proto.Thought
}
var file_thought_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_thought_proto_init() }
func file_thought_proto_init() {
	if File_thought_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thought_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thought); i {
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
			RawDescriptor: file_thought_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thought_proto_goTypes,
		DependencyIndexes: file_thought_proto_depIdxs,
		MessageInfos:      file_thought_proto_msgTypes,
	}.Build()
	File_thought_proto = out.File
	file_thought_proto_rawDesc = nil
	file_thought_proto_goTypes = nil
	file_thought_proto_depIdxs = nil
}
