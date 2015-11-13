package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/essentialkaos/ek/fsutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DecodeFile reads and decode json file
func DecodeFile(file string, v interface{}) error {
	if !fsutil.CheckPerms("FR", file) {
		return errors.New("File " + file + " not exist or not readable")
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)

	if err != nil {
		return err
	}

	return nil
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

	_, err = fd.Write(jsonData)

	if err != nil {
		return err
	}
	_, err = fd.WriteString("\n")

	if err != nil {
		return err
	}

	return nil
}
