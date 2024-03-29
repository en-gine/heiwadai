// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/admin/Messages.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	shared "server/api/v1/shared"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *MessageIDRequest) Reset() {
	*x = MessageIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIDRequest) ProtoMessage() {}

func (x *MessageIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIDRequest.ProtoReflect.Descriptor instead.
func (*MessageIDRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{0}
}

func (x *MessageIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type MessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages     []*MessageResponse   `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	PageResponse *shared.PageResponse `protobuf:"bytes,2,opt,name=PageResponse,proto3" json:"PageResponse,omitempty"`
}

func (x *MessagesResponse) Reset() {
	*x = MessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesResponse) ProtoMessage() {}

func (x *MessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesResponse.ProtoReflect.Descriptor instead.
func (*MessagesResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{1}
}

func (x *MessagesResponse) GetMessages() []*MessageResponse {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *MessagesResponse) GetPageResponse() *shared.PageResponse {
	if x != nil {
		return x.PageResponse
	}
	return nil
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pager *shared.Pager `protobuf:"bytes,1,opt,name=Pager,proto3" json:"Pager,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequest) GetPager() *shared.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

type MessageCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                 `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content     string                 `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	DisplayDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=DisplayDate,proto3" json:"DisplayDate,omitempty"`
}

func (x *MessageCreateRequest) Reset() {
	*x = MessageCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageCreateRequest) ProtoMessage() {}

func (x *MessageCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageCreateRequest.ProtoReflect.Descriptor instead.
func (*MessageCreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{3}
}

func (x *MessageCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageCreateRequest) GetDisplayDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DisplayDate
	}
	return nil
}

type MessageUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       *string                `protobuf:"bytes,2,opt,name=Title,proto3,oneof" json:"Title,omitempty"`
	Content     *string                `protobuf:"bytes,3,opt,name=Content,proto3,oneof" json:"Content,omitempty"`
	DisplayDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=DisplayDate,proto3,oneof" json:"DisplayDate,omitempty"`
}

func (x *MessageUpdateRequest) Reset() {
	*x = MessageUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageUpdateRequest) ProtoMessage() {}

func (x *MessageUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageUpdateRequest.ProtoReflect.Descriptor instead.
func (*MessageUpdateRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{4}
}

func (x *MessageUpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MessageUpdateRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *MessageUpdateRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *MessageUpdateRequest) GetDisplayDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DisplayDate
	}
	return nil
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content     string                 `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	AuthorID    string                 `protobuf:"bytes,4,opt,name=AuthorID,proto3" json:"AuthorID,omitempty"`
	DisplayDate *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=DisplayDate,proto3" json:"DisplayDate,omitempty"`
	CreateAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Messages_proto_rawDescGZIP(), []int{5}
}

func (x *MessageResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *MessageResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageResponse) GetAuthorID() string {
	if x != nil {
		return x.AuthorID
	}
	return ""
}

func (x *MessageResponse) GetDisplayDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DisplayDate
	}
	return nil
}

func (x *MessageResponse) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

var File_v1_admin_Messages_proto protoreflect.FileDescriptor

var file_v1_admin_Messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x50, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x3f, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x14, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0b,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52,
	0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32, 0x8f, 0x03, 0x0a, 0x11,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x87, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x18, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_admin_Messages_proto_rawDescOnce sync.Once
	file_v1_admin_Messages_proto_rawDescData = file_v1_admin_Messages_proto_rawDesc
)

func file_v1_admin_Messages_proto_rawDescGZIP() []byte {
	file_v1_admin_Messages_proto_rawDescOnce.Do(func() {
		file_v1_admin_Messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_admin_Messages_proto_rawDescData)
	})
	return file_v1_admin_Messages_proto_rawDescData
}

var file_v1_admin_Messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_admin_Messages_proto_goTypes = []interface{}{
	(*MessageIDRequest)(nil),      // 0: server.admin.MessageIDRequest
	(*MessagesResponse)(nil),      // 1: server.admin.MessagesResponse
	(*GetMessageRequest)(nil),     // 2: server.admin.GetMessageRequest
	(*MessageCreateRequest)(nil),  // 3: server.admin.MessageCreateRequest
	(*MessageUpdateRequest)(nil),  // 4: server.admin.MessageUpdateRequest
	(*MessageResponse)(nil),       // 5: server.admin.MessageResponse
	(*shared.PageResponse)(nil),   // 6: server.shared.PageResponse
	(*shared.Pager)(nil),          // 7: server.shared.Pager
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_v1_admin_Messages_proto_depIdxs = []int32{
	5,  // 0: server.admin.MessagesResponse.messages:type_name -> server.admin.MessageResponse
	6,  // 1: server.admin.MessagesResponse.PageResponse:type_name -> server.shared.PageResponse
	7,  // 2: server.admin.GetMessageRequest.Pager:type_name -> server.shared.Pager
	8,  // 3: server.admin.MessageCreateRequest.DisplayDate:type_name -> google.protobuf.Timestamp
	8,  // 4: server.admin.MessageUpdateRequest.DisplayDate:type_name -> google.protobuf.Timestamp
	8,  // 5: server.admin.MessageResponse.DisplayDate:type_name -> google.protobuf.Timestamp
	8,  // 6: server.admin.MessageResponse.CreateAt:type_name -> google.protobuf.Timestamp
	0,  // 7: server.admin.MessageController.GetByID:input_type -> server.admin.MessageIDRequest
	2,  // 8: server.admin.MessageController.GetList:input_type -> server.admin.GetMessageRequest
	3,  // 9: server.admin.MessageController.Create:input_type -> server.admin.MessageCreateRequest
	4,  // 10: server.admin.MessageController.Update:input_type -> server.admin.MessageUpdateRequest
	0,  // 11: server.admin.MessageController.Delete:input_type -> server.admin.MessageIDRequest
	5,  // 12: server.admin.MessageController.GetByID:output_type -> server.admin.MessageResponse
	1,  // 13: server.admin.MessageController.GetList:output_type -> server.admin.MessagesResponse
	5,  // 14: server.admin.MessageController.Create:output_type -> server.admin.MessageResponse
	5,  // 15: server.admin.MessageController.Update:output_type -> server.admin.MessageResponse
	9,  // 16: server.admin.MessageController.Delete:output_type -> google.protobuf.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_admin_Messages_proto_init() }
func file_v1_admin_Messages_proto_init() {
	if File_v1_admin_Messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_admin_Messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageIDRequest); i {
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
		file_v1_admin_Messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagesResponse); i {
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
		file_v1_admin_Messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_v1_admin_Messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageCreateRequest); i {
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
		file_v1_admin_Messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageUpdateRequest); i {
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
		file_v1_admin_Messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
	file_v1_admin_Messages_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_admin_Messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_admin_Messages_proto_goTypes,
		DependencyIndexes: file_v1_admin_Messages_proto_depIdxs,
		MessageInfos:      file_v1_admin_Messages_proto_msgTypes,
	}.Build()
	File_v1_admin_Messages_proto = out.File
	file_v1_admin_Messages_proto_rawDesc = nil
	file_v1_admin_Messages_proto_goTypes = nil
	file_v1_admin_Messages_proto_depIdxs = nil
}
