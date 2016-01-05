// +build !linux, !darwin, windows

package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// Env is map with environment values
type Env map[string]string

// ////////////////////////////////////////////////////////////////////////////////// //

// Get return key-value map with environment values
func Get() Env {
	return Env{}
}

// Which find full path to some app
func Which(name string) string {
	return ""
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Path return path as string slice
func (e Env) Path() []string {
	return []string{}
}
