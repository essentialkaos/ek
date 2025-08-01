//go:build linux
// +build linux

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"syscall"
	"time"

	PATH "github.com/essentialkaos/ek/v13/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetTimes returns time of access, modification, and creation at once
func GetTimes(path string) (time.Time, time.Time, time.Time, error) {
	if path == "" {
		return time.Time{}, time.Time{}, time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, fmt.Errorf("Can't get file info for %q: %w", path, err)
	}

	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)),
		time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec)),
		time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)),
		nil
}

// GetTimestamps returns time of access, modification, and creation at once as unix timestamp
func GetTimestamps(path string) (int64, int64, int64, error) {
	if path == "" {
		return -1, -1, -1, ErrEmptyPath
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, -1, -1, fmt.Errorf("Can't get file info for %q: %w", path, err)
	}

	return int64(stat.Atim.Sec),
		int64(stat.Mtim.Sec),
		int64(stat.Ctim.Sec),
		nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// isEmptyDirent checks if the dirent shows that the directory is empty
func isEmptyDirent(n int) bool {
	return n <= 0x40
}
