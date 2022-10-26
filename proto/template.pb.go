// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: proto/template.proto

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

type Lamport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock   int32  `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Lamport) Reset() {
	*x = Lamport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lamport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lamport) ProtoMessage() {}

func (x *Lamport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lamport.ProtoReflect.Descriptor instead.
func (*Lamport) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{0}
}

func (x *Lamport) GetClock() int32 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *Lamport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lamport) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[1]
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
	return file_proto_template_proto_rawDescGZIP(), []int{1}
}

var File_proto_template_proto protoreflect.FileDescriptor

var file_proto_template_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0x81, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x09, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x47, 0x6d, 0x61, 0x72, 0x73, 0x65, 0x6e, 0x2f, 0x48, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x2d, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_template_proto_rawDescOnce sync.Once
	file_proto_template_proto_rawDescData = file_proto_template_proto_rawDesc
)

func file_proto_template_proto_rawDescGZIP() []byte {
	file_proto_template_proto_rawDescOnce.Do(func() {
		file_proto_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_template_proto_rawDescData)
	})
	return file_proto_template_proto_rawDescData
}

var file_proto_template_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_template_proto_goTypes = []interface{}{
	(*Lamport)(nil), // 0: proto.Lamport
	(*Empty)(nil),   // 1: proto.Empty
}
var file_proto_template_proto_depIdxs = []int32{
	1, // 0: proto.Template.Join:input_type -> proto.Empty
	0, // 1: proto.Template.Send:input_type -> proto.Lamport
	0, // 2: proto.Template.NotifyAll:input_type -> proto.Lamport
	0, // 3: proto.Template.Join:output_type -> proto.Lamport
	1, // 4: proto.Template.Send:output_type -> proto.Empty
	1, // 5: proto.Template.NotifyAll:output_type -> proto.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_template_proto_init() }
func file_proto_template_proto_init() {
	if File_proto_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lamport); i {
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
		file_proto_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_template_proto_goTypes,
		DependencyIndexes: file_proto_template_proto_depIdxs,
		MessageInfos:      file_proto_template_proto_msgTypes,
	}.Build()
	File_proto_template_proto = out.File
	file_proto_template_proto_rawDesc = nil
	file_proto_template_proto_goTypes = nil
	file_proto_template_proto_depIdxs = nil
}
