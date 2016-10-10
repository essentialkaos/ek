// +build linux

// Package process provides methods for getting information about active processes
package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"pkg.re/essentialkaos/ek.v5/fsutil"
	"pkg.re/essentialkaos/ek.v5/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcessInfo contains basic info about process
type ProcessInfo struct {
	Command  string         // Full command
	User     string         // Username
	PID      int            // PID
	IsThread bool           // True if process is thread
	Parent   int            // Parent process PID
	Childs   []*ProcessInfo // Slice with child processes
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetTree return root process with all subprocesses on system
func GetTree() (*ProcessInfo, error) {
	list, err := findInfo("/proc", make(map[int]string))

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errors.New("Can't find any processes")
	}

	processMap := processListToMap(list)

	for _, process := range processMap {
		if process.Parent < 0 {
			continue
		}

		parentProcess := processMap[process.Parent]

		if parentProcess == nil {
			continue
		}

		parentProcess.Childs = append(parentProcess.Childs, process)
	}

	return processMap[1], nil
}

// GetList return slice with all active processes on system
func GetList() ([]*ProcessInfo, error) {
	return findInfo("/proc", make(map[int]string))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func findInfo(dir string, userMap map[int]string) ([]*ProcessInfo, error) {
	var result []*ProcessInfo

	dirs := fsutil.List(dir, true, &fsutil.ListingFilter{Perms: "DRX"})

	for _, pidDir := range dirs {
		if !isPID(pidDir) {
			continue
		}

		taskDir := dir + "/" + pidDir + "/task"

		if fsutil.IsExist(taskDir) {
			threads, err := findInfo(taskDir, userMap)

			if err != nil {
				return nil, err
			}

			if len(threads) == 0 {
				continue
			}

			processThreads(threads)

			result = append(result, threads...)

			continue
		}

		info, err := readProcessInfo(dir+"/"+pidDir, pidDir, userMap)

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
	cmd, err := ioutil.ReadFile(dir + "/cmdline")

	if len(cmd) == 0 {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	uid, _, err := fsutil.GetOwner(dir)

	if err != nil {
		return nil, err
	}

	username, err := getProcessUser(uid, userMap)

	if err != nil {
		return nil, err
	}

	pidInt, err := strconv.Atoi(pid)

	if err != nil {
		return nil, err
	}

	return &ProcessInfo{
		Command: formatCommand(string(cmd)),
		User:    username,
		PID:     pidInt,
		Parent:  getProcessParent(dir, pidInt),
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

func processThreads(threads []*ProcessInfo) {
	for _, info := range threads {
		info.IsThread = true
	}
}

func getProcessParent(pidDir string, pid int) int {
	tgid, ppid := getParentPIDs(pidDir)

	if tgid != pid {
		return tgid
	}

	return ppid
}

func getParentPIDs(pidDir string) (int, int) {
	data, err := ioutil.ReadFile(pidDir + "/status")

	if err != nil {
		return -1, -1
	}

	var (
		ppid string
		tgid string
	)

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

	tgidInt, _ := strconv.Atoi(tgid)
	ppidInt, _ := strconv.Atoi(ppid)

	return tgidInt, ppidInt
}

func formatCommand(cmd string) string {
	// Normalize delimiters
	command := strings.Replace(cmd, "\000", " ", -1)

	// Remove space on the end of command
	command = strings.TrimSpace(command)

	return command
}

func processListToMap(processes []*ProcessInfo) map[int]*ProcessInfo {
	var result = make(map[int]*ProcessInfo)

	for _, info := range processes {
		result[info.PID] = info
	}

	return result
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
