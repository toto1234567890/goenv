// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: backendDb.proto

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

type LogDbMsg_LevelNumber int32

const (
	LogDbMsg_unloggable LogDbMsg_LevelNumber = 0
	LogDbMsg_notset     LogDbMsg_LevelNumber = 1
	LogDbMsg_debug      LogDbMsg_LevelNumber = 2
	LogDbMsg_stream     LogDbMsg_LevelNumber = 3
	LogDbMsg_info       LogDbMsg_LevelNumber = 4
	LogDbMsg_logon      LogDbMsg_LevelNumber = 5
	LogDbMsg_logout     LogDbMsg_LevelNumber = 6
	LogDbMsg_trade      LogDbMsg_LevelNumber = 7
	LogDbMsg_schedule   LogDbMsg_LevelNumber = 8
	LogDbMsg_report     LogDbMsg_LevelNumber = 9
	LogDbMsg_warning    LogDbMsg_LevelNumber = 10
	LogDbMsg_error      LogDbMsg_LevelNumber = 11
	LogDbMsg_critical   LogDbMsg_LevelNumber = 12
)

// Enum value maps for LogDbMsg_LevelNumber.
var (
	LogDbMsg_LevelNumber_name = map[int32]string{
		0:  "unloggable",
		1:  "notset",
		2:  "debug",
		3:  "stream",
		4:  "info",
		5:  "logon",
		6:  "logout",
		7:  "trade",
		8:  "schedule",
		9:  "report",
		10: "warning",
		11: "error",
		12: "critical",
	}
	LogDbMsg_LevelNumber_value = map[string]int32{
		"unloggable": 0,
		"notset":     1,
		"debug":      2,
		"stream":     3,
		"info":       4,
		"logon":      5,
		"logout":     6,
		"trade":      7,
		"schedule":   8,
		"report":     9,
		"warning":    10,
		"error":      11,
		"critical":   12,
	}
)

func (x LogDbMsg_LevelNumber) Enum() *LogDbMsg_LevelNumber {
	p := new(LogDbMsg_LevelNumber)
	*p = x
	return p
}

func (x LogDbMsg_LevelNumber) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogDbMsg_LevelNumber) Descriptor() protoreflect.EnumDescriptor {
	return file_backendDb_proto_enumTypes[0].Descriptor()
}

func (LogDbMsg_LevelNumber) Type() protoreflect.EnumType {
	return &file_backendDb_proto_enumTypes[0]
}

func (x LogDbMsg_LevelNumber) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogDbMsg_LevelNumber.Descriptor instead.
func (LogDbMsg_LevelNumber) EnumDescriptor() ([]byte, []int) {
	return file_backendDb_proto_rawDescGZIP(), []int{0, 0}
}

type LogDbMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// recorded by log_server :
	Timestamp    string               `protobuf:"bytes,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`                               // When the event occurred
	Hostname     string               `protobuf:"bytes,2,opt,name=Hostname,proto3" json:"Hostname,omitempty"`                                 // Host/machine name
	LoggerName   string               `protobuf:"bytes,3,opt,name=LoggerName,proto3" json:"LoggerName,omitempty"`                             // Name of the logger (usually __name__)
	Module       string               `protobuf:"bytes,4,opt,name=Module,proto3" json:"Module,omitempty"`                                     // Module (name portion of filename)
	Level        LogDbMsg_LevelNumber `protobuf:"varint,5,opt,name=Level,proto3,enum=MyDatabase.LogDbMsg_LevelNumber" json:"Level,omitempty"` // Logging level/severity
	Filename     string               `protobuf:"bytes,6,opt,name=Filename,proto3" json:"Filename,omitempty"`                                 // Filename portion of pathname
	FunctionName string               `protobuf:"bytes,7,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`                         // Function name
	LineNumber   string               `protobuf:"bytes,8,opt,name=LineNumber,proto3" json:"LineNumber,omitempty"`                             // Source line number
	Message      string               `protobuf:"bytes,9,opt,name=Message,proto3" json:"Message,omitempty"`                                   // The log message
	// others
	// path of the file
	PathName string `protobuf:"bytes,10,opt,name=PathName,proto3" json:"PathName,omitempty"` // Full pathname of the source file
	// Process information
	ProcessId   string `protobuf:"bytes,11,opt,name=ProcessId,proto3" json:"ProcessId,omitempty"`     // Process ID
	ProcessName string `protobuf:"bytes,12,opt,name=ProcessName,proto3" json:"ProcessName,omitempty"` // Process name
	// Thread information
	ThreadId   string `protobuf:"bytes,13,opt,name=ThreadId,proto3" json:"ThreadId,omitempty"`     // Thread ID
	ThreadName string `protobuf:"bytes,14,opt,name=ThreadName,proto3" json:"ThreadName,omitempty"` // Thread name
	// Additional requested fields
	ServiceName string `protobuf:"bytes,15,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"` // Name of the service generating the log
	// Optional stack trace for errors
	StackTrace string `protobuf:"bytes,16,opt,name=StackTrace,proto3" json:"StackTrace,omitempty"` // Stack trace if available
}

