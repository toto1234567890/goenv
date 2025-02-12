package MyLogger

// used for backend services -> capnp proto
import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	loggerMsg "govenv/api/capnp/loggerMsg"
	"govenv/pkg/common/Config"
	"govenv/pkg/common/PyHelpers"

	capnplib "capnproto.org/go/capnp/v3"
)

const (
	NOTSET     int8 = 0
	UNLOGGABLE int8 = 1
	DEBUG      int8 = 10
	STREAM     int8 = 11
	INFO       int8 = 20
	LOGON      int8 = 21
	LOGOUT     int8 = 22
	TRADE      int8 = 23
	SCHEDULE   int8 = 24
	REPORT     int8 = 25
	WARNING    int8 = 30
	ERROR      int8 = 40
	CRITICAL   int8 = 50

	StrNOTSET     string = "NOTSET"
	StrUNLOGGABLE string = "UNLOGGABLE"
	StrDEBUG      string = "DEBUG"
	StrSTREAM     string = "STREAM"
	StrINFO       string = "INFO"
	StrLOGON      string = "LOGON"
	StrLOGOUT     string = "LOGOUT"
	StrTRADE      string = "TRADE"
	StrSCHEDULE   string = "SCHEDULE"
	StrREPORT     string = "REPORT"
	StrWARNING    string = "WARNING"
	StrERROR      string = "ERROR"
	StrCRITICAL   string = "CRITICAL"
)

var (
	LogLevelToStr = map[int8]string{NOTSET: StrNOTSET, UNLOGGABLE: StrUNLOGGABLE, DEBUG: StrDEBUG, STREAM: StrSTREAM, INFO: StrINFO, LOGON: StrLOGON, LOGOUT: StrLOGOUT, TRADE: StrTRADE, SCHEDULE: StrSCHEDULE, REPORT: StrREPORT, WARNING: StrWARNING, ERROR: StrERROR, CRITICAL: StrCRITICAL}
	LogLevelToInt = map[string]int8{StrNOTSET: NOTSET, StrUNLOGGABLE: UNLOGGABLE, StrDEBUG: DEBUG, StrSTREAM: STREAM, StrINFO: INFO, StrLOGON: LOGON, StrLOGOUT: LOGOUT, StrTRADE: TRADE, StrSCHEDULE: SCHEDULE, StrREPORT: REPORT, StrWARNING: WARNING, StrERROR: ERROR, StrCRITICAL: CRITICAL}
)

type LoggerMessage struct {
	Timestamp    string // When the event occurred
	Hostname     string // Host/machine name
	LoggerName   string // Name of the logger (usually __name__)
	Module       string // Module (name portion of filename)
	Level        string // Logging level/severity
	Filename     string // Filename portion of pathname
	FunctionName string // Function name
	LineNumber   string // Source line number
	Message      string // The log message
	// extra informations...
	PathName    string // Full pathname of the source file
	ProcessId   string // Process ID
	ProcessName string // Process name
	ThreadId    string // Thread ID
	ThreadName  string // Thread name
	ServiceName string // Name of the service generating the log
	StackTrace  string // Stack trace if available
}

var (
	hostName   string = func() string { host, _ := os.Hostname(); return host }()
	ProcessId  string = strconv.Itoa(os.Getgid())
	ThreadName string = "GoRoutine" // go : no thread -> goroutine
	ThreadId   string = ""          // go : no thread -> goroutine
)

func NewLoggerMessage(parentClassName string, serviceName string) *LoggerMessage {
	loggerMessage := &LoggerMessage{
		// will be set/forced automatically
		Hostname:    hostName,
		ProcessId:   ProcessId,
		ProcessName: parentClassName,
		ThreadName:  ThreadName,
		ThreadId:    ThreadId,
		LoggerName:  fmt.Sprintf("%s %s %s", parentClassName, ProcessId, serviceName),
		ServiceName: serviceName,
	}
	return loggerMessage
}

func (loggerMessage *LoggerMessage) createNewMessage(level int8, msg, filepath, module, file, function, line, StackTrace string) *LoggerMessage {
	// required
	loggerMessage.Timestamp = time.Now().UTC().Format("2006-01-02 15:04:05.000000")
	loggerMessage.Level = LogLevelToStr[level]
	loggerMessage.Message = msg
	// extra informations...
	loggerMessage.Module = module
	loggerMessage.PathName = filepath
	loggerMessage.Filename = file
	loggerMessage.FunctionName = function
	loggerMessage.LineNumber = line
	loggerMessage.StackTrace = StackTrace
	return loggerMessage
}

