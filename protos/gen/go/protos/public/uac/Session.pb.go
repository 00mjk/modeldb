// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: uac/Session.proto

package uac

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VertaUserId string `protobuf:"bytes,2,opt,name=verta_user_id,json=vertaUserId,proto3" json:"verta_user_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Time after which the session is invalid, in seconds since epoch in GMT
	TtlEpoch uint64 `protobuf:"varint,4,opt,name=ttl_epoch,json=ttlEpoch,proto3" json:"ttl_epoch,omitempty"`
	// Equivalent of a developer key
	SessionSecretKey string `protobuf:"bytes,5,opt,name=session_secret_key,json=sessionSecretKey,proto3" json:"session_secret_key,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_Session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_uac_Session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_uac_Session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Session) GetVertaUserId() string {
	if x != nil {
		return x.VertaUserId
	}
	return ""
}

func (x *Session) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Session) GetTtlEpoch() uint64 {
	if x != nil {
		return x.TtlEpoch
	}
	return 0
}

func (x *Session) GetSessionSecretKey() string {
	if x != nil {
		return x.SessionSecretKey
	}
	return ""
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VertaUserId string `protobuf:"bytes,1,opt,name=verta_user_id,json=vertaUserId,proto3" json:"verta_user_id,omitempty"`
	// Session names must be unique. If we create another session with the same name, the original one will be invalidated.
	SessionName string `protobuf:"bytes,2,opt,name=session_name,json=sessionName,proto3" json:"session_name,omitempty"`
	// Number of seconds to keep this session active from the moment of creation.
	TtlSeconds uint64 `protobuf:"varint,3,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_Session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uac_Session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_uac_Session_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionRequest) GetVertaUserId() string {
	if x != nil {
		return x.VertaUserId
	}
	return ""
}

func (x *CreateSessionRequest) GetSessionName() string {
	if x != nil {
		return x.SessionName
	}
	return ""
}

func (x *CreateSessionRequest) GetTtlSeconds() uint64 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

type FindSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids         []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	VertaUserId []string `protobuf:"bytes,2,rep,name=verta_user_id,json=vertaUserId,proto3" json:"verta_user_id,omitempty"`
	Name        []string `protobuf:"bytes,3,rep,name=name,proto3" json:"name,omitempty"`
}

func (x *FindSessionRequest) Reset() {
	*x = FindSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_Session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSessionRequest) ProtoMessage() {}

func (x *FindSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uac_Session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSessionRequest.ProtoReflect.Descriptor instead.
func (*FindSessionRequest) Descriptor() ([]byte, []int) {
	return file_uac_Session_proto_rawDescGZIP(), []int{2}
}

func (x *FindSessionRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FindSessionRequest) GetVertaUserId() []string {
	if x != nil {
		return x.VertaUserId
	}
	return nil
}

func (x *FindSessionRequest) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

type DeleteSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteSessionRequest) Reset() {
	*x = DeleteSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_Session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionRequest) ProtoMessage() {}

func (x *DeleteSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uac_Session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return file_uac_Session_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSessionRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FindSessionRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *FindSessionRequest_Response) Reset() {
	*x = FindSessionRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uac_Session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSessionRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSessionRequest_Response) ProtoMessage() {}

func (x *FindSessionRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_uac_Session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSessionRequest_Response.ProtoReflect.Descriptor instead.
func (*FindSessionRequest_Response) Descriptor() ([]byte, []int) {
	return file_uac_Session_proto_rawDescGZIP(), []int{2, 0}
}

func (x *FindSessionRequest_Response) GetSessions() []*Session {
	if x != nil {
		return x.Sessions
	}
	return nil
}

var File_uac_Session_proto protoreflect.FileDescriptor

var file_uac_Session_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x61, 0x63, 0x2f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61,
	0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x75, 0x61, 0x63, 0x2f, 0x55, 0x41, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x74, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x74, 0x61, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x74, 0x6c,
	0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x74,
	0x6c, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x22, 0x7e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x76, 0x65, 0x72, 0x74, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x22, 0x0a,
	0x0d, 0x76, 0x65, 0x72, 0x74, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x74, 0x61, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75,
	0x61, 0x63, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0xf2,
	0x02, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x70, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74,
	0x61, 0x2e, 0x75, 0x61, 0x63, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e, 0x75, 0x61,
	0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e,
	0x75, 0x61, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x74, 0x61, 0x2e,
	0x75, 0x61, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x65,
	0x72, 0x74, 0x61, 0x2e, 0x75, 0x61, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x3e, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x74, 0x61, 0x41, 0x49, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x75, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uac_Session_proto_rawDescOnce sync.Once
	file_uac_Session_proto_rawDescData = file_uac_Session_proto_rawDesc
)

func file_uac_Session_proto_rawDescGZIP() []byte {
	file_uac_Session_proto_rawDescOnce.Do(func() {
		file_uac_Session_proto_rawDescData = protoimpl.X.CompressGZIP(file_uac_Session_proto_rawDescData)
	})
	return file_uac_Session_proto_rawDescData
}

var file_uac_Session_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_uac_Session_proto_goTypes = []interface{}{
	(*Session)(nil),                     // 0: ai.verta.uac.Session
	(*CreateSessionRequest)(nil),        // 1: ai.verta.uac.CreateSessionRequest
	(*FindSessionRequest)(nil),          // 2: ai.verta.uac.FindSessionRequest
	(*DeleteSessionRequest)(nil),        // 3: ai.verta.uac.DeleteSessionRequest
	(*FindSessionRequest_Response)(nil), // 4: ai.verta.uac.FindSessionRequest.Response
	(*Empty)(nil),                       // 5: ai.verta.uac.Empty
}
var file_uac_Session_proto_depIdxs = []int32{
	0, // 0: ai.verta.uac.FindSessionRequest.Response.sessions:type_name -> ai.verta.uac.Session
	1, // 1: ai.verta.uac.SessionService.createSession:input_type -> ai.verta.uac.CreateSessionRequest
	2, // 2: ai.verta.uac.SessionService.findSession:input_type -> ai.verta.uac.FindSessionRequest
	3, // 3: ai.verta.uac.SessionService.deleteSession:input_type -> ai.verta.uac.DeleteSessionRequest
	0, // 4: ai.verta.uac.SessionService.createSession:output_type -> ai.verta.uac.Session
	4, // 5: ai.verta.uac.SessionService.findSession:output_type -> ai.verta.uac.FindSessionRequest.Response
	5, // 6: ai.verta.uac.SessionService.deleteSession:output_type -> ai.verta.uac.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uac_Session_proto_init() }
func file_uac_Session_proto_init() {
	if File_uac_Session_proto != nil {
		return
	}
	file_uac_UACService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_uac_Session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_uac_Session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_uac_Session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSessionRequest); i {
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
		file_uac_Session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSessionRequest); i {
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
		file_uac_Session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSessionRequest_Response); i {
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
			RawDescriptor: file_uac_Session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uac_Session_proto_goTypes,
		DependencyIndexes: file_uac_Session_proto_depIdxs,
		MessageInfos:      file_uac_Session_proto_msgTypes,
	}.Build()
	File_uac_Session_proto = out.File
	file_uac_Session_proto_rawDesc = nil
	file_uac_Session_proto_goTypes = nil
	file_uac_Session_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error)
	FindSession(ctx context.Context, in *FindSessionRequest, opts ...grpc.CallOption) (*FindSessionRequest_Response, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.SessionService/createSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) FindSession(ctx context.Context, in *FindSessionRequest, opts ...grpc.CallOption) (*FindSessionRequest_Response, error) {
	out := new(FindSessionRequest_Response)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.SessionService/findSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ai.verta.uac.SessionService/deleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	CreateSession(context.Context, *CreateSessionRequest) (*Session, error)
	FindSession(context.Context, *FindSessionRequest) (*FindSessionRequest_Response, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*Empty, error)
}

// UnimplementedSessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (*UnimplementedSessionServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedSessionServiceServer) FindSession(context.Context, *FindSessionRequest) (*FindSessionRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSession not implemented")
}
func (*UnimplementedSessionServiceServer) DeleteSession(context.Context, *DeleteSessionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_FindSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).FindSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.SessionService/FindSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).FindSession(ctx, req.(*FindSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.verta.uac.SessionService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.verta.uac.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "findSession",
			Handler:    _SessionService_FindSession_Handler,
		},
		{
			MethodName: "deleteSession",
			Handler:    _SessionService_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uac/Session.proto",
}