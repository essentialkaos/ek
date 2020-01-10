// +build linux
// +build 386, amd64

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// byteSliceToString convert byte slice to string
func byteSliceToString(s [65]int8) string {
	result := ""

	for _, r := range s {
		if r == 0 {
			break
		}

		result += string(r)
	}

	return result
}
