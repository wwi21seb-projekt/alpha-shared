// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: server_alpha/user/v1/subscription.proto

package userv1

import (
	v11 "github.com/wwi21seb-projekt/alpha-shared/gen/server_alpha/common/v1"
	v1 "github.com/wwi21seb-projekt/alpha-shared/gen/server_alpha/image/v1"
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

type SubscriptionType int32

const (
	SubscriptionType_SUBSCRIPTION_TYPE_UNSPECIFIED SubscriptionType = 0 // default by setting to 0
	SubscriptionType_SUBSCRIPTION_TYPE_FOLLOWER    SubscriptionType = 1
	SubscriptionType_SUBSCRIPTION_TYPE_FOLLOWING   SubscriptionType = 2
)

// Enum value maps for SubscriptionType.
var (
	SubscriptionType_name = map[int32]string{
		0: "SUBSCRIPTION_TYPE_UNSPECIFIED",
		1: "SUBSCRIPTION_TYPE_FOLLOWER",
		2: "SUBSCRIPTION_TYPE_FOLLOWING",
	}
	SubscriptionType_value = map[string]int32{
		"SUBSCRIPTION_TYPE_UNSPECIFIED": 0,
		"SUBSCRIPTION_TYPE_FOLLOWER":    1,
		"SUBSCRIPTION_TYPE_FOLLOWING":   2,
	}
)

func (x SubscriptionType) Enum() *SubscriptionType {
	p := new(SubscriptionType)
	*p = x
	return p
}

func (x SubscriptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_alpha_user_v1_subscription_proto_enumTypes[0].Descriptor()
}

func (SubscriptionType) Type() protoreflect.EnumType {
	return &file_server_alpha_user_v1_subscription_proto_enumTypes[0]
}

func (x SubscriptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriptionType.Descriptor instead.
func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerSubscriptionId string      `protobuf:"bytes,1,opt,name=follower_subscription_id,json=followerSubscriptionId,proto3" json:"follower_subscription_id,omitempty"` // populated when the other user follows the current user
	FollowedSubscriptionId string      `protobuf:"bytes,2,opt,name=followed_subscription_id,json=followedSubscriptionId,proto3" json:"followed_subscription_id,omitempty"` // populated when the current user follows the other user
	Username               string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`                                                             // the username of the other user
	Nickname               string      `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`                                                             // the nickname of the other user
	Picture                *v1.Picture `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`                                                               // the picture of the other user
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetFollowerSubscriptionId() string {
	if x != nil {
		return x.FollowerSubscriptionId
	}
	return ""
}

func (x *Subscription) GetFollowedSubscriptionId() string {
	if x != nil {
		return x.FollowedSubscriptionId
	}
	return ""
}

func (x *Subscription) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Subscription) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Subscription) GetPicture() *v1.Picture {
	if x != nil {
		return x.Picture
	}
	return nil
}

type ListSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination       *v11.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	SubscriptionType SubscriptionType       `protobuf:"varint,2,opt,name=subscription_type,json=subscriptionType,proto3,enum=server_alpha.user.v1.SubscriptionType" json:"subscription_type,omitempty"`
	Username         string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ListSubscriptionsRequest) Reset() {
	*x = ListSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsRequest) ProtoMessage() {}

func (x *ListSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubscriptionsRequest) GetPagination() *v11.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListSubscriptionsRequest) GetSubscriptionType() SubscriptionType {
	if x != nil {
		return x.SubscriptionType
	}
	return SubscriptionType_SUBSCRIPTION_TYPE_UNSPECIFIED
}

func (x *ListSubscriptionsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowedUsername string `protobuf:"bytes,1,opt,name=followed_username,json=followedUsername,proto3" json:"followed_username,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSubscriptionRequest) GetFollowedUsername() string {
	if x != nil {
		return x.FollowedUsername
	}
	return ""
}

type DeleteSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *DeleteSubscriptionRequest) Reset() {
	*x = DeleteSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriptionRequest) ProtoMessage() {}

