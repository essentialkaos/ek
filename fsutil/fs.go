// +build !windows

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"time"

	PATH "pkg.re/essentialkaos/ek.v11/path"
	"pkg.re/essentialkaos/ek.v11/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_IFMT   = 0XF000
	_IFSOCK = 0XC000
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

// ErrEmptyPath error
var ErrEmptyPath = errors.New("Path is empty")

// ////////////////////////////////////////////////////////////////////////////////// //

// CheckPerms check many props at once
//
// F - is file
// D - is directory
// X - is executable
// L - is link
// W - is writable
// R - is readable
// S - not empty (only for files)
//
func CheckPerms(props, path string) bool {
	if len(props) == 0 || path == "" {
		return false
	}

	path = PATH.Clean(path)
	props = strings.ToUpper(props)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	for _, k := range props {
		switch k {

		case 'F':
			if stat.Mode&_IFMT != _IFREG {
				return false
			}

		case 'D':
			if stat.Mode&_IFMT != _IFDIR {
				return false
			}

		case 'L':
			if !IsLink(path) {
				return false
			}

		case 'X':
			if !isExecutableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'W':
			if !isWritableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'R':
			if !isReadableStat(stat, user.UID, getGIDList(user)) {
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

// ProperPath returns the first proper path from a given slice
func ProperPath(props string, paths []string) string {
	for _, path := range paths {
		path = PATH.Clean(path)

		if CheckPerms(props, path) {
			return path
		}
	}

	return ""
}

// IsExist returns true if the given object is exist
func IsExist(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return syscall.Access(path, syscall.F_OK) == nil
}

// IsRegular returns true if the given object is a regular file
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

// IsSocket returns true if the given object is a socket
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

// IsBlockDevice returns true if the given object is a device
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

// IsCharacterDevice returns true if the given object is a character device
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

// IsDir returns true if the given object is a directory
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

// IsLink returns true if the given object is a link
func IsLink(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	var buf = make([]byte, 1)
	_, err := syscall.Readlink(path, buf)

	return err == nil
}

// IsReadable returns true if given object is readable by current user
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

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isReadableStat(stat, user.UID, getGIDList(user))
}

// IsReadableByUser returns true if given object is readable by some user
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

// IsWritable returns true if given object is writable by current user
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

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isWritableStat(stat, user.UID, getGIDList(user))
}

// IsWritableByUser returns true if given object is writable by some user
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

// IsExecutable returns true if given object is executable by current user
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

	user, err := system.CurrentUser()

	if err != nil {
		return false
	}

	return isExecutableStat(stat, user.UID, getGIDList(user))
}

// IsExecutableByUser returns true if given object is executable by some user
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

// IsNonEmpty returns true if given file is not empty
func IsNonEmpty(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return GetSize(path) > 0
}

// IsEmptyDir returns true if given directory es empty
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

// GetOwner returns object owner UID and GID
func GetOwner(path string) (int, int, error) {
	if path == "" {
		return -1, -1, ErrEmptyPath
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, -1, err
	}

	return int(stat.Uid), int(stat.Gid), nil
}

// GetATime returns time of last access
func GetATime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	atime, _, _, err := GetTimes(path)

	return atime, err
}

// GetCTime returns time of creation
func GetCTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, _, ctime, err := GetTimes(path)

	return ctime, err
}

// GetMTime returns time of modification
func GetMTime(path string) (time.Time, error) {
	if path == "" {
		return time.Time{}, ErrEmptyPath
	}

	path = PATH.Clean(path)

	_, mtime, _, err := GetTimes(path)

	return mtime, err
}

// GetSize returns file size in bytes
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

// GetMode returns file mode bits
func GetMode(path string) os.FileMode {
	if path == "" {
		return 0
	}

	path = PATH.Clean(path)

	return os.FileMode(getMode(path) & 0777)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMode(path string) uint32 {
	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return 0
	}

	return uint32(stat.Mode)
}

func isReadableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
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

func isWritableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
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

func isExecutableStat(stat *syscall.Stat_t, uid int, gids []int) bool {
	if uid == 0 {
		return true
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
