package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with LA info in procfs
var procLoadAvgFile = "/proc/loadavg"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetLA returns loadavg
func GetLA() (*LoadAvg, error) {
	data, err := os.ReadFile(procLoadAvgFile)

	if err != nil {
		return nil, err
	}

	return parseLAInfo(string(data))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseLAInfo parses loadavg data
func parseLAInfo(text string) (*LoadAvg, error) {
	var err error

	text = strings.Trim(text, " \n\r")

	if len(text) < 18 {
		return nil, fmt.Errorf("file %s is not valid loadavg source", procLoadAvgFile)
	}

	la := &LoadAvg{}

	la.Min1, err = strconv.ParseFloat(strutil.ReadField(text, 0, true), 64)

	if err != nil {
		return nil, errors.New("can't parse field 0 as float number in " + procLoadAvgFile)
	}

	la.Min5, err = strconv.ParseFloat(strutil.ReadField(text, 1, true), 64)

	if err != nil {
		return nil, errors.New("can't parse field 1 as float number in " + procLoadAvgFile)
	}

	la.Min15, err = strconv.ParseFloat(strutil.ReadField(text, 2, true), 64)

	if err != nil {
		return nil, errors.New("can't parse field 2 as float number in " + procLoadAvgFile)
	}

	rproc, tproc, ok := strings.Cut(strutil.ReadField(text, 3, true), "/")

	if !ok {
		return nil, errors.New("can't parse field 3 in " + procLoadAvgFile)
	}

	la.RProc, err = strconv.Atoi(rproc)

	if err != nil {
		return nil, errors.New("can't parse processes number in " + procLoadAvgFile)
	}

	la.TProc, err = strconv.Atoi(tproc)

	if err != nil {
		return nil, errors.New("can't parse processes number in " + procLoadAvgFile)
	}

	return la, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //
