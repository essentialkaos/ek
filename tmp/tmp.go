// Package tmp provides methods and structs for working with temporary data
package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path"

	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/rand"
	"github.com/essentialkaos/ek/v12/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Temp is basic temp struct
type Temp struct {
	Dir       string
	DirPerms  os.FileMode
	FilePerms os.FileMode

	objects []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is path to temporary directory
var Dir = "/tmp"

// DefaultDirPerms is default permissions for directories
var DefaultDirPerms = os.FileMode(0750)

// DefaultFilePerms is default permissions for files
var DefaultFilePerms = os.FileMode(0640)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTemp creates new Temp structure
func NewTemp(dir ...string) (*Temp, error) {
	tempDir := path.Clean(Dir)

	if len(dir) != 0 {
		tempDir = path.Clean(dir[0])
	}

	if !fsutil.IsExist(tempDir) {
		return nil, fmt.Errorf("Directory %s does not exist", tempDir)
	}

	if !fsutil.IsDir(tempDir) {
		return nil, fmt.Errorf("%s is not a directory", tempDir)
	}

	if !fsutil.IsWritable(tempDir) {
		return nil, fmt.Errorf("Directory %s is not writable", tempDir)
	}

	return &Temp{
		Dir:       tempDir,
		DirPerms:  DefaultDirPerms,
		FilePerms: DefaultFilePerms,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// MkDir creates temporary directory
func (t *Temp) MkDir(nameSuffix ...string) (string, error) {
	if t == nil {
		return "", fmt.Errorf("Temp struct is nil")
	}

	name := strutil.Q(nameSuffix...)
	tmpDir := getTempName(t.Dir, name)
	err := os.MkdirAll(tmpDir, t.DirPerms)

	if err != nil {
		return "", err
	}

	t.objects = append(t.objects, tmpDir)

	return tmpDir, err
}

// MkFile creates temporary file
func (t *Temp) MkFile(nameSuffix ...string) (*os.File, string, error) {
	if t == nil {
		return nil, "", fmt.Errorf("Temp struct is nil")
	}

	name := strutil.Q(nameSuffix...)
	tmpFile := getTempName(t.Dir, name)
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, t.FilePerms)

	if err != nil {
		return nil, "", err
	}

	t.objects = append(t.objects, tmpFile)

	return fd, tmpFile, nil
}

// MkName returns name for temporary object (file or directory)
func (t *Temp) MkName(nameSuffix ...string) string {
	name := strutil.Q(nameSuffix...)
	tmpObj := getTempName(t.Dir, name)
	t.objects = append(t.objects, tmpObj)

	return tmpObj
}

// Clean removes all temporary objects (files and directories)
func (t *Temp) Clean() {
	if t == nil || t.objects == nil || len(t.objects) == 0 {
		return
	}

	for _, object := range t.objects {
		os.RemoveAll(object)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getTempName returns name of temporary file
func getTempName(dir, name string) string {
	var result string

	for {
		if name != "" {
			result = path.Join(dir, "_"+rand.String(12)+"_"+name)
		} else {
			result = path.Join(dir, "_tmp_"+rand.String(12))
		}

		if !fsutil.IsExist(result) {
			break
		}
	}

	return result
}
