//go:build darwin

package tty

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// isTmuxAncestor walks the process tree to check whether any ancestor is a tmux server
func isTmuxAncestor() (bool, error) {
	return false, nil
}
