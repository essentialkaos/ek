package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ProcessInfo contains basic info about process
type ProcessInfo struct {
	Command  string         // ❗ Full command
	User     string         // ❗ Username
	PID      int            // ❗ PID
	IsThread bool           // ❗ True if process is thread
	Parent   int            // ❗ Parent process PID
	Childs   []*ProcessInfo // ❗ Slice with child processes
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ GetTree returns root process with all subprocesses on the system
func GetTree(pid ...int) (*ProcessInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}

// ❗ GetList returns slice with all active processes on the system
func GetList() ([]*ProcessInfo, error) {
	panic("UNSUPPORTED")
	return nil, nil
}