// Package jsonutil provides methods for working with json data
package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DecodeFile reads and decode json file
func DecodeFile(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// EncodeToFile encode data to json and save to file
func EncodeToFile(file string, v interface{}) error {
	fd, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer fd.Close()

	jsonData, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		return err
	}

	jsonData = append(jsonData, byte('\n'))

	_, err = fd.Write(jsonData)

	if err != nil {
		return err
	}

	return nil
}
