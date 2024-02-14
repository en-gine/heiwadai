// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/shared/Coupon.proto

package shared

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type CouponType int32

const (
	CouponType_COUPON_STANDARD CouponType = 0
	CouponType_COUPON_CUSTOM   CouponType = 1
	CouponType_COUPON_BIRTHDAY CouponType = 2
)

// Enum value maps for CouponType.
var (
	CouponType_name = map[int32]string{
		0: "COUPON_STANDARD",
		1: "COUPON_CUSTOM",
		2: "COUPON_BIRTHDAY",
	}
	CouponType_value = map[string]int32{
		"COUPON_STANDARD": 0,
		"COUPON_CUSTOM":   1,
		"COUPON_BIRTHDAY": 2,
	}
)

func (x CouponType) Enum() *CouponType {
	p := new(CouponType)
	*p = x
	return p
}

func (x CouponType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_Coupon_proto_enumTypes[0].Descriptor()
}

func (CouponType) Type() protoreflect.EnumType {
	return &file_v1_shared_Coupon_proto_enumTypes[0]
}

func (x CouponType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponType.Descriptor instead.
func (CouponType) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_Coupon_proto_rawDescGZIP(), []int{0}
}

type CouponStatus int32

const (
	CouponStatus_COUPON_DRAFT   CouponStatus = 0
	CouponStatus_COUPON_CREATED CouponStatus = 1
	CouponStatus_COUPON_ISSUED  CouponStatus = 2
)

// Enum value maps for CouponStatus.
var (
	CouponStatus_name = map[int32]string{
		0: "COUPON_DRAFT",
		1: "COUPON_CREATED",
		2: "COUPON_ISSUED",
	}
	CouponStatus_value = map[string]int32{
		"COUPON_DRAFT":   0,
		"COUPON_CREATED": 1,
		"COUPON_ISSUED":  2,
	}
)

func (x CouponStatus) Enum() *CouponStatus {
	p := new(CouponStatus)
	*p = x
	return p
}

func (x CouponStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_Coupon_proto_enumTypes[1].Descriptor()
}

func (CouponStatus) Type() protoreflect.EnumType {
	return &file_v1_shared_Coupon_proto_enumTypes[1]
}

func (x CouponStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponStatus.Descriptor instead.
func (CouponStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_Coupon_proto_rawDescGZIP(), []int{1}
}

type Coupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CouponType        CouponType             `protobuf:"varint,3,opt,name=CouponType,proto3,enum=server.shared.CouponType" json:"CouponType,omitempty"`
	DiscountAmount    uint32                 `protobuf:"varint,4,opt,name=DiscountAmount,proto3" json:"DiscountAmount,omitempty"`
	ExpireAt          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ExpireAt,proto3" json:"ExpireAt,omitempty"` // time.Time
	IsCombinationable bool                   `protobuf:"varint,6,opt,name=IsCombinationable,proto3" json:"IsCombinationable,omitempty"`
	Notices           []string               `protobuf:"bytes,7,rep,name=Notices,proto3" json:"Notices,omitempty"`
	TargetStore       []*Store               `protobuf:"bytes,8,rep,name=TargetStore,proto3" json:"TargetStore,omitempty"`
	CreateAt          *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"` // time.Time
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_Coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_Coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_v1_shared_Coupon_proto_rawDescGZIP(), []int{0}
}

func (x *Coupon) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Coupon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coupon) GetCouponType() CouponType {
	if x != nil {
		return x.CouponType
	}
	return CouponType_COUPON_STANDARD
}

func (x *Coupon) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *Coupon) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *Coupon) GetIsCombinationable() bool {
	if x != nil {
		return x.IsCombinationable
	}
	return false
}

func (x *Coupon) GetNotices() []string {
	if x != nil {
		return x.Notices
	}
	return nil
}

func (x *Coupon) GetTargetStore() []*Store {
	if x != nil {
		return x.TargetStore
	}
	return nil
}

