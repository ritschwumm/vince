// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: staples/v1/scan.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filters_Projection int32

const (
	Filters_Timestamp      Filters_Projection = 0
	Filters_ID             Filters_Projection = 1
	Filters_Bounce         Filters_Projection = 2
	Filters_Session        Filters_Projection = 3
	Filters_Duration       Filters_Projection = 4
	Filters_Browser        Filters_Projection = 5
	Filters_BrowserVersion Filters_Projection = 6
	Filters_City           Filters_Projection = 7
	Filters_Country        Filters_Projection = 8
	Filters_Domain         Filters_Projection = 9
	Filters_EntryPage      Filters_Projection = 10
	Filters_ExitPage       Filters_Projection = 11
	Filters_Host           Filters_Projection = 12
	Filters_Event          Filters_Projection = 13
	Filters_Os             Filters_Projection = 14
	Filters_OsVersion      Filters_Projection = 15
	Filters_Path           Filters_Projection = 16
	Filters_Referrer       Filters_Projection = 17
	Filters_ReferrerSource Filters_Projection = 18
	Filters_Region         Filters_Projection = 19
	Filters_Screen         Filters_Projection = 20
	Filters_UtmCampaign    Filters_Projection = 21
	Filters_UtmContent     Filters_Projection = 22
	Filters_UtmMedium      Filters_Projection = 23
	Filters_UtmSource      Filters_Projection = 24
	Filters_UtmTerm        Filters_Projection = 25
)

// Enum value maps for Filters_Projection.
var (
	Filters_Projection_name = map[int32]string{
		0:  "Timestamp",
		1:  "ID",
		2:  "Bounce",
		3:  "Session",
		4:  "Duration",
		5:  "Browser",
		6:  "BrowserVersion",
		7:  "City",
		8:  "Country",
		9:  "Domain",
		10: "EntryPage",
		11: "ExitPage",
		12: "Host",
		13: "Event",
		14: "Os",
		15: "OsVersion",
		16: "Path",
		17: "Referrer",
		18: "ReferrerSource",
		19: "Region",
		20: "Screen",
		21: "UtmCampaign",
		22: "UtmContent",
		23: "UtmMedium",
		24: "UtmSource",
		25: "UtmTerm",
	}
	Filters_Projection_value = map[string]int32{
		"Timestamp":      0,
		"ID":             1,
		"Bounce":         2,
		"Session":        3,
		"Duration":       4,
		"Browser":        5,
		"BrowserVersion": 6,
		"City":           7,
		"Country":        8,
		"Domain":         9,
		"EntryPage":      10,
		"ExitPage":       11,
		"Host":           12,
		"Event":          13,
		"Os":             14,
		"OsVersion":      15,
		"Path":           16,
		"Referrer":       17,
		"ReferrerSource": 18,
		"Region":         19,
		"Screen":         20,
		"UtmCampaign":    21,
		"UtmContent":     22,
		"UtmMedium":      23,
		"UtmSource":      24,
		"UtmTerm":        25,
	}
)

func (x Filters_Projection) Enum() *Filters_Projection {
	p := new(Filters_Projection)
	*p = x
	return p
}

func (x Filters_Projection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Filters_Projection) Descriptor() protoreflect.EnumDescriptor {
	return file_staples_v1_scan_proto_enumTypes[0].Descriptor()
}

func (Filters_Projection) Type() protoreflect.EnumType {
	return &file_staples_v1_scan_proto_enumTypes[0]
}

func (x Filters_Projection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Filters_Projection.Descriptor instead.
func (Filters_Projection) EnumDescriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{4, 0}
}

type ScanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *TimeRange `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Filters   *Filters   `protobuf:"bytes,2,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *ScanRequest) Reset() {
	*x = ScanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRequest) ProtoMessage() {}

func (x *ScanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRequest.ProtoReflect.Descriptor instead.
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{0}
}

