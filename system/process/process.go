package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strconv"

	"pkg.re/essentialkaos/ek.v9/errutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// parseSize parse size in kB
func parseSize(v string, errs *errutil.Errors) uint64 {
	size, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return size * 1024
}

// parseUint parse uint value
func parseUint(v string, errs *errutil.Errors) uint64 {
	value, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return value
}

// parseFloat parse float value
func parseFloat(v string, errs *errutil.Errors) float64 {
	value, err := strconv.ParseFloat(v, 64)

	if err != nil {
		errs.Add(err)
		return 0.0
	}

	return value
}

// parseInt parse int value
func parseInt(v string, errs *errutil.Errors) int {
	value, err := strconv.ParseInt(v, 10, 64)

	if err != nil {
		errs.Add(err)
		return 0
	}

	return int(value)
}

// readField read field from data
func readField(data string, index int) string {
	if data == "" {
		return ""
	}

	curIndex, startPointer := -1, -1

	for i, r := range data {
		if r == ' ' || r == '\t' {
			if curIndex == index {
				return data[startPointer:i]
			}

			startPointer = -1
			continue
		}

		if startPointer == -1 {
			startPointer = i
			curIndex++
		}
	}

	if index > curIndex {
		return ""
	}

	return data[startPointer:]
}
