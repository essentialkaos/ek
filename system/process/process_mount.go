//go:build linux

package process

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// MountInfo contains information about mounts
// https://www.kernel.org/doc/Documentation/filesystems/proc.txt
type MountInfo struct {
	MountID        uint16   `json:"mount_id"`        // Unique identifier of the mount (may be reused after umount)
	ParentID       uint16   `json:"parent_id"`       // ID of parent (or of self for the top of the mount tree)
	StDevMajor     uint16   `json:"stdev_major"`     // Major value of st_dev for files on filesystem
	StDevMinor     uint16   `json:"stdev_minor"`     // Minor value of st_dev for files on filesystem
	Root           string   `json:"root"`            // Root of the mount within the filesystem
	MountPoint     string   `json:"mount_point"`     // Mount point relative to the process's root
	MountOptions   []string `json:"mount_options"`   // Per mount options
	OptionalFields []string `json:"optional_fields"` // Zero or more fields of the form "tag[:value]"
	FSType         string   `json:"fs_type"`         // Name of filesystem of the form "type[.subtype]"
	MountSource    string   `json:"mount_source"`    // Filesystem specific information or "none"
	SuperOptions   []string `json:"super_options"`   // Per super block options
}

// ////////////////////////////////////////////////////////////////////////////////// //

// GetMountInfo returns info about process mounts
func GetMountInfo(pid int) ([]*MountInfo, error) {
	fd, err := os.OpenFile(procFS+"/"+strconv.Itoa(pid)+"/mountinfo", os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	var result []*MountInfo

	for s.Scan() {
		info, err := parseMountInfoLine(s.Text())

		if err != nil {
			return nil, err
		}

		result = append(result, info)
	}

	return result, nil
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
			info.MountID, err = parseFieldUint16(value, "MountID")
		case 1:
			info.ParentID, err = parseFieldUint16(value, "MountID")
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
		return 0, fmt.Errorf("Can't parse field %s: %w", field, err)
	}

	return uint16(u), nil
}
