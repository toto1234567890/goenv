//go:build windows
// +build windows

package Helpers

import (
	"os/exec"
	"syscall"
)

// ####################################################################
// ############ Start independant process (not a child) ###############
// # https://stackoverflow.com/questions/13243807/popen-waiting-for-child-process-even-when-the-immediate-child-has-terminated/13256908#13256908
func StartIndependantProcessSpecific(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	//# for windows from msdn [1]
	var CREATE_NEW_PROCESS_GROUP uint32 = 0x00000200 //# note: could get it from subprocess
	var DETACHED_PROCESS uint32 = 0x00000008         //# 0x8 | 0x200 == 0x208
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: CREATE_NEW_PROCESS_GROUP | DETACHED_PROCESS,
	}
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd
}

//############ Start independant process (not a child) ###############
//####################################################################
