package MyLogger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"govenv/pkg/common/Config"
	"govenv/pkg/common/Helpers"
	"govenv/pkg/common/Notifie"
	"govenv/pkg/common/PyHelpers"
)

var loggerMu sync.Mutex

// ///////////////////////////////
// used to know from where the notification has been triggered, duplicated with createLogRecord, but easier and more natural (follow 'my logic')
var WHOAMI string = fmt.Sprintf("+++++++++++++ " + func() string { host, _ := os.Hostname(); return host }() + " +++++++++++++\n")

// configure logger
/////////////////////////////////
// main Logger

type MyLogger struct {
	Name            string
	config          *Config.Config
	Logger          *Logger
	withNotif       bool
	withNotifServer bool
	NotifChan       chan *Notifie.NotifMessage
	localNotifier   *Notifie.Notifie
	NotifLogLevel   map[string][]string
}

func NewMyLogger(name string, config *Config.Config, onlyLogger bool, withLogServer bool, withNotif bool, withNotifServer bool, withTeleCommand bool, logLevel string, logDir string, serviceName string) *MyLogger {
	Name := "MyLogger"
	if name == "" {
		Name = name
	}

	// logger setup
	if logDir == "" {
		curDir, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("Unable to get current working directory : '%v'\n", err))
		}
		logDir = filepath.Join(curDir, "logs")
	}

	if !PyHelpers.FileExists(logDir) {
		err := PyHelpers.MakeDirectoryIfNotExists(logDir)
		if err != nil {
			panic(fmt.Sprintf("Error while trying to create log folder '%s' : '%v'\n", logDir, err))
		}
	}

	myLogger := &MyLogger{
		Name:          Name,
		config:        config,
		Logger:        NewLogger(name, logLevel, logDir, config, withLogServer, serviceName),
		NotifLogLevel: map[string][]string{},
	}

	// if only_logger do not notify even if True
	myLogger.withNotif = false
	myLogger.withNotifServer = false

	// onlyLogger should be used for independant process or testing...
	if !onlyLogger {
		// prepare wait time use to wait for started service to listen
		wait, err := strconv.ParseFloat(config.MAIN_WAIT_BEAT, 64)
		if err != nil {
			wait = 0.1
		}

		// if not onlyLogger, start the log server, which also strat the DBerror if not running...
		if withLogServer {
			_logServer := "log_server"
			// check if log_server server already run
			running, _ := Helpers.IsProcessRunning("", _logServer, "", false, 0)
			if !running {
				Helpers.LaunchLogServer("", "", config.LG_IP, config.LG_PORT, config.COMMON_FILE_PATH, logLevel)
			}
			// wait until service listen on his port
			for {
				if Helpers.IsServiceListen(config.LG_IP, config.LG_PORT, 1, false) {
					break
				}
				time.Sleep(time.Duration(wait) * time.Second)
			}
			// FIXME log server should also start backendDb if not running
			myLogger.Info("Main Logger : Main log_server is starting.. .  . ")
		}

		if withNotif {
			// if notification enable, a callback is needed to update myLogger.NotifLogLevel map...
			myLogger.config.Handler.SetLoggerCallBack(myLogger.setNotifLogLevel)
		}

		if withNotif && withNotifServer {
			myLogger.withNotif = true
			myLogger.withNotifServer = true
			_notifServer := "notif_server"
			// check if notification server already run
			running, _ := Helpers.IsProcessRunning("", _notifServer, "", false, 0)
			if !running {
				Helpers.LaunchNotifServer("", "", config.NT_IP, config.NT_PORT, config.COMMON_FILE_PATH, logLevel)
			}
			for {
				// check if notif server is listening
				if Helpers.IsServiceListen(config.NT_IP, config.NT_PORT, 1, false) {
					break
				}
				time.Sleep(time.Duration(wait) * time.Second)
			}
			myLogger.Info("Main Notifier : Main notif_server is starting.. .  . ")

			myLogger.NotifChan = make(chan *Notifie.NotifMessage)
			go myLogger.backgroundNotifSender()
		}

		if withNotif && !withNotifServer {
			myLogger.withNotif = true
			myLogger.localNotifier = Notifie.NewNotifie(myLogger.config, myLogger.Name)
			myLogger.NotifChan = myLogger.localNotifier.NotifChan
		}

		if withTeleCommand {
			_teleRemote := "tele_remote"
			// check if tele_remote server already run
			running, _ := Helpers.IsProcessRunning("", _teleRemote, "", false, 0)
			if !running {
				// FIXME only logger ?!?
				Helpers.LaunchTeleremoteServer("", "", config.TB_IP, config.TB_PORT, config.COMMON_FILE_PATH, logLevel, strconv.FormatBool(onlyLogger))
			}
			for {
				if Helpers.IsServiceListen(config.NT_IP, config.NT_PORT, 1, false) {
					break
				}
				time.Sleep(time.Duration(wait) * time.Second)
			}
			myLogger.Info("Main Telecommand : Main tele_remote is starting.. .  . ")
		}
	}

	return myLogger
}

