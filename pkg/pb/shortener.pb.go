// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: proto/shortener.proto

package pb

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

type GetOriginalURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *GetOriginalURLRequest) Reset() {
	*x = GetOriginalURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLRequest) ProtoMessage() {}

func (x *GetOriginalURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetOriginalURLRequest) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetOriginalURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
}

func (x *GetOriginalURLResponse) Reset() {
	*x = GetOriginalURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginalURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLResponse) ProtoMessage() {}

func (x *GetOriginalURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLResponse.ProtoReflect.Descriptor instead.
func (*GetOriginalURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetOriginalURLResponse) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortURLRequest) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

type CreateShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURL string `protobuf:"bytes,1,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *CreateShortURLResponse) Reset() {
	*x = CreateShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponse) ProtoMessage() {}

func (x *CreateShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponse.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponse) Descriptor() ([]byte, []int) {
	return file_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *CreateShortURLResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

var File_proto_shortener_proto protoreflect.FileDescriptor

var file_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22,
	0x39, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x34, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x32, 0xad, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_shortener_proto_rawDescOnce sync.Once
	file_proto_shortener_proto_rawDescData = file_proto_shortener_proto_rawDesc
)

func file_proto_shortener_proto_rawDescGZIP() []byte {
	file_proto_shortener_proto_rawDescOnce.Do(func() {
		file_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortener_proto_rawDescData)
	})
	return file_proto_shortener_proto_rawDescData
}

var file_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shortener_proto_goTypes = []interface{}{
	(*GetOriginalURLRequest)(nil),  // 0: proto.GetOriginalURLRequest
	(*GetOriginalURLResponse)(nil), // 1: proto.GetOriginalURLResponse
	(*CreateShortURLRequest)(nil),  // 2: proto.CreateShortURLRequest
	(*CreateShortURLResponse)(nil), // 3: proto.CreateShortURLResponse
}
var file_proto_shortener_proto_depIdxs = []int32{
	0, // 0: proto.Shortener.GetOriginalURL:input_type -> proto.GetOriginalURLRequest
	2, // 1: proto.Shortener.CreateShortURL:input_type -> proto.CreateShortURLRequest
	1, // 2: proto.Shortener.GetOriginalURL:output_type -> proto.GetOriginalURLResponse
	3, // 3: proto.Shortener.CreateShortURL:output_type -> proto.CreateShortURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_shortener_proto_init() }
func file_proto_shortener_proto_init() {
	if File_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOriginalURLResponse); i {
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
		file_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLRequest); i {
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
		file_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLResponse); i {
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
			RawDescriptor: file_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortener_proto_goTypes,
		DependencyIndexes: file_proto_shortener_proto_depIdxs,
		MessageInfos:      file_proto_shortener_proto_msgTypes,
	}.Build()
	File_proto_shortener_proto = out.File
	file_proto_shortener_proto_rawDesc = nil
	file_proto_shortener_proto_goTypes = nil
	file_proto_shortener_proto_depIdxs = nil
}
