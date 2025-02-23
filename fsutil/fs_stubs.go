//go:build !linux && !darwin && !freebsd
// +build !linux,!darwin,!freebsd

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	S_IFMT   = 0
	S_IFSOCK = 0
	S_IFLNK  = 0
	S_IFREG  = 0
	S_IFBLK  = 0
	S_IFDIR  = 0
	S_IFCHR  = 0
	S_IFIFO  = 0
	S_ISUID  = 0
	S_ISGID  = 0
	S_ISVTX  = 0
	S_IRWXU  = 0
	S_IRUSR  = 0
	S_IWUSR  = 0
	S_IXUSR  = 0
	S_IRWXG  = 0
	S_IRGRP  = 0
	S_IWGRP  = 0
	S_IXGRP  = 0
	S_IRWXO  = 0
	S_IROTH  = 0
	S_IWOTH  = 0
	S_IXOTH  = 0
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ ErrEmptyPath can be returned by different methods if given path is empty and
// can't be used
var ErrEmptyPath = errors.New("Path is empty")

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ CheckPerms checks many permissions at once
//
// Permissions:
//
//   - F: is file
//   - D: is directory
//   - X: is executable
//   - L: is link
//   - W: is writable
//   - R: is readable
//   - B: is block device
//   - C: is character device
//   - S: not empty (only for files)
func CheckPerms(perms, path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ ValidatePerms validates permissions for file or directory
//
// Permissions:
//
//   - F: is file
//   - D: is directory
//   - X: is executable
//   - L: is link
//   - W: is writable
//   - R: is readable
//   - B: is block device
//   - C: is character device
//   - S: not empty (only for files)
func ValidatePerms(perms, path string) error {
	panic("UNSUPPORTED")
	return nil
}

// ❗ ProperPath returns the first proper path from a given slice
//
// Permissions:
//
//   - F: is file
//   - D: is directory
//   - X: is executable
//   - L: is link
//   - W: is writable
//   - R: is readable
//   - B: is block device
//   - C: is character device
//   - S: not empty (only for files)
func ProperPath(perms string, paths []string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ IsExist returns true if the given object is exist
func IsExist(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsRegular returns true if the given object is a regular file
func IsRegular(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsSocket returns true if the given object is a socket
func IsSocket(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsBlockDevice returns true if the given object is a device
func IsBlockDevice(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsCharacterDevice returns true if the given object is a character device
func IsCharacterDevice(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsDir returns true if the given object is a directory
func IsDir(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsLink returns true if the given object is a link
func IsLink(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsReadable returns true if given object is readable by current user
func IsReadable(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsReadableByUser returns true if given object is readable by some user
func IsReadableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsWritable returns true if given object is writable by current user
func IsWritable(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsWritableByUser returns true if given object is writable by some user
func IsWritableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsExecutable returns true if given object is executable by current user
func IsExecutable(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsExecutableByUser returns true if given object is executable by some user
func IsExecutableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsEmpty returns true if given file is empty
func IsEmpty(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ IsEmptyDir returns true if given directory es empty
func IsEmptyDir(path string) bool {
	panic("UNSUPPORTED")
	return false
}

// ❗ GetOwner returns object owner UID and GID
func GetOwner(path string) (int, int, error) {
	panic("UNSUPPORTED")
	return 0, 0, nil
}

// ❗ GetATime returns time of last access
func GetATime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
	return time.Time{}, nil
}

// ❗ GetCTime returns time of creation
func GetCTime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
	return time.Time{}, nil
}

// ❗ GetMTime returns time of modification
func GetMTime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
	return time.Time{}, nil
}

// ❗ GetSize returns file size in bytes
func GetSize(path string) int64 {
	panic("UNSUPPORTED")
	return -1
}

// ❗ GetMode returns file mode bits
func GetMode(path string) os.FileMode {
	panic("UNSUPPORTED")
	return 0
}

// ❗ GetModeOctal returns file mode bits in octal form (like 0644)
func GetModeOctal(path string) string {
	panic("UNSUPPORTED")
	return ""
}

// ❗ GetTimes returns time of access, modification, and creation at once
func GetTimes(path string) (time.Time, time.Time, time.Time, error) {
	panic("UNSUPPORTED")
	return time.Time{}, time.Time{}, time.Time{}, nil
}

// ❗ GetTimestamps returns time of access, modification, and creation at once as unix timestamp
func GetTimestamps(path string) (int64, int64, int64, error) {
	panic("UNSUPPORTED")
	return -1, -1, -1, nil
}
