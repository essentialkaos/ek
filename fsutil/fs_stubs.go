//go:build !linux && !darwin && !freebsd

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
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

var (
	// ❗ ErrEmptyPath is returned by methods when the given path is empty and cannot be used
	ErrEmptyPath = errors.New("path is empty")

	// ❗ ErrEmptyPerms is returned by methods when the given permissions is empty
	ErrEmptyPerms = errors.New("permissions is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ CheckPerms checks multiple filesystem permissions for the given path at once
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
}

// ❗ ValidatePerms validates filesystem permissions for the given path and returns
// a descriptive error if any check fails
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
}

// ❗ ProperPath returns the first path from the slice that satisfies the given
// permission checks, or an empty string if none match
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
func ProperPath(perms string, paths ...string) string {
	panic("UNSUPPORTED")
}

// ❗ IsExist returns true if the given path exists on the filesystem
func IsExist(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsRegular returns true if the given path is a regular file
func IsRegular(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsSocket returns true if the given path is a Unix domain socket
func IsSocket(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsBlockDevice returns true if the given path is a block device
func IsBlockDevice(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsCharacterDevice returns true if the given path is a character device
func IsCharacterDevice(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsDir returns true if the given path is a directory
func IsDir(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsLink returns true if the given path is a symbolic link
func IsLink(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsReadable returns true if the given path is readable by the current user
func IsReadable(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsReadableByUser returns true if the given path is readable by the named user
func IsReadableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsWritable returns true if the given path is writable by the current user
func IsWritable(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsWritableByUser returns true if the given path is writable by the named user
func IsWritableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsExecutable returns true if the given path is executable by the current user
func IsExecutable(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsExecutableByUser returns true if the given path is executable by the named user
func IsExecutableByUser(path, userName string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsEmpty returns true if the given file exists and has a size of zero bytes
func IsEmpty(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ IsEmptyDir returns true if the given directory exists and contains no entries
func IsEmptyDir(path string) bool {
	panic("UNSUPPORTED")
}

// ❗ GetOwner returns the UID and GID of the given path's owner
func GetOwner(path string) (int, int, error) {
	panic("UNSUPPORTED")
}

// ❗ GetATime returns the time of last access for the given path
func GetATime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
}

// ❗ GetCTime returns the time of creation (inode change) for the given path
func GetCTime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
}

// ❗ GetMTime returns the time of last modification for the given path
func GetMTime(path string) (time.Time, error) {
	panic("UNSUPPORTED")
}

// ❗ GetSize returns the size of the given file in bytes, or -1 on error
func GetSize(path string) int64 {
	panic("UNSUPPORTED")
}

// ❗ GetMode returns the permission bits of the given path as an [os.FileMode]
func GetMode(path string) os.FileMode {
	panic("UNSUPPORTED")
}

// ❗ GetModeOctal returns the permission bits of the given path as an octal
// string (e.g. "0644")
func GetModeOctal(path string) string {
	panic("UNSUPPORTED")
}

// ❗ GetTimes returns the access, modification, and creation times of the given path
// at once
func GetTimes(path string) (time.Time, time.Time, time.Time, error) {
	panic("UNSUPPORTED")
}

// ❗ GetTimestamps returns the access, modification, and creation times of the given
// path as Unix timestamps
func GetTimestamps(path string) (int64, int64, int64, error) {
	panic("UNSUPPORTED")
}
