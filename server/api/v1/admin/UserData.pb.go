// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/admin/UserData.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	shared "server/api/v1/shared"
	user "server/api/v1/user"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserUpdateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User            *user.UserUpdateDataRequest `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	InnerNote       string                      `protobuf:"bytes,2,opt,name=InnerNote,proto3" json:"InnerNote,omitempty"`
	IsBlackCustomer bool                        `protobuf:"varint,3,opt,name=IsBlackCustomer,proto3" json:"IsBlackCustomer,omitempty"`
}

func (x *UserUpdateDataRequest) Reset() {
	*x = UserUpdateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateDataRequest) ProtoMessage() {}

func (x *UserUpdateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateDataRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateDataRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{0}
}

func (x *UserUpdateDataRequest) GetUser() *user.UserUpdateDataRequest {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserUpdateDataRequest) GetInnerNote() string {
	if x != nil {
		return x.InnerNote
	}
	return ""
}

func (x *UserUpdateDataRequest) GetIsBlackCustomer() bool {
	if x != nil {
		return x.IsBlackCustomer
	}
	return false
}

type UserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User            *user.UserDataResponse `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	InnerNote       string                 `protobuf:"bytes,2,opt,name=InnerNote,proto3" json:"InnerNote,omitempty"`
	IsBlackCustomer bool                   `protobuf:"varint,3,opt,name=IsBlackCustomer,proto3" json:"IsBlackCustomer,omitempty"`
}

func (x *UserDataResponse) Reset() {
	*x = UserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDataResponse) ProtoMessage() {}

func (x *UserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDataResponse.ProtoReflect.Descriptor instead.
func (*UserDataResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{1}
}

func (x *UserDataResponse) GetUser() *user.UserDataResponse {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserDataResponse) GetInnerNote() string {
	if x != nil {
		return x.InnerNote
	}
	return ""
}

func (x *UserDataResponse) GetIsBlackCustomer() bool {
	if x != nil {
		return x.IsBlackCustomer
	}
	return false
}

type UserWithCheckIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User          *user.UserDataResponse `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	LastCheckinAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=LastCheckinAt,proto3,oneof" json:"LastCheckinAt,omitempty"`
}

func (x *UserWithCheckIn) Reset() {
	*x = UserWithCheckIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWithCheckIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithCheckIn) ProtoMessage() {}

func (x *UserWithCheckIn) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithCheckIn.ProtoReflect.Descriptor instead.
func (*UserWithCheckIn) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{2}
}

func (x *UserWithCheckIn) GetUser() *user.UserDataResponse {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserWithCheckIn) GetLastCheckinAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastCheckinAt
	}
	return nil
}

type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users        []*UserWithCheckIn   `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	PageResponse *shared.PageResponse `protobuf:"bytes,3,opt,name=PageResponse,proto3" json:"PageResponse,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{3}
}

func (x *UserListResponse) GetUsers() []*UserWithCheckIn {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UserListResponse) GetPageResponse() *shared.PageResponse {
	if x != nil {
		return x.PageResponse
	}
	return nil
}

type UserGetIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserGetIDRequest) Reset() {
	*x = UserGetIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetIDRequest) ProtoMessage() {}

func (x *UserGetIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetIDRequest.ProtoReflect.Descriptor instead.
func (*UserGetIDRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{4}
}

func (x *UserGetIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{5}
}

func (x *UserDeleteRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type SearchFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName      *string            `protobuf:"bytes,1,opt,name=LastName,proto3,oneof" json:"LastName,omitempty"`
	FirstName     *string            `protobuf:"bytes,2,opt,name=FirstName,proto3,oneof" json:"FirstName,omitempty"`
	LastNameKana  *string            `protobuf:"bytes,3,opt,name=LastNameKana,proto3,oneof" json:"LastNameKana,omitempty"`
	FirstNameKana *string            `protobuf:"bytes,4,opt,name=FirstNameKana,proto3,oneof" json:"FirstNameKana,omitempty"`
	Prefecture    *shared.Prefecture `protobuf:"varint,5,opt,name=Prefecture,proto3,enum=server.shared.Prefecture,oneof" json:"Prefecture,omitempty"`
}

func (x *SearchFilter) Reset() {
	*x = SearchFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFilter) ProtoMessage() {}

func (x *SearchFilter) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFilter.ProtoReflect.Descriptor instead.
func (*SearchFilter) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{6}
}

func (x *SearchFilter) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *SearchFilter) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *SearchFilter) GetLastNameKana() string {
	if x != nil && x.LastNameKana != nil {
		return *x.LastNameKana
	}
	return ""
}

func (x *SearchFilter) GetFirstNameKana() string {
	if x != nil && x.FirstNameKana != nil {
		return *x.FirstNameKana
	}
	return ""
}

func (x *SearchFilter) GetPrefecture() shared.Prefecture {
	if x != nil && x.Prefecture != nil {
		return *x.Prefecture
	}
	return shared.Prefecture(0)
}

type UserListFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search *SearchFilter `protobuf:"bytes,1,opt,name=Search,proto3" json:"Search,omitempty"`
	Pager  *shared.Pager `protobuf:"bytes,2,opt,name=Pager,proto3" json:"Pager,omitempty"`
}

func (x *UserListFilterRequest) Reset() {
	*x = UserListFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_UserData_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListFilterRequest) ProtoMessage() {}

