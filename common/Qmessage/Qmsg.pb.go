// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: Qmsg.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamMsg_StreamAction int32

const (
	StreamMsg_QUOTE       StreamMsg_StreamAction = 0
	StreamMsg_UNQUOTE     StreamMsg_StreamAction = 1
	StreamMsg_TICK        StreamMsg_StreamAction = 2
	StreamMsg_UNTICK      StreamMsg_StreamAction = 3
	StreamMsg_ORDERBOOK   StreamMsg_StreamAction = 4
	StreamMsg_UNORDERBOOK StreamMsg_StreamAction = 5
)

// Enum value maps for StreamMsg_StreamAction.
var (
	StreamMsg_StreamAction_name = map[int32]string{
		0: "QUOTE",
		1: "UNQUOTE",
		2: "TICK",
		3: "UNTICK",
		4: "ORDERBOOK",
		5: "UNORDERBOOK",
	}
	StreamMsg_StreamAction_value = map[string]int32{
		"QUOTE":       0,
		"UNQUOTE":     1,
		"TICK":        2,
		"UNTICK":      3,
		"ORDERBOOK":   4,
		"UNORDERBOOK": 5,
	}
)

func (x StreamMsg_StreamAction) Enum() *StreamMsg_StreamAction {
	p := new(StreamMsg_StreamAction)
	*p = x
	return p
}

func (x StreamMsg_StreamAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamMsg_StreamAction) Descriptor() protoreflect.EnumDescriptor {
	return file_Qmsg_proto_enumTypes[0].Descriptor()
}

func (StreamMsg_StreamAction) Type() protoreflect.EnumType {
	return &file_Qmsg_proto_enumTypes[0]
}

func (x StreamMsg_StreamAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamMsg_StreamAction.Descriptor instead.
func (StreamMsg_StreamAction) EnumDescriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{1, 0}
}

type TradeMsg_TradeAction int32

const (
	TradeMsg_BUY  TradeMsg_TradeAction = 0
	TradeMsg_SELL TradeMsg_TradeAction = 1
	TradeMsg_CALL TradeMsg_TradeAction = 2
	TradeMsg_PUT  TradeMsg_TradeAction = 3
)

// Enum value maps for TradeMsg_TradeAction.
var (
	TradeMsg_TradeAction_name = map[int32]string{
		0: "BUY",
		1: "SELL",
		2: "CALL",
		3: "PUT",
	}
	TradeMsg_TradeAction_value = map[string]int32{
		"BUY":  0,
		"SELL": 1,
		"CALL": 2,
		"PUT":  3,
	}
)

func (x TradeMsg_TradeAction) Enum() *TradeMsg_TradeAction {
	p := new(TradeMsg_TradeAction)
	*p = x
	return p
}

func (x TradeMsg_TradeAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TradeMsg_TradeAction) Descriptor() protoreflect.EnumDescriptor {
	return file_Qmsg_proto_enumTypes[1].Descriptor()
}

func (TradeMsg_TradeAction) Type() protoreflect.EnumType {
	return &file_Qmsg_proto_enumTypes[1]
}

func (x TradeMsg_TradeAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TradeMsg_TradeAction.Descriptor instead.
func (TradeMsg_TradeAction) EnumDescriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{2, 0}
}

type HelloMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LocalServer string `protobuf:"bytes,2,opt,name=localServer,proto3" json:"localServer,omitempty"`
	LocalPort   string `protobuf:"bytes,3,opt,name=localPort,proto3" json:"localPort,omitempty"`
}

func (x *HelloMsg) Reset() {
	*x = HelloMsg{}
	mi := &file_Qmsg_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMsg) ProtoMessage() {}

func (x *HelloMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Qmsg_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMsg.ProtoReflect.Descriptor instead.
func (*HelloMsg) Descriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{0}
}

func (x *HelloMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloMsg) GetLocalServer() string {
	if x != nil {
		return x.LocalServer
	}
	return ""
}

func (x *HelloMsg) GetLocalPort() string {
	if x != nil {
		return x.LocalPort
	}
	return ""
}

type StreamMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BROKER        string                 `protobuf:"bytes,1,opt,name=BROKER,proto3" json:"BROKER,omitempty"`
	STREAM_ACTION StreamMsg_StreamAction `protobuf:"varint,2,opt,name=STREAM_ACTION,json=STREAMACTION,proto3,enum=Qmsg.StreamMsg_StreamAction" json:"STREAM_ACTION,omitempty"`
	TICKER        string                 `protobuf:"bytes,3,opt,name=TICKER,proto3" json:"TICKER,omitempty"`
}

func (x *StreamMsg) Reset() {
	*x = StreamMsg{}
	mi := &file_Qmsg_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMsg) ProtoMessage() {}

func (x *StreamMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Qmsg_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMsg.ProtoReflect.Descriptor instead.
func (*StreamMsg) Descriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{1}
}

func (x *StreamMsg) GetBROKER() string {
	if x != nil {
		return x.BROKER
	}
	return ""
}

func (x *StreamMsg) GetSTREAM_ACTION() StreamMsg_StreamAction {
	if x != nil {
		return x.STREAM_ACTION
	}
	return StreamMsg_QUOTE
}

func (x *StreamMsg) GetTICKER() string {
	if x != nil {
		return x.TICKER
	}
	return ""
}

type TradeMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ACCOUNT_ID  int32                 `protobuf:"varint,1,opt,name=ACCOUNT_ID,json=ACCOUNTID,proto3" json:"ACCOUNT_ID,omitempty"`
	TRADEACTION TradeMsg_TradeAction  `protobuf:"varint,2,opt,name=TRADEACTION,proto3,enum=Qmsg.TradeMsg_TradeAction" json:"TRADEACTION,omitempty"`
	TICKER      string                `protobuf:"bytes,3,opt,name=TICKER,proto3" json:"TICKER,omitempty"`
	QUANTITY    string                `protobuf:"bytes,4,opt,name=QUANTITY,proto3" json:"QUANTITY,omitempty"`
	TradeParams map[string]*anypb.Any `protobuf:"bytes,5,rep,name=TradeParams,proto3" json:"TradeParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TradeMsg) Reset() {
	*x = TradeMsg{}
	mi := &file_Qmsg_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeMsg) ProtoMessage() {}

func (x *TradeMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Qmsg_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeMsg.ProtoReflect.Descriptor instead.
func (*TradeMsg) Descriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{2}
}

func (x *TradeMsg) GetACCOUNT_ID() int32 {
	if x != nil {
		return x.ACCOUNT_ID
	}
	return 0
}

func (x *TradeMsg) GetTRADEACTION() TradeMsg_TradeAction {
	if x != nil {
		return x.TRADEACTION
	}
	return TradeMsg_BUY
}

func (x *TradeMsg) GetTICKER() string {
	if x != nil {
		return x.TICKER
	}
	return ""
}

func (x *TradeMsg) GetQUANTITY() string {
	if x != nil {
		return x.QUANTITY
	}
	return ""
}

func (x *TradeMsg) GetTradeParams() map[string]*anypb.Any {
	if x != nil {
		return x.TradeParams
	}
	return nil
}

type QMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FROME    string `protobuf:"bytes,2,opt,name=FROME,proto3" json:"FROME,omitempty"`
	TOO      string `protobuf:"bytes,3,opt,name=TOO,proto3" json:"TOO,omitempty"`
	ACKW     bool   `protobuf:"varint,4,opt,name=ACKW,proto3" json:"ACKW,omitempty"`
	PRIORITY bool   `protobuf:"varint,5,opt,name=PRIORITY,proto3" json:"PRIORITY,omitempty"`
	MESSAGE  []byte `protobuf:"bytes,6,opt,name=MESSAGE,proto3" json:"MESSAGE,omitempty"`
}

func (x *QMsg) Reset() {
	*x = QMsg{}
	mi := &file_Qmsg_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMsg) ProtoMessage() {}

func (x *QMsg) ProtoReflect() protoreflect.Message {
	mi := &file_Qmsg_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMsg.ProtoReflect.Descriptor instead.
func (*QMsg) Descriptor() ([]byte, []int) {
	return file_Qmsg_proto_rawDescGZIP(), []int{3}
}

func (x *QMsg) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *QMsg) GetFROME() string {
	if x != nil {
		return x.FROME
	}
	return ""
}

func (x *QMsg) GetTOO() string {
	if x != nil {
		return x.TOO
	}
	return ""
}

func (x *QMsg) GetACKW() bool {
	if x != nil {
		return x.ACKW
	}
	return false
}

func (x *QMsg) GetPRIORITY() bool {
	if x != nil {
		return x.PRIORITY
	}
	return false
}

func (x *QMsg) GetMESSAGE() []byte {
	if x != nil {
		return x.MESSAGE
	}
	return nil
}

var File_Qmsg_proto protoreflect.FileDescriptor

var file_Qmsg_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x51, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x51, 0x6d,
	0x73, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e,
	0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xdc,
	0x01, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x42, 0x52, 0x4f, 0x4b, 0x45, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x52,
	0x4f, 0x4b, 0x45, 0x52, 0x12, 0x41, 0x0a, 0x0d, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x51, 0x6d,
	0x73, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x73, 0x67, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x53, 0x54, 0x52, 0x45, 0x41,
	0x4d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x52, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x52, 0x22,
	0x5c, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x09, 0x0a, 0x05, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x51, 0x55, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x43, 0x4b, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x54, 0x49, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x05, 0x22, 0xe9, 0x02,
	0x0a, 0x08, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x44, 0x12, 0x3c, 0x0a, 0x0b, 0x54, 0x52, 0x41,
	0x44, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x51, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x52, 0x41, 0x44,
	0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x52, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x52, 0x12,
	0x1a, 0x0a, 0x08, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x12, 0x41, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x51, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4d, 0x73, 0x67,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x54,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x33, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x03, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x51, 0x4d,
	0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x52, 0x4f, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x46, 0x52, 0x4f, 0x4d, 0x45, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x4f, 0x4f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x4f, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x43,
	0x4b, 0x57, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x41, 0x43, 0x4b, 0x57, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Qmsg_proto_rawDescOnce sync.Once
	file_Qmsg_proto_rawDescData = file_Qmsg_proto_rawDesc
)

func file_Qmsg_proto_rawDescGZIP() []byte {
	file_Qmsg_proto_rawDescOnce.Do(func() {
		file_Qmsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qmsg_proto_rawDescData)
	})
	return file_Qmsg_proto_rawDescData
}

var file_Qmsg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Qmsg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Qmsg_proto_goTypes = []any{
	(StreamMsg_StreamAction)(0), // 0: Qmsg.StreamMsg.StreamAction
	(TradeMsg_TradeAction)(0),   // 1: Qmsg.TradeMsg.TradeAction
	(*HelloMsg)(nil),            // 2: Qmsg.HelloMsg
	(*StreamMsg)(nil),           // 3: Qmsg.StreamMsg
	(*TradeMsg)(nil),            // 4: Qmsg.TradeMsg
	(*QMsg)(nil),                // 5: Qmsg.QMsg
	nil,                         // 6: Qmsg.TradeMsg.TradeParamsEntry
	(*anypb.Any)(nil),           // 7: google.protobuf.Any
}
var file_Qmsg_proto_depIdxs = []int32{
	0, // 0: Qmsg.StreamMsg.STREAM_ACTION:type_name -> Qmsg.StreamMsg.StreamAction
	1, // 1: Qmsg.TradeMsg.TRADEACTION:type_name -> Qmsg.TradeMsg.TradeAction
	6, // 2: Qmsg.TradeMsg.TradeParams:type_name -> Qmsg.TradeMsg.TradeParamsEntry
	7, // 3: Qmsg.TradeMsg.TradeParamsEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Qmsg_proto_init() }
func file_Qmsg_proto_init() {
	if File_Qmsg_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Qmsg_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qmsg_proto_goTypes,
		DependencyIndexes: file_Qmsg_proto_depIdxs,
		EnumInfos:         file_Qmsg_proto_enumTypes,
		MessageInfos:      file_Qmsg_proto_msgTypes,
	}.Build()
	File_Qmsg_proto = out.File
	file_Qmsg_proto_rawDesc = nil
	file_Qmsg_proto_goTypes = nil
	file_Qmsg_proto_depIdxs = nil
}
