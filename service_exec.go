package utils

import (
	"fmt"
	"os/exec"
)

type ExecResult struct {
	ExitCode  int
	Error     error
	Output    []byte
	HasErrors bool
}

type ExecService interface {
	RunCmd(program string, args []string, options ...*exec.Cmd) *ExecResult
}

type DefaultExecService struct {
}

func GetExecService() ExecService {
	return &DefaultExecService{}
}

func (*DefaultExecService) RunCmd(program string, args []string, options ...*exec.Cmd) *ExecResult {
	cmd := exec.Command(program, args...)
	if len(options) > 0 {
		opt := options[0]
		cmd.Dir = opt.Dir
		cmd.Env = opt.Env
		cmd.WaitDelay = opt.WaitDelay
	}

	output, err := cmd.CombinedOutput()
	success := cmd.ProcessState.Success()
	if err == nil && !success {
		err = fmt.Errorf("program exited with error code %d and output '%s'", 
			cmd.ProcessState.ExitCode(),
			string(output))
	}
	return &ExecResult{
		ExitCode:  cmd.ProcessState.ExitCode(),
		Error:     err,
		Output:    output,
		HasErrors: !success || err != nil,
	}
}
