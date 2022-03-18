//go:build freebsd
// +build freebsd

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"
	"time"

	PATH "github.com/essentialkaos/ek/path"
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
		return time.Time{}, time.Time{}, time.Time{}, err
	}

	return time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec)),
		time.Unix(int64(stat.Mtimespec.Sec), int64(stat.Mtimespec.Nsec)),
		time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec)),
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
		return -1, -1, -1, err
	}

	return int64(stat.Atimespec.Sec),
		int64(stat.Mtimespec.Sec),
		int64(stat.Ctimespec.Sec),
		nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isEmptyDirent(n int) bool {
	return n == 0x18
}
