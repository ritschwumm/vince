// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vince/store/v1/store.proto

package v1

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

// StorePrefix defines different kind of data we are storing in the key value
// store.
type StorePrefix int32

const (
	StorePrefix_SITES           StorePrefix = 0
	StorePrefix_ALERTS          StorePrefix = 2
	StorePrefix_ACCOUNT         StorePrefix = 3
	StorePrefix_TOKEN           StorePrefix = 4
	StorePrefix_CLUSTER         StorePrefix = 8
	StorePrefix_BLOCK_METADATA  StorePrefix = 9
	StorePrefix_BLOCK_INDEX     StorePrefix = 10
	StorePrefix_OAUTH_CLIENT    StorePrefix = 11
	StorePrefix_OAUTH_AUTHORIZE StorePrefix = 12
	StorePrefix_OAUTH_ACCESS    StorePrefix = 13
	StorePrefix_OAUTH_REFRESH   StorePrefix = 14
	StorePrefix_SNIPPET         StorePrefix = 15
	StorePrefix_IMPORT          StorePrefix = 16
)

// Enum value maps for StorePrefix.
var (
	StorePrefix_name = map[int32]string{
		0:  "SITES",
		2:  "ALERTS",
		3:  "ACCOUNT",
		4:  "TOKEN",
		8:  "CLUSTER",
		9:  "BLOCK_METADATA",
		10: "BLOCK_INDEX",
		11: "OAUTH_CLIENT",
		12: "OAUTH_AUTHORIZE",
		13: "OAUTH_ACCESS",
		14: "OAUTH_REFRESH",
		15: "SNIPPET",
		16: "IMPORT",
	}
	StorePrefix_value = map[string]int32{
		"SITES":           0,
		"ALERTS":          2,
		"ACCOUNT":         3,
		"TOKEN":           4,
		"CLUSTER":         8,
		"BLOCK_METADATA":  9,
		"BLOCK_INDEX":     10,
		"OAUTH_CLIENT":    11,
		"OAUTH_AUTHORIZE": 12,
		"OAUTH_ACCESS":    13,
		"OAUTH_REFRESH":   14,
		"SNIPPET":         15,
		"IMPORT":          16,
	}
)

func (x StorePrefix) Enum() *StorePrefix {
	p := new(StorePrefix)
	*p = x
	return p
}

func (x StorePrefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorePrefix) Descriptor() protoreflect.EnumDescriptor {
	return file_vince_store_v1_store_proto_enumTypes[0].Descriptor()
}

func (StorePrefix) Type() protoreflect.EnumType {
	return &file_vince_store_v1_store_proto_enumTypes[0]
}

func (x StorePrefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorePrefix.Descriptor instead.
func (StorePrefix) EnumDescriptor() ([]byte, []int) {
	return file_vince_store_v1_store_proto_rawDescGZIP(), []int{0}
}

// Parquet fields used to store events. There are only two phisical data types
// int64 and string.
//
// The columns are grouped by data types for easy access
type Column int32

const (
	Column_bounce          Column = 0
	Column_duration        Column = 1
	Column_id              Column = 2
	Column_session         Column = 3
	Column_timestamp       Column = 4
	Column_browser         Column = 5
	Column_browser_version Column = 6
	Column_city            Column = 7
	Column_country         Column = 8
	Column_entry_page      Column = 9
	Column_event           Column = 10
	Column_exit_page       Column = 11
	Column_host            Column = 12
	Column_os              Column = 13
	Column_os_version      Column = 14
	Column_path            Column = 15
	Column_referrer        Column = 16
	Column_referrer_source Column = 17
	Column_region          Column = 18
	Column_screen          Column = 19
	Column_utm_campaign    Column = 20
	Column_utm_content     Column = 21
	Column_utm_medium      Column = 22
	Column_utm_source      Column = 23
	Column_utm_term        Column = 24
)

// Enum value maps for Column.
var (
	Column_name = map[int32]string{
		0:  "bounce",
		1:  "duration",
		2:  "id",
		3:  "session",
		4:  "timestamp",
		5:  "browser",
		6:  "browser_version",
		7:  "city",
		8:  "country",
		9:  "entry_page",
		10: "event",
		11: "exit_page",
		12: "host",
		13: "os",
		14: "os_version",
		15: "path",
		16: "referrer",
		17: "referrer_source",
		18: "region",
		19: "screen",
		20: "utm_campaign",
		21: "utm_content",
		22: "utm_medium",
		23: "utm_source",
		24: "utm_term",
	}
	Column_value = map[string]int32{
		"bounce":          0,
		"duration":        1,
		"id":              2,
		"session":         3,
		"timestamp":       4,
		"browser":         5,
		"browser_version": 6,
		"city":            7,
		"country":         8,
		"entry_page":      9,
		"event":           10,
		"exit_page":       11,
		"host":            12,
		"os":              13,
		"os_version":      14,
		"path":            15,
		"referrer":        16,
		"referrer_source": 17,
		"region":          18,
		"screen":          19,
		"utm_campaign":    20,
		"utm_content":     21,
		"utm_medium":      22,
		"utm_source":      23,
		"utm_term":        24,
	}
)

func (x Column) Enum() *Column {
	p := new(Column)
	*p = x
	return p
}

func (x Column) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Column) Descriptor() protoreflect.EnumDescriptor {
	return file_vince_store_v1_store_proto_enumTypes[1].Descriptor()
}

