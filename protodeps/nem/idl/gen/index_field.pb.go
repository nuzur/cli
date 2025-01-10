// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.28.3
// source: index_field.proto

package gen

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

type IndexFieldOrder int32

const (
	IndexFieldOrder_INDEX_FIELD_ORDER_INVALID IndexFieldOrder = 0
	IndexFieldOrder_INDEX_FIELD_ORDER_ASC     IndexFieldOrder = 1
	IndexFieldOrder_INDEX_FIELD_ORDER_DESC    IndexFieldOrder = 2
)

// Enum value maps for IndexFieldOrder.
var (
	IndexFieldOrder_name = map[int32]string{
		0: "INDEX_FIELD_ORDER_INVALID",
		1: "INDEX_FIELD_ORDER_ASC",
		2: "INDEX_FIELD_ORDER_DESC",
	}
	IndexFieldOrder_value = map[string]int32{
		"INDEX_FIELD_ORDER_INVALID": 0,
		"INDEX_FIELD_ORDER_ASC":     1,
		"INDEX_FIELD_ORDER_DESC":    2,
	}
)

func (x IndexFieldOrder) Enum() *IndexFieldOrder {
	p := new(IndexFieldOrder)
	*p = x
	return p
}

func (x IndexFieldOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexFieldOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_index_field_proto_enumTypes[0].Descriptor()
}

func (IndexFieldOrder) Type() protoreflect.EnumType {
	return &file_index_field_proto_enumTypes[0]
}

func (x IndexFieldOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexFieldOrder.Descriptor instead.
func (IndexFieldOrder) EnumDescriptor() ([]byte, []int) {
	return file_index_field_proto_rawDescGZIP(), []int{0}
}

type IndexField struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FieldUuid     string                 `protobuf:"bytes,1,opt,name=field_uuid,json=fieldUuid,proto3" json:"field_uuid,omitempty"`
	Priority      int64                  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Order         IndexFieldOrder        `protobuf:"varint,3,opt,name=order,proto3,enum=nem.IndexFieldOrder" json:"order,omitempty"`
	Length        int64                  `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexField) Reset() {
	*x = IndexField{}
	mi := &file_index_field_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexField) ProtoMessage() {}

func (x *IndexField) ProtoReflect() protoreflect.Message {
	mi := &file_index_field_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexField.ProtoReflect.Descriptor instead.
func (*IndexField) Descriptor() ([]byte, []int) {
	return file_index_field_proto_rawDescGZIP(), []int{0}
}

func (x *IndexField) GetFieldUuid() string {
	if x != nil {
		return x.FieldUuid
	}
	return ""
}

func (x *IndexField) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *IndexField) GetOrder() IndexFieldOrder {
	if x != nil {
		return x.Order
	}
	return IndexFieldOrder_INDEX_FIELD_ORDER_INVALID
}

func (x *IndexField) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

var File_index_field_proto protoreflect.FileDescriptor

var file_index_field_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x8b, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2a, 0x67, 0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x44,
	0x45, 0x58, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x44, 0x45,
	0x58, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53,
	0x43, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42,
	0x20, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_field_proto_rawDescOnce sync.Once
	file_index_field_proto_rawDescData = file_index_field_proto_rawDesc
)

func file_index_field_proto_rawDescGZIP() []byte {
	file_index_field_proto_rawDescOnce.Do(func() {
		file_index_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_field_proto_rawDescData)
	})
	return file_index_field_proto_rawDescData
}

var file_index_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_index_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_index_field_proto_goTypes = []any{
	(IndexFieldOrder)(0), // 0: nem.IndexFieldOrder
	(*IndexField)(nil),   // 1: nem.IndexField
}
var file_index_field_proto_depIdxs = []int32{
	0, // 0: nem.IndexField.order:type_name -> nem.IndexFieldOrder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_index_field_proto_init() }
func file_index_field_proto_init() {
	if File_index_field_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_index_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_index_field_proto_goTypes,
		DependencyIndexes: file_index_field_proto_depIdxs,
		EnumInfos:         file_index_field_proto_enumTypes,
		MessageInfos:      file_index_field_proto_msgTypes,
	}.Build()
	File_index_field_proto = out.File
	file_index_field_proto_rawDesc = nil
	file_index_field_proto_goTypes = nil
	file_index_field_proto_depIdxs = nil
}
