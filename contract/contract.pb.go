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

type CompositeRules_Operator int32

const (
	CompositeRules_AND CompositeRules_Operator = 0
	CompositeRules_OR  CompositeRules_Operator = 1
)

// Enum value maps for CompositeRules_Operator.
var (
	CompositeRules_Operator_name = map[int32]string{
		0: "AND",
		1: "OR",
	}
	CompositeRules_Operator_value = map[string]int32{
		"AND": 0,
		"OR":  1,
	}
)

func (x CompositeRules_Operator) Enum() *CompositeRules_Operator {
	p := new(CompositeRules_Operator)
	*p = x
	return p
}

func (x CompositeRules_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompositeRules_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_contract_proto_enumTypes[0].Descriptor()
}

func (CompositeRules_Operator) Type() protoreflect.EnumType {
	return &file_contract_contract_proto_enumTypes[0]
}

func (x CompositeRules_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompositeRules_Operator.Descriptor instead.
func (CompositeRules_Operator) EnumDescriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{2, 0}
}

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

	Name          string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Method        string          `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Request       *anypb.Any      `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Response      *anypb.Any      `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	WantError     bool            `protobuf:"varint,5,opt,name=wantError,proto3" json:"wantError,omitempty"`
	Rules         *CompositeRules `protobuf:"bytes,6,opt,name=rules,proto3" json:"rules,omitempty"`
	Preconditions []string        `protobuf:"bytes,7,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
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

func (x *Interaction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

func (x *Interaction) GetWantError() bool {
	if x != nil {
		return x.WantError
	}
	return false
}

func (x *Interaction) GetRules() *CompositeRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Interaction) GetPreconditions() []string {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

type CompositeRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator    CompositeRules_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=contract.CompositeRules_Operator" json:"operator,omitempty"`
	IntRules    []*IntRule              `protobuf:"bytes,2,rep,name=intRules,proto3" json:"intRules,omitempty"`
	StringRules []*StringRule           `protobuf:"bytes,3,rep,name=stringRules,proto3" json:"stringRules,omitempty"`
	DoubleRules []*DoubleRule           `protobuf:"bytes,4,rep,name=doubleRules,proto3" json:"doubleRules,omitempty"`
	NestedRules []*CompositeRules       `protobuf:"bytes,5,rep,name=nestedRules,proto3" json:"nestedRules,omitempty"`
}

func (x *CompositeRules) Reset() {
	*x = CompositeRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositeRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositeRules) ProtoMessage() {}

func (x *CompositeRules) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CompositeRules.ProtoReflect.Descriptor instead.
func (*CompositeRules) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{2}
}

func (x *CompositeRules) GetOperator() CompositeRules_Operator {
	if x != nil {
		return x.Operator
	}
	return CompositeRules_AND
}

func (x *CompositeRules) GetIntRules() []*IntRule {
	if x != nil {
		return x.IntRules
	}
	return nil
}

func (x *CompositeRules) GetStringRules() []*StringRule {
	if x != nil {
		return x.StringRules
	}
	return nil
}

func (x *CompositeRules) GetDoubleRules() []*DoubleRule {
	if x != nil {
		return x.DoubleRules
	}
	return nil
}

func (x *CompositeRules) GetNestedRules() []*CompositeRules {
	if x != nil {
		return x.NestedRules
	}
	return nil
}

type IntRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string  `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Min      int64   `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	CheckMin bool    `protobuf:"varint,3,opt,name=checkMin,proto3" json:"checkMin,omitempty"`
	Max      int64   `protobuf:"varint,4,opt,name=max,proto3" json:"max,omitempty"`
	ChechMax bool    `protobuf:"varint,5,opt,name=chechMax,proto3" json:"chechMax,omitempty"`
	Allowed  []int64 `protobuf:"varint,6,rep,packed,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *IntRule) Reset() {
	*x = IntRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntRule) ProtoMessage() {}

func (x *IntRule) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IntRule.ProtoReflect.Descriptor instead.
func (*IntRule) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{3}
}

func (x *IntRule) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *IntRule) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *IntRule) GetCheckMin() bool {
	if x != nil {
		return x.CheckMin
	}
	return false
}

func (x *IntRule) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *IntRule) GetChechMax() bool {
	if x != nil {
		return x.ChechMax
	}
	return false
}

func (x *IntRule) GetAllowed() []int64 {
	if x != nil {
		return x.Allowed
	}
	return nil
}

type DoubleRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    string  `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Min      float64 `protobuf:"fixed64,2,opt,name=min,proto3" json:"min,omitempty"`
	CheckMin bool    `protobuf:"varint,3,opt,name=checkMin,proto3" json:"checkMin,omitempty"`
	Max      float64 `protobuf:"fixed64,4,opt,name=max,proto3" json:"max,omitempty"`
	ChechMax bool    `protobuf:"varint,5,opt,name=chechMax,proto3" json:"chechMax,omitempty"`
}

