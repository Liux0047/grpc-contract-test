// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: contract/contract.proto

package contract

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service      string         `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Consumer     string         `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Interactions []*Interaction `protobuf:"bytes,3,rep,name=interactions,proto3" json:"interactions,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Contract) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

func (x *Contract) GetInteractions() []*Interaction {
	if x != nil {
		return x.Interactions
	}
	return nil
}

type Interaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method           string             `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Request          *anypb.Any         `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response         *anypb.Any         `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	IntConditions    []*IntCondition    `protobuf:"bytes,4,rep,name=intConditions,proto3" json:"intConditions,omitempty"`
	StringConditions []*StringCondition `protobuf:"bytes,5,rep,name=stringConditions,proto3" json:"stringConditions,omitempty"`
}

func (x *Interaction) Reset() {
	*x = Interaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interaction) ProtoMessage() {}

func (x *Interaction) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interaction.ProtoReflect.Descriptor instead.
func (*Interaction) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Interaction) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Interaction) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Interaction) GetResponse() *anypb.Any {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Interaction) GetIntConditions() []*IntCondition {
	if x != nil {
		return x.IntConditions
	}
	return nil
}

func (x *Interaction) GetStringConditions() []*StringCondition {
	if x != nil {
		return x.StringConditions
	}
	return nil
}

type IntCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field   string  `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Min     int64   `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max     int64   `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Allowed []int64 `protobuf:"varint,4,rep,packed,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *IntCondition) Reset() {
	*x = IntCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntCondition) ProtoMessage() {}

func (x *IntCondition) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntCondition.ProtoReflect.Descriptor instead.
func (*IntCondition) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{2}
}

func (x *IntCondition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *IntCondition) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *IntCondition) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *IntCondition) GetAllowed() []int64 {
	if x != nil {
		return x.Allowed
	}
	return nil
}

type StringCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field      string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	MatchRegex string   `protobuf:"bytes,2,opt,name=matchRegex,proto3" json:"matchRegex,omitempty"`
	Allowed    []string `protobuf:"bytes,3,rep,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *StringCondition) Reset() {
	*x = StringCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringCondition) ProtoMessage() {}

func (x *StringCondition) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringCondition.ProtoReflect.Descriptor instead.
func (*StringCondition) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{3}
}

func (x *StringCondition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *StringCondition) GetMatchRegex() string {
	if x != nil {
		return x.MatchRegex
	}
	return ""
}

func (x *StringCondition) GetAllowed() []string {
	if x != nil {
		return x.Allowed
	}
	return nil
}

var File_contract_contract_proto protoreflect.FileDescriptor

var file_contract_contract_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x12, 0x39, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8c, 0x02, 0x0a, 0x0b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x62, 0x0a, 0x0c, 0x49, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x61,
	0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x75, 0x78, 0x30, 0x30, 0x34, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_contract_proto_rawDescOnce sync.Once
	file_contract_contract_proto_rawDescData = file_contract_contract_proto_rawDesc
)

func file_contract_contract_proto_rawDescGZIP() []byte {
	file_contract_contract_proto_rawDescOnce.Do(func() {
		file_contract_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_contract_proto_rawDescData)
	})
	return file_contract_contract_proto_rawDescData
}

var file_contract_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contract_contract_proto_goTypes = []interface{}{
	(*Contract)(nil),        // 0: contract.Contract
	(*Interaction)(nil),     // 1: contract.Interaction
	(*IntCondition)(nil),    // 2: contract.IntCondition
	(*StringCondition)(nil), // 3: contract.StringCondition
	(*anypb.Any)(nil),       // 4: google.protobuf.Any
}
var file_contract_contract_proto_depIdxs = []int32{
	1, // 0: contract.Contract.interactions:type_name -> contract.Interaction
	4, // 1: contract.Interaction.request:type_name -> google.protobuf.Any
	4, // 2: contract.Interaction.response:type_name -> google.protobuf.Any
	2, // 3: contract.Interaction.intConditions:type_name -> contract.IntCondition
	3, // 4: contract.Interaction.stringConditions:type_name -> contract.StringCondition
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_contract_contract_proto_init() }
func file_contract_contract_proto_init() {
	if File_contract_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_contract_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interaction); i {
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
		file_contract_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntCondition); i {
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
		file_contract_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringCondition); i {
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
			RawDescriptor: file_contract_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_contract_proto_goTypes,
		DependencyIndexes: file_contract_contract_proto_depIdxs,
		MessageInfos:      file_contract_contract_proto_msgTypes,
	}.Build()
	File_contract_contract_proto = out.File
	file_contract_contract_proto_rawDesc = nil
	file_contract_contract_proto_goTypes = nil
	file_contract_contract_proto_depIdxs = nil
}
