// +build windows

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

// GetTree return root process with all subprocesses on system
func GetTree() (*ProcessInfo, error) {
	return nil, nil
}

// GetList return slice with all active processes on system
func GetList() ([]*ProcessInfo, error) {
	return nil, nil
}