// close the file handler
func (myLogger *MyLogger) Stop() {
	defer myLogger.Logger.FileWriter.Close()
}

// main Logger
// ///////////////////////////////
// callback function for child process

func (myLogger *MyLogger) setNotifLogLevel(_NotiflogLevel map[string][]string) {
	loggerMu.Lock()
	myLogger.NotifLogLevel = _NotiflogLevel
	loggerMu.Unlock()
}

func (myLogger *MyLogger) Log(msg string, level int8) {
	switch level {
	case DEBUG:
		myLogger.Debug(msg)
	case INFO:
		myLogger.Info(msg)
	case WARNING:
		myLogger.Warning(msg)
	case ERROR:
		myLogger.Error(msg)
	case CRITICAL:
		myLogger.Critical(msg)
	default:
		// other log level should not be called here !!
		_msg := fmt.Sprintf("%s : wrong Log LEVEL callback triggered in logger : level -> %d, log msg -> %s", myLogger.Name, level, msg)
		myLogger.Error(_msg)
	}
}

// callback function for child process
/////////////////////////////////
// notification from notif server or locally

func (myLogger *MyLogger) notifie(message string, attachment string, tags []string) {
	if !myLogger.withNotif {
		return
	}
	myLogger.NotifChan <- &Notifie.NotifMessage{Message: message, Attachment: attachment, Tags: tags}
}

func (myLogger *MyLogger) backgroundNotifSender() {
	notifHandler := Notifie.NewNotifHandler(myLogger.Name, myLogger.config)
	notifSock := Helpers.MySocket(myLogger.Name, myLogger.config.NT_IP, myLogger.config.NT_PORT, 1)
	for {
		notifSock.SendData(notifHandler.NotifNcapSerialize(<-myLogger.NotifChan))
	}
}

// notification from server or locally
/////////////////////////////////
// std log method

func (myLogger *MyLogger) Debug(msg string) {
	if myLogger.Logger.IsLevelEnabled(DEBUG) {
		myLogger.Logger.createLogRecord(DEBUG, msg, "")
	}
}

func (myLogger *MyLogger) Info(msg string) {
	if myLogger.Logger.IsLevelEnabled(INFO) {
		myLogger.Logger.createLogRecord(INFO, msg, "")
	}
}

func (myLogger *MyLogger) Warning(msg string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrWARNING]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(WARNING, msg, "")
}

func (myLogger *MyLogger) Error(msg string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrERROR]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(ERROR, msg, "")
}

// std log method
/////////////////////////////////
// custom log method

func (myLogger *MyLogger) Stream(msg string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrSTREAM]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(STREAM, msg, "")
}

func (myLogger *MyLogger) Logon(msg string) {
	// FIXME replace sensitive data
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrLOGON]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(LOGON, msg, "")
}

func (myLogger *MyLogger) Logout(msg string) {
	// FIXME replace sensitive data
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrLOGOUT]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(LOGOUT, msg, "")
}

func (myLogger *MyLogger) Trade(msg string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrTRADE]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(TRADE, msg, "")
}

func (myLogger *MyLogger) Schedule(msg string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrSCHEDULE]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(SCHEDULE, msg, "")
}

func (myLogger *MyLogger) Report(msg string, attachment string) {
	if myLogger.withNotif {
		if Tags, ok := myLogger.NotifLogLevel[StrREPORT]; ok {
			myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: attachment, Tags: Tags}
		}
	}
	myLogger.Logger.createLogRecord(REPORT, msg, "")
}

// custom log method
/////////////////////////////////

// custom  log method
/////////////////////////////////
// if critical <-> panic <-> exit program <-> get stackStrace <-> wait for writers...

func (myLogger *MyLogger) waitForEmptyChan(wait float64) {
	for {
		if len(myLogger.NotifChan) == 0 && len(myLogger.Logger.LogFileChan) == 0 {
			break
		}
		time.Sleep(time.Duration(wait) * time.Second)
	}
	// wait for the sending
	time.Sleep(time.Duration(wait) * time.Second)
}

func (myLogger *MyLogger) Critical(msg string) {
	// get stack trace
	loggerMu.Lock()
	stackTrace := string(debug.Stack())
	loggerMu.Unlock()

	// create logger record with stacktrace
	myLogger.Logger.createLogRecord(CRITICAL, msg, stackTrace)

	// Fatal should close the program 'os.Exit(1)' -> Critical in python
	if Tags, ok := myLogger.NotifLogLevel[StrCRITICAL]; ok {
		myLogger.NotifChan <- &Notifie.NotifMessage{Message: msg, Attachment: "", Tags: Tags}
	}

	// all going wrong, can wait for the error sending...
	myLogger.waitForEmptyChan(1)
}

// if critical <-> panic <-> exit program <-> get stackStrace <-> wait for writers...
/////////////////////////////////
