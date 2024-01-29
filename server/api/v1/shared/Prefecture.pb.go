// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: v1/shared/Prefecture.proto

package shared

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

type Prefecture int32

const (
	Prefecture_Unspecified Prefecture = 0
	Prefecture_Hokkaido    Prefecture = 1
	Prefecture_Aomori      Prefecture = 2
	Prefecture_Iwate       Prefecture = 3
	Prefecture_Miyagi      Prefecture = 4
	Prefecture_Akita       Prefecture = 5
	Prefecture_Yamagata    Prefecture = 6
	Prefecture_Fukushima   Prefecture = 7
	Prefecture_Ibaraki     Prefecture = 8
	Prefecture_Tochigi     Prefecture = 9
	Prefecture_Gunma       Prefecture = 10
	Prefecture_Saitama     Prefecture = 11
	Prefecture_Chiba       Prefecture = 12
	Prefecture_Tokyo       Prefecture = 13
	Prefecture_Kanagawa    Prefecture = 14
	Prefecture_Niigata     Prefecture = 15
	Prefecture_Toyama      Prefecture = 16
	Prefecture_Ishikawa    Prefecture = 17
	Prefecture_Fukui       Prefecture = 18
	Prefecture_Yamanashi   Prefecture = 19
	Prefecture_Nagano      Prefecture = 20
	Prefecture_Gifu        Prefecture = 21
	Prefecture_Shizuoka    Prefecture = 22
	Prefecture_Aichi       Prefecture = 23
	Prefecture_Mie         Prefecture = 24
	Prefecture_Shiga       Prefecture = 25
	Prefecture_Kyoto       Prefecture = 26
	Prefecture_Osaka       Prefecture = 27
	Prefecture_Hyogo       Prefecture = 28
	Prefecture_Nara        Prefecture = 29
	Prefecture_Wakayama    Prefecture = 30
	Prefecture_Tottori     Prefecture = 31
	Prefecture_Shimane     Prefecture = 32
	Prefecture_Okayama     Prefecture = 33
	Prefecture_Hiroshima   Prefecture = 34
	Prefecture_Yamaguchi   Prefecture = 35
	Prefecture_Tokushima   Prefecture = 36
	Prefecture_Kagawa      Prefecture = 37
	Prefecture_Ehime       Prefecture = 38
	Prefecture_Kochi       Prefecture = 39
	Prefecture_Fukuoka     Prefecture = 40
	Prefecture_Saga        Prefecture = 41
	Prefecture_Nagasaki    Prefecture = 42
	Prefecture_Kumamoto    Prefecture = 43
	Prefecture_Oita        Prefecture = 44
	Prefecture_Miyazaki    Prefecture = 45
	Prefecture_Kagoshima   Prefecture = 46
	Prefecture_Okinawa     Prefecture = 47
)

