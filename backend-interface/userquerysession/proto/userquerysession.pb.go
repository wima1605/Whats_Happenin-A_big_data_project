// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: userquerysession.proto

package proto

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

type UserQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location  string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	DateStart string `protobuf:"bytes,3,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd   string `protobuf:"bytes,4,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *UserQuery) Reset() {
	*x = UserQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQuery) ProtoMessage() {}

func (x *UserQuery) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQuery.ProtoReflect.Descriptor instead.
func (*UserQuery) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{0}
}

func (x *UserQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserQuery) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *UserQuery) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *UserQuery) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Overallresponse []byte `protobuf:"bytes,2,opt,name=overallresponse,proto3" json:"overallresponse,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{1}
}

func (x *UserID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserID) GetOverallresponse() []byte {
	if x != nil {
		return x.Overallresponse
	}
	return nil
}

type OverallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid string `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *OverallRequest) Reset() {
	*x = OverallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallRequest) ProtoMessage() {}

func (x *OverallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallRequest.ProtoReflect.Descriptor instead.
func (*OverallRequest) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{2}
}

func (x *OverallRequest) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

type OverallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OidResponse []byte `protobuf:"bytes,1,opt,name=oidResponse,proto3" json:"oidResponse,omitempty"`
}

func (x *OverallResponse) Reset() {
	*x = OverallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallResponse) ProtoMessage() {}

func (x *OverallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallResponse.ProtoReflect.Descriptor instead.
func (*OverallResponse) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{3}
}

func (x *OverallResponse) GetOidResponse() []byte {
	if x != nil {
		return x.OidResponse
	}
	return nil
}

type TopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Oid   string `protobuf:"bytes,2,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *TopicRequest) Reset() {
	*x = TopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicRequest) ProtoMessage() {}

func (x *TopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicRequest.ProtoReflect.Descriptor instead.
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{4}
}

func (x *TopicRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicRequest) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

type TopicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticlesResponse []byte `protobuf:"bytes,1,opt,name=articlesResponse,proto3" json:"articlesResponse,omitempty"`
}

func (x *TopicResponse) Reset() {
	*x = TopicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userquerysession_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicResponse) ProtoMessage() {}

func (x *TopicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userquerysession_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicResponse.ProtoReflect.Descriptor instead.
func (*TopicResponse) Descriptor() ([]byte, []int) {
	return file_userquerysession_proto_rawDescGZIP(), []int{5}
}

func (x *TopicResponse) GetArticlesResponse() []byte {
	if x != nil {
		return x.ArticlesResponse
	}
	return nil
}

var File_userquerysession_proto protoreflect.FileDescriptor

var file_userquerysession_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x22, 0x42, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x6c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6f,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x8b, 0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x5a, 0x0a, 0x13, 0x53,
	0x65, 0x6e, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x65, 0x5a,
	0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x6e,
	0x79, 0x61, 0x63, 0x68, 0x61, 0x75, 0x62, 0x65, 0x79, 0x2f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x5f,
	0x48, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x2d, 0x41, 0x5f, 0x62, 0x69, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userquerysession_proto_rawDescOnce sync.Once
	file_userquerysession_proto_rawDescData = file_userquerysession_proto_rawDesc
)

func file_userquerysession_proto_rawDescGZIP() []byte {
	file_userquerysession_proto_rawDescOnce.Do(func() {
		file_userquerysession_proto_rawDescData = protoimpl.X.CompressGZIP(file_userquerysession_proto_rawDescData)
	})
	return file_userquerysession_proto_rawDescData
}

var file_userquerysession_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_userquerysession_proto_goTypes = []interface{}{
	(*UserQuery)(nil),       // 0: userquerysession.UserQuery
	(*UserID)(nil),          // 1: userquerysession.UserID
	(*OverallRequest)(nil),  // 2: userquerysession.OverallRequest
	(*OverallResponse)(nil), // 3: userquerysession.OverallResponse
	(*TopicRequest)(nil),    // 4: userquerysession.TopicRequest
	(*TopicResponse)(nil),   // 5: userquerysession.TopicResponse
}
var file_userquerysession_proto_depIdxs = []int32{
	0, // 0: userquerysession.UserQueryService.StartSession:input_type -> userquerysession.UserQuery
	2, // 1: userquerysession.UserQueryService.SendOverallResponse:input_type -> userquerysession.OverallRequest
	4, // 2: userquerysession.UserQueryService.SendTopicResponse:input_type -> userquerysession.TopicRequest
	1, // 3: userquerysession.UserQueryService.StartSession:output_type -> userquerysession.UserID
	3, // 4: userquerysession.UserQueryService.SendOverallResponse:output_type -> userquerysession.OverallResponse
	5, // 5: userquerysession.UserQueryService.SendTopicResponse:output_type -> userquerysession.TopicResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userquerysession_proto_init() }
func file_userquerysession_proto_init() {
	if File_userquerysession_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userquerysession_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQuery); i {
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
		file_userquerysession_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_userquerysession_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverallRequest); i {
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
		file_userquerysession_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverallResponse); i {
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
		file_userquerysession_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicRequest); i {
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
		file_userquerysession_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicResponse); i {
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
			RawDescriptor: file_userquerysession_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userquerysession_proto_goTypes,
		DependencyIndexes: file_userquerysession_proto_depIdxs,
		MessageInfos:      file_userquerysession_proto_msgTypes,
	}.Build()
	File_userquerysession_proto = out.File
	file_userquerysession_proto_rawDesc = nil
	file_userquerysession_proto_goTypes = nil
	file_userquerysession_proto_depIdxs = nil
}