func (x *LogDbMsg) Reset() {
	*x = LogDbMsg{}
	mi := &file_backendDb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogDbMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDbMsg) ProtoMessage() {}

func (x *LogDbMsg) ProtoReflect() protoreflect.Message {
	mi := &file_backendDb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDbMsg.ProtoReflect.Descriptor instead.
func (*LogDbMsg) Descriptor() ([]byte, []int) {
	return file_backendDb_proto_rawDescGZIP(), []int{0}
}

func (x *LogDbMsg) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *LogDbMsg) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *LogDbMsg) GetLoggerName() string {
	if x != nil {
		return x.LoggerName
	}
	return ""
}

func (x *LogDbMsg) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *LogDbMsg) GetLevel() LogDbMsg_LevelNumber {
	if x != nil {
		return x.Level
	}
	return LogDbMsg_unloggable
}

func (x *LogDbMsg) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *LogDbMsg) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *LogDbMsg) GetLineNumber() string {
	if x != nil {
		return x.LineNumber
	}
	return ""
}

func (x *LogDbMsg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogDbMsg) GetPathName() string {
	if x != nil {
		return x.PathName
	}
	return ""
}

func (x *LogDbMsg) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *LogDbMsg) GetProcessName() string {
	if x != nil {
		return x.ProcessName
	}
	return ""
}

func (x *LogDbMsg) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *LogDbMsg) GetThreadName() string {
	if x != nil {
		return x.ThreadName
	}
	return ""
}

func (x *LogDbMsg) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *LogDbMsg) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

var File_backendDb_proto protoreflect.FileDescriptor

var file_backendDb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x4d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0xb7, 0x05,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x44, 0x62, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x4d, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x62, 0x4d, 0x73,
	0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x75, 0x6e, 0x6c, 0x6f,
	0x67, 0x67, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x73,
	0x65, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x6f, 0x6e, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x10,
	0x09, 0x12, 0x0b, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x09,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x72, 0x69,
	0x74, 0x69, 0x63, 0x61, 0x6c, 0x10, 0x0c, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backendDb_proto_rawDescOnce sync.Once
	file_backendDb_proto_rawDescData = file_backendDb_proto_rawDesc
)

func file_backendDb_proto_rawDescGZIP() []byte {
	file_backendDb_proto_rawDescOnce.Do(func() {
		file_backendDb_proto_rawDescData = protoimpl.X.CompressGZIP(file_backendDb_proto_rawDescData)
	})
	return file_backendDb_proto_rawDescData
}

var file_backendDb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backendDb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_backendDb_proto_goTypes = []any{
	(LogDbMsg_LevelNumber)(0), // 0: MyDatabase.LogDbMsg.LevelNumber
	(*LogDbMsg)(nil),          // 1: MyDatabase.LogDbMsg
}
var file_backendDb_proto_depIdxs = []int32{
	0, // 0: MyDatabase.LogDbMsg.Level:type_name -> MyDatabase.LogDbMsg.LevelNumber
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_backendDb_proto_init() }
func file_backendDb_proto_init() {
	if File_backendDb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backendDb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backendDb_proto_goTypes,
		DependencyIndexes: file_backendDb_proto_depIdxs,
		EnumInfos:         file_backendDb_proto_enumTypes,
		MessageInfos:      file_backendDb_proto_msgTypes,
	}.Build()
	File_backendDb_proto = out.File
	file_backendDb_proto_rawDesc = nil
	file_backendDb_proto_goTypes = nil
	file_backendDb_proto_depIdxs = nil
}