func (x *Coupon) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

type CustomCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CouponType        CouponType             `protobuf:"varint,3,opt,name=CouponType,proto3,enum=server.shared.CouponType" json:"CouponType,omitempty"`
	DiscountAmount    uint32                 `protobuf:"varint,4,opt,name=DiscountAmount,proto3" json:"DiscountAmount,omitempty"`
	ExpireAt          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ExpireAt,proto3" json:"ExpireAt,omitempty"` // time.Time
	IsCombinationable bool                   `protobuf:"varint,6,opt,name=IsCombinationable,proto3" json:"IsCombinationable,omitempty"`
	Notices           []string               `protobuf:"bytes,7,rep,name=Notices,proto3" json:"Notices,omitempty"`
	TargetStore       []*Store               `protobuf:"bytes,8,rep,name=TargetStore,proto3" json:"TargetStore,omitempty"`
	Status            CouponStatus           `protobuf:"varint,9,opt,name=Status,proto3,enum=server.shared.CouponStatus" json:"Status,omitempty"`
	CreateAt          *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"` // time.Time
	IssueCount        *uint32                `protobuf:"varint,11,opt,name=IssueCount,proto3,oneof" json:"IssueCount,omitempty"`
	IssueAt           *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=IssueAt,proto3,oneof" json:"IssueAt,omitempty"` // time.Time
}

func (x *CustomCoupon) Reset() {
	*x = CustomCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_Coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomCoupon) ProtoMessage() {}

func (x *CustomCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_Coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomCoupon.ProtoReflect.Descriptor instead.
func (*CustomCoupon) Descriptor() ([]byte, []int) {
	return file_v1_shared_Coupon_proto_rawDescGZIP(), []int{1}
}

func (x *CustomCoupon) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CustomCoupon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomCoupon) GetCouponType() CouponType {
	if x != nil {
		return x.CouponType
	}
	return CouponType_COUPON_STANDARD
}

func (x *CustomCoupon) GetDiscountAmount() uint32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *CustomCoupon) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *CustomCoupon) GetIsCombinationable() bool {
	if x != nil {
		return x.IsCombinationable
	}
	return false
}

func (x *CustomCoupon) GetNotices() []string {
	if x != nil {
		return x.Notices
	}
	return nil
}

func (x *CustomCoupon) GetTargetStore() []*Store {
	if x != nil {
		return x.TargetStore
	}
	return nil
}

func (x *CustomCoupon) GetStatus() CouponStatus {
	if x != nil {
		return x.Status
	}
	return CouponStatus_COUPON_DRAFT
}

func (x *CustomCoupon) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *CustomCoupon) GetIssueCount() uint32 {
	if x != nil && x.IssueCount != nil {
		return *x.IssueCount
	}
	return 0
}

func (x *CustomCoupon) GetIssueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueAt
	}
	return nil
}

type UserAttachedCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Coupon *Coupon                `protobuf:"bytes,2,opt,name=Coupon,proto3" json:"Coupon,omitempty"`
	UsedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=UsedAt,proto3" json:"UsedAt,omitempty"` // time.Time
}

func (x *UserAttachedCoupon) Reset() {
	*x = UserAttachedCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_Coupon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAttachedCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttachedCoupon) ProtoMessage() {}

func (x *UserAttachedCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_Coupon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttachedCoupon.ProtoReflect.Descriptor instead.
func (*UserAttachedCoupon) Descriptor() ([]byte, []int) {
	return file_v1_shared_Coupon_proto_rawDescGZIP(), []int{2}
}

func (x *UserAttachedCoupon) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserAttachedCoupon) GetCoupon() *Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

func (x *UserAttachedCoupon) GetUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UsedAt
	}
	return nil
}

var File_v1_shared_Coupon_proto protoreflect.FileDescriptor

