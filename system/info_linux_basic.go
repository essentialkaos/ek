//go:build linux && (386 || arm64 || amd64)
// +build linux
// +build 386 arm64 amd64

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// byteSliceToString convert byte slice to string
func byteSliceToString(s [65]int8) string {
	result := ""

	for _, r := range s {
		if r == 0 {
			break
		}

		result += string(rune(r))
	}

	return result
}