func (x *ScanRequest) GetTimestamp() *TimeRange {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ScanRequest) GetFilters() *Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ScanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record []byte `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *ScanResponse) Reset() {
	*x = ScanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanResponse) ProtoMessage() {}

func (x *ScanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanResponse.ProtoReflect.Descriptor instead.
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{1}
}

func (x *ScanResponse) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{2}
}

func (x *TimeRange) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TimeRange) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{3}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Filter `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// columns returned
	Projection []Filters_Projection `protobuf:"varint,2,rep,packed,name=projection,proto3,enum=v1.Filters_Projection" json:"projection,omitempty"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{4}
}

func (x *Filters) GetList() []*Filter {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Filters) GetProjection() []Filters_Projection {
	if x != nil {
		return x.Projection
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Metadata_Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Min     uint64             `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max     uint64             `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{5}
}

func (x *Metadata) GetColumns() []*Metadata_Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Metadata) GetMin() uint64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Metadata) GetMax() uint64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type Metadata_Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumRows       uint32   `protobuf:"varint,2,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	BitmapsOffset []uint32 `protobuf:"varint,3,rep,packed,name=bitmaps_offset,json=bitmapsOffset,proto3" json:"bitmaps_offset,omitempty"`
	FstOffset     uint32   `protobuf:"varint,4,opt,name=fst_offset,json=fstOffset,proto3" json:"fst_offset,omitempty"`
	RawSize       uint32   `protobuf:"varint,5,opt,name=raw_size,json=rawSize,proto3" json:"raw_size,omitempty"`
	Offset        uint64   `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Metadata_Column) Reset() {
	*x = Metadata_Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staples_v1_scan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata_Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata_Column) ProtoMessage() {}

func (x *Metadata_Column) ProtoReflect() protoreflect.Message {
	mi := &file_staples_v1_scan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata_Column.ProtoReflect.Descriptor instead.
func (*Metadata_Column) Descriptor() ([]byte, []int) {
	return file_staples_v1_scan_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Metadata_Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata_Column) GetNumRows() uint32 {
	if x != nil {
		return x.NumRows
	}
	return 0
}

func (x *Metadata_Column) GetBitmapsOffset() []uint32 {
	if x != nil {
		return x.BitmapsOffset
	}
	return nil
}

func (x *Metadata_Column) GetFstOffset() uint32 {
	if x != nil {
		return x.FstOffset
	}
	return 0
}

func (x *Metadata_Column) GetRawSize() uint32 {
	if x != nil {
		return x.RawSize
	}
	return 0
}

func (x *Metadata_Column) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_staples_v1_scan_proto protoreflect.FileDescriptor

var file_staples_v1_scan_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x61, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x71, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x33, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x7b, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xcb, 0x03,
	0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xe7, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x42,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x78, 0x69, 0x74, 0x50, 0x61, 0x67, 0x65, 0x10, 0x0b, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x10, 0x0d, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x73, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x0f, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x61, 0x74, 0x68, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x72, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x12, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x10, 0x14, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x74, 0x6d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x10, 0x15,
	0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x74, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x10, 0x16,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x74, 0x6d, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x17, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x74, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x18, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x74, 0x6d, 0x54, 0x65, 0x72, 0x6d, 0x10, 0x19, 0x22, 0x90, 0x02, 0x0a, 0x08,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x1a, 0xb0, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75,
	0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x73,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x62,
	0x69, 0x74, 0x6d, 0x61, 0x70, 0x73, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x66, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x61, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x6c,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x74, 0x73, 0x75, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56,
	0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staples_v1_scan_proto_rawDescOnce sync.Once
	file_staples_v1_scan_proto_rawDescData = file_staples_v1_scan_proto_rawDesc
)

func file_staples_v1_scan_proto_rawDescGZIP() []byte {
	file_staples_v1_scan_proto_rawDescOnce.Do(func() {
		file_staples_v1_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_staples_v1_scan_proto_rawDescData)
	})
	return file_staples_v1_scan_proto_rawDescData
}

var file_staples_v1_scan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_staples_v1_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_staples_v1_scan_proto_goTypes = []interface{}{
	(Filters_Projection)(0),       // 0: v1.Filters.Projection
	(*ScanRequest)(nil),           // 1: v1.ScanRequest
	(*ScanResponse)(nil),          // 2: v1.ScanResponse
	(*TimeRange)(nil),             // 3: v1.TimeRange
	(*KeyValue)(nil),              // 4: v1.KeyValue
	(*Filters)(nil),               // 5: v1.Filters
	(*Metadata)(nil),              // 6: v1.Metadata
	(*Metadata_Column)(nil),       // 7: v1.Metadata.Column
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*Filter)(nil),                // 9: v1.Filter
}
var file_staples_v1_scan_proto_depIdxs = []int32{
	3, // 0: v1.ScanRequest.timestamp:type_name -> v1.TimeRange
	5, // 1: v1.ScanRequest.filters:type_name -> v1.Filters
	8, // 2: v1.TimeRange.start:type_name -> google.protobuf.Timestamp
	8, // 3: v1.TimeRange.end:type_name -> google.protobuf.Timestamp
	9, // 4: v1.Filters.list:type_name -> v1.Filter
	0, // 5: v1.Filters.projection:type_name -> v1.Filters.Projection
	7, // 6: v1.Metadata.columns:type_name -> v1.Metadata.Column
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_staples_v1_scan_proto_init() }
func file_staples_v1_scan_proto_init() {
	if File_staples_v1_scan_proto != nil {
		return
	}
	file_staples_v1_stats_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staples_v1_scan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRequest); i {
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
		file_staples_v1_scan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanResponse); i {
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
		file_staples_v1_scan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
		file_staples_v1_scan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_staples_v1_scan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_staples_v1_scan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_staples_v1_scan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata_Column); i {
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
			RawDescriptor: file_staples_v1_scan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staples_v1_scan_proto_goTypes,
		DependencyIndexes: file_staples_v1_scan_proto_depIdxs,
		EnumInfos:         file_staples_v1_scan_proto_enumTypes,
		MessageInfos:      file_staples_v1_scan_proto_msgTypes,
	}.Build()
	File_staples_v1_scan_proto = out.File
	file_staples_v1_scan_proto_rawDesc = nil
	file_staples_v1_scan_proto_goTypes = nil
	file_staples_v1_scan_proto_depIdxs = nil
}