func (x *UserListFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_UserData_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListFilterRequest.ProtoReflect.Descriptor instead.
func (*UserListFilterRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_UserData_proto_rawDescGZIP(), []int{7}
}

func (x *UserListFilterRequest) GetSearch() *SearchFilter {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *UserListFilterRequest) GetPager() *shared.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

var File_v1_admin_UserData_proto protoreflect.FileDescriptor

var file_v1_admin_UserData_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x50, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x01, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x49, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x49, 0x73, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x49, 0x73, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x49, 0x73, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x31,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x4c, 0x61, 0x73,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x41, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xb3,
	0x02, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4b,
	0x61, 0x6e, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4b, 0x61, 0x6e, 0x61, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4b, 0x61, 0x6e, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0d, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x4b, 0x61, 0x6e, 0x61, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x48, 0x04, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x66, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4b,
	0x61, 0x6e, 0x61, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x4b, 0x61, 0x6e, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x50, 0x72, 0x65, 0x66, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x77, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x2a, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x32, 0xc9, 0x02,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x87, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x18, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_admin_UserData_proto_rawDescOnce sync.Once
	file_v1_admin_UserData_proto_rawDescData = file_v1_admin_UserData_proto_rawDesc
)

func file_v1_admin_UserData_proto_rawDescGZIP() []byte {
	file_v1_admin_UserData_proto_rawDescOnce.Do(func() {
		file_v1_admin_UserData_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_admin_UserData_proto_rawDescData)
	})
	return file_v1_admin_UserData_proto_rawDescData
}

var file_v1_admin_UserData_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_admin_UserData_proto_goTypes = []interface{}{
	(*UserUpdateDataRequest)(nil),      // 0: server.admin.UserUpdateDataRequest
	(*UserDataResponse)(nil),           // 1: server.admin.UserDataResponse
	(*UserWithCheckIn)(nil),            // 2: server.admin.UserWithCheckIn
	(*UserListResponse)(nil),           // 3: server.admin.UserListResponse
	(*UserGetIDRequest)(nil),           // 4: server.admin.UserGetIDRequest
	(*UserDeleteRequest)(nil),          // 5: server.admin.UserDeleteRequest
	(*SearchFilter)(nil),               // 6: server.admin.SearchFilter
	(*UserListFilterRequest)(nil),      // 7: server.admin.UserListFilterRequest
	(*user.UserUpdateDataRequest)(nil), // 8: server.user.UserUpdateDataRequest
	(*user.UserDataResponse)(nil),      // 9: server.user.UserDataResponse
	(*timestamppb.Timestamp)(nil),      // 10: google.protobuf.Timestamp
	(*shared.PageResponse)(nil),        // 11: server.shared.PageResponse
	(shared.Prefecture)(0),             // 12: server.shared.Prefecture
	(*shared.Pager)(nil),               // 13: server.shared.Pager
	(*emptypb.Empty)(nil),              // 14: google.protobuf.Empty
}
var file_v1_admin_UserData_proto_depIdxs = []int32{
	8,  // 0: server.admin.UserUpdateDataRequest.User:type_name -> server.user.UserUpdateDataRequest
	9,  // 1: server.admin.UserDataResponse.User:type_name -> server.user.UserDataResponse
	9,  // 2: server.admin.UserWithCheckIn.User:type_name -> server.user.UserDataResponse
	10, // 3: server.admin.UserWithCheckIn.LastCheckinAt:type_name -> google.protobuf.Timestamp
	2,  // 4: server.admin.UserListResponse.Users:type_name -> server.admin.UserWithCheckIn
	11, // 5: server.admin.UserListResponse.PageResponse:type_name -> server.shared.PageResponse
	12, // 6: server.admin.SearchFilter.Prefecture:type_name -> server.shared.Prefecture
	6,  // 7: server.admin.UserListFilterRequest.Search:type_name -> server.admin.SearchFilter
	13, // 8: server.admin.UserListFilterRequest.Pager:type_name -> server.shared.Pager
	4,  // 9: server.admin.UserDataController.GetByID:input_type -> server.admin.UserGetIDRequest
	0,  // 10: server.admin.UserDataController.Update:input_type -> server.admin.UserUpdateDataRequest
	5,  // 11: server.admin.UserDataController.Delete:input_type -> server.admin.UserDeleteRequest
	7,  // 12: server.admin.UserDataController.GetList:input_type -> server.admin.UserListFilterRequest
	1,  // 13: server.admin.UserDataController.GetByID:output_type -> server.admin.UserDataResponse
	1,  // 14: server.admin.UserDataController.Update:output_type -> server.admin.UserDataResponse
	14, // 15: server.admin.UserDataController.Delete:output_type -> google.protobuf.Empty
	3,  // 16: server.admin.UserDataController.GetList:output_type -> server.admin.UserListResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_v1_admin_UserData_proto_init() }
func file_v1_admin_UserData_proto_init() {
	if File_v1_admin_UserData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_admin_UserData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateDataRequest); i {
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
		file_v1_admin_UserData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDataResponse); i {
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
		file_v1_admin_UserData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWithCheckIn); i {
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
		file_v1_admin_UserData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
		file_v1_admin_UserData_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetIDRequest); i {
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
		file_v1_admin_UserData_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRequest); i {
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
		file_v1_admin_UserData_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFilter); i {
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
		file_v1_admin_UserData_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListFilterRequest); i {
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
	file_v1_admin_UserData_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_v1_admin_UserData_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_admin_UserData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_admin_UserData_proto_goTypes,
		DependencyIndexes: file_v1_admin_UserData_proto_depIdxs,
		MessageInfos:      file_v1_admin_UserData_proto_msgTypes,
	}.Build()
	File_v1_admin_UserData_proto = out.File
	file_v1_admin_UserData_proto_rawDesc = nil
	file_v1_admin_UserData_proto_goTypes = nil
	file_v1_admin_UserData_proto_depIdxs = nil
}
