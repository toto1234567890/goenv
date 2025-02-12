package MyLogger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"govenv/pkg/common/Config"
	"govenv/pkg/common/Helpers"
	"govenv/pkg/common/PyHelpers"
)

/////////////////////////////////
// configure logger
// FIXME should also start backendDb service if not running....

type Logger struct {
	Name                string
	Level               int8
	config              *Config.Config
	LogRecord           *LoggerMessage
	FileWriter          *os.File
	withLogServer       bool
	LoggerSocket        *Helpers.SafeSocket
	LogRecordSerializer *LoggerNcapSerializer
	LogFileChan         chan []byte
}

func NewLogger(name string, level string, logDir string, config *Config.Config, withLogServer bool, serviceName string) *Logger {
	// set level
	intLevel := LogLevelToInt[strings.ToUpper(level)]

	// Create local log file and folder if not exists
	// file handler closed in parent
	logFile, err := PyHelpers.CreateFile(false, filepath.Join(logDir, fmt.Sprintf("%s.log", name)), []byte(""), true)
	if err != nil {
		panic(fmt.Sprintf("Unable to create log file or folder : '%v'\n", err))
	}

	// create logger object
	logger := &Logger{Name: name, Level: intLevel, config: config, FileWriter: logFile, LogRecord: NewLoggerMessage(name, serviceName), LogFileChan: make(chan []byte)}
	go logger.writeLogFile()

	// init socket and serializer if withLogServer
	logger.withLogServer = false
	if withLogServer {
		logger.withLogServer = true
		logger.LogRecordSerializer = NewLoggerSerializer(logger.Name, config)
		logger.LoggerSocket = Helpers.MySocket(logger.Name, config.LG_IP, config.LG_PORT, 1)
	}

	return logger
}

func (logger *Logger) IsLevelEnabled(level int8) bool {
	if level+1 > logger.Level {
		return true
	}
	return false
}

func (logger *Logger) GetCallerInfo(depth int) (filePath, module, file, function, line string) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	pc, fullFilePath, _line, ok := runtime.Caller(depth)
	if !ok {
		return "unknown", "unknown", "unknown", "unknown", "0"
	}
	file = filepath.Base(fullFilePath)
	filePath = filepath.Dir(fullFilePath)
	fullFuncName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndex(fullFuncName, "/")
	lastDot := strings.LastIndex(fullFuncName, ".")
	if lastDot > lastSlash {
		module = fullFuncName[:lastDot]
		function = fullFuncName[lastDot+1:]
	} else {
		module = "unknown"
		function = fullFuncName
	}

	return filePath, module, file, function, strconv.Itoa(_line)
}

func (logger *Logger) writeStdOut(message []byte) {
	os.Stdout.Write(message)
}

func (logger *Logger) writeLogFile() {
	for {
		_, err := logger.FileWriter.Write(<-logger.LogFileChan)
		if err != nil {
			println(fmt.Printf("%v", err))
		}
	}
}

func truncateIfNeeded(s string, lenght int) string {
	if len(s) <= lenght {
		return s
	}
	return s[:lenght]
}

func (logger *Logger) LogWriterMessage(Timestamp, Level, Filename, FunctionName, LineNumber, Message, StackTrace string) []byte {
	var strBuild bytes.Buffer
	strBuild.WriteString(Timestamp)
	strBuild.WriteString(strings.Repeat(" ", 27-len(truncateIfNeeded(Timestamp, 26))))
	strBuild.WriteString(Level)
	strBuild.WriteString(strings.Repeat(" ", 11-len(truncateIfNeeded(Level, 10))))
	strBuild.WriteString(Filename)
	strBuild.WriteString(strings.Repeat(" ", 24-len(truncateIfNeeded(Filename, 23))))
	strBuild.WriteString(FunctionName)
	strBuild.WriteString(strings.Repeat(" ", 24-len(truncateIfNeeded(FunctionName, 23))))
	strBuild.WriteString(LineNumber)
	strBuild.WriteString(strings.Repeat(" ", 6-len(truncateIfNeeded(LineNumber, 5))))
	strBuild.WriteString(Message + "\n")
	if StackTrace != "" {
		strBuild.WriteString(StackTrace + "\n")
	}
	return []byte(strBuild.String())
}

func (logger *Logger) createLogRecord(level int8, msg, stackTrace string) {
	filePath, module, file, function, line := logger.GetCallerInfo(3)
	logRecord := logger.LogRecord.createNewMessage(level, msg, filePath, module, file, function, line, stackTrace)
	byteMsg := logger.LogWriterMessage(logRecord.Timestamp, logRecord.Level, logRecord.Filename, logRecord.FunctionName, logRecord.LineNumber, logRecord.Message, logRecord.StackTrace)
	logger.writeStdOut(byteMsg)
	logger.LogFileChan <- byteMsg
	if logger.withLogServer {
		go logger.LoggerSocket.SendData(logger.LogRecordSerializer.LoggerNcapSerialize(logRecord))
	}
}
