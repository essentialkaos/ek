//go:build !windows

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
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	PATH "github.com/essentialkaos/ek/v13/path"
	"github.com/essentialkaos/ek/v13/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_IFMT   = 0xF000
	_IFSOCK = 0xC000
	_IFREG  = 0x8000
	_IFBLK  = 0x6000
	_IFDIR  = 0x4000
	_IFCHR  = 0x2000
	_IRUSR  = 0x100
	_IWUSR  = 0x80
	_IXUSR  = 0x40
	_IRGRP  = 0x20
	_IWGRP  = 0x10
	_IXGRP  = 0x8
	_IROTH  = 0x4
	_IWOTH  = 0x2
	_IXOTH  = 0x1
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyPath is returned by methods when the given path is empty and cannot be used
	ErrEmptyPath = errors.New("path is empty")

	// ErrEmptyPerms is returned by methods when the given permissions is empty
	ErrEmptyPerms = errors.New("permissions is empty")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CheckPerms checks multiple filesystem permissions for the given path at once
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
	if perms == "" || path == "" {
		return false
	}

	path = PATH.Clean(path)
	perms = strings.ToUpper(perms)

	var gidList []int
	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := getCurrentUser()

	if err != nil {
		return false
	}

	if strings.ContainsAny(perms, "XWR") {
		gidList = getGIDList(user)
	}

	for _, k := range perms {
		switch k {

		case 'F':
			if stat.Mode&_IFMT != _IFREG {
				return false
			}

		case 'D':
			if stat.Mode&_IFMT != _IFDIR {
				return false
			}

		case 'B':
			if stat.Mode&_IFMT != _IFBLK {
				return false
			}

		case 'C':
			if stat.Mode&_IFMT != _IFCHR {
				return false
			}

		case 'L':
			if !IsLink(path) {
				return false
			}

		case 'X':
			if !isExecutableStat(stat, user.UID, gidList) {
				return false
			}

		case 'W':
			if !isWritableStat(stat, user.UID, gidList) {
				return false
			}

		case 'R':
			if !isReadableStat(stat, user.UID, gidList) {
				return false
			}

		case 'S':
			if stat.Size == 0 {
				return false
			}
		}
	}

	return true
}

// ValidatePerms validates filesystem permissions for the given path and returns
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
	switch {
	case perms == "":
		return ErrEmptyPerms
	case path == "":
		return ErrEmptyPath
	}

	path = PATH.Clean(path)
	perms = strings.ToUpper(perms)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		switch {
		case strings.ContainsRune(perms, 'F'):
			return fmt.Errorf("file %s doesn't exist or not accessible", path)
		case strings.ContainsRune(perms, 'D'):
			return fmt.Errorf("directory %s doesn't exist or not accessible", path)
		case strings.ContainsRune(perms, 'B'):
			return fmt.Errorf("block device %s doesn't exist or not accessible", path)
		case strings.ContainsRune(perms, 'C'):
			return fmt.Errorf("character device %s doesn't exist or not accessible", path)
		case strings.ContainsRune(perms, 'L'):
			return fmt.Errorf("link %s doesn't exist or not accessible", path)
		}

		return fmt.Errorf("object %s doesn't exist or not accessible", path)
	}

	user, err := getCurrentUser()

	if err != nil {
		return errors.New("can't get information about the current user")
	}

	for _, k := range perms {
		switch k {

		case 'F':
			if stat.Mode&_IFMT != _IFREG {
				return fmt.Errorf("%s is not a file", path)
			}

		case 'D':
			if stat.Mode&_IFMT != _IFDIR {
				return fmt.Errorf("%s is not a directory", path)
			}

		case 'B':
			if stat.Mode&_IFMT != _IFBLK {
				return fmt.Errorf("%s is not a block device", path)
			}

		case 'C':
			if stat.Mode&_IFMT != _IFCHR {
				return fmt.Errorf("%s is not a character device", path)
			}

		case 'L':
			if !IsLink(path) {
				return fmt.Errorf("%s is not a link", path)
			}

		case 'X':
			if !isExecutableStat(stat, user.UID, getGIDList(user)) {
				return fmt.Errorf(
					"%s %s is not executable",
					getObjectType(stat), path,
				)
			}

		case 'W':
			if !isWritableStat(stat, user.UID, getGIDList(user)) {
				return fmt.Errorf(
					"%s %s is not writable",
					getObjectType(stat), path,
				)
			}

		case 'R':
			if !isReadableStat(stat, user.UID, getGIDList(user)) {
				return fmt.Errorf(
					"%s %s is not readable",
					getObjectType(stat), path,
				)
			}

		case 'S':
			if stat.Size == 0 {
				return fmt.Errorf("%s %s is empty", getObjectType(stat), path)
			}
		}
	}

	return nil
}

// ProperPath returns the first path from the slice that satisfies the given
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
	for _, path := range paths {
		if strings.TrimSpace(path) == "" {
			continue
		}

		path = PATH.Clean(path)

		if CheckPerms(perms, path) {
			return path
		}
	}

	return ""
}

// IsExist returns true if the given path exists on the filesystem
func IsExist(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return syscall.Access(path, syscall.F_OK) == nil
}

// IsRegular returns true if the given path is a regular file
func IsRegular(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFREG
}

// IsSocket returns true if the given path is a Unix domain socket
func IsSocket(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFSOCK
}

// IsBlockDevice returns true if the given path is a block device
func IsBlockDevice(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFBLK
}

// IsCharacterDevice returns true if the given path is a character device
func IsCharacterDevice(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFCHR
}

// IsDir returns true if the given path is a directory
func IsDir(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFDIR
}

// IsLink returns true if the given path is a symbolic link
func IsLink(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var buf = make([]byte, 1)
	_, err := syscall.Readlink(path, buf)

	return err == nil
}

