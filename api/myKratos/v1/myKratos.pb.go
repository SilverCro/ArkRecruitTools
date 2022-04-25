// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: myKratos/v1/myKratos.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// SayHello api 请求格式
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SayHello api 返回格式
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// ArkRecruitImageTool api 请求格式
type ArkRecruitRecommendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ArkRecruitRecommendRequest) Reset() {
	*x = ArkRecruitRecommendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArkRecruitRecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArkRecruitRecommendRequest) ProtoMessage() {}

func (x *ArkRecruitRecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArkRecruitRecommendRequest.ProtoReflect.Descriptor instead.
func (*ArkRecruitRecommendRequest) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{2}
}

func (x *ArkRecruitRecommendRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// ArkRecruitImageTool api 返回格式
type ArkRecruitRecommendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                string                                      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	RecommendTags         []string                                    `protobuf:"bytes,2,rep,name=recommendTags,proto3" json:"recommendTags,omitempty"`
	RecommendOperatorInfo []*ArkRecruitRecommendReply_ArkOperatorInfo `protobuf:"bytes,3,rep,name=recommendOperatorInfo,proto3" json:"recommendOperatorInfo,omitempty"`
}

func (x *ArkRecruitRecommendReply) Reset() {
	*x = ArkRecruitRecommendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArkRecruitRecommendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArkRecruitRecommendReply) ProtoMessage() {}

func (x *ArkRecruitRecommendReply) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArkRecruitRecommendReply.ProtoReflect.Descriptor instead.
func (*ArkRecruitRecommendReply) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{3}
}

func (x *ArkRecruitRecommendReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ArkRecruitRecommendReply) GetRecommendTags() []string {
	if x != nil {
		return x.RecommendTags
	}
	return nil
}

func (x *ArkRecruitRecommendReply) GetRecommendOperatorInfo() []*ArkRecruitRecommendReply_ArkOperatorInfo {
	if x != nil {
		return x.RecommendOperatorInfo
	}
	return nil
}

// GetArkOperatorInfo api 请求格式
type GetArkOperatorInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetArkOperatorInfoRequest) Reset() {
	*x = GetArkOperatorInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArkOperatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArkOperatorInfoRequest) ProtoMessage() {}

func (x *GetArkOperatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArkOperatorInfoRequest.ProtoReflect.Descriptor instead.
func (*GetArkOperatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{4}
}

func (x *GetArkOperatorInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GetArkOperatorInfo api 返回格式
type GetArkOperatorInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArkOperatorInfo *GetArkOperatorInfoReply_ArkOperatorInfo `protobuf:"bytes,1,opt,name=arkOperatorInfo,proto3" json:"arkOperatorInfo,omitempty"`
}

func (x *GetArkOperatorInfoReply) Reset() {
	*x = GetArkOperatorInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArkOperatorInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArkOperatorInfoReply) ProtoMessage() {}

func (x *GetArkOperatorInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArkOperatorInfoReply.ProtoReflect.Descriptor instead.
func (*GetArkOperatorInfoReply) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{5}
}

func (x *GetArkOperatorInfoReply) GetArkOperatorInfo() *GetArkOperatorInfoReply_ArkOperatorInfo {
	if x != nil {
		return x.ArkOperatorInfo
	}
	return nil
}

type ArkRecruitRecommendReply_ArkOperatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TarList    []string `protobuf:"bytes,2,rep,name=tarList,proto3" json:"tarList,omitempty"`
	Profession string   `protobuf:"bytes,3,opt,name=profession,proto3" json:"profession,omitempty"`
	Rarity     int32    `protobuf:"varint,4,opt,name=rarity,proto3" json:"rarity,omitempty"`
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) Reset() {
	*x = ArkRecruitRecommendReply_ArkOperatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArkRecruitRecommendReply_ArkOperatorInfo) ProtoMessage() {}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArkRecruitRecommendReply_ArkOperatorInfo.ProtoReflect.Descriptor instead.
func (*ArkRecruitRecommendReply_ArkOperatorInfo) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) GetTarList() []string {
	if x != nil {
		return x.TarList
	}
	return nil
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *ArkRecruitRecommendReply_ArkOperatorInfo) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

type GetArkOperatorInfoReply_ArkOperatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TarList    []string `protobuf:"bytes,2,rep,name=tarList,proto3" json:"tarList,omitempty"`
	Profession string   `protobuf:"bytes,3,opt,name=profession,proto3" json:"profession,omitempty"`
	Rarity     int32    `protobuf:"varint,4,opt,name=rarity,proto3" json:"rarity,omitempty"`
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) Reset() {
	*x = GetArkOperatorInfoReply_ArkOperatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myKratos_v1_myKratos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArkOperatorInfoReply_ArkOperatorInfo) ProtoMessage() {}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_myKratos_v1_myKratos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArkOperatorInfoReply_ArkOperatorInfo.ProtoReflect.Descriptor instead.
func (*GetArkOperatorInfoReply_ArkOperatorInfo) Descriptor() ([]byte, []int) {
	return file_myKratos_v1_myKratos_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) GetTarList() []string {
	if x != nil {
		return x.TarList
	}
	return nil
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) GetProfession() string {
	if x != nil {
		return x.Profession
	}
	return ""
}

