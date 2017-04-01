// Package jsonutil provides methods for working with json data
package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// EncodeToFile encode data to json and save to file
func EncodeToFile(file string, v interface{}, perms ...os.FileMode) error {
	jsonData, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		return err
	}

	if jsonData[len(jsonData)-1] != '\n' {
		jsonData = append(jsonData, byte('\n'))
	}

	var perm os.FileMode = 0644

	if len(perms) > 0 {
		perm = perms[0]
	}

	return ioutil.WriteFile(file, jsonData, perm)
}

// DecodeFile reads and decode json file
func DecodeFile(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