var file_v1_shared_Coupon_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x02, 0x0a,
	0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36,
	0x0a, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x12, 0x36,
	0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0xb5,
	0x04, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x2c,
	0x0a, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x49, 0x73, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x39, 0x0a, 0x07, 0x49, 0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52,
	0x07, 0x49, 0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x49, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f, 0x42, 0x49, 0x52, 0x54, 0x48, 0x44, 0x41,
	0x59, 0x10, 0x02, 0x2a, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f, 0x44, 0x52,
	0x41, 0x46, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x55,
	0x50, 0x4f, 0x4e, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x44, 0x10, 0x02, 0x42, 0x8b, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x42, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xca, 0x02, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xe2, 0x02, 0x19,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_shared_Coupon_proto_rawDescOnce sync.Once
	file_v1_shared_Coupon_proto_rawDescData = file_v1_shared_Coupon_proto_rawDesc
)

func file_v1_shared_Coupon_proto_rawDescGZIP() []byte {
	file_v1_shared_Coupon_proto_rawDescOnce.Do(func() {
		file_v1_shared_Coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_shared_Coupon_proto_rawDescData)
	})
	return file_v1_shared_Coupon_proto_rawDescData
}

var file_v1_shared_Coupon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_shared_Coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_shared_Coupon_proto_goTypes = []interface{}{
	(CouponType)(0),               // 0: server.shared.CouponType
	(CouponStatus)(0),             // 1: server.shared.CouponStatus
	(*Coupon)(nil),                // 2: server.shared.Coupon
	(*CustomCoupon)(nil),          // 3: server.shared.CustomCoupon
	(*UserAttachedCoupon)(nil),    // 4: server.shared.UserAttachedCoupon
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Store)(nil),                 // 6: server.shared.Store
}
var file_v1_shared_Coupon_proto_depIdxs = []int32{
	0,  // 0: server.shared.Coupon.CouponType:type_name -> server.shared.CouponType
	5,  // 1: server.shared.Coupon.ExpireAt:type_name -> google.protobuf.Timestamp
	6,  // 2: server.shared.Coupon.TargetStore:type_name -> server.shared.Store
	5,  // 3: server.shared.Coupon.CreateAt:type_name -> google.protobuf.Timestamp
	0,  // 4: server.shared.CustomCoupon.CouponType:type_name -> server.shared.CouponType
	5,  // 5: server.shared.CustomCoupon.ExpireAt:type_name -> google.protobuf.Timestamp
	6,  // 6: server.shared.CustomCoupon.TargetStore:type_name -> server.shared.Store
	1,  // 7: server.shared.CustomCoupon.Status:type_name -> server.shared.CouponStatus
	5,  // 8: server.shared.CustomCoupon.CreateAt:type_name -> google.protobuf.Timestamp
	5,  // 9: server.shared.CustomCoupon.IssueAt:type_name -> google.protobuf.Timestamp
	2,  // 10: server.shared.UserAttachedCoupon.Coupon:type_name -> server.shared.Coupon
	5,  // 11: server.shared.UserAttachedCoupon.UsedAt:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_v1_shared_Coupon_proto_init() }
func file_v1_shared_Coupon_proto_init() {
	if File_v1_shared_Coupon_proto != nil {
		return
	}
	file_v1_shared_Store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_shared_Coupon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coupon); i {
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
		file_v1_shared_Coupon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomCoupon); i {
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
		file_v1_shared_Coupon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAttachedCoupon); i {
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
	file_v1_shared_Coupon_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_shared_Coupon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_shared_Coupon_proto_goTypes,
		DependencyIndexes: file_v1_shared_Coupon_proto_depIdxs,
		EnumInfos:         file_v1_shared_Coupon_proto_enumTypes,
		MessageInfos:      file_v1_shared_Coupon_proto_msgTypes,
	}.Build()
	File_v1_shared_Coupon_proto = out.File
	file_v1_shared_Coupon_proto_rawDesc = nil
	file_v1_shared_Coupon_proto_goTypes = nil
	file_v1_shared_Coupon_proto_depIdxs = nil
}
