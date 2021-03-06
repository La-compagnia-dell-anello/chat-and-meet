// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MatchRequest_Gender int32

const (
	MatchRequest_Male      MatchRequest_Gender = 0
	MatchRequest_Female    MatchRequest_Gender = 1
	MatchRequest_NonBinary MatchRequest_Gender = 2
)

// Enum value maps for MatchRequest_Gender.
var (
	MatchRequest_Gender_name = map[int32]string{
		0: "Male",
		1: "Female",
		2: "NonBinary",
	}
	MatchRequest_Gender_value = map[string]int32{
		"Male":      0,
		"Female":    1,
		"NonBinary": 2,
	}
)

func (x MatchRequest_Gender) Enum() *MatchRequest_Gender {
	p := new(MatchRequest_Gender)
	*p = x
	return p
}

func (x MatchRequest_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchRequest_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (MatchRequest_Gender) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x MatchRequest_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchRequest_Gender.Descriptor instead.
func (MatchRequest_Gender) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0, 0}
}

type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyInfo      *MatchRequest_MyInfo      `protobuf:"bytes,1,opt,name=myInfo,proto3" json:"myInfo,omitempty"`
	Preferences *MatchRequest_Preferences `protobuf:"bytes,2,opt,name=preferences,proto3" json:"preferences,omitempty"`
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRequest) GetMyInfo() *MatchRequest_MyInfo {
	if x != nil {
		return x.MyInfo
	}
	return nil
}

func (x *MatchRequest) GetPreferences() *MatchRequest_Preferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

type MatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ChatKey string `protobuf:"bytes,2,opt,name=chatKey,proto3" json:"chatKey,omitempty"`
}

func (x *MatchResponse) Reset() {
	*x = MatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResponse) ProtoMessage() {}

func (x *MatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResponse.ProtoReflect.Descriptor instead.
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *MatchResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MatchResponse) GetChatKey() string {
	if x != nil {
		return x.ChatKey
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MatchRequest_MyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gender    MatchRequest_Gender `protobuf:"varint,1,opt,name=gender,proto3,enum=MatchRequest_Gender" json:"gender,omitempty"`
	Age       uint32              `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Latitude  float64             `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64             `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *MatchRequest_MyInfo) Reset() {
	*x = MatchRequest_MyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest_MyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest_MyInfo) ProtoMessage() {}

func (x *MatchRequest_MyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest_MyInfo.ProtoReflect.Descriptor instead.
func (*MatchRequest_MyInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchRequest_MyInfo) GetGender() MatchRequest_Gender {
	if x != nil {
		return x.Gender
	}
	return MatchRequest_Male
}

func (x *MatchRequest_MyInfo) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *MatchRequest_MyInfo) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *MatchRequest_MyInfo) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type MatchRequest_Preferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KilometersRange int32               `protobuf:"varint,1,opt,name=kilometersRange,proto3" json:"kilometersRange,omitempty"`
	Gender          MatchRequest_Gender `protobuf:"varint,2,opt,name=gender,proto3,enum=MatchRequest_Gender" json:"gender,omitempty"`
	MinAge          uint32              `protobuf:"varint,3,opt,name=minAge,proto3" json:"minAge,omitempty"`
	MaxAge          uint32              `protobuf:"varint,4,opt,name=maxAge,proto3" json:"maxAge,omitempty"`
}

func (x *MatchRequest_Preferences) Reset() {
	*x = MatchRequest_Preferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest_Preferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest_Preferences) ProtoMessage() {}

func (x *MatchRequest_Preferences) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest_Preferences.ProtoReflect.Descriptor instead.
func (*MatchRequest_Preferences) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MatchRequest_Preferences) GetKilometersRange() int32 {
	if x != nil {
		return x.KilometersRange
	}
	return 0
}

func (x *MatchRequest_Preferences) GetGender() MatchRequest_Gender {
	if x != nil {
		return x.Gender
	}
	return MatchRequest_Male
}

func (x *MatchRequest_Preferences) GetMinAge() uint32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *MatchRequest_Preferences) GetMaxAge() uint32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc5, 0x03, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b,
	0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0b,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x82, 0x01, 0x0a, 0x06,
	0x4d, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x1a, 0x95, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x28, 0x0a, 0x0f, 0x6b, 0x69, 0x6c, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6b, 0x69, 0x6c, 0x6f, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x41,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x41, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x6e, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x10, 0x02, 0x22, 0x43, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x1d, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x5c, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x0d, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x25, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(MatchRequest_Gender)(0),         // 0: MatchRequest.Gender
	(*MatchRequest)(nil),             // 1: MatchRequest
	(*MatchResponse)(nil),            // 2: MatchResponse
	(*Message)(nil),                  // 3: Message
	(*MatchRequest_MyInfo)(nil),      // 4: MatchRequest.MyInfo
	(*MatchRequest_Preferences)(nil), // 5: MatchRequest.Preferences
}
var file_service_proto_depIdxs = []int32{
	4, // 0: MatchRequest.myInfo:type_name -> MatchRequest.MyInfo
	5, // 1: MatchRequest.preferences:type_name -> MatchRequest.Preferences
	0, // 2: MatchRequest.MyInfo.gender:type_name -> MatchRequest.Gender
	0, // 3: MatchRequest.Preferences.gender:type_name -> MatchRequest.Gender
	1, // 4: Service.Match:input_type -> MatchRequest
	3, // 5: Service.StartChat:input_type -> Message
	2, // 6: Service.Match:output_type -> MatchResponse
	3, // 7: Service.StartChat:output_type -> Message
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest_MyInfo); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest_Preferences); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (Service_MatchClient, error)
	StartChat(ctx context.Context, opts ...grpc.CallOption) (Service_StartChatClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Match(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (Service_MatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/Service/Match", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceMatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_MatchClient interface {
	Recv() (*MatchResponse, error)
	grpc.ClientStream
}

type serviceMatchClient struct {
	grpc.ClientStream
}

func (x *serviceMatchClient) Recv() (*MatchResponse, error) {
	m := new(MatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) StartChat(ctx context.Context, opts ...grpc.CallOption) (Service_StartChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[1], "/Service/StartChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceStartChatClient{stream}
	return x, nil
}

type Service_StartChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type serviceStartChatClient struct {
	grpc.ClientStream
}

func (x *serviceStartChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceStartChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	Match(*MatchRequest, Service_MatchServer) error
	StartChat(Service_StartChatServer) error
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Match(*MatchRequest, Service_MatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (*UnimplementedServiceServer) StartChat(Service_StartChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StartChat not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Match_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Match(m, &serviceMatchServer{stream})
}

type Service_MatchServer interface {
	Send(*MatchResponse) error
	grpc.ServerStream
}

type serviceMatchServer struct {
	grpc.ServerStream
}

func (x *serviceMatchServer) Send(m *MatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_StartChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).StartChat(&serviceStartChatServer{stream})
}

type Service_StartChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type serviceStartChatServer struct {
	grpc.ServerStream
}

func (x *serviceStartChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceStartChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Match",
			Handler:       _Service_Match_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StartChat",
			Handler:       _Service_StartChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
