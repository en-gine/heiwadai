// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/admin/Coupon.proto

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

type CouponIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CouponIDRequest) Reset() {
	*x = CouponIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponIDRequest) ProtoMessage() {}

func (x *CouponIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponIDRequest.ProtoReflect.Descriptor instead.
func (*CouponIDRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{0}
}

func (x *CouponIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Pager *shared.Pager `protobuf:"bytes,2,opt,name=Pager,proto3" json:"Pager,omitempty"`
}

func (x *UserCouponRequest) Reset() {
	*x = UserCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCouponRequest) ProtoMessage() {}

func (x *UserCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCouponRequest.ProtoReflect.Descriptor instead.
func (*UserCouponRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{1}
}

func (x *UserCouponRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserCouponRequest) GetPager() *shared.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

type CreateCustomCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	DiscountAmount    uint32                 `protobuf:"varint,2,opt,name=DiscountAmount,proto3" json:"DiscountAmount,omitempty"`
	ExpireAt          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ExpireAt,proto3" json:"ExpireAt,omitempty"`
	IsCombinationable bool                   `protobuf:"varint,4,opt,name=IsCombinationable,proto3" json:"IsCombinationable,omitempty"`
	Notices           []string               `protobuf:"bytes,5,rep,name=Notices,proto3" json:"Notices,omitempty"`
	TargetStoresID    []string               `protobuf:"bytes,6,rep,name=TargetStoresID,proto3" json:"TargetStoresID,omitempty"`
}

func (x *CreateCustomCouponRequest) Reset() {
	*x = CreateCustomCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomCouponRequest) ProtoMessage() {}

func (x *CreateCustomCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomCouponRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomCouponRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomCouponRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCustomCouponRequest) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *CreateCustomCouponRequest) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *CreateCustomCouponRequest) GetIsCombinationable() bool {
	if x != nil {
		return x.IsCombinationable
	}
	return false
}

func (x *CreateCustomCouponRequest) GetNotices() []string {
	if x != nil {
		return x.Notices
	}
	return nil
}

func (x *CreateCustomCouponRequest) GetTargetStoresID() []string {
	if x != nil {
		return x.TargetStoresID
	}
	return nil
}

type AffectedCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AffectedUserCount uint32 `protobuf:"varint,1,opt,name=AffectedUserCount,proto3" json:"AffectedUserCount,omitempty"`
}

func (x *AffectedCountResponse) Reset() {
	*x = AffectedCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedCountResponse) ProtoMessage() {}

func (x *AffectedCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedCountResponse.ProtoReflect.Descriptor instead.
func (*AffectedCountResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{3}
}

func (x *AffectedCountResponse) GetAffectedUserCount() uint32 {
	if x != nil {
		return x.AffectedUserCount
	}
	return 0
}

type UserAttachedCouponsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAttachedCoupons []*shared.UserAttachedCoupon `protobuf:"bytes,1,rep,name=UserAttachedCoupons,proto3" json:"UserAttachedCoupons,omitempty"`
	PageResponse        *shared.PageResponse         `protobuf:"bytes,2,opt,name=PageResponse,proto3" json:"PageResponse,omitempty"`
}

func (x *UserAttachedCouponsResponse) Reset() {
	*x = UserAttachedCouponsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAttachedCouponsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttachedCouponsResponse) ProtoMessage() {}

func (x *UserAttachedCouponsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttachedCouponsResponse.ProtoReflect.Descriptor instead.
func (*UserAttachedCouponsResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{4}
}

func (x *UserAttachedCouponsResponse) GetUserAttachedCoupons() []*shared.UserAttachedCoupon {
	if x != nil {
		return x.UserAttachedCoupons
	}
	return nil
}

func (x *UserAttachedCouponsResponse) GetPageResponse() *shared.PageResponse {
	if x != nil {
		return x.PageResponse
	}
	return nil
}

type DefaultCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupon *shared.CustomCoupon `protobuf:"bytes,1,opt,name=Coupon,proto3" json:"Coupon,omitempty"`
}

func (x *DefaultCouponResponse) Reset() {
	*x = DefaultCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultCouponResponse) ProtoMessage() {}

func (x *DefaultCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultCouponResponse.ProtoReflect.Descriptor instead.
func (*DefaultCouponResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{5}
}

func (x *DefaultCouponResponse) GetCoupon() *shared.CustomCoupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

type CouponListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupons      []*shared.CustomCoupon `protobuf:"bytes,1,rep,name=Coupons,proto3" json:"Coupons,omitempty"`
	PageResponse *shared.PageResponse   `protobuf:"bytes,2,opt,name=PageResponse,proto3" json:"PageResponse,omitempty"`
}