// Enum value maps for Prefecture.
var (
	Prefecture_name = map[int32]string{
		0:  "Unspecified",
		1:  "Hokkaido",
		2:  "Aomori",
		3:  "Iwate",
		4:  "Miyagi",
		5:  "Akita",
		6:  "Yamagata",
		7:  "Fukushima",
		8:  "Ibaraki",
		9:  "Tochigi",
		10: "Gunma",
		11: "Saitama",
		12: "Chiba",
		13: "Tokyo",
		14: "Kanagawa",
		15: "Niigata",
		16: "Toyama",
		17: "Ishikawa",
		18: "Fukui",
		19: "Yamanashi",
		20: "Nagano",
		21: "Gifu",
		22: "Shizuoka",
		23: "Aichi",
		24: "Mie",
		25: "Shiga",
		26: "Kyoto",
		27: "Osaka",
		28: "Hyogo",
		29: "Nara",
		30: "Wakayama",
		31: "Tottori",
		32: "Shimane",
		33: "Okayama",
		34: "Hiroshima",
		35: "Yamaguchi",
		36: "Tokushima",
		37: "Kagawa",
		38: "Ehime",
		39: "Kochi",
		40: "Fukuoka",
		41: "Saga",
		42: "Nagasaki",
		43: "Kumamoto",
		44: "Oita",
		45: "Miyazaki",
		46: "Kagoshima",
		47: "Okinawa",
	}
	Prefecture_value = map[string]int32{
		"Unspecified": 0,
		"Hokkaido":    1,
		"Aomori":      2,
		"Iwate":       3,
		"Miyagi":      4,
		"Akita":       5,
		"Yamagata":    6,
		"Fukushima":   7,
		"Ibaraki":     8,
		"Tochigi":     9,
		"Gunma":       10,
		"Saitama":     11,
		"Chiba":       12,
		"Tokyo":       13,
		"Kanagawa":    14,
		"Niigata":     15,
		"Toyama":      16,
		"Ishikawa":    17,
		"Fukui":       18,
		"Yamanashi":   19,
		"Nagano":      20,
		"Gifu":        21,
		"Shizuoka":    22,
		"Aichi":       23,
		"Mie":         24,
		"Shiga":       25,
		"Kyoto":       26,
		"Osaka":       27,
		"Hyogo":       28,
		"Nara":        29,
		"Wakayama":    30,
		"Tottori":     31,
		"Shimane":     32,
		"Okayama":     33,
		"Hiroshima":   34,
		"Yamaguchi":   35,
		"Tokushima":   36,
		"Kagawa":      37,
		"Ehime":       38,
		"Kochi":       39,
		"Fukuoka":     40,
		"Saga":        41,
		"Nagasaki":    42,
		"Kumamoto":    43,
		"Oita":        44,
		"Miyazaki":    45,
		"Kagoshima":   46,
		"Okinawa":     47,
	}
)

func (x Prefecture) Enum() *Prefecture {
	p := new(Prefecture)
	*p = x
	return p
}

func (x Prefecture) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prefecture) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_Prefecture_proto_enumTypes[0].Descriptor()
}

func (Prefecture) Type() protoreflect.EnumType {
	return &file_v1_shared_Prefecture_proto_enumTypes[0]
}

func (x Prefecture) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prefecture.Descriptor instead.
func (Prefecture) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_Prefecture_proto_rawDescGZIP(), []int{0}
}

var File_v1_shared_Prefecture_proto protoreflect.FileDescriptor

