// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"pkg.re/essentialkaos/ek.v9/errutil"
	"pkg.re/essentialkaos/ek.v9/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Path to file with LA info in procfs
var procLoadAvgFile = "/proc/loadavg"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetLA return loadavg
func GetLA() (*LoadAvg, error) {
	fd, err := os.OpenFile(procLoadAvgFile, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	text, _ := r.ReadString('\n')

	if len(text) < 20 {
		return nil, errors.New("Can't parse file " + procLoadAvgFile)
	}

	la := &LoadAvg{}
	errs := errutil.NewErrors()

	la.Min1 = parseFloat(strutil.ReadField(text, 0, true), errs)
	la.Min5 = parseFloat(strutil.ReadField(text, 1, true), errs)
	la.Min15 = parseFloat(strutil.ReadField(text, 2, true), errs)

	if errs.HasErrors() {
		return nil, errs.Last()
	}

	procs := strutil.ReadField(text, 3, true)
	delimPosition := strings.IndexRune(procs, '/')

	if delimPosition == -1 {
		return nil, errors.New("Can't parse file " + procLoadAvgFile)
	}

	la.RProc = parseInt(procs[:delimPosition], errs)
	la.TProc = parseInt(procs[delimPosition+1:], errs)

	if errs.HasErrors() {
		return nil, errs.Last()
	}

	return la, nil
}