// IsReadable returns true if the given path is readable by the current user
func IsReadable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := getCurrentUser()

	if err != nil {
		return false
	}

	return isReadableStat(stat, user.UID, getGIDList(user))
}

// IsReadableByUser returns true if the given path is readable by the named user
func IsReadableByUser(path, userName string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.LookupUser(userName)

	if err != nil {
		return false
	}

	return isReadableStat(stat, user.UID, getGIDList(user))
}

// IsWritable returns true if the given path is writable by the current user
func IsWritable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := getCurrentUser()

	if err != nil {
		return false
	}

	return isWritableStat(stat, user.UID, getGIDList(user))
}

// IsWritableByUser returns true if the given path is writable by the named user
func IsWritableByUser(path, userName string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.LookupUser(userName)

	if err != nil {
		return false
	}

	return isWritableStat(stat, user.UID, getGIDList(user))
}

// IsExecutable returns true if the given path is executable by the current user
func IsExecutable(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := getCurrentUser()

	if err != nil {
		return false
	}

	return isExecutableStat(stat, user.UID, getGIDList(user))
}

// IsExecutableByUser returns true if the given path is executable by the named user
func IsExecutableByUser(path, userName string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.LookupUser(userName)

	if err != nil {
		return false
	}

	return isExecutableStat(stat, user.UID, getGIDList(user))
}

// IsEmpty returns true if the given file exists and has a size of zero bytes
func IsEmpty(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return GetSize(path) == 0
}

// IsEmptyDir returns true if the given directory exists and contains no entries
func IsEmptyDir(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	fd, err := syscall.Open(path, syscall.O_RDONLY, 0)

	if err != nil {
		return false
	}

	defer syscall.Close(fd)

	n, err := syscall.ReadDirent(fd, make([]byte, 4096))

	if isEmptyDirent(n) || err != nil {
		return true
	}

	return false
}

// GetOwner returns the UID and GID of the given path's owner
func GetOwner(path string) (int, int, error) {
	if path == "" {
		return -1, -1, ErrEmptyPath
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, -1, fmt.Errorf("can't get owner info for %q: %w", path, err)
	}

	return int(stat.Uid), int(stat.Gid), nil
}

// GetATime returns the time of last access for the given path
func GetATime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	atime, _, _, err := GetTimes(path)

	return atime, err
}

// GetCTime returns the time of creation (inode change) for the given path
func GetCTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, _, ctime, err := GetTimes(path)

	return ctime, err
}

// GetMTime returns the time of last modification for the given path
func GetMTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, mtime, _, err := GetTimes(path)

	return mtime, err
}

// GetSize returns the size of the given file in bytes, or -1 on error
func GetSize(path string) int64 {
	if path == "" {
		return -1
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1
	}

	return stat.Size
}

// GetMode returns the permission bits of the given path as an [os.FileMode]
func GetMode(path string) os.FileMode {
	if path == "" {
		return 0
	}

	path = PATH.Clean(path)

	return getMode(path) & 0777
}

// GetModeOctal returns the permission bits of the given path as an octal
// string (e.g. "0644")
func GetModeOctal(path string) string {
	if path == "" {
		return ""
	}

	path = PATH.Clean(path)

	return getModeOctal(path)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getMode returns file mode bits
func getMode(path string) os.FileMode {
	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return 0
	}

	return os.FileMode(stat.Mode)
}

// getModeOctal returns file mode bits in octal form (like 0644)
func getModeOctal(path string) string {
	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return ""
	}

	m := strconv.FormatUint(uint64(stat.Mode&0777), 8)
	s := 0

	if stat.Mode&syscall.S_ISVTX != 0 {
		s += 1
	}

	if stat.Mode&syscall.S_ISGID != 0 {
		s += 2
	}

	if stat.Mode&syscall.S_ISUID != 0 {
		s += 4
	}

	return strconv.Itoa(s) + m
}

// isReadableStat checks if the object stat info indicates that the object is readable
func isReadableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if stat == nil {
		return false
	}

	if uid == 0 {
		return true
	}

	if stat.Mode&_IROTH == _IROTH {
		return true
	}

	if stat.Mode&_IRUSR == _IRUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IRGRP == _IRGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

// isWritableStat checks if the object stat info indicates that the object is writable
func isWritableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if stat == nil {
		return false
	}

	if uid == 0 {
		return true
	}

	if stat.Mode&_IWOTH == _IWOTH {
		return true
	}

	if stat.Mode&_IWUSR == _IWUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IWGRP == _IWGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

// isExecutableStat checks if the object stat info indicates that the object is executable
func isExecutableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if stat == nil {
		return false
	}

	if uid == 0 {
		return stat.Mode&(_IXUSR|_IXGRP|_IXOTH) != 0
	}

	if stat.Mode&_IXOTH == _IXOTH {
		return true
	}

	if stat.Mode&_IXUSR == _IXUSR && uid == int(stat.Uid) {
		return true
	}

	for _, gid := range gids {
		if stat.Mode&_IXGRP == _IXGRP && gid == int(stat.Gid) {
			return true
		}
	}

	return false
}

// getGIDList returns a list of group IDs for the given user
func getGIDList(user *system.User) []int {
	if user == nil {
		return nil
	}

	var result []int

	for _, group := range user.Groups {
		result = append(result, group.GID)
	}

	return result
}

// getObjectType returns a string representation of the object type based on its
// mode bits
func getObjectType(stat *syscall.Stat_t) string {
	switch stat.Mode & _IFMT {
	case _IFREG:
		return "file"
	case _IFDIR:
		return "directory"
	case _IFBLK:
		return "block device"
	case _IFCHR:
		return "character device"
	}

	return "object"
}
