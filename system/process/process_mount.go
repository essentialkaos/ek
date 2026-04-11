//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMountInfo returns info about process mounts
func GetMountInfo(pid int) ([]*MountInfo, error) {
	mountFile := path.Join(procFS, strconv.Itoa(pid), "mountinfo")
	s, closeFunc, err := getFileScanner(mountFile)

	if err != nil {
		return nil, err
	}

	defer closeFunc()

	var result []*MountInfo

	for s.Scan() {
		info, err := parseMountInfoLine(s.Text())

		if err != nil {
			return nil, err
		}

		result = append(result, info)
	}

	return result, s.Err()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseMountInfoLine parses a single line from /proc/[pid]/mountinfo
func parseMountInfoLine(data string) (*MountInfo, error) {
	var err error

	info := &MountInfo{}

	optFieldsNum := 0
	optFieldParsed := false

LOOP:
	for i := range 128 {
		pseudoIndex := i - optFieldsNum
		value := strutil.ReadField(data, i, false, ' ')

		if i >= 6 && !optFieldParsed {
			if value != "-" {
				info.OptionalFields = append(info.OptionalFields, value)
			} else {
				optFieldParsed = true
			}

			optFieldsNum++
			continue
		}

		switch pseudoIndex {
		case 0:
			info.ID, err = parseFieldUint16(value, "MountID")
		case 1:
			info.ParentID, err = parseFieldUint16(value, "ParentID")
		case 2:
			info.StDevMajor, info.StDevMinor, err = parseStDevValue(value)
		case 3:
			info.Root = value
		case 4:
			info.MountPoint = value
		case 5:
			info.MountOptions = strings.Split(value, ",")
		case 6:
			info.FSType = value
		case 7:
			info.MountSource = value
		case 8:
			info.SuperOptions = strings.Split(value, ",")
		default:
			break LOOP
		}

		if err != nil {
			return nil, err
		}
	}

	return info, nil
}

// parseStDevValue parses st_dev value from mount info line
func parseStDevValue(data string) (uint16, uint16, error) {
	major, err := parseFieldUint16(strutil.ReadField(data, 0, false, ':'), "StDevMajor")

	if err != nil {
		return 0, 0, err
	}

	minor, err := parseFieldUint16(strutil.ReadField(data, 1, false, ':'), "StDevMinor")

	if err != nil {
		return 0, 0, err
	}

	return major, minor, nil
}

// parseFieldUint16 parses a field as uint16
func parseFieldUint16(s, field string) (uint16, error) {
	u, err := strconv.ParseUint(s, 10, 16)

	if err != nil {
		return 0, fmt.Errorf("can't parse field %s: %w", field, err)
	}

	return uint16(u), nil
}
