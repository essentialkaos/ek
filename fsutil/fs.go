// +build !windows

// Package fsutil provides methods for working with files on posix compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"time"

	PATH "pkg.re/essentialkaos/ek.v5/path"
	"pkg.re/essentialkaos/ek.v5/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_IFMT   = 0170000
	_IFSOCK = 0140000
	_IFLNK  = 0120000
	_IFREG  = 0100000
	_IFBLK  = 0060000
	_IFDIR  = 0040000
	_IFCHR  = 0020000
	_IRUSR  = 00400
	_IWUSR  = 00200
	_IXUSR  = 00100
	_IRGRP  = 00040
	_IWGRP  = 00020
	_IXGRP  = 00010
	_IROTH  = 00004
	_IWOTH  = 00002
	_IXOTH  = 00001
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CheckPerms check many props at once.
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

	var user *system.User

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
			if stat.Mode&_IFMT != _IFLNK {
				return false
			}

		case 'X':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

			if !isExecutableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'W':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

			if !isWritableStat(stat, user.UID, getGIDList(user)) {
				return false
			}

		case 'R':
			if user == nil {
				user, err = system.CurrentUser()

				if err != nil {
					return false
				}
			}

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

// ProperPath return first proper path from given slice
func ProperPath(props string, paths []string) string {
	for _, path := range paths {
		path = PATH.Clean(path)

		if CheckPerms(props, path) {
			return path
		}
	}

	return ""
}

// IsExist check if target is exist in fs or not
func IsExist(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return syscall.Access(path, syscall.F_OK) == nil
}

// IsRegular check if target is regular file or not
func IsRegular(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFREG
}

// IsSocket check if target is socket or not
func IsSocket(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFSOCK
}

// IsBlockDevice check if target is block device or not
func IsBlockDevice(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFBLK
}

// IsCharacterDevice check if target is character device or not
func IsCharacterDevice(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFCHR
}

// IsDir check if target is directory or not
func IsDir(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFDIR
}

// IsLink check if file is link or not
func IsLink(path string) bool {
	path = PATH.Clean(path)
	mode := getMode(path)

	if mode == 0 {
		return false
	}

	return mode&_IFMT == _IFLNK
}

// IsReadable check if file is readable or not
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

// IsWritable check if file is writable or not
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

// IsExecutable check if file is executable or not
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

// IsNonEmpty check if file is empty or not
func IsNonEmpty(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return GetSize(path) != 0
}

// IsEmptyDir check if directory empty or not
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

	if n == 0x30 || err != nil {
		return true
	}

	return false
}

// GetOwner return object owner pid and gid
func GetOwner(path string) (int, int, error) {
	if path == "" {
		return -1, -1, errors.New("Path is empty")
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return -1, -1, err
	}

	return int(stat.Uid), int(stat.Gid), nil
}

// GetATime return time of last access
func GetATime(path string) (time.Time, error) {
	path = PATH.Clean(path)

	atime, _, _, err := GetTimes(path)

	return atime, err
}

// GetCTime return time of creation
func GetCTime(path string) (time.Time, error) {
	path = PATH.Clean(path)

	_, _, ctime, err := GetTimes(path)

	return ctime, err
}

// GetMTime return time of modification
func GetMTime(path string) (time.Time, error) {
	path = PATH.Clean(path)

	_, mtime, _, err := GetTimes(path)

	return mtime, err
}

// GetSize return file size in bytes
func GetSize(path string) int64 {
	if path == "" {
		return 0
	}

	path = PATH.Clean(path)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return 0
	}

	return stat.Size
}

// GetPerm return file permissions
func GetPerm(path string) os.FileMode {
	path = PATH.Clean(path)
	return os.FileMode(getMode(path) & 0777)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getMode(path string) uint32 {
	if path == "" {
		return 0
	}

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

	switch {
	case stat.Mode&_IROTH == _IROTH:
		return true
	case stat.Mode&_IRUSR == _IRUSR && uid == int(stat.Uid):
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

	switch {
	case stat.Mode&_IWOTH == _IWOTH:
		return true
	case stat.Mode&_IWUSR == _IWUSR && uid == int(stat.Uid):
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

	switch {
	case stat.Mode&_IXOTH == _IXOTH:
		return true
	case stat.Mode&_IXUSR == _IXUSR && uid == int(stat.Uid):
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
		return []int{}
	}

	var result []int

	for _, group := range user.Groups {
		result = append(result, group.GID)
	}

	return result
}
