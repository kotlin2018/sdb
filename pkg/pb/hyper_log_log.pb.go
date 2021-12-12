// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/protobuf-spec/hyper_log_log.proto

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

type HLLCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *HLLCreateRequest) Reset() {
	*x = HLLCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLCreateRequest) ProtoMessage() {}

func (x *HLLCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLCreateRequest.ProtoReflect.Descriptor instead.
func (*HLLCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *HLLCreateRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type HLLCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HLLCreateResponse) Reset() {
	*x = HLLCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLCreateResponse) ProtoMessage() {}

func (x *HLLCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLCreateResponse.ProtoReflect.Descriptor instead.
func (*HLLCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *HLLCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HLLDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *HLLDelRequest) Reset() {
	*x = HLLDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLDelRequest) ProtoMessage() {}

func (x *HLLDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLDelRequest.ProtoReflect.Descriptor instead.
func (*HLLDelRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *HLLDelRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type HLLDelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HLLDelResponse) Reset() {
	*x = HLLDelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLDelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLDelResponse) ProtoMessage() {}

func (x *HLLDelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLDelResponse.ProtoReflect.Descriptor instead.
func (*HLLDelResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{3}
}

func (x *HLLDelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HLLAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *HLLAddRequest) Reset() {
	*x = HLLAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLAddRequest) ProtoMessage() {}

func (x *HLLAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLAddRequest.ProtoReflect.Descriptor instead.
func (*HLLAddRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *HLLAddRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *HLLAddRequest) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type HLLAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *HLLAddResponse) Reset() {
	*x = HLLAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLAddResponse) ProtoMessage() {}

func (x *HLLAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLAddResponse.ProtoReflect.Descriptor instead.
func (*HLLAddResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{5}
}

func (x *HLLAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HLLCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *HLLCountRequest) Reset() {
	*x = HLLCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLCountRequest) ProtoMessage() {}

func (x *HLLCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLCountRequest.ProtoReflect.Descriptor instead.
func (*HLLCountRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{6}
}

func (x *HLLCountRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type HLLCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *HLLCountResponse) Reset() {
	*x = HLLCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLLCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLLCountResponse) ProtoMessage() {}

func (x *HLLCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_spec_hyper_log_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLLCountResponse.ProtoReflect.Descriptor instead.
func (*HLLCountResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP(), []int{7}
}

func (x *HLLCountResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_api_protobuf_spec_hyper_log_log_proto protoreflect.FileDescriptor

var file_api_protobuf_spec_hyper_log_log_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x10, 0x48, 0x4c, 0x4c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x11, 0x48, 0x4c, 0x4c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x48, 0x4c, 0x4c, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2a, 0x0a, 0x0e, 0x48, 0x4c, 0x4c, 0x44, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x39, 0x0a, 0x0d, 0x48, 0x4c, 0x4c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x2a, 0x0a,
	0x0e, 0x48, 0x4c, 0x4c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x23, 0x0a, 0x0f, 0x48, 0x4c, 0x4c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28,
	0x0a, 0x10, 0x48, 0x4c, 0x4c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_spec_hyper_log_log_proto_rawDescOnce sync.Once
	file_api_protobuf_spec_hyper_log_log_proto_rawDescData = file_api_protobuf_spec_hyper_log_log_proto_rawDesc
)

func file_api_protobuf_spec_hyper_log_log_proto_rawDescGZIP() []byte {
	file_api_protobuf_spec_hyper_log_log_proto_rawDescOnce.Do(func() {
		file_api_protobuf_spec_hyper_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_spec_hyper_log_log_proto_rawDescData)
	})
	return file_api_protobuf_spec_hyper_log_log_proto_rawDescData
}

var file_api_protobuf_spec_hyper_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_protobuf_spec_hyper_log_log_proto_goTypes = []interface{}{
	(*HLLCreateRequest)(nil),  // 0: proto.HLLCreateRequest
	(*HLLCreateResponse)(nil), // 1: proto.HLLCreateResponse
	(*HLLDelRequest)(nil),     // 2: proto.HLLDelRequest
	(*HLLDelResponse)(nil),    // 3: proto.HLLDelResponse
	(*HLLAddRequest)(nil),     // 4: proto.HLLAddRequest
	(*HLLAddResponse)(nil),    // 5: proto.HLLAddResponse
	(*HLLCountRequest)(nil),   // 6: proto.HLLCountRequest
	(*HLLCountResponse)(nil),  // 7: proto.HLLCountResponse
}
var file_api_protobuf_spec_hyper_log_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protobuf_spec_hyper_log_log_proto_init() }
func file_api_protobuf_spec_hyper_log_log_proto_init() {
	if File_api_protobuf_spec_hyper_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLCreateRequest); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLCreateResponse); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLDelRequest); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLDelResponse); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLAddRequest); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLAddResponse); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLCountRequest); i {
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
		file_api_protobuf_spec_hyper_log_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLLCountResponse); i {
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
			RawDescriptor: file_api_protobuf_spec_hyper_log_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protobuf_spec_hyper_log_log_proto_goTypes,
		DependencyIndexes: file_api_protobuf_spec_hyper_log_log_proto_depIdxs,
		MessageInfos:      file_api_protobuf_spec_hyper_log_log_proto_msgTypes,
	}.Build()
	File_api_protobuf_spec_hyper_log_log_proto = out.File
	file_api_protobuf_spec_hyper_log_log_proto_rawDesc = nil
	file_api_protobuf_spec_hyper_log_log_proto_goTypes = nil
	file_api_protobuf_spec_hyper_log_log_proto_depIdxs = nil
}
