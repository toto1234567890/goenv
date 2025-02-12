package Helpers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

var NotPython string = "!=python"

// ####################################################################
// ############ Start independant process (not a child) ###############
// # https://stackoverflow.com/questions/13243807/popen-waiting-for-child-process-even-when-the-immediate-child-has-terminated/13256908#13256908
func StartIndependantProcess(command string, args ...string) error {
	if runtime.GOOS == "windows" {
		return StartIndependantProcessSpecific(command, args...)
	} else if runtime.GOOS == "linux" {
		return StartIndependantProcessSpecific(command, args...)
	} else if runtime.GOOS == "darwin" {
		return StartIndependantProcessSpecific(command, args...)
	}
	// check python executable and caffeinate..
	return errors.New(fmt.Sprintf("unable to start process, undefine OS type : %s", runtime.GOOS))
}

// ############ Start independant process (not a child) ###############
// ####################################################################
// ########## Check if process/service is running by name #############
func IsProcessRunning(cmdlinePatt string, processName string, argvPatt string, getPid bool, exceptThisPid int32) (bool, int32) {
	procs, err := process.Processes()
	if err != nil {
		goto err
	}
	for _, proc := range procs {
		name, err := proc.Name()
		if err != nil {
			goto err
		}
		if strings.ToLower(processName) == strings.ToLower(name) {
			cmdline, err := proc.Cmdline()
			if err != nil {
				goto err
			}
			cmdline = strings.ToLower(cmdline)
			if strings.Contains(cmdline, strings.ToLower(cmdlinePatt)) {
				if argvPatt != "" {
					if strings.Contains(cmdline, strings.ToLower(argvPatt)) {
						if exceptThisPid != proc.Pid {
							return true, proc.Pid
						}
					}
				} else {
					return true, proc.Pid
				}
			}
		}
	}
	return false, 0
err:
	return false, -9999 // error ??
}

// ########## Check if process/service is running by name #############
// ####################################################################
// ############# Get current script executed directory ################
func GetExecutedScriptDir() string {
	selfExePath, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("Unable to get executable parent directory: '%v'\n", err))
	}
	return filepath.Dir(selfExePath)
}

// ############# Get current script executed directory ################
// ####################################################################
// ####################### Load default args ##########################
func LoadDefaultArgs(port string, conf string, host string, name string, log_level string) []string {
	args := []string{}
	if port != "" {
		args = append(args, "--port", port)
	}
	if conf != "" {
		args = append(args, "--conf", conf)
	}
	if name != "" {
		args = append(args, "--name", name)
	}
	if host != "" {
		args = append(args, "--host", host)
	}
	if log_level != "" {
		args = append(args, "--log_level", log_level)
	}
	return args
}

// ####################### Load default args ##########################
// ####################################################################
// ####################### Get conf section ###########################
func getConfSec(conf string) string {
	if conf != "" {
		if strings.ContainsRune(conf, filepath.Separator) {
			config_files_map := LoadConfigFiles("", []string{}, "")
			for key, val := range config_files_map {
				if key == conf {
					return val
				}
			}
		}
	}
	return "error you must provide a valid conf section..."
}

//####################### Get conf section ###########################
// ####################################################################
