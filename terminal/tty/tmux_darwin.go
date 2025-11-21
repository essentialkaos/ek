//go:build darwin

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// isTmuxAncestor returns true if the current process is an ancestor of tmux
func isTmuxAncestor() (bool, error) {
	return false, nil
}
