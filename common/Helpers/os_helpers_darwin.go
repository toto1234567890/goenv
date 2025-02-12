//go:build darwin
// +build darwin

package Helpers

import (
	"os/exec"
	"syscall"
)

// ####################################################################
// ############ Start independant process (not a child) ###############
// # https://stackoverflow.com/questions/13243807/popen-waiting-for-child-process-even-when-the-immediate-child-has-terminated/13256908#13256908
func StartIndependantProcessSpecific(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	// for mac
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Start()
}

//############ Start independant process (not a child) ###############
//####################################################################
