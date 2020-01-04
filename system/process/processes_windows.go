// +build windows

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
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

// GetTree returns root process with all subprocesses on the system
func GetTree(pid ...int) (*ProcessInfo, error) {
	return nil, nil
}

// GetList returns slice with all active processes on the system
func GetList() ([]*ProcessInfo, error) {
	return nil, nil
}
