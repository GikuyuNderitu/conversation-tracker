// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: service/v1/list_conversations.proto

package servicev1

import (
	v1 "atypicaldev.com/conversation/api/conversation/v1"
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

type ListConversationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListConversationsRequest) Reset() {
	*x = ListConversationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_list_conversations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConversationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConversationsRequest) ProtoMessage() {}

func (x *ListConversationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_list_conversations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConversationsRequest.ProtoReflect.Descriptor instead.
func (*ListConversationsRequest) Descriptor() ([]byte, []int) {
	return file_service_v1_list_conversations_proto_rawDescGZIP(), []int{0}
}

type ListConversationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conversations []*v1.Conversation `protobuf:"bytes,1,rep,name=conversations,proto3" json:"conversations,omitempty"`
}

func (x *ListConversationsResponse) Reset() {
	*x = ListConversationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_list_conversations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConversationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConversationsResponse) ProtoMessage() {}

func (x *ListConversationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_list_conversations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConversationsResponse.ProtoReflect.Descriptor instead.
func (*ListConversationsResponse) Descriptor() ([]byte, []int) {
	return file_service_v1_list_conversations_proto_rawDescGZIP(), []int{1}
}

func (x *ListConversationsResponse) GetConversations() []*v1.Conversation {
	if x != nil {
		return x.Conversations
	}
	return nil
}

var File_service_v1_list_conversations_proto protoreflect.FileDescriptor

var file_service_v1_list_conversations_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x23, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x61, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x61, 0x74, 0x79, 0x70, 0x69, 0x63, 0x61,
	0x6c, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_v1_list_conversations_proto_rawDescOnce sync.Once
	file_service_v1_list_conversations_proto_rawDescData = file_service_v1_list_conversations_proto_rawDesc
)

func file_service_v1_list_conversations_proto_rawDescGZIP() []byte {
	file_service_v1_list_conversations_proto_rawDescOnce.Do(func() {
		file_service_v1_list_conversations_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_v1_list_conversations_proto_rawDescData)
	})
	return file_service_v1_list_conversations_proto_rawDescData
}

var file_service_v1_list_conversations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_v1_list_conversations_proto_goTypes = []interface{}{
	(*ListConversationsRequest)(nil),  // 0: service.v1.ListConversationsRequest
	(*ListConversationsResponse)(nil), // 1: service.v1.ListConversationsResponse
	(*v1.Conversation)(nil),           // 2: conversations.v1.Conversation
}
var file_service_v1_list_conversations_proto_depIdxs = []int32{
	2, // 0: service.v1.ListConversationsResponse.conversations:type_name -> conversations.v1.Conversation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_v1_list_conversations_proto_init() }
func file_service_v1_list_conversations_proto_init() {
	if File_service_v1_list_conversations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_v1_list_conversations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConversationsRequest); i {
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
		file_service_v1_list_conversations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConversationsResponse); i {
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
			RawDescriptor: file_service_v1_list_conversations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_v1_list_conversations_proto_goTypes,
		DependencyIndexes: file_service_v1_list_conversations_proto_depIdxs,
		MessageInfos:      file_service_v1_list_conversations_proto_msgTypes,
	}.Build()
	File_service_v1_list_conversations_proto = out.File
	file_service_v1_list_conversations_proto_rawDesc = nil
	file_service_v1_list_conversations_proto_goTypes = nil
	file_service_v1_list_conversations_proto_depIdxs = nil
}
