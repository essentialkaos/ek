// +build gofuzz

package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
