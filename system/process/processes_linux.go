//go:build linux
// +build linux

// Package process provides methods for gathering information about active processes
package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/fsutil"
	"github.com/essentialkaos/ek/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcessInfo contains basic info about process
type ProcessInfo struct {
	Command  string         // Full command
	User     string         // Username
	PID      int            // PID
	Parent   int            // Parent process PID
	Childs   []*ProcessInfo // Slice with child processes
	IsThread bool           // True if process is thread
}

// ////////////////////////////////////////////////////////////////////////////////// //

// procFS is path to procfs
var procFS = "/proc"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetTree returns root process with all subprocesses on the system
func GetTree(pid ...int) (*ProcessInfo, error) {
	root := 1

	if len(pid) != 0 {
		root = pid[0]
	}

	if !fsutil.IsExist(procFS + "/" + strconv.Itoa(root)) {
		return nil, fmt.Errorf("Process with PID %d doesn't exist", pid)
	}

	list, err := findInfo(procFS, make(map[int]string))

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("Can't find any processes")
	}

	return processListToTree(list, root), nil
}

// GetList returns slice with all active processes on the system
func GetList() ([]*ProcessInfo, error) {
	return findInfo(procFS, make(map[int]string))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func findInfo(dir string, userMap map[int]string) ([]*ProcessInfo, error) {
	var result []*ProcessInfo

	dirs := fsutil.List(dir, true, fsutil.ListingFilter{Perms: "DRX"})

	for _, pid := range dirs {
		if !isPID(pid) {
			continue
		}

		taskDir := dir + "/" + pid + "/task"

		if fsutil.IsExist(taskDir) {
			threads, err := findInfo(taskDir, userMap)

			if err != nil {
				return nil, err
			}

			if len(threads) == 0 {
				continue
			}

			result = append(result, threads...)

			continue
		}

		info, err := readProcessInfo(dir+"/"+pid, pid, userMap)

		if err != nil {
			return nil, err
		}

		if info == nil {
			continue
		}

		result = append(result, info)
	}

	return result, nil
}

func readProcessInfo(dir, pid string, userMap map[int]string) (*ProcessInfo, error) {
	pidInt, err := strconv.Atoi(pid)

	if err != nil {
		return nil, err
	}

	// The process had died after the moment when we have created a list of processes
	if !fsutil.IsExist(dir + "/cmdline") {
		return nil, nil
	}

	cmd, err := ioutil.ReadFile(dir + "/cmdline")

	if err != nil {
		return nil, err
	}

	if len(cmd) == 0 {
		return nil, nil
	}

	uid, _, err := fsutil.GetOwner(dir)

	if err != nil {
		return nil, err
	}

	username, err := getProcessUser(uid, userMap)

	if err != nil {
		return nil, err
	}

	ppid, isThread := getProcessParent(dir, pidInt)

	return &ProcessInfo{
		Command:  formatCommand(string(cmd)),
		User:     username,
		PID:      pidInt,
		Parent:   ppid,
		IsThread: isThread,
	}, nil
}

func getProcessUser(uid int, userMap map[int]string) (string, error) {
	if uid == 0 {
		return "root", nil
	}

	if userMap[uid] != "" {
		return userMap[uid], nil
	}

	user, err := system.LookupUser(strconv.Itoa(uid))

	if err != nil {
		return "", err
	}

	userMap[uid] = user.Name

	return user.Name, nil
}

func getProcessParent(pidDir string, pid int) (int, bool) {
	tgid, ppid := getParentPIDs(pidDir)

	if tgid != pid {
		return tgid, true
	}

	return ppid, false
}

func getParentPIDs(pidDir string) (int, int) {
	data, err := ioutil.ReadFile(pidDir + "/status")

	if err != nil {
		return -1, -1
	}

	var ppid, tgid string

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "Tgid:") {
			tgid = strings.TrimSpace(line[5:])
		}

		if strings.HasPrefix(line, "PPid:") {
			ppid = strings.TrimSpace(line[5:])
		}

		if ppid != "" && tgid != "" {
			break
		}
	}

	if tgid == "" || ppid == "" {
		return -1, -1
	}

	tgidInt, tgidErr := strconv.Atoi(tgid)
	ppidInt, ppidErr := strconv.Atoi(ppid)

	if tgidErr != nil || ppidErr != nil {
		return -1, -1
	}

	return tgidInt, ppidInt
}

func formatCommand(cmd string) string {
	// Normalize delimiters
	command := strings.Replace(cmd, "\000", " ", -1)

	// Remove space on the end of command
	command = strings.TrimSpace(command)

	return command
}

func processListToTree(processes []*ProcessInfo, root int) *ProcessInfo {
	var result = make(map[int]*ProcessInfo)

	for _, info := range processes {
		result[info.PID] = info
	}

	for _, process := range result {
		if process.Parent < 0 {
			continue
		}

		parentProcess := result[process.Parent]

		if parentProcess == nil {
			continue
		}

		parentProcess.Childs = append(parentProcess.Childs, process)
	}

	return result[root]
}

func isPID(pid string) bool {
	if pid == "" {
		return false
	}

	// Pid must start from number
	switch pid[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}

	return false
}
