//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v14/fsutil"
	"github.com/essentialkaos/ek/v14/mathutil"
	"github.com/essentialkaos/ek/v14/system"
)

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

	rootDir := path.Join(procFS, strconv.Itoa(root))

	if !fsutil.IsExist(rootDir) {
		return nil, fmt.Errorf("process with PID %d doesn't exist", root)
	}

	list, err := findInfo(procFS, make(map[int]string))

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("can't find any processes")
	}

	return processListToTree(list, root), nil
}

// GetList returns slice with all active processes on the system
func GetList() ([]*ProcessInfo, error) {
	return findInfo(procFS, make(map[int]string))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getFileScanner opens file and creates scanner for reading text files line by line
func getFileScanner(file string) (*bufio.Scanner, func() error, error) {
	fd, err := os.Open(file)

	if err != nil {
		return nil, nil, err
	}

	s := bufio.NewScanner(fd)

	return s, fd.Close, nil
}

// findInfo recursively searches for process info in the specified directory
func findInfo(dir string, userMap map[int]string) ([]*ProcessInfo, error) {
	var result []*ProcessInfo

	dirs := fsutil.List(dir, true, fsutil.ListingFilter{Perms: "DRX"})

	for _, pid := range dirs {
		if !isPID(pid) {
			continue
		}

		taskDir := path.Join(dir, pid, "task")

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

		info, err := readProcessInfo(path.Join(dir, pid), pid, userMap)

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

// readProcessInfo reads process info from the specified directory
func readProcessInfo(dir, pid string, userMap map[int]string) (*ProcessInfo, error) {
	pidInt, err := strconv.Atoi(pid)

	if err != nil {
		return nil, err
	}

	cmdlineFile := path.Join(dir, "cmdline")

	// The process had died after the moment when we have created a list of processes
	if !fsutil.IsExist(cmdlineFile) {
		return nil, nil
	}

	cmd, err := os.ReadFile(cmdlineFile)

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

// getProcessUser returns username by UID
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

// getProcessParent returns parent PID and true if process is thread
func getProcessParent(pidDir string, pid int) (int, bool) {
	tgid, ppid := getParentPIDs(pidDir)

	if tgid != pid {
		return tgid, true
	}

	return ppid, false
}

// getParentPIDs reads /proc/[pid]/status file and returns Tgid and PPid
func getParentPIDs(pidDir string) (int, int) {
	statusFile := path.Join(pidDir, "status")
	data, err := os.ReadFile(statusFile)

	if err != nil {
		return -1, -1
	}

	var ppid, tgid string

	for line := range strings.SplitSeq(string(data), "\n") {
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

// formatCommand formats command string by normalizing delimiters and trimming spaces
func formatCommand(cmd string) string {
	// Normalize delimiters
	command := strings.ReplaceAll(cmd, "\000", " ")

	// Remove space on the end of command
	command = strings.TrimSpace(command)

	return command
}

// processListToTree converts a flat list of processes into a tree structure
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

		parentProcess.Children = append(parentProcess.Children, process)
	}

	return result[root]
}

// isPID checks if the given string is a valid PID
func isPID(pid string) bool {
	return pid != "" && mathutil.IsInt(pid)
}
