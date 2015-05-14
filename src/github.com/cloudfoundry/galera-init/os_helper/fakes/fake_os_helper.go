// This file was generated by counterfeiter
package fakes

import (
	"os/exec"
	"sync"
	"time"

	"github.com/cloudfoundry/mariadb_ctrl/os_helper"
)

type FakeOsHelper struct {
	RunCommandStub        func(executable string, args ...string) (string, error)
	runCommandMutex       sync.RWMutex
	runCommandArgsForCall []struct {
		executable string
		args       []string
	}
	runCommandReturns struct {
		result1 string
		result2 error
	}
	RunCommandWithTimeoutStub        func(timeout int, logFileName string, executable string, args ...string) error
	runCommandWithTimeoutMutex       sync.RWMutex
	runCommandWithTimeoutArgsForCall []struct {
		timeout     int
		logFileName string
		executable  string
		args        []string
	}
	runCommandWithTimeoutReturns struct {
		result1 error
	}
	StartCommandStub        func(logFileName string, executable string, args ...string) (*exec.Cmd, error)
	startCommandMutex       sync.RWMutex
	startCommandArgsForCall []struct {
		logFileName string
		executable  string
		args        []string
	}
	startCommandReturns struct {
		result1 *exec.Cmd
		result2 error
	}
	FileExistsStub        func(filename string) bool
	fileExistsMutex       sync.RWMutex
	fileExistsArgsForCall []struct {
		filename string
	}
	fileExistsReturns struct {
		result1 bool
	}
	ReadFileStub        func(filename string) (string, error)
	readFileMutex       sync.RWMutex
	readFileArgsForCall []struct {
		filename string
	}
	readFileReturns struct {
		result1 string
		result2 error
	}
	WriteStringToFileStub        func(filename string, contents string) error
	writeStringToFileMutex       sync.RWMutex
	writeStringToFileArgsForCall []struct {
		filename string
		contents string
	}
	writeStringToFileReturns struct {
		result1 error
	}
	SleepStub        func(duration time.Duration)
	sleepMutex       sync.RWMutex
	sleepArgsForCall []struct {
		duration time.Duration
	}
}

func (fake *FakeOsHelper) RunCommand(executable string, args ...string) (string, error) {
	fake.runCommandMutex.Lock()
	fake.runCommandArgsForCall = append(fake.runCommandArgsForCall, struct {
		executable string
		args       []string
	}{executable, args})
	fake.runCommandMutex.Unlock()
	if fake.RunCommandStub != nil {
		return fake.RunCommandStub(executable, args...)
	} else {
		return fake.runCommandReturns.result1, fake.runCommandReturns.result2
	}
}

func (fake *FakeOsHelper) RunCommandCallCount() int {
	fake.runCommandMutex.RLock()
	defer fake.runCommandMutex.RUnlock()
	return len(fake.runCommandArgsForCall)
}

func (fake *FakeOsHelper) RunCommandArgsForCall(i int) (string, []string) {
	fake.runCommandMutex.RLock()
	defer fake.runCommandMutex.RUnlock()
	return fake.runCommandArgsForCall[i].executable, fake.runCommandArgsForCall[i].args
}

func (fake *FakeOsHelper) RunCommandReturns(result1 string, result2 error) {
	fake.RunCommandStub = nil
	fake.runCommandReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) RunCommandWithTimeout(timeout int, logFileName string, executable string, args ...string) error {
	fake.runCommandWithTimeoutMutex.Lock()
	fake.runCommandWithTimeoutArgsForCall = append(fake.runCommandWithTimeoutArgsForCall, struct {
		timeout     int
		logFileName string
		executable  string
		args        []string
	}{timeout, logFileName, executable, args})
	fake.runCommandWithTimeoutMutex.Unlock()
	if fake.RunCommandWithTimeoutStub != nil {
		return fake.RunCommandWithTimeoutStub(timeout, logFileName, executable, args...)
	} else {
		return fake.runCommandWithTimeoutReturns.result1
	}
}

func (fake *FakeOsHelper) RunCommandWithTimeoutCallCount() int {
	fake.runCommandWithTimeoutMutex.RLock()
	defer fake.runCommandWithTimeoutMutex.RUnlock()
	return len(fake.runCommandWithTimeoutArgsForCall)
}

func (fake *FakeOsHelper) RunCommandWithTimeoutArgsForCall(i int) (int, string, string, []string) {
	fake.runCommandWithTimeoutMutex.RLock()
	defer fake.runCommandWithTimeoutMutex.RUnlock()
	return fake.runCommandWithTimeoutArgsForCall[i].timeout, fake.runCommandWithTimeoutArgsForCall[i].logFileName, fake.runCommandWithTimeoutArgsForCall[i].executable, fake.runCommandWithTimeoutArgsForCall[i].args
}