func (loggerMessage *LoggerMessage) GenerateSqlLite3Table(tableName string) string {
	createTableStmt := "BEGIN;\n"
	createTableStmt += fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", strings.ToLower(tableName))
	createTableStmt += "\tid INTEGER PRIMARY KEY AUTOINCREMENT,\n"
	t := reflect.TypeOf(LoggerMessage{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbFieldName := strings.ToLower(field.Name)
		if dbFieldName == "timestamp" {
			createTableStmt += fmt.Sprintf("\t%s DATE,\n", dbFieldName)
		} else {
			createTableStmt += fmt.Sprintf("\t%s TEXT,\n", dbFieldName)
		}
	}
	createTableStmt += "\rcreationrelativetime DATETIME\n"
	createTableStmt = PyHelpers.SubstringRunes(createTableStmt, 0, len(createTableStmt)-3)
	createTableStmt += ");\n"
	createTableStmt += fmt.Sprintf("CREATE INDEX %s_timestamp_idx ON %s (timestamp);\n", tableName, tableName)
	createTableStmt += fmt.Sprintf("CREATE INDEX %s_level_idx ON %s (level);\n", tableName, tableName)
	createTableStmt = "COMMIT;"
	return createTableStmt
}

// /////////////////////////////////
// Ncapnp proto
type LoggerNcapSerializer struct {
	Name              string
	parentClassConfig *Config.Config
	loggerMessage     *loggerMsg.LoggerMsg
	memSeg            *capnplib.Segment
	msgSerDeSer       *capnplib.Message
}

func NewLoggerSerializer(name string, parentClassConfig *Config.Config) *LoggerNcapSerializer {
	capnplibMsg, memSeg, err := capnplib.NewMessage(capnplib.SingleSegment(nil))
	if err != nil {
		panic(fmt.Sprintf("Error while trying to initialize Logger serializer :'%v'\n", err))
	}

	loggerObj, err := loggerMsg.NewRootLoggerMsg(memSeg)
	if err != nil {
		panic(fmt.Sprintf("Error while trying to initialize Logger serializer :'%v'\n", err))
	}
	return &LoggerNcapSerializer{Name: name, parentClassConfig: parentClassConfig, memSeg: memSeg, loggerMessage: &loggerObj, msgSerDeSer: capnplibMsg}
}

func (loggerNcapSerializer *LoggerNcapSerializer) LoggerNcapSerialize(loggerMessage *LoggerMessage) []byte {
	loggerNcapSerializer.loggerMessage.SetTimestamp(loggerMessage.Timestamp)
	loggerNcapSerializer.loggerMessage.SetHostname(loggerMessage.Hostname)
	loggerNcapSerializer.loggerMessage.SetLoggerName(loggerMessage.LoggerName)
	loggerNcapSerializer.loggerMessage.SetModule(loggerMessage.Module)
	// the proto capnp used the proto format for enum (0, 1, 2, 3, 4, 5)
	loggerNcapSerializer.loggerMessage.SetLevel(loggerMsg.Level(LogLevelToInt[loggerMessage.Level] / 10))
	loggerNcapSerializer.loggerMessage.SetFilename(loggerMessage.Filename)
	loggerNcapSerializer.loggerMessage.SetFunctionName(loggerMessage.FunctionName)
	loggerNcapSerializer.loggerMessage.SetLineNumber(loggerMessage.LineNumber)
	loggerNcapSerializer.loggerMessage.SetMessage_(loggerMessage.Message)
	// extra log information
	loggerNcapSerializer.loggerMessage.SetPathName(loggerMessage.PathName)
	loggerNcapSerializer.loggerMessage.SetProcessId(loggerMessage.ProcessId)
	loggerNcapSerializer.loggerMessage.SetProcessName(loggerMessage.ProcessName)
	loggerNcapSerializer.loggerMessage.SetThreadId(loggerMessage.ThreadId)
	loggerNcapSerializer.loggerMessage.SetThreadName(loggerMessage.ThreadName)
	loggerNcapSerializer.loggerMessage.SetServiceName(loggerMessage.ServiceName)
	loggerNcapSerializer.loggerMessage.SetStackTrace(loggerMessage.StackTrace)
	byteMsg, _ := loggerNcapSerializer.msgSerDeSer.MarshalPacked()
	return byteMsg
}

func (loggerNcapSerializer *LoggerNcapSerializer) LoggerNcapDeSerialize(data []byte) *LoggerMessage {
	capnpMessage, _ := capnplib.UnmarshalPacked(data)
	goObj, _ := loggerMsg.ReadRootLoggerMsg(capnpMessage)
	loggerMessage := &LoggerMessage{}
	loggerMessage.Timestamp, _ = goObj.Timestamp()
	loggerMessage.Hostname, _ = goObj.Hostname()
	loggerMessage.LoggerName, _ = goObj.LoggerName()
	loggerMessage.Module, _ = goObj.Module()
	// the proto capnp used the proto format for enum (0, 1, 2, 3, 4, 5)
	loggerMessage.Level = strings.ToUpper(goObj.Level().String())
	loggerMessage.Filename, _ = goObj.Filename()
	loggerMessage.FunctionName, _ = goObj.FunctionName()
	loggerMessage.LineNumber, _ = goObj.LineNumber()
	loggerMessage.Message, _ = goObj.Message_()
	// extra log information
	loggerMessage.PathName, _ = goObj.PathName()
	loggerMessage.ProcessId, _ = goObj.ProcessId()
	loggerMessage.ProcessName, _ = goObj.ProcessName()
	loggerMessage.ThreadId, _ = goObj.ThreadId()
	loggerMessage.ThreadName, _ = goObj.ThreadName()
	loggerMessage.ServiceName, _ = goObj.ServiceName()
	loggerMessage.StackTrace, _ = goObj.StackTrace()
	return loggerMessage
}

// Ncapnp proto
///////////////////////////////////
