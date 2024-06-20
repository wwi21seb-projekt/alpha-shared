// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: server_alpha/mail/mail.proto

package mail

import (
	common "github.com/wwi21seb-projekt/alpha-shared/proto/server_alpha/common"
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

type TokenMailType int32

const (
	TokenMailType_TOKENMAILTYPE_REGISTRATION   TokenMailType = 0
	TokenMailType_TOKENMAILTYPE_PASSWORD_RESET TokenMailType = 1
)

// Enum value maps for TokenMailType.
var (
	TokenMailType_name = map[int32]string{
		0: "TOKENMAILTYPE_REGISTRATION",
		1: "TOKENMAILTYPE_PASSWORD_RESET",
	}
	TokenMailType_value = map[string]int32{
		"TOKENMAILTYPE_REGISTRATION":   0,
		"TOKENMAILTYPE_PASSWORD_RESET": 1,
	}
)

func (x TokenMailType) Enum() *TokenMailType {
	p := new(TokenMailType)
	*p = x
	return p
}

func (x TokenMailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenMailType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_alpha_mail_mail_proto_enumTypes[0].Descriptor()
}

func (TokenMailType) Type() protoreflect.EnumType {
	return &file_server_alpha_mail_mail_proto_enumTypes[0]
}

func (x TokenMailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenMailType.Descriptor instead.
func (TokenMailType) EnumDescriptor() ([]byte, []int) {
	return file_server_alpha_mail_mail_proto_rawDescGZIP(), []int{0}
}

type UserInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInformation) Reset() {
	*x = UserInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInformation) ProtoMessage() {}

func (x *UserInformation) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_mail_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInformation.ProtoReflect.Descriptor instead.
func (*UserInformation) Descriptor() ([]byte, []int) {
	return file_server_alpha_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *UserInformation) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInformation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TokenMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *UserInformation `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token string           `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Type  TokenMailType    `protobuf:"varint,3,opt,name=type,proto3,enum=server_alpha.mail.TokenMailType" json:"type,omitempty"`
}

func (x *TokenMailRequest) Reset() {
	*x = TokenMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_mail_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMailRequest) ProtoMessage() {}

func (x *TokenMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_mail_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMailRequest.ProtoReflect.Descriptor instead.
func (*TokenMailRequest) Descriptor() ([]byte, []int) {
	return file_server_alpha_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *TokenMailRequest) GetUser() *UserInformation {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *TokenMailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TokenMailRequest) GetType() TokenMailType {
	if x != nil {
		return x.Type
	}
	return TokenMailType_TOKENMAILTYPE_REGISTRATION
}

type ConfirmationMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserInformation `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ConfirmationMailRequest) Reset() {
	*x = ConfirmationMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_alpha_mail_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationMailRequest) ProtoMessage() {}

func (x *ConfirmationMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_alpha_mail_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationMailRequest.ProtoReflect.Descriptor instead.
func (*ConfirmationMailRequest) Descriptor() ([]byte, []int) {
	return file_server_alpha_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmationMailRequest) GetUser() *UserInformation {
	if x != nil {
		return x.User
	}
	return nil
}

var File_server_alpha_mail_mail_proto protoreflect.FileDescriptor

var file_server_alpha_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d,
	0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x1a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x51, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x2a, 0x51, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x69,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x4d, 0x41,
	0x49, 0x4c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x4d, 0x41,
	0x49, 0x4c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f,
	0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x01, 0x32, 0xbf, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5e, 0x0a, 0x14, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x69,
	0x6c, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x77, 0x69, 0x32, 0x31, 0x73, 0x65, 0x62,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x6b, 0x74, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2d, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_alpha_mail_mail_proto_rawDescOnce sync.Once
	file_server_alpha_mail_mail_proto_rawDescData = file_server_alpha_mail_mail_proto_rawDesc
)

func file_server_alpha_mail_mail_proto_rawDescGZIP() []byte {
	file_server_alpha_mail_mail_proto_rawDescOnce.Do(func() {
		file_server_alpha_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_alpha_mail_mail_proto_rawDescData)
	})
	return file_server_alpha_mail_mail_proto_rawDescData
}

var file_server_alpha_mail_mail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_alpha_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_alpha_mail_mail_proto_goTypes = []any{
	(TokenMailType)(0),              // 0: server_alpha.mail.TokenMailType
	(*UserInformation)(nil),         // 1: server_alpha.mail.UserInformation
	(*TokenMailRequest)(nil),        // 2: server_alpha.mail.TokenMailRequest
	(*ConfirmationMailRequest)(nil), // 3: server_alpha.mail.ConfirmationMailRequest
	(*common.Empty)(nil),            // 4: server_alpha.common.Empty
}
var file_server_alpha_mail_mail_proto_depIdxs = []int32{
	1, // 0: server_alpha.mail.TokenMailRequest.user:type_name -> server_alpha.mail.UserInformation
	0, // 1: server_alpha.mail.TokenMailRequest.type:type_name -> server_alpha.mail.TokenMailType
	1, // 2: server_alpha.mail.ConfirmationMailRequest.user:type_name -> server_alpha.mail.UserInformation
	2, // 3: server_alpha.mail.MailService.SendTokenMail:input_type -> server_alpha.mail.TokenMailRequest
	3, // 4: server_alpha.mail.MailService.SendConfirmationMail:input_type -> server_alpha.mail.ConfirmationMailRequest
	4, // 5: server_alpha.mail.MailService.SendTokenMail:output_type -> server_alpha.common.Empty
	4, // 6: server_alpha.mail.MailService.SendConfirmationMail:output_type -> server_alpha.common.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_alpha_mail_mail_proto_init() }
func file_server_alpha_mail_mail_proto_init() {
	if File_server_alpha_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_alpha_mail_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserInformation); i {
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
		file_server_alpha_mail_mail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TokenMailRequest); i {
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
		file_server_alpha_mail_mail_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ConfirmationMailRequest); i {
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
			RawDescriptor: file_server_alpha_mail_mail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_alpha_mail_mail_proto_goTypes,
		DependencyIndexes: file_server_alpha_mail_mail_proto_depIdxs,
		EnumInfos:         file_server_alpha_mail_mail_proto_enumTypes,
		MessageInfos:      file_server_alpha_mail_mail_proto_msgTypes,
	}.Build()
	File_server_alpha_mail_mail_proto = out.File
	file_server_alpha_mail_mail_proto_rawDesc = nil
	file_server_alpha_mail_mail_proto_goTypes = nil
	file_server_alpha_mail_mail_proto_depIdxs = nil
}