func (x *DoubleRule) Reset() {
	*x = DoubleRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleRule) ProtoMessage() {}

func (x *DoubleRule) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleRule.ProtoReflect.Descriptor instead.
func (*DoubleRule) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{4}
}

func (x *DoubleRule) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *DoubleRule) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DoubleRule) GetCheckMin() bool {
	if x != nil {
		return x.CheckMin
	}
	return false
}

func (x *DoubleRule) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *DoubleRule) GetChechMax() bool {
	if x != nil {
		return x.ChechMax
	}
	return false
}

type StringRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field      string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	MatchRegex string   `protobuf:"bytes,2,opt,name=matchRegex,proto3" json:"matchRegex,omitempty"`
	Allowed    []string `protobuf:"bytes,3,rep,name=allowed,proto3" json:"allowed,omitempty"`
}

func (x *StringRule) Reset() {
	*x = StringRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringRule) ProtoMessage() {}

func (x *StringRule) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringRule.ProtoReflect.Descriptor instead.
func (*StringRule) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{5}
}

func (x *StringRule) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *StringRule) GetMatchRegex() string {
	if x != nil {
		return x.MatchRegex
	}
	return ""
}

func (x *StringRule) GetAllowed() []string {
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
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x0b,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x6e,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x61,
	0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xc7, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x2d, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x36,
	0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x0b, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x0b, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x08, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x44, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x01, 0x22, 0x95, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x4d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x4d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x68, 0x4d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x68, 0x65,
	0x63, 0x68, 0x4d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22,
	0x7e, 0x0a, 0x0a, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6d, 0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x68, 0x4d, 0x61, 0x78, 0x22,
	0x5c, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x67, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x75, 0x78,
	0x30, 0x30, 0x34, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_contract_contract_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contract_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_contract_contract_proto_goTypes = []interface{}{
	(CompositeRules_Operator)(0), // 0: contract.CompositeRules.Operator
	(*Contract)(nil),             // 1: contract.Contract
	(*Interaction)(nil),          // 2: contract.Interaction
	(*CompositeRules)(nil),       // 3: contract.CompositeRules
	(*IntRule)(nil),              // 4: contract.IntRule
	(*DoubleRule)(nil),           // 5: contract.DoubleRule
	(*StringRule)(nil),           // 6: contract.StringRule
	(*anypb.Any)(nil),            // 7: google.protobuf.Any
}
var file_contract_contract_proto_depIdxs = []int32{
	2, // 0: contract.Contract.interactions:type_name -> contract.Interaction
	7, // 1: contract.Interaction.request:type_name -> google.protobuf.Any
	7, // 2: contract.Interaction.response:type_name -> google.protobuf.Any
	3, // 3: contract.Interaction.rules:type_name -> contract.CompositeRules
	0, // 4: contract.CompositeRules.operator:type_name -> contract.CompositeRules.Operator
	4, // 5: contract.CompositeRules.intRules:type_name -> contract.IntRule
	6, // 6: contract.CompositeRules.stringRules:type_name -> contract.StringRule
	5, // 7: contract.CompositeRules.doubleRules:type_name -> contract.DoubleRule
	3, // 8: contract.CompositeRules.nestedRules:type_name -> contract.CompositeRules
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
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
			switch v := v.(*CompositeRules); i {
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
			switch v := v.(*IntRule); i {
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
		file_contract_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleRule); i {
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
		file_contract_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringRule); i {
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
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_contract_proto_goTypes,
		DependencyIndexes: file_contract_contract_proto_depIdxs,
		EnumInfos:         file_contract_contract_proto_enumTypes,
		MessageInfos:      file_contract_contract_proto_msgTypes,
	}.Build()
	File_contract_contract_proto = out.File
	file_contract_contract_proto_rawDesc = nil
	file_contract_contract_proto_goTypes = nil
	file_contract_contract_proto_depIdxs = nil
}