func (x *CouponListResponse) Reset() {
	*x = CouponListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponListResponse) ProtoMessage() {}

func (x *CouponListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponListResponse.ProtoReflect.Descriptor instead.
func (*CouponListResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{6}
}

func (x *CouponListResponse) GetCoupons() []*shared.CustomCoupon {
	if x != nil {
		return x.Coupons
	}
	return nil
}

func (x *CouponListResponse) GetPageResponse() *shared.PageResponse {
	if x != nil {
		return x.PageResponse
	}
	return nil
}

type SaveCustomCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	DiscountAmount    uint32                 `protobuf:"varint,3,opt,name=DiscountAmount,proto3" json:"DiscountAmount,omitempty"`
	ExpireAt          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ExpireAt,proto3" json:"ExpireAt,omitempty"`
	IsCombinationable bool                   `protobuf:"varint,5,opt,name=IsCombinationable,proto3" json:"IsCombinationable,omitempty"`
	Notices           []string               `protobuf:"bytes,6,rep,name=Notices,proto3" json:"Notices,omitempty"`
	TargetStoresID    []string               `protobuf:"bytes,7,rep,name=TargetStoresID,proto3" json:"TargetStoresID,omitempty"`
}

func (x *SaveCustomCouponRequest) Reset() {
	*x = SaveCustomCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_admin_Coupon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCustomCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCustomCouponRequest) ProtoMessage() {}

func (x *SaveCustomCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_Coupon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCustomCouponRequest.ProtoReflect.Descriptor instead.
func (*SaveCustomCouponRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_Coupon_proto_rawDescGZIP(), []int{7}
}

func (x *SaveCustomCouponRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SaveCustomCouponRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveCustomCouponRequest) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *SaveCustomCouponRequest) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *SaveCustomCouponRequest) GetIsCombinationable() bool {
	if x != nil {
		return x.IsCombinationable
	}
	return false
}

func (x *SaveCustomCouponRequest) GetNotices() []string {
	if x != nil {
		return x.Notices
	}
	return nil
}

func (x *SaveCustomCouponRequest) GetTargetStoresID() []string {
	if x != nil {
		return x.TargetStoresID
	}
	return nil
}

var File_v1_admin_Coupon_proto protoreflect.FileDescriptor

var file_v1_admin_Coupon_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x50, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x50, 0x61, 0x67, 0x65, 0x72, 0x22, 0xff, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x49, 0x73, 0x43, 0x6f,
	0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x49, 0x44, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x49, 0x44, 0x22,
	0x45, 0x0a, 0x15, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb3, 0x01, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x13, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x15,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x02, 0x0a, 0x17, 0x53, 0x61,
	0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x73, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x49, 0x44, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x49, 0x44, 0x32, 0x90, 0x05, 0x0a, 0x15, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x1b, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x6f, 0x41,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x85, 0x01, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x42, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x18, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_admin_Coupon_proto_rawDescOnce sync.Once
	file_v1_admin_Coupon_proto_rawDescData = file_v1_admin_Coupon_proto_rawDesc
)

func file_v1_admin_Coupon_proto_rawDescGZIP() []byte {
	file_v1_admin_Coupon_proto_rawDescOnce.Do(func() {
		file_v1_admin_Coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_admin_Coupon_proto_rawDescData)
	})
	return file_v1_admin_Coupon_proto_rawDescData
}

