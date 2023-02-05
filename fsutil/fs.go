//go:build !windows
// +build !windows

// Package fsutil provides methods for working with files on POSIX compatible systems
package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	PATH "github.com/essentialkaos/ek/v12/path"
	"github.com/essentialkaos/ek/v12/system"
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

// ErrEmptyPath can be returned by different methods if given path is empty and can't be
// used
var ErrEmptyPath = errors.New("Path is empty")

// ////////////////////////////////////////////////////////////////////////////////// //

// CheckPerms checks many props at once
//
//    * F - is file
//    * D - is directory
//    * X - is executable
//    * L - is link
//    * W - is writable
//    * R - is readable
//    * B - is block device
//    * C - is character device
//    * S - not empty (only for files)
//
func CheckPerms(props, path string) bool {
	if props == "" || path == "" {
		return false
	}

	path = PATH.Clean(path)
	props = strings.ToUpper(props)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		return false
	}

	user, err := getCurrentUser()

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

// ValidatePerms validates permissions for file or directory
func ValidatePerms(props, path string) error {
	switch {
	case props == "":
		return errors.New("Props is empty")
	case path == "":
		return errors.New("Path is empty")
	}

	path = PATH.Clean(path)
	props = strings.ToUpper(props)

	var stat = &syscall.Stat_t{}

	err := syscall.Stat(path, stat)

	if err != nil {
		switch {
		case strings.ContainsRune(props, 'F'):
			return fmt.Errorf("File %s doesn't exist or not accessible", path)
		case strings.ContainsRune(props, 'D'):
			return fmt.Errorf("Directory %s doesn't exist or not accessible", path)
		case strings.ContainsRune(props, 'B'):
			return fmt.Errorf("Block device %s doesn't exist or not accessible", path)
		case strings.ContainsRune(props, 'C'):
			return fmt.Errorf("Character device %s doesn't exist or not accessible", path)
		case strings.ContainsRune(props, 'L'):
			return fmt.Errorf("Link %s doesn't exist or not accessible", path)
		}

		return fmt.Errorf("Object %s doesn't exist or not accessible", path)
	}

	user, err := getCurrentUser()

	if err != nil {
		return errors.New("Can't get information about the current user")
	}

	for _, k := range props {
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

	user, err := getCurrentUser()

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

	user, err := getCurrentUser()

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

	user, err := getCurrentUser()

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

// IsEmpty returns true if given file is empty
func IsEmpty(path string) bool {
	if path == "" {
		return false
	}

	path = PATH.Clean(path)

	return GetSize(path) == 0
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
		return -1, -1, fmt.Errorf("Can't get owner info for %q: %w", path, err)
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

func getObjectType(stat *syscall.Stat_t) string {
	switch {
	case stat.Mode&_IFMT == _IFREG:
		return "File"
	case stat.Mode&_IFMT == _IFDIR:
		return "Directory"
	case stat.Mode&_IFMT != _IFBLK:
		return "Block device"
	case stat.Mode&_IFMT != _IFCHR:
		return "Character device"
	}

	return "Object"
}
