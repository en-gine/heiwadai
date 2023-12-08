// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/user/Checkin.proto

package user

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

type CheckinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QrHash string `protobuf:"bytes,1,opt,name=qrHash,proto3" json:"qrHash,omitempty"`
}

func (x *CheckinRequest) Reset() {
	*x = CheckinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinRequest) ProtoMessage() {}

func (x *CheckinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinRequest.ProtoReflect.Descriptor instead.
func (*CheckinRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_Checkin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckinRequest) GetQrHash() string {
	if x != nil {
		return x.QrHash
	}
	return ""
}

type CheckinStamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StoreID         string                 `protobuf:"bytes,2,opt,name=StoreID,proto3" json:"StoreID,omitempty"`
	StoreName       string                 `protobuf:"bytes,3,opt,name=StoreName,proto3" json:"StoreName,omitempty"`
	StoreStampImage string                 `protobuf:"bytes,4,opt,name=StoreStampImage,proto3" json:"StoreStampImage,omitempty"`
	CheckInAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CheckInAt,proto3" json:"CheckInAt,omitempty"`
}

func (x *CheckinStamp) Reset() {
	*x = CheckinStamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinStamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinStamp) ProtoMessage() {}

func (x *CheckinStamp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinStamp.ProtoReflect.Descriptor instead.
func (*CheckinStamp) Descriptor() ([]byte, []int) {
	return file_v1_user_Checkin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckinStamp) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CheckinStamp) GetStoreID() string {
	if x != nil {
		return x.StoreID
	}
	return ""
}

func (x *CheckinStamp) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *CheckinStamp) GetStoreStampImage() string {
	if x != nil {
		return x.StoreStampImage
	}
	return ""
}

func (x *CheckinStamp) GetCheckInAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CheckInAt
	}
	return nil
}

type StampCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stamps []*CheckinStamp `protobuf:"bytes,1,rep,name=stamps,proto3" json:"stamps,omitempty"`
}

func (x *StampCardResponse) Reset() {
	*x = StampCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Checkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StampCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StampCardResponse) ProtoMessage() {}

func (x *StampCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Checkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StampCardResponse.ProtoReflect.Descriptor instead.
func (*StampCardResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_Checkin_proto_rawDescGZIP(), []int{2}
}

func (x *StampCardResponse) GetStamps() []*CheckinStamp {
	if x != nil {
		return x.Stamps
	}
	return nil
}

type CheckinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MayCoupon *shared.Coupon `protobuf:"bytes,1,opt,name=MayCoupon,proto3,oneof" json:"MayCoupon,omitempty"`
}

func (x *CheckinResponse) Reset() {
	*x = CheckinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Checkin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinResponse) ProtoMessage() {}

func (x *CheckinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Checkin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinResponse.ProtoReflect.Descriptor instead.
func (*CheckinResponse) Descriptor() ([]byte, []int) {
	return file_v1_user_Checkin_proto_rawDescGZIP(), []int{3}
}

func (x *CheckinResponse) GetMayCoupon() *shared.Coupon {
	if x != nil {
		return x.MayCoupon
	}
	return nil
}

var File_v1_user_Checkin_proto protoreflect.FileDescriptor

var file_v1_user_Checkin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x71, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x41,
	0x74, 0x22, 0x46, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x06, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x22, 0x59, 0x0a, 0x0f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x4d, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x4d, 0x61, 0x79, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x4d, 0x61, 0x79, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x32, 0xa5, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x80, 0x01, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5c, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c,
	0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_Checkin_proto_rawDescOnce sync.Once
	file_v1_user_Checkin_proto_rawDescData = file_v1_user_Checkin_proto_rawDesc
)

func file_v1_user_Checkin_proto_rawDescGZIP() []byte {
	file_v1_user_Checkin_proto_rawDescOnce.Do(func() {
		file_v1_user_Checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_Checkin_proto_rawDescData)
	})
	return file_v1_user_Checkin_proto_rawDescData
}

var file_v1_user_Checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_user_Checkin_proto_goTypes = []interface{}{
	(*CheckinRequest)(nil),        // 0: server.user.CheckinRequest
	(*CheckinStamp)(nil),          // 1: server.user.CheckinStamp
	(*StampCardResponse)(nil),     // 2: server.user.StampCardResponse
	(*CheckinResponse)(nil),       // 3: server.user.CheckinResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*shared.Coupon)(nil),         // 5: server.shared.Coupon
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_v1_user_Checkin_proto_depIdxs = []int32{
	4, // 0: server.user.CheckinStamp.CheckInAt:type_name -> google.protobuf.Timestamp
	1, // 1: server.user.StampCardResponse.stamps:type_name -> server.user.CheckinStamp
	5, // 2: server.user.CheckinResponse.MayCoupon:type_name -> server.shared.Coupon
	6, // 3: server.user.CheckinController.GetStampCard:input_type -> google.protobuf.Empty
	0, // 4: server.user.CheckinController.Checkin:input_type -> server.user.CheckinRequest
	2, // 5: server.user.CheckinController.GetStampCard:output_type -> server.user.StampCardResponse
	3, // 6: server.user.CheckinController.Checkin:output_type -> server.user.CheckinResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_user_Checkin_proto_init() }
func file_v1_user_Checkin_proto_init() {
	if File_v1_user_Checkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_Checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinRequest); i {
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
		file_v1_user_Checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinStamp); i {
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
		file_v1_user_Checkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StampCardResponse); i {
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
		file_v1_user_Checkin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinResponse); i {
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
	file_v1_user_Checkin_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_user_Checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_Checkin_proto_goTypes,
		DependencyIndexes: file_v1_user_Checkin_proto_depIdxs,
		MessageInfos:      file_v1_user_Checkin_proto_msgTypes,
	}.Build()
	File_v1_user_Checkin_proto = out.File
	file_v1_user_Checkin_proto_rawDesc = nil
	file_v1_user_Checkin_proto_goTypes = nil
	file_v1_user_Checkin_proto_depIdxs = nil
}
