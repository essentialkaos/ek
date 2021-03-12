// +build gofuzz

package version

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

func Fuzz(data []byte) int {
	_, err := Parse(string(data))

	if err != nil {
		return 0
	}

	return 1
}