func (x *GetArkOperatorInfoReply_ArkOperatorInfo) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

var File_myKratos_v1_myKratos_proto protoreflect.FileDescriptor

var file_myKratos_v1_myKratos_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x79,
	0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x79,
	0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x1a, 0x41, 0x72, 0x6b, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xbe, 0x02, 0x0a, 0x18, 0x41, 0x72, 0x6b, 0x52, 0x65, 0x63,
	0x72, 0x75, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x6b, 0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x6b, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x77, 0x0a,
	0x0f, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6b,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0f, 0x61, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6d,
	0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x61, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x77, 0x0a, 0x0f, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x32, 0x8c, 0x03, 0x0a,
	0x08, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x5a, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x96, 0x01, 0x0a, 0x17, 0x41, 0x72, 0x6b, 0x52, 0x65, 0x63,
	0x72, 0x75, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6f,
	0x6c, 0x12, 0x27, 0x2e, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x6b, 0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x79, 0x4b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x6b, 0x52, 0x65, 0x63, 0x72,
	0x75, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x72, 0x6b, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69,
	0x74, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x8a,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6d, 0x79, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x72, 0x6b, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x1f, 0x5a, 0x1d, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x54, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x79,
	0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myKratos_v1_myKratos_proto_rawDescOnce sync.Once
	file_myKratos_v1_myKratos_proto_rawDescData = file_myKratos_v1_myKratos_proto_rawDesc
)

func file_myKratos_v1_myKratos_proto_rawDescGZIP() []byte {
	file_myKratos_v1_myKratos_proto_rawDescOnce.Do(func() {
		file_myKratos_v1_myKratos_proto_rawDescData = protoimpl.X.CompressGZIP(file_myKratos_v1_myKratos_proto_rawDescData)
	})
	return file_myKratos_v1_myKratos_proto_rawDescData
}

var file_myKratos_v1_myKratos_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_myKratos_v1_myKratos_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),                             // 0: myKratos.v1.HelloRequest
	(*HelloReply)(nil),                               // 1: myKratos.v1.HelloReply
	(*ArkRecruitRecommendRequest)(nil),               // 2: myKratos.v1.ArkRecruitRecommendRequest
	(*ArkRecruitRecommendReply)(nil),                 // 3: myKratos.v1.ArkRecruitRecommendReply
	(*GetArkOperatorInfoRequest)(nil),                // 4: myKratos.v1.GetArkOperatorInfoRequest
	(*GetArkOperatorInfoReply)(nil),                  // 5: myKratos.v1.GetArkOperatorInfoReply
	(*ArkRecruitRecommendReply_ArkOperatorInfo)(nil), // 6: myKratos.v1.ArkRecruitRecommendReply.ArkOperatorInfo
	(*GetArkOperatorInfoReply_ArkOperatorInfo)(nil),  // 7: myKratos.v1.GetArkOperatorInfoReply.ArkOperatorInfo
}
var file_myKratos_v1_myKratos_proto_depIdxs = []int32{
	6, // 0: myKratos.v1.ArkRecruitRecommendReply.recommendOperatorInfo:type_name -> myKratos.v1.ArkRecruitRecommendReply.ArkOperatorInfo
	7, // 1: myKratos.v1.GetArkOperatorInfoReply.arkOperatorInfo:type_name -> myKratos.v1.GetArkOperatorInfoReply.ArkOperatorInfo
	0, // 2: myKratos.v1.myKratos.SayHello:input_type -> myKratos.v1.HelloRequest
	2, // 3: myKratos.v1.myKratos.ArkRecruitRecommendTool:input_type -> myKratos.v1.ArkRecruitRecommendRequest
	4, // 4: myKratos.v1.myKratos.GetArkOperatorInfo:input_type -> myKratos.v1.GetArkOperatorInfoRequest
	1, // 5: myKratos.v1.myKratos.SayHello:output_type -> myKratos.v1.HelloReply
	3, // 6: myKratos.v1.myKratos.ArkRecruitRecommendTool:output_type -> myKratos.v1.ArkRecruitRecommendReply
	5, // 7: myKratos.v1.myKratos.GetArkOperatorInfo:output_type -> myKratos.v1.GetArkOperatorInfoReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_myKratos_v1_myKratos_proto_init() }
func file_myKratos_v1_myKratos_proto_init() {
	if File_myKratos_v1_myKratos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myKratos_v1_myKratos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArkRecruitRecommendRequest); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArkRecruitRecommendReply); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArkOperatorInfoRequest); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArkOperatorInfoReply); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArkRecruitRecommendReply_ArkOperatorInfo); i {
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
		file_myKratos_v1_myKratos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArkOperatorInfoReply_ArkOperatorInfo); i {
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
			RawDescriptor: file_myKratos_v1_myKratos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_myKratos_v1_myKratos_proto_goTypes,
		DependencyIndexes: file_myKratos_v1_myKratos_proto_depIdxs,
		MessageInfos:      file_myKratos_v1_myKratos_proto_msgTypes,
	}.Build()
	File_myKratos_v1_myKratos_proto = out.File
	file_myKratos_v1_myKratos_proto_rawDesc = nil
	file_myKratos_v1_myKratos_proto_goTypes = nil
	file_myKratos_v1_myKratos_proto_depIdxs = nil
}
