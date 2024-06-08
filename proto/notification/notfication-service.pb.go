// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: notification/notfication-service.proto

package serveralpha

import (
	common "github.com/wwi21seb-projekt/alpha-shared/proto/common"
	user "github.com/wwi21seb-projekt/alpha-shared/proto/user"
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

type PushSubscriptionType int32

const (
	PushSubscriptionType_WEB  PushSubscriptionType = 0
	PushSubscriptionType_EXPO PushSubscriptionType = 1
)

// Enum value maps for PushSubscriptionType.
var (
	PushSubscriptionType_name = map[int32]string{
		0: "WEB",
		1: "EXPO",
	}
	PushSubscriptionType_value = map[string]int32{
		"WEB":  0,
		"EXPO": 1,
	}
)

func (x PushSubscriptionType) Enum() *PushSubscriptionType {
	p := new(PushSubscriptionType)
	*p = x
	return p
}

func (x PushSubscriptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushSubscriptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_notfication_service_proto_enumTypes[0].Descriptor()
}

func (PushSubscriptionType) Type() protoreflect.EnumType {
	return &file_notification_notfication_service_proto_enumTypes[0]
}

func (x PushSubscriptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushSubscriptionType.Descriptor instead.
func (PushSubscriptionType) EnumDescriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{0}
}

type NotificationType int32

const (
	NotificationType_FOLLOW NotificationType = 0
	NotificationType_REPOST NotificationType = 1
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "FOLLOW",
		1: "REPOST",
	}
	NotificationType_value = map[string]int32{
		"FOLLOW": 0,
		"REPOST": 1,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_notfication_service_proto_enumTypes[1].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_notification_notfication_service_proto_enumTypes[1]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{1}
}

type PublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *PublicKeyResponse) Reset() {
	*x = PublicKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKeyResponse) ProtoMessage() {}

func (x *PublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKeyResponse.ProtoReflect.Descriptor instead.
func (*PublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKeyResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type CreatePushSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint       string               `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	P256Dh         string               `protobuf:"bytes,2,opt,name=p256dh,proto3" json:"p256dh,omitempty"`
	Auth           string               `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Type           PushSubscriptionType `protobuf:"varint,4,opt,name=type,proto3,enum=serveralpha.notification.PushSubscriptionType" json:"type,omitempty"`
	Token          string               `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	ExpirationTime string               `protobuf:"bytes,6,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
}

func (x *CreatePushSubscriptionRequest) Reset() {
	*x = CreatePushSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePushSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePushSubscriptionRequest) ProtoMessage() {}

func (x *CreatePushSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePushSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreatePushSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePushSubscriptionRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *CreatePushSubscriptionRequest) GetP256Dh() string {
	if x != nil {
		return x.P256Dh
	}
	return ""
}

func (x *CreatePushSubscriptionRequest) GetAuth() string {
	if x != nil {
		return x.Auth
	}
	return ""
}

func (x *CreatePushSubscriptionRequest) GetType() PushSubscriptionType {
	if x != nil {
		return x.Type
	}
	return PushSubscriptionType_WEB
}

func (x *CreatePushSubscriptionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreatePushSubscriptionRequest) GetExpirationTime() string {
	if x != nil {
		return x.ExpirationTime
	}
	return ""
}

type CreatePushSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *CreatePushSubscriptionResponse) Reset() {
	*x = CreatePushSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePushSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePushSubscriptionResponse) ProtoMessage() {}

func (x *CreatePushSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePushSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CreatePushSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePushSubscriptionResponse) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}


type GetNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId  string           `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	Timestamp       string           `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NotficationType NotificationType `protobuf:"varint,3,opt,name=notfication_type,json=notficationType,proto3,enum=serveralpha.notification.NotificationType" json:"notfication_type,omitempty"`
	User            *user.User       `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetNotificationsResponse) Reset() {
	*x = GetNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsResponse) ProtoMessage() {}

func (x *GetNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}


// Deprecated: Use GetNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetNotificationsResponse) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}


func (x *GetNotificationsResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}


func (x *GetNotificationsResponse) GetNotficationType() NotificationType {
	if x != nil {
		return x.NotficationType
	}
	return NotificationType_FOLLOW
}


func (x *GetNotificationsResponse) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields


	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
}

func (x *DeleteNotificationRequest) Reset() {
	*x = DeleteNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}


func (x *DeleteNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotificationRequest) ProtoMessage() {}

func (x *DeleteNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}


// Deprecated: Use DeleteNotificationRequest.ProtoReflect.Descriptor instead.
func (*DeleteNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type SendNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields


	NotificationType string `protobuf:"bytes,1,opt,name=NotificationType,proto3" json:"NotificationType,omitempty"`
	Sender           string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *SendNotificationRequest) Reset() {
	*x = SendNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_notfication_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationRequest) ProtoMessage() {}

func (x *SendNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_notfication_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationRequest) Descriptor() ([]byte, []int) {
	return file_notification_notfication_service_proto_rawDescGZIP(), []int{5}
}

func (x *SendNotificationRequest) GetNotificationType() string {
	if x != nil {
		return x.NotificationType
	}
	return ""
}

func (x *SendNotificationRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

var File_notification_notfication_service_proto protoreflect.FileDescriptor

var file_notification_notfication_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e,
	0x6f, 0x74, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x22, 0xe9, 0x01, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x32, 0x35, 0x36, 0x64, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x32, 0x35, 0x36, 0x64, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x42,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x49, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xe4, 0x01, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x55, 0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x44, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2a, 0x29, 0x0a, 0x14, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x58, 0x50, 0x4f,
	0x10, 0x01, 0x2a, 0x2a, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x32, 0xf3,
	0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x73, 0x68, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x73, 0x68,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc0, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x64, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x77, 0x69, 0x32, 0x31, 0x73, 0x65, 0x62, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x6b, 0x74, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_notfication_service_proto_rawDescOnce sync.Once
	file_notification_notfication_service_proto_rawDescData = file_notification_notfication_service_proto_rawDesc
)

func file_notification_notfication_service_proto_rawDescGZIP() []byte {
	file_notification_notfication_service_proto_rawDescOnce.Do(func() {
		file_notification_notfication_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_notfication_service_proto_rawDescData)
	})
	return file_notification_notfication_service_proto_rawDescData
}

var file_notification_notfication_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_notification_notfication_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_notification_notfication_service_proto_goTypes = []interface{}{
	(PushSubscriptionType)(0),              // 0: serveralpha.notification.PushSubscriptionType
	(NotificationType)(0),                  // 1: serveralpha.notification.NotificationType
	(*PublicKeyResponse)(nil),              // 2: serveralpha.notification.PublicKeyResponse
	(*CreatePushSubscriptionRequest)(nil),  // 3: serveralpha.notification.CreatePushSubscriptionRequest
	(*CreatePushSubscriptionResponse)(nil), // 4: serveralpha.notification.CreatePushSubscriptionResponse
	(*GetNotificationsResponse)(nil),       // 5: serveralpha.notification.GetNotificationsResponse
	(*DeleteNotificationRequest)(nil),      // 6: serveralpha.notification.DeleteNotificationRequest
	(*SendNotificationRequest)(nil),        // 7: serveralpha.notification.SendNotificationRequest
	(*user.User)(nil),                      // 8: serveralpha.user.User
	(*common.Empty)(nil),                   // 9: serveralpha.common.Empty
}
var file_notification_notfication_service_proto_depIdxs = []int32{
	0, // 0: serveralpha.notification.CreatePushSubscriptionRequest.type:type_name -> serveralpha.notification.PushSubscriptionType
	1, // 1: serveralpha.notification.GetNotificationsResponse.notfication_type:type_name -> serveralpha.notification.NotificationType
	8, // 2: serveralpha.notification.GetNotificationsResponse.user:type_name -> serveralpha.user.User
	9, // 3: serveralpha.notification.PushService.GetPublicKey:input_type -> serveralpha.common.Empty
	3, // 4: serveralpha.notification.PushService.CreatePushSubscription:input_type -> serveralpha.notification.CreatePushSubscriptionRequest
	9, // 5: serveralpha.notification.NotificationService.GetNotifications:input_type -> serveralpha.common.Empty
	6, // 6: serveralpha.notification.NotificationService.DeleteNotification:input_type -> serveralpha.notification.DeleteNotificationRequest
	7, // 7: serveralpha.notification.NotificationService.SendNotification:input_type -> serveralpha.notification.SendNotificationRequest
	2, // 8: serveralpha.notification.PushService.GetPublicKey:output_type -> serveralpha.notification.PublicKeyResponse
	4, // 9: serveralpha.notification.PushService.CreatePushSubscription:output_type -> serveralpha.notification.CreatePushSubscriptionResponse
	5, // 10: serveralpha.notification.NotificationService.GetNotifications:output_type -> serveralpha.notification.GetNotificationsResponse
	9, // 11: serveralpha.notification.NotificationService.DeleteNotification:output_type -> serveralpha.common.Empty
	9, // 12: serveralpha.notification.NotificationService.SendNotification:output_type -> serveralpha.common.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notification_notfication_service_proto_init() }
func file_notification_notfication_service_proto_init() {
	if File_notification_notfication_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_notfication_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKeyResponse); i {
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
		file_notification_notfication_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePushSubscriptionRequest); i {
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
		file_notification_notfication_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePushSubscriptionResponse); i {
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
		file_notification_notfication_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationsResponse); i {
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
		file_notification_notfication_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {

			switch v := v.(*DeleteNotificationRequest); i {
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
		file_notification_notfication_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {

			switch v := v.(*SendNotificationRequest); i {
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
			RawDescriptor: file_notification_notfication_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_notification_notfication_service_proto_goTypes,
		DependencyIndexes: file_notification_notfication_service_proto_depIdxs,
		EnumInfos:         file_notification_notfication_service_proto_enumTypes,
		MessageInfos:      file_notification_notfication_service_proto_msgTypes,
	}.Build()
	File_notification_notfication_service_proto = out.File
	file_notification_notfication_service_proto_rawDesc = nil
	file_notification_notfication_service_proto_goTypes = nil
	file_notification_notfication_service_proto_depIdxs = nil
}