func (x *DeleteSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSubscriptionRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type ListSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription         `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Pagination    *v11.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListSubscriptionsResponse) Reset() {
	*x = ListSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsResponse) ProtoMessage() {}

func (x *ListSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (x *ListSubscriptionsResponse) GetPagination() *v11.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId   string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	SubscriptionDate string `protobuf:"bytes,2,opt,name=subscription_date,json=subscriptionDate,proto3" json:"subscription_date,omitempty"`
	FollowerUsername string `protobuf:"bytes,3,opt,name=follower_username,json=followerUsername,proto3" json:"follower_username,omitempty"`
	FollowedUsername string `protobuf:"bytes,4,opt,name=followed_username,json=followedUsername,proto3" json:"followed_username,omitempty"`
}

func (x *CreateSubscriptionResponse) Reset() {
	*x = CreateSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionResponse) ProtoMessage() {}

func (x *CreateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSubscriptionResponse) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *CreateSubscriptionResponse) GetSubscriptionDate() string {
	if x != nil {
		return x.SubscriptionDate
	}
	return ""
}

func (x *CreateSubscriptionResponse) GetFollowerUsername() string {
	if x != nil {
		return x.FollowerUsername
	}
	return ""
}

func (x *CreateSubscriptionResponse) GetFollowedUsername() string {
	if x != nil {
		return x.FollowedUsername
	}
	return ""
}

type DeleteSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteSubscriptionResponse) Reset() {
	*x = DeleteSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriptionResponse) ProtoMessage() {}

func (x *DeleteSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_user_v1_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*DeleteSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_server_alpha_user_v1_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSubscriptionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_server_alpha_user_v1_subscription_proto protoreflect.FileDescriptor

var file_server_alpha_user_v1_subscription_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x23, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0xd6,
	0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x44, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x01, 0x0a, 0x1a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x1a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x76, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4f, 0x4c,
	0x4c, 0x4f, 0x57, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0xfd, 0x02, 0x0a, 0x13, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x74, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x77, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe5, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x77, 0x69, 0x32, 0x31, 0x73, 0x65, 0x62, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x6b, 0x74, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41,
	0x6c, 0x70, 0x68, 0x61, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_alpha_user_v1_subscription_proto_rawDescOnce sync.Once
	file_server_alpha_user_v1_subscription_proto_rawDescData = file_server_alpha_user_v1_subscription_proto_rawDesc
)

func file_server_alpha_user_v1_subscription_proto_rawDescGZIP() []byte {
	file_server_alpha_user_v1_subscription_proto_rawDescOnce.Do(func() {
		file_server_alpha_user_v1_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_alpha_user_v1_subscription_proto_rawDescData)
	})
	return file_server_alpha_user_v1_subscription_proto_rawDescData
}

var file_server_alpha_user_v1_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_alpha_user_v1_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_server_alpha_user_v1_subscription_proto_goTypes = []any{
	(SubscriptionType)(0),              // 0: server_alpha.user.v1.SubscriptionType
	(*Subscription)(nil),               // 1: server_alpha.user.v1.Subscription
	(*ListSubscriptionsRequest)(nil),   // 2: server_alpha.user.v1.ListSubscriptionsRequest
	(*CreateSubscriptionRequest)(nil),  // 3: server_alpha.user.v1.CreateSubscriptionRequest
	(*DeleteSubscriptionRequest)(nil),  // 4: server_alpha.user.v1.DeleteSubscriptionRequest
	(*ListSubscriptionsResponse)(nil),  // 5: server_alpha.user.v1.ListSubscriptionsResponse
	(*CreateSubscriptionResponse)(nil), // 6: server_alpha.user.v1.CreateSubscriptionResponse
	(*DeleteSubscriptionResponse)(nil), // 7: server_alpha.user.v1.DeleteSubscriptionResponse
	(*v1.Picture)(nil),                 // 8: server_alpha.image.v1.Picture
	(*v11.PaginationRequest)(nil),      // 9: server_alpha.common.v1.PaginationRequest
	(*v11.PaginationResponse)(nil),     // 10: server_alpha.common.v1.PaginationResponse
}
var file_server_alpha_user_v1_subscription_proto_depIdxs = []int32{
	8,  // 0: server_alpha.user.v1.Subscription.picture:type_name -> server_alpha.image.v1.Picture
	9,  // 1: server_alpha.user.v1.ListSubscriptionsRequest.pagination:type_name -> server_alpha.common.v1.PaginationRequest
	0,  // 2: server_alpha.user.v1.ListSubscriptionsRequest.subscription_type:type_name -> server_alpha.user.v1.SubscriptionType
	1,  // 3: server_alpha.user.v1.ListSubscriptionsResponse.subscriptions:type_name -> server_alpha.user.v1.Subscription
	10, // 4: server_alpha.user.v1.ListSubscriptionsResponse.pagination:type_name -> server_alpha.common.v1.PaginationResponse
	2,  // 5: server_alpha.user.v1.SubscriptionService.ListSubscriptions:input_type -> server_alpha.user.v1.ListSubscriptionsRequest
	3,  // 6: server_alpha.user.v1.SubscriptionService.CreateSubscription:input_type -> server_alpha.user.v1.CreateSubscriptionRequest
	4,  // 7: server_alpha.user.v1.SubscriptionService.DeleteSubscription:input_type -> server_alpha.user.v1.DeleteSubscriptionRequest
	5,  // 8: server_alpha.user.v1.SubscriptionService.ListSubscriptions:output_type -> server_alpha.user.v1.ListSubscriptionsResponse
	6,  // 9: server_alpha.user.v1.SubscriptionService.CreateSubscription:output_type -> server_alpha.user.v1.CreateSubscriptionResponse
	7,  // 10: server_alpha.user.v1.SubscriptionService.DeleteSubscription:output_type -> server_alpha.user.v1.DeleteSubscriptionResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_server_alpha_user_v1_subscription_proto_init() }
func file_server_alpha_user_v1_subscription_proto_init() {
	if File_server_alpha_user_v1_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_alpha_user_v1_subscription_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Subscription); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubscriptionsRequest); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubscriptionRequest); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubscriptionRequest); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubscriptionsResponse); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubscriptionResponse); i {
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
		file_server_alpha_user_v1_subscription_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubscriptionResponse); i {
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
			RawDescriptor: file_server_alpha_user_v1_subscription_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_alpha_user_v1_subscription_proto_goTypes,
		DependencyIndexes: file_server_alpha_user_v1_subscription_proto_depIdxs,
		EnumInfos:         file_server_alpha_user_v1_subscription_proto_enumTypes,
		MessageInfos:      file_server_alpha_user_v1_subscription_proto_msgTypes,
	}.Build()
	File_server_alpha_user_v1_subscription_proto = out.File
	file_server_alpha_user_v1_subscription_proto_rawDesc = nil
	file_server_alpha_user_v1_subscription_proto_goTypes = nil
	file_server_alpha_user_v1_subscription_proto_depIdxs = nil
}