func (fake *FakeOsHelper) RunCommandWithTimeoutReturns(result1 error) {
	fake.RunCommandWithTimeoutStub = nil
	fake.runCommandWithTimeoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) StartCommand(logFileName string, executable string, args ...string) (*exec.Cmd, error) {
	fake.startCommandMutex.Lock()
	fake.startCommandArgsForCall = append(fake.startCommandArgsForCall, struct {
		logFileName string
		executable  string
		args        []string
	}{logFileName, executable, args})
	fake.startCommandMutex.Unlock()
	if fake.StartCommandStub != nil {
		return fake.StartCommandStub(logFileName, executable, args...)
	} else {
		return fake.startCommandReturns.result1, fake.startCommandReturns.result2
	}
}

func (fake *FakeOsHelper) StartCommandCallCount() int {
	fake.startCommandMutex.RLock()
	defer fake.startCommandMutex.RUnlock()
	return len(fake.startCommandArgsForCall)
}

func (fake *FakeOsHelper) StartCommandArgsForCall(i int) (string, string, []string) {
	fake.startCommandMutex.RLock()
	defer fake.startCommandMutex.RUnlock()
	return fake.startCommandArgsForCall[i].logFileName, fake.startCommandArgsForCall[i].executable, fake.startCommandArgsForCall[i].args
}

func (fake *FakeOsHelper) StartCommandReturns(result1 *exec.Cmd, result2 error) {
	fake.StartCommandStub = nil
	fake.startCommandReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) FileExists(filename string) bool {
	fake.fileExistsMutex.Lock()
	fake.fileExistsArgsForCall = append(fake.fileExistsArgsForCall, struct {
		filename string
	}{filename})
	fake.fileExistsMutex.Unlock()
	if fake.FileExistsStub != nil {
		return fake.FileExistsStub(filename)
	} else {
		return fake.fileExistsReturns.result1
	}
}

func (fake *FakeOsHelper) FileExistsCallCount() int {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return len(fake.fileExistsArgsForCall)
}

func (fake *FakeOsHelper) FileExistsArgsForCall(i int) string {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return fake.fileExistsArgsForCall[i].filename
}

func (fake *FakeOsHelper) FileExistsReturns(result1 bool) {
	fake.FileExistsStub = nil
	fake.fileExistsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOsHelper) ReadFile(filename string) (string, error) {
	fake.readFileMutex.Lock()
	fake.readFileArgsForCall = append(fake.readFileArgsForCall, struct {
		filename string
	}{filename})
	fake.readFileMutex.Unlock()
	if fake.ReadFileStub != nil {
		return fake.ReadFileStub(filename)
	} else {
		return fake.readFileReturns.result1, fake.readFileReturns.result2
	}
}

func (fake *FakeOsHelper) ReadFileCallCount() int {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return len(fake.readFileArgsForCall)
}

func (fake *FakeOsHelper) ReadFileArgsForCall(i int) string {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return fake.readFileArgsForCall[i].filename
}

func (fake *FakeOsHelper) ReadFileReturns(result1 string, result2 error) {
	fake.ReadFileStub = nil
	fake.readFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsHelper) WriteStringToFile(filename string, contents string) error {
	fake.writeStringToFileMutex.Lock()
	fake.writeStringToFileArgsForCall = append(fake.writeStringToFileArgsForCall, struct {
		filename string
		contents string
	}{filename, contents})
	fake.writeStringToFileMutex.Unlock()
	if fake.WriteStringToFileStub != nil {
		return fake.WriteStringToFileStub(filename, contents)
	} else {
		return fake.writeStringToFileReturns.result1
	}
}

func (fake *FakeOsHelper) WriteStringToFileCallCount() int {
	fake.writeStringToFileMutex.RLock()
	defer fake.writeStringToFileMutex.RUnlock()
	return len(fake.writeStringToFileArgsForCall)
}

func (fake *FakeOsHelper) WriteStringToFileArgsForCall(i int) (string, string) {
	fake.writeStringToFileMutex.RLock()
	defer fake.writeStringToFileMutex.RUnlock()
	return fake.writeStringToFileArgsForCall[i].filename, fake.writeStringToFileArgsForCall[i].contents
}

func (fake *FakeOsHelper) WriteStringToFileReturns(result1 error) {
	fake.WriteStringToFileStub = nil
	fake.writeStringToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsHelper) Sleep(duration time.Duration) {
	fake.sleepMutex.Lock()
	fake.sleepArgsForCall = append(fake.sleepArgsForCall, struct {
		duration time.Duration
	}{duration})
	fake.sleepMutex.Unlock()
	if fake.SleepStub != nil {
		fake.SleepStub(duration)
	}
}

func (fake *FakeOsHelper) SleepCallCount() int {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return len(fake.sleepArgsForCall)
}

func (fake *FakeOsHelper) SleepArgsForCall(i int) time.Duration {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.sleepArgsForCall[i].duration
}

var _ os_helper.OsHelper = new(FakeOsHelper)
