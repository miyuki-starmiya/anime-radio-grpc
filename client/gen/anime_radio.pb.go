// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: api/proto/anime_radio.proto

package __

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

type YouTubeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *YouTubeInfo) Reset() {
	*x = YouTubeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_anime_radio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YouTubeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YouTubeInfo) ProtoMessage() {}

func (x *YouTubeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_anime_radio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YouTubeInfo.ProtoReflect.Descriptor instead.
func (*YouTubeInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_anime_radio_proto_rawDescGZIP(), []int{0}
}

func (x *YouTubeInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *YouTubeInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type SlackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SlackResponse) Reset() {
	*x = SlackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_anime_radio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackResponse) ProtoMessage() {}

func (x *SlackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_anime_radio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackResponse.ProtoReflect.Descriptor instead.
func (*SlackResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_anime_radio_proto_rawDescGZIP(), []int{1}
}

func (x *SlackResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_api_proto_anime_radio_proto protoreflect.FileDescriptor

var file_api_proto_anime_radio_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x0b, 0x59, 0x6f, 0x75, 0x54, 0x75, 0x62, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x49, 0x0a,
	0x11, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x64, 0x69, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x59, 0x6f, 0x75, 0x54, 0x75,
	0x62, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_anime_radio_proto_rawDescOnce sync.Once
	file_api_proto_anime_radio_proto_rawDescData = file_api_proto_anime_radio_proto_rawDesc
)

func file_api_proto_anime_radio_proto_rawDescGZIP() []byte {
	file_api_proto_anime_radio_proto_rawDescOnce.Do(func() {
		file_api_proto_anime_radio_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_anime_radio_proto_rawDescData)
	})
	return file_api_proto_anime_radio_proto_rawDescData
}

var file_api_proto_anime_radio_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_anime_radio_proto_goTypes = []interface{}{
	(*YouTubeInfo)(nil),   // 0: YouTubeInfo
	(*SlackResponse)(nil), // 1: SlackResponse
}
var file_api_proto_anime_radio_proto_depIdxs = []int32{
	0, // 0: AnimeRadioService.SendAnimeRadioInfo:input_type -> YouTubeInfo
	1, // 1: AnimeRadioService.SendAnimeRadioInfo:output_type -> SlackResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_anime_radio_proto_init() }
func file_api_proto_anime_radio_proto_init() {
	if File_api_proto_anime_radio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_anime_radio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YouTubeInfo); i {
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
		file_api_proto_anime_radio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackResponse); i {
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
			RawDescriptor: file_api_proto_anime_radio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_anime_radio_proto_goTypes,
		DependencyIndexes: file_api_proto_anime_radio_proto_depIdxs,
		MessageInfos:      file_api_proto_anime_radio_proto_msgTypes,
	}.Build()
	File_api_proto_anime_radio_proto = out.File
	file_api_proto_anime_radio_proto_rawDesc = nil
	file_api_proto_anime_radio_proto_goTypes = nil
	file_api_proto_anime_radio_proto_depIdxs = nil
}