func (Column) Type() protoreflect.EnumType {
	return &file_vince_store_v1_store_proto_enumTypes[1]
}

func (x Column) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Column.Descriptor instead.
func (Column) EnumDescriptor() ([]byte, []int) {
	return file_vince_store_v1_store_proto_rawDescGZIP(), []int{1}
}

type StoreKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string      `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Prefix    StorePrefix `protobuf:"varint,2,opt,name=prefix,proto3,enum=v1.StorePrefix" json:"prefix,omitempty"`
}

func (x *StoreKey) Reset() {
	*x = StoreKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_store_v1_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreKey) ProtoMessage() {}

func (x *StoreKey) ProtoReflect() protoreflect.Message {
	mi := &file_vince_store_v1_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreKey.ProtoReflect.Descriptor instead.
func (*StoreKey) Descriptor() ([]byte, []int) {
	return file_vince_store_v1_store_proto_rawDescGZIP(), []int{0}
}

func (x *StoreKey) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *StoreKey) GetPrefix() StorePrefix {
	if x != nil {
		return x.Prefix
	}
	return StorePrefix_SITES
}

var File_vince_store_v1_store_proto protoreflect.FileDescriptor

var file_vince_store_v1_store_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x51, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x2a, 0xd3, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x49, 0x54, 0x45, 0x53, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x08, 0x12,
	0x12, 0x0a, 0x0e, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54,
	0x41, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x4e, 0x44,
	0x45, 0x58, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x4f,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x0d, 0x12, 0x11, 0x0a,
	0x0d, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x10, 0x0e,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4e, 0x49, 0x50, 0x50, 0x45, 0x54, 0x10, 0x0f, 0x12, 0x0a, 0x0a,
	0x06, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x10, 0x2a, 0xe1, 0x02, 0x0a, 0x06, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x06,
	0x0a, 0x02, 0x69, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x10, 0x07, 0x12, 0x0b,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x10, 0x0c, 0x12,
	0x06, 0x0a, 0x02, 0x6f, 0x73, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x0e, 0x12, 0x08, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x10,
	0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x10, 0x10, 0x12,
	0x13, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x10, 0x12,
	0x12, 0x0a, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c,
	0x75, 0x74, 0x6d, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x10, 0x14, 0x12, 0x0f,
	0x0a, 0x0b, 0x75, 0x74, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x15, 0x12,
	0x0e, 0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x16, 0x12,
	0x0e, 0x0a, 0x0a, 0x75, 0x74, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x17, 0x12,
	0x0c, 0x0a, 0x08, 0x75, 0x74, 0x6d, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x10, 0x18, 0x42, 0x79, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02,
	0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vince_store_v1_store_proto_rawDescOnce sync.Once
	file_vince_store_v1_store_proto_rawDescData = file_vince_store_v1_store_proto_rawDesc
)

func file_vince_store_v1_store_proto_rawDescGZIP() []byte {
	file_vince_store_v1_store_proto_rawDescOnce.Do(func() {
		file_vince_store_v1_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_vince_store_v1_store_proto_rawDescData)
	})
	return file_vince_store_v1_store_proto_rawDescData
}

var file_vince_store_v1_store_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vince_store_v1_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vince_store_v1_store_proto_goTypes = []interface{}{
	(StorePrefix)(0), // 0: v1.StorePrefix
	(Column)(0),      // 1: v1.Column
	(*StoreKey)(nil), // 2: v1.StoreKey
}
var file_vince_store_v1_store_proto_depIdxs = []int32{
	0, // 0: v1.StoreKey.prefix:type_name -> v1.StorePrefix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vince_store_v1_store_proto_init() }
func file_vince_store_v1_store_proto_init() {
	if File_vince_store_v1_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vince_store_v1_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreKey); i {
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
			RawDescriptor: file_vince_store_v1_store_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vince_store_v1_store_proto_goTypes,
		DependencyIndexes: file_vince_store_v1_store_proto_depIdxs,
		EnumInfos:         file_vince_store_v1_store_proto_enumTypes,
		MessageInfos:      file_vince_store_v1_store_proto_msgTypes,
	}.Build()
	File_vince_store_v1_store_proto = out.File
	file_vince_store_v1_store_proto_rawDesc = nil
	file_vince_store_v1_store_proto_goTypes = nil
	file_vince_store_v1_store_proto_depIdxs = nil
}