var file_v1_shared_Prefecture_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2a, 0xe6, 0x04, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x48,
	0x6f, 0x6b, 0x6b, 0x61, 0x69, 0x64, 0x6f, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6f, 0x6d,
	0x6f, 0x72, 0x69, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x77, 0x61, 0x74, 0x65, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x69, 0x79, 0x61, 0x67, 0x69, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x6b, 0x69, 0x74, 0x61, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x59, 0x61, 0x6d, 0x61, 0x67,
	0x61, 0x74, 0x61, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x75, 0x6b, 0x75, 0x73, 0x68, 0x69,
	0x6d, 0x61, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x62, 0x61, 0x72, 0x61, 0x6b, 0x69, 0x10,
	0x08, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x6f, 0x63, 0x68, 0x69, 0x67, 0x69, 0x10, 0x09, 0x12, 0x09,
	0x0a, 0x05, 0x47, 0x75, 0x6e, 0x6d, 0x61, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x61, 0x69,
	0x74, 0x61, 0x6d, 0x61, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x62, 0x61, 0x10,
	0x0c, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x79, 0x6f, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08,
	0x4b, 0x61, 0x6e, 0x61, 0x67, 0x61, 0x77, 0x61, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x69,
	0x69, 0x67, 0x61, 0x74, 0x61, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x6f, 0x79, 0x61, 0x6d,
	0x61, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x73, 0x68, 0x69, 0x6b, 0x61, 0x77, 0x61, 0x10,
	0x11, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x75, 0x6b, 0x75, 0x69, 0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09,
	0x59, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x73, 0x68, 0x69, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x4e,
	0x61, 0x67, 0x61, 0x6e, 0x6f, 0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x69, 0x66, 0x75, 0x10,
	0x15, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x7a, 0x75, 0x6f, 0x6b, 0x61, 0x10, 0x16, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x69, 0x63, 0x68, 0x69, 0x10, 0x17, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x69,
	0x65, 0x10, 0x18, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x68, 0x69, 0x67, 0x61, 0x10, 0x19, 0x12, 0x09,
	0x0a, 0x05, 0x4b, 0x79, 0x6f, 0x74, 0x6f, 0x10, 0x1a, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x73, 0x61,
	0x6b, 0x61, 0x10, 0x1b, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x79, 0x6f, 0x67, 0x6f, 0x10, 0x1c, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x61, 0x72, 0x61, 0x10, 0x1d, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x61, 0x6b,
	0x61, 0x79, 0x61, 0x6d, 0x61, 0x10, 0x1e, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x6f, 0x74, 0x74, 0x6f,
	0x72, 0x69, 0x10, 0x1f, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x68, 0x69, 0x6d, 0x61, 0x6e, 0x65, 0x10,
	0x20, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x6b, 0x61, 0x79, 0x61, 0x6d, 0x61, 0x10, 0x21, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x69, 0x72, 0x6f, 0x73, 0x68, 0x69, 0x6d, 0x61, 0x10, 0x22, 0x12, 0x0d, 0x0a,
	0x09, 0x59, 0x61, 0x6d, 0x61, 0x67, 0x75, 0x63, 0x68, 0x69, 0x10, 0x23, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x6f, 0x6b, 0x75, 0x73, 0x68, 0x69, 0x6d, 0x61, 0x10, 0x24, 0x12, 0x0a, 0x0a, 0x06, 0x4b,
	0x61, 0x67, 0x61, 0x77, 0x61, 0x10, 0x25, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x68, 0x69, 0x6d, 0x65,
	0x10, 0x26, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x6f, 0x63, 0x68, 0x69, 0x10, 0x27, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x75, 0x6b, 0x75, 0x6f, 0x6b, 0x61, 0x10, 0x28, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x61,
	0x67, 0x61, 0x10, 0x29, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x61, 0x67, 0x61, 0x73, 0x61, 0x6b, 0x69,
	0x10, 0x2a, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x75, 0x6d, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x10, 0x2b,
	0x12, 0x08, 0x0a, 0x04, 0x4f, 0x69, 0x74, 0x61, 0x10, 0x2c, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x69,
	0x79, 0x61, 0x7a, 0x61, 0x6b, 0x69, 0x10, 0x2d, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x61, 0x67, 0x6f,
	0x73, 0x68, 0x69, 0x6d, 0x61, 0x10, 0x2e, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x6b, 0x69, 0x6e, 0x61,
	0x77, 0x61, 0x10, 0x2f, 0x42, 0x8f, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x42, 0x0f, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xca, 0x02, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xe2, 0x02, 0x19, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_shared_Prefecture_proto_rawDescOnce sync.Once
	file_v1_shared_Prefecture_proto_rawDescData = file_v1_shared_Prefecture_proto_rawDesc
)

func file_v1_shared_Prefecture_proto_rawDescGZIP() []byte {
	file_v1_shared_Prefecture_proto_rawDescOnce.Do(func() {
		file_v1_shared_Prefecture_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_shared_Prefecture_proto_rawDescData)
	})
	return file_v1_shared_Prefecture_proto_rawDescData
}

var file_v1_shared_Prefecture_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_shared_Prefecture_proto_goTypes = []interface{}{
	(Prefecture)(0), // 0: server.shared.Prefecture
}
var file_v1_shared_Prefecture_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_shared_Prefecture_proto_init() }
func file_v1_shared_Prefecture_proto_init() {
	if File_v1_shared_Prefecture_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_shared_Prefecture_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_shared_Prefecture_proto_goTypes,
		DependencyIndexes: file_v1_shared_Prefecture_proto_depIdxs,
		EnumInfos:         file_v1_shared_Prefecture_proto_enumTypes,
	}.Build()
	File_v1_shared_Prefecture_proto = out.File
	file_v1_shared_Prefecture_proto_rawDesc = nil
	file_v1_shared_Prefecture_proto_goTypes = nil
	file_v1_shared_Prefecture_proto_depIdxs = nil
}
