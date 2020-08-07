// +build windows

package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

// CheckPerms check many props at once
func CheckPerms(perms, path string) bool {
	return false
}

// ProperPath returns the first proper path from a given slice
func ProperPath(props string, paths []string) string {
	return ""
}

// IsExist returns true if the given object is exist
func IsExist(path string) bool {
	return false
}

// IsRegular returns true if the given object is a regular file
func IsRegular(path string) bool {
	return false
}

// IsSocket returns true if the given object is a socket
func IsSocket(path string) bool {
	return false
}

// IsBlockDevice returns true if the given object is a device
func IsBlockDevice(path string) bool {
	return false
}

// IsCharacterDevice returns true if the given object is a character device
func IsCharacterDevice(path string) bool {
	return false
}

// IsDir returns true if the given object is a directory
func IsDir(path string) bool {
	return false
}

// IsLink returns true if the given object is a link
func IsLink(path string) bool {
	return false
}

// IsReadable returns true if given object is readable by current user
func IsReadable(path string) bool {
	return false
}

// IsReadableByUser returns true if given object is readable by some user
func IsReadableByUser(path, userName string) bool {
	return false
}

// IsWritable returns true if given object is writable by current user
func IsWritable(path string) bool {
	return false
}

// IsWritableByUser returns true if given object is writable by some user
func IsWritableByUser(path, userName string) bool {
	return false
}

// IsExecutable returns true if given object is executable by current user
func IsExecutable(path string) bool {
	return false
}

// IsExecutableByUser returns true if given object is executable by some user
func IsExecutableByUser(path, userName string) bool {
	return false
}

// IsNonEmpty returns true if given file is not empty
func IsNonEmpty(path string) bool {
	return false
}

// IsEmptyDir returns true if given directory es empty
func IsEmptyDir(path string) bool {
	return false
}

// GetOwner returns object owner UID and GID
func GetOwner(path string) (int, int, error) {
	return 0, 0, nil
}

// GetATime returns time of last access
func GetATime(path string) (time.Time, error) {
	return time.Time{}, nil
}

// GetCTime returns time of creation
func GetCTime(path string) (time.Time, error) {
	return time.Time{}, nil
}

// GetMTime returns time of modification
func GetMTime(path string) (time.Time, error) {
	return time.Time{}, nil
}

// GetSize returns file size in bytes
func GetSize(path string) int64 {
	return -1
}

// GetMode returns file mode bits
func GetMode(path string) os.FileMode {
	return 0
}

// GetTimes returns time of access, modification, and creation at once
func GetTimes(path string) (time.Time, time.Time, time.Time, error) {
	return time.Time{}, time.Time{}, time.Time{}, nil
}

// GetTimestamps returns time of access, modification, and creation at once as unix timestamp
func GetTimestamps(path string) (int64, int64, int64, error) {
	return -1, -1, -1, nil
}
