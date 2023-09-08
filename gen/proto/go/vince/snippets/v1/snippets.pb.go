// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vince/snippets/v1/snippets.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/vinceanalytics/vince/gen/proto/go/vince/query/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Snippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Query     string                 `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	Params    []*v1.QueryParam       `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Snippet) Reset() {
	*x = Snippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snippet) ProtoMessage() {}

func (x *Snippet) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snippet.ProtoReflect.Descriptor instead.
func (*Snippet) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{0}
}

func (x *Snippet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Snippet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Snippet) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Snippet) GetParams() []*v1.QueryParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Snippet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Snippet) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Query  string           `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Params []*v1.QueryParam `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *CreateSnippetRequest) Reset() {
	*x = CreateSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetRequest) ProtoMessage() {}

func (x *CreateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetRequest.ProtoReflect.Descriptor instead.
func (*CreateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSnippetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSnippetRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CreateSnippetRequest) GetParams() []*v1.QueryParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type UpdateSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Query  string           `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	Params []*v1.QueryParam `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *UpdateSnippetRequest) Reset() {
	*x = UpdateSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnippetRequest) ProtoMessage() {}

func (x *UpdateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnippetRequest.ProtoReflect.Descriptor instead.
func (*UpdateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSnippetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSnippetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSnippetRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *UpdateSnippetRequest) GetParams() []*v1.QueryParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type DeleteSnippetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSnippetRequest) Reset() {
	*x = DeleteSnippetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnippetRequest) ProtoMessage() {}

func (x *DeleteSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnippetRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnippetRequest) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSnippetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListSnippetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSnippetsRequest) Reset() {
	*x = ListSnippetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnippetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnippetsRequest) ProtoMessage() {}

func (x *ListSnippetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnippetsRequest.ProtoReflect.Descriptor instead.
func (*ListSnippetsRequest) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{4}
}

type ListSnippetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snippets []*Snippet `protobuf:"bytes,1,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *ListSnippetsResponse) Reset() {
	*x = ListSnippetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_snippets_v1_snippets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSnippetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSnippetsResponse) ProtoMessage() {}

func (x *ListSnippetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_snippets_v1_snippets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSnippetsResponse.ProtoReflect.Descriptor instead.
func (*ListSnippetsResponse) Descriptor() ([]byte, []int) {
	return file_vince_snippets_v1_snippets_proto_rawDescGZIP(), []int{5}
}

func (x *ListSnippetsResponse) GetSnippets() []*Snippet {
	if x != nil {
		return x.Snippets
	}
	return nil
}

var File_vince_snippets_v1_snippets_proto protoreflect.FileDescriptor

var file_vince_snippets_v1_snippets_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1a, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x07,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x78, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x73,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x73, 0x32, 0xd8, 0x02, 0x0a, 0x08, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12,
	0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x1a,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12, 0x57, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12, 0x17, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x2a, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x42,
	0x7f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58,
	0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vince_snippets_v1_snippets_proto_rawDescOnce sync.Once
	file_vince_snippets_v1_snippets_proto_rawDescData = file_vince_snippets_v1_snippets_proto_rawDesc
)

func file_vince_snippets_v1_snippets_proto_rawDescGZIP() []byte {
	file_vince_snippets_v1_snippets_proto_rawDescOnce.Do(func() {
		file_vince_snippets_v1_snippets_proto_rawDescData = protoimpl.X.CompressGZIP(file_vince_snippets_v1_snippets_proto_rawDescData)
	})
	return file_vince_snippets_v1_snippets_proto_rawDescData
}

var file_vince_snippets_v1_snippets_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vince_snippets_v1_snippets_proto_goTypes = []interface{}{
	(*Snippet)(nil),               // 0: v1.Snippet
	(*CreateSnippetRequest)(nil),  // 1: v1.CreateSnippetRequest
	(*UpdateSnippetRequest)(nil),  // 2: v1.UpdateSnippetRequest
	(*DeleteSnippetRequest)(nil),  // 3: v1.DeleteSnippetRequest
	(*ListSnippetsRequest)(nil),   // 4: v1.ListSnippetsRequest
	(*ListSnippetsResponse)(nil),  // 5: v1.ListSnippetsResponse
	(*v1.QueryParam)(nil),         // 6: v1.QueryParam
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_vince_snippets_v1_snippets_proto_depIdxs = []int32{
	6,  // 0: v1.Snippet.params:type_name -> v1.QueryParam
	7,  // 1: v1.Snippet.created_at:type_name -> google.protobuf.Timestamp
	7,  // 2: v1.Snippet.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 3: v1.CreateSnippetRequest.params:type_name -> v1.QueryParam
	6,  // 4: v1.UpdateSnippetRequest.params:type_name -> v1.QueryParam
	0,  // 5: v1.ListSnippetsResponse.snippets:type_name -> v1.Snippet
	1,  // 6: v1.Snippets.CreateSnippet:input_type -> v1.CreateSnippetRequest
	2,  // 7: v1.Snippets.UpdateSnippet:input_type -> v1.UpdateSnippetRequest
	4,  // 8: v1.Snippets.ListSnippets:input_type -> v1.ListSnippetsRequest
	3,  // 9: v1.Snippets.DeleteSnippet:input_type -> v1.DeleteSnippetRequest
	0,  // 10: v1.Snippets.CreateSnippet:output_type -> v1.Snippet
	0,  // 11: v1.Snippets.UpdateSnippet:output_type -> v1.Snippet
	5,  // 12: v1.Snippets.ListSnippets:output_type -> v1.ListSnippetsResponse
	8,  // 13: v1.Snippets.DeleteSnippet:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_vince_snippets_v1_snippets_proto_init() }
func file_vince_snippets_v1_snippets_proto_init() {
	if File_vince_snippets_v1_snippets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vince_snippets_v1_snippets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snippet); i {
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
		file_vince_snippets_v1_snippets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSnippetRequest); i {
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
		file_vince_snippets_v1_snippets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSnippetRequest); i {
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
		file_vince_snippets_v1_snippets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnippetRequest); i {
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
		file_vince_snippets_v1_snippets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnippetsRequest); i {
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
		file_vince_snippets_v1_snippets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSnippetsResponse); i {
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
			RawDescriptor: file_vince_snippets_v1_snippets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vince_snippets_v1_snippets_proto_goTypes,
		DependencyIndexes: file_vince_snippets_v1_snippets_proto_depIdxs,
		MessageInfos:      file_vince_snippets_v1_snippets_proto_msgTypes,
	}.Build()
	File_vince_snippets_v1_snippets_proto = out.File
	file_vince_snippets_v1_snippets_proto_rawDesc = nil
	file_vince_snippets_v1_snippets_proto_goTypes = nil
	file_vince_snippets_v1_snippets_proto_depIdxs = nil
}