var file_v1_admin_Coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_admin_Coupon_proto_goTypes = []interface{}{
	(*CouponIDRequest)(nil),             // 0: server.admin.CouponIDRequest
	(*UserCouponRequest)(nil),           // 1: server.admin.UserCouponRequest
	(*CreateCustomCouponRequest)(nil),   // 2: server.admin.CreateCustomCouponRequest
	(*AffectedCountResponse)(nil),       // 3: server.admin.AffectedCountResponse
	(*UserAttachedCouponsResponse)(nil), // 4: server.admin.UserAttachedCouponsResponse
	(*DefaultCouponResponse)(nil),       // 5: server.admin.DefaultCouponResponse
	(*CouponListResponse)(nil),          // 6: server.admin.CouponListResponse
	(*SaveCustomCouponRequest)(nil),     // 7: server.admin.SaveCustomCouponRequest
	(*shared.Pager)(nil),                // 8: server.shared.Pager
	(*timestamppb.Timestamp)(nil),       // 9: google.protobuf.Timestamp
	(*shared.UserAttachedCoupon)(nil),   // 10: server.shared.UserAttachedCoupon
	(*shared.PageResponse)(nil),         // 11: server.shared.PageResponse
	(*shared.CustomCoupon)(nil),         // 12: server.shared.CustomCoupon
	(*emptypb.Empty)(nil),               // 13: google.protobuf.Empty
}
var file_v1_admin_Coupon_proto_depIdxs = []int32{
	8,  // 0: server.admin.UserCouponRequest.Pager:type_name -> server.shared.Pager
	9,  // 1: server.admin.CreateCustomCouponRequest.ExpireAt:type_name -> google.protobuf.Timestamp
	10, // 2: server.admin.UserAttachedCouponsResponse.UserAttachedCoupons:type_name -> server.shared.UserAttachedCoupon
	11, // 3: server.admin.UserAttachedCouponsResponse.PageResponse:type_name -> server.shared.PageResponse
	12, // 4: server.admin.DefaultCouponResponse.Coupon:type_name -> server.shared.CustomCoupon
	12, // 5: server.admin.CouponListResponse.Coupons:type_name -> server.shared.CustomCoupon
	11, // 6: server.admin.CouponListResponse.PageResponse:type_name -> server.shared.PageResponse
	9,  // 7: server.admin.SaveCustomCouponRequest.ExpireAt:type_name -> google.protobuf.Timestamp
	1,  // 8: server.admin.AdminCouponController.GetUserCouponList:input_type -> server.admin.UserCouponRequest
	2,  // 9: server.admin.AdminCouponController.CreateCustomCoupon:input_type -> server.admin.CreateCustomCouponRequest
	7,  // 10: server.admin.AdminCouponController.SaveCustomCoupon:input_type -> server.admin.SaveCustomCouponRequest
	13, // 11: server.admin.AdminCouponController.GetDefaultEmptyCoupon:input_type -> google.protobuf.Empty
	0,  // 12: server.admin.AdminCouponController.GetCustomCouponByID:input_type -> server.admin.CouponIDRequest
	8,  // 13: server.admin.AdminCouponController.GetCustomCouponList:input_type -> server.shared.Pager
	0,  // 14: server.admin.AdminCouponController.AttachCustomCouponToAllUser:input_type -> server.admin.CouponIDRequest
	4,  // 15: server.admin.AdminCouponController.GetUserCouponList:output_type -> server.admin.UserAttachedCouponsResponse
	12, // 16: server.admin.AdminCouponController.CreateCustomCoupon:output_type -> server.shared.CustomCoupon
	13, // 17: server.admin.AdminCouponController.SaveCustomCoupon:output_type -> google.protobuf.Empty
	5,  // 18: server.admin.AdminCouponController.GetDefaultEmptyCoupon:output_type -> server.admin.DefaultCouponResponse
	12, // 19: server.admin.AdminCouponController.GetCustomCouponByID:output_type -> server.shared.CustomCoupon
	6,  // 20: server.admin.AdminCouponController.GetCustomCouponList:output_type -> server.admin.CouponListResponse
	3,  // 21: server.admin.AdminCouponController.AttachCustomCouponToAllUser:output_type -> server.admin.AffectedCountResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_v1_admin_Coupon_proto_init() }
func file_v1_admin_Coupon_proto_init() {
	if File_v1_admin_Coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_admin_Coupon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponIDRequest); i {
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
		file_v1_admin_Coupon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCouponRequest); i {
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
		file_v1_admin_Coupon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomCouponRequest); i {
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
		file_v1_admin_Coupon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedCountResponse); i {
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
		file_v1_admin_Coupon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAttachedCouponsResponse); i {
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
		file_v1_admin_Coupon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultCouponResponse); i {
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
		file_v1_admin_Coupon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponListResponse); i {
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
		file_v1_admin_Coupon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveCustomCouponRequest); i {
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
			RawDescriptor: file_v1_admin_Coupon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_admin_Coupon_proto_goTypes,
		DependencyIndexes: file_v1_admin_Coupon_proto_depIdxs,
		MessageInfos:      file_v1_admin_Coupon_proto_msgTypes,
	}.Build()
	File_v1_admin_Coupon_proto = out.File
	file_v1_admin_Coupon_proto_rawDesc = nil
	file_v1_admin_Coupon_proto_goTypes = nil
	file_v1_admin_Coupon_proto_depIdxs = nil
}
