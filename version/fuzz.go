// +build gofuzz

package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func Fuzz(data []byte) int {
	_, err := Parse(string(data))

	if err != nil {
		return 0
	}

	return 1
}
