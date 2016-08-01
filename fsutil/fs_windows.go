// +build windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	S_IFMT   = 0170000
	S_IFSOCK = 0140000
	S_IFLNK  = 0120000
	S_IFREG  = 0100000
	S_IFBLK  = 0060000
	S_IFDIR  = 0040000
	S_IFCHR  = 0020000
	S_IFIFO  = 0010000
	S_ISUID  = 0004000
	S_ISGID  = 0002000
	S_ISVTX  = 0001000
	S_IRWXU  = 00700
	S_IRUSR  = 00400
	S_IWUSR  = 00200
	S_IXUSR  = 00100
	S_IRWXG  = 00070
	S_IRGRP  = 00040
	S_IWGRP  = 00020
	S_IXGRP  = 00010
	S_IRWXO  = 00007
	S_IROTH  = 00004
	S_IWOTH  = 00002
	S_IXOTH  = 00001
)

// ////////////////////////////////////////////////////////////////////////////////// //

func CheckPerms(perms, path string) bool {
	return false
}

func ProperPath(props string, paths []string) string {
	return ""
}

func IsExist(path string) bool {
	return false
}

func IsRegular(path string) bool {
	return false
}

func IsSocket(path string) bool {
	return false
}

func IsBlockDevice(path string) bool {
	return false
}

func IsCharacterDevice(path string) bool {
	return false
}

func IsDir(path string) bool {
	return false
}

func IsLink(path string) bool {
	return false
}

func IsReadable(path string) bool {
	return false
}

func IsWritable(path string) bool {
	return false
}

func IsExecutable(path string) bool {
	return false
}

func IsNonEmpty(path string) bool {
	return false
}

func IsEmptyDir(path string) bool {
	return false
}

func GetOwner(path string) (int, int, error) {
	return 0, 0, nil
}

func GetATime(path string) (time.Time, error) {
	return time.Time{}, nil
}

func GetCTime(path string) (time.Time, error) {
	return time.Time{}, nil
}

func GetMTime(path string) (time.Time, error) {
	return time.Time{}, nil
}

func GetSize(path string) int64 {
	return 0
}

func GetPerm(path string) os.FileMode {
	return 0644
}

func GetTimes(path string) (time.Time, time.Time, time.Time, error) {
	return time.Time{}, time.Time{}, time.Time{}, nil
}

func GetTimestamps(path string) (int64, int64, int64, error) {
	return -1, -1, -1, nil
}
