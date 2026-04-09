// Package tmp provides methods and structs for working with temporary data
package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/essentialkaos/ek/v13/fsutil"
	"github.com/essentialkaos/ek/v13/rand"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Temp holds the configuration and the list of temporary objects created during
// the session, allowing them to be removed all at once via Clean
type Temp struct {
	Dir       string
	DirPerms  os.FileMode
	FilePerms os.FileMode

	mu      sync.Mutex
	objects []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ErrNilTemp is returned when a method is called on a nil Temp pointer
var ErrNilTemp = fmt.Errorf("temp struct is nil")

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is the default path used as the root for new Temp instances
var Dir = os.TempDir()

// DefaultDirPerms is the default permission mode applied to temporary directories
var DefaultDirPerms = os.FileMode(0750)

// DefaultFilePerms is the default permission mode applied to temporary files
var DefaultFilePerms = os.FileMode(0640)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTemp creates a new [Temp] instance rooted at the given directory, or at the
// package-level Dir if no argument is provided
func NewTemp(dir ...string) (*Temp, error) {
	tempDir := filepath.Clean(Dir)

	if len(dir) != 0 {
		tempDir = filepath.Clean(dir[0])
	}

	if !fsutil.IsExist(tempDir) {
		return nil, fmt.Errorf("directory %s does not exist", tempDir)
	}

	if !fsutil.IsDir(tempDir) {
		return nil, fmt.Errorf("%s is not a directory", tempDir)
	}

	if !fsutil.IsWritable(tempDir) {
		return nil, fmt.Errorf("directory %s is not writable", tempDir)
	}

	return &Temp{
		Dir:       tempDir,
		DirPerms:  DefaultDirPerms,
		FilePerms: DefaultFilePerms,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// MkDir creates a temporary directory inside t.Dir and registers it for cleanup.
// An optional nameSuffix is appended to the auto-generated name.
func (t *Temp) MkDir(nameSuffix ...string) (string, error) {
	if t == nil {
		return "", ErrNilTemp
	}

	name := strutil.Q(nameSuffix...)
	tmpDir := getTempName(t.Dir, name)
	err := os.MkdirAll(tmpDir, t.DirPerms)

	if err != nil {
		return "", err
	}

	t.addObject(tmpDir)

	return tmpDir, nil
}

// MkFile creates a temporary file inside t.Dir and registers it for cleanup.
// An optional nameSuffix is appended to the auto-generated name.
func (t *Temp) MkFile(nameSuffix ...string) (*os.File, string, error) {
	if t == nil {
		return nil, "", ErrNilTemp
	}

	name := strutil.Q(nameSuffix...)
	tmpFile := getTempName(t.Dir, name)
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, t.FilePerms)

	if err != nil {
		return nil, "", err
	}

	t.addObject(tmpFile)

	return fd, tmpFile, nil
}

// MkName generates and registers a unique name for a temporary object inside
// t.Dir without creating the object itself. An optional nameSuffix is appended.
func (t *Temp) MkName(nameSuffix ...string) string {
	if t == nil {
		return ""
	}

	name := strutil.Q(nameSuffix...)
	tmpObj := getTempName(t.Dir, name)

	t.addObject(tmpObj)

	return tmpObj
}

// Clean removes all temporary objects (files and directories) registered in
// this [Temp] instance
func (t *Temp) Clean() {
	if t == nil || len(t.objects) == 0 {
		return
	}

	t.mu.Lock()
	objects := t.objects
	t.objects = nil
	t.mu.Unlock()

	for _, object := range objects {
		os.RemoveAll(object) //nolint:errcheck
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// addObject adds object to temp objects list
func (t *Temp) addObject(path string) {
	t.mu.Lock()
	t.objects = append(t.objects, path)
	t.mu.Unlock()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getTempName generates a unique temporary name based on the current time and
// a random string
func getTempName(dir, name string) string {
	var result string

	for {
		if name != "" {
			result = filepath.Join(dir, fmt.Sprintf(
				"tmp_%d_%s_%s", time.Now().UnixMilli(), rand.String(8), name,
			))
		} else {
			result = filepath.Join(dir, fmt.Sprintf(
				"tmp_%d_%s", time.Now().UnixMilli(), rand.String(8),
			))
		}

		if !fsutil.IsExist(result) {
			break
		}
	}

	return result
}
