// +build linux

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"syscall"
	"time"

	PATH "pkg.re/essentialkaos/ek.v9/path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetTimes return time of access, modification, and creation at once
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

	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)),
		time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec)),
		time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)),
		nil
}

// GetTimestamps return time of access, modification, and creation at once as unix timestamp
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

	return int64(stat.Atim.Sec),
		int64(stat.Mtim.Sec),
		int64(stat.Ctim.Sec),
		nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func isEmptyDirent(n int) bool {
	return n == 0x30
}
