// +build gofuzz

package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Fuzz(data []byte) int {
	config := &Config{
		data:     make(map[string]string),
		sections: make([]string, 0),
		props:    make([]string, 0),
		file:     "",
	}

	err := readConfigData(config, bytes.NewReader(data), "")

	if err != nil {
		return 0
	}

	return 1
}
