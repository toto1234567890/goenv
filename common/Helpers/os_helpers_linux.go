//go:build linux
// +build linux

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
	// For Linux, we can use the session
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // Create process group
	}
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd
}

//############ Start independant process (not a child) ###############
//####################################################################
