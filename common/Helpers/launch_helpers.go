package Helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"govenv/pkg/common/PyHelpers"
)

// ###################################################################
// ############## Launch CONFIG server if not exist... ###############
func LaunchConfigServer(name string, curDir string, host string, port string, conf string, logLevel string) error {
	if curDir == "" {
		curDir, _ = os.Getwd()
	}
	cmdline := fmt.Sprintf("%s", filepath.Join(curDir, "common", "Config", "config_server.exe"))
	_conf := getConfSec(conf)
	args := LoadDefaultArgs(port, _conf, host, name, logLevel)
	return StartIndependantProcess(cmdline, args...)
}

// ############## Launch CONFIG server if not exist... ###############
// ###################################################################
// ############## Launch NOTIF server if not exist... ################
func LaunchNotifServer(name string, curDir string, host string, port string, conf string, logLevel string) error {
	if curDir == "" {
		curDir, _ = os.Getwd()
	}
	cmdline := fmt.Sprintf("%s", filepath.Join(curDir, "common", "Notifie", "notif_server.exe"))
	_conf := getConfSec(conf)
	args := LoadDefaultArgs(port, _conf, host, name, logLevel)
	return StartIndependantProcess(cmdline, args...)
}

// ############## Launch NOTIF server if not exist... ################
// ###################################################################
// ############### Launch LOG server if not exist... #################
func LaunchLogServer(name string, curDir string, host string, port string, conf string, logLevel string) error {
	if curDir == "" {
		curDir, _ = os.Getwd()
	}
	cmdLine := fmt.Sprintf("%s", filepath.Join(curDir, "common", "MyLogger", "log_server.exe"))
	_conf := getConfSec(conf)
	args := LoadDefaultArgs(port, _conf, host, name, logLevel)
	return StartIndependantProcess(cmdLine, args...)
}

// ############### Launch LOG server if not exist... ##################
// ####################################################################
// ############ Launch TELEREMOTE server if not exist... ##############
func LaunchTeleremoteServer(name string, curDir string, host string, port string, conf string, logLevel string, onlyLogger string) error {
	if curDir == "" {
		curDir, _ = os.Getwd()
	}
	cmdLine := fmt.Sprintf("%s", filepath.Join(curDir, "common", "TeleRemote", "tele_remote.exe"))
	_conf := getConfSec(conf)
	args := LoadDefaultArgs(port, _conf, host, name, logLevel)
	if onlyLogger != "" {
		args = append(args, "--only_logger", PyHelpers.Capitalize(onlyLogger))
	}
	return StartIndependantProcess(cmdLine, args...)
}

//############ Launch TELEREMOTE server if not exist... ##############
