package utils

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	var service = GetExecService()
	service.RunCmd("pwd", []string{})
}

func TestExecWithPwd(t *testing.T) {
	var service = GetExecService()
	res := service.RunCmd("pwd", []string{}, &exec.Cmd{Dir: "/"})

	assert.NotNil(t, res)
	assert.Nil(t, res.Error)
	assert.Equal(t, "/", strings.TrimSpace(string(res.Output)))
}

func TestExecWithError(t *testing.T) {
	var service = GetExecService()
	res := service.RunCmd("ls", []string{"non_existing_dir"})

	assert.NotNil(t, res)
	assert.NotNil(t, res.Error)
	assert.True(t, res.HasErrors)
	assert.NotZero(t, res.ExitCode)
}
