// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: image/image.proto

package image

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base64Image string `protobuf:"bytes,1,opt,name=base64image,proto3" json:"base64image,omitempty"`
	ImageType   string `protobuf:"bytes,2,opt,name=imageType,proto3" json:"imageType,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetBase64Image() string {
	if x != nil {
		return x.Base64Image
	}
	return ""
}

func (x *Image) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image         string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`                                      // base64 encoded image
	ContextString string `protobuf:"bytes,2,opt,name=context_string,json=contextString,proto3" json:"context_string,omitempty"` // postId or username for file naming
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UploadImageRequest) GetContextString() string {
	if x != nil {
		return x.ContextString
	}
	return ""
}

type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName string `protobuf:"bytes,1,opt,name=imageName,proto3" json:"imageName,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetImageRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

type ListImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageNames []string `protobuf:"bytes,1,rep,name=imageNames,proto3" json:"imageNames,omitempty"`
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListImagesRequest) GetImageNames() []string {
	if x != nil {
		return x.ImageNames
	}
	return nil
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_image_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_image_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_image_image_service_proto_rawDescGZIP(), []int{4}
}

func (x *UploadImageResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadImageResponse) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *UploadImageResponse) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_image_image_service_proto protoreflect.FileDescriptor

var file_image_image_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x47,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x36,
	0x34, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x36, 0x34, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x55, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xb6, 0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0xc3, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x11, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x77, 0x69, 0x32,
	0x31, 0x73, 0x65, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x6b, 0x74, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x49, 0x58, 0xaa, 0x02, 0x11, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0xca, 0x02, 0x11,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0xe2, 0x02, 0x1d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a,
	0x3a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_image_service_proto_rawDescOnce sync.Once
	file_image_image_service_proto_rawDescData = file_image_image_service_proto_rawDesc
)

func file_image_image_service_proto_rawDescGZIP() []byte {
	file_image_image_service_proto_rawDescOnce.Do(func() {
		file_image_image_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_image_service_proto_rawDescData)
	})
	return file_image_image_service_proto_rawDescData
}

var file_image_image_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_image_image_service_proto_goTypes = []any{
	(*Image)(nil),               // 0: serveralpha.image.Image
	(*UploadImageRequest)(nil),  // 1: serveralpha.image.UploadImageRequest
	(*GetImageRequest)(nil),     // 2: serveralpha.image.GetImageRequest
	(*ListImagesRequest)(nil),   // 3: serveralpha.image.ListImagesRequest
	(*UploadImageResponse)(nil), // 4: serveralpha.image.UploadImageResponse
}
var file_image_image_service_proto_depIdxs = []int32{
	1, // 0: serveralpha.image.ImageService.UploadImage:input_type -> serveralpha.image.UploadImageRequest
	2, // 1: serveralpha.image.ImageService.GetImage:input_type -> serveralpha.image.GetImageRequest
	4, // 2: serveralpha.image.ImageService.UploadImage:output_type -> serveralpha.image.UploadImageResponse
	0, // 3: serveralpha.image.ImageService.GetImage:output_type -> serveralpha.image.Image
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_image_image_service_proto_init() }
func file_image_image_service_proto_init() {
	if File_image_image_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_image_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Image); i {
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
		file_image_image_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UploadImageRequest); i {
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
		file_image_image_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetImageRequest); i {
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
		file_image_image_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListImagesRequest); i {
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
		file_image_image_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UploadImageResponse); i {
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
			RawDescriptor: file_image_image_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_image_service_proto_goTypes,
		DependencyIndexes: file_image_image_service_proto_depIdxs,
		MessageInfos:      file_image_image_service_proto_msgTypes,
	}.Build()
	File_image_image_service_proto = out.File
	file_image_image_service_proto_rawDesc = nil
	file_image_image_service_proto_goTypes = nil
	file_image_image_service_proto_depIdxs = nil
}