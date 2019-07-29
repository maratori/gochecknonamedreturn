package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainNoArgs(t *testing.T) {
	if isTest() {
		main()
		return
	}
	stdOut, stdErr, err := runTest(t)
	assert.Empty(t, stdOut)
	assert.Contains(t, stdErr, "gochecknonamedreturn: report usage of functions with named return value")
	assert.Error(t, err)
	require.IsType(t, &exec.ExitError{}, err)
	assert.Equal(t, 1, err.(*exec.ExitError).ExitCode())
}

func TestMainCheckTestData(t *testing.T) {
	if isTest() {
		main()
		return
	}
	stdOut, stdErr, err := runTest(t, "./gochecknonamedreturn/testdata")
	assert.Empty(t, stdOut)
	assert.Equal(t, 4, strings.Count(stdErr, "\n"), "Wrong number of lines")
	assert.Contains(t, stdErr, "testdata/named_return_function_declaration.go:3:27: don't use named return values")
	assert.Contains(t, stdErr, "testdata/named_return_function_declaration.go:7:37: don't use named return values")
	assert.Contains(t, stdErr, "testdata/named_return_function_declaration.go:11:35: don't use named return values")
	assert.Contains(t, stdErr, "testdata/named_return_function_declaration.go:15:45: don't use named return values")
	assert.Error(t, err)
	require.IsType(t, &exec.ExitError{}, err)
	assert.Equal(t, 3, err.(*exec.ExitError).ExitCode())
}

func TestMainCheckMain(t *testing.T) {
	if isTest() {
		main()
		return
	}
	stdOut, stdErr, err := runTest(t, ".")
	assert.NoError(t, err)
	assert.Empty(t, stdOut)
	assert.Empty(t, stdErr)
}

func isTest() bool {
	return os.Getenv("RUN_TEST") == "1"
}

func runTest(t *testing.T, args ...string) (string, string, error) {
	allArgs := []string{"-test.run=" + t.Name()}
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.coverprofile") {
			// TODO: coverage does not work if test uses os.Exit()
			//  because test runner does not finalize
			allArgs = append(allArgs, "-test.coverprofile=coverage"+t.Name()+".out")
			break
		}
	}
	allArgs = append(allArgs, args...)
	cmd := exec.Command(os.Args[0], allArgs...)
	cmd.Env = []string{"RUN_TEST=1"}
	for _, env := range []string{"HOME", "PATH"} {
		if val, ok := os.LookupEnv(env); ok {
			cmd.Env = append(cmd.Env, env+"="+val)
		}
	}
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	return outBuf.String(), errBuf.String(), err
}
