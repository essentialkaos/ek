// Package tmp provides methods and structs for working with temporary data
package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"path"

	"pkg.re/essentialkaos/ek.v9/fsutil"
	"pkg.re/essentialkaos/ek.v9/rand"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Temp is basic temp struct
type Temp struct {
	Dir       string
	DirPerms  os.FileMode
	FilePerms os.FileMode

	targets []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is default temporary directory
var Dir = "/tmp"

// DefaultDirPerms is default permissions for directories
var DefaultDirPerms os.FileMode = 0750

// DefaultFilePerms is default permissions for files
var DefaultFilePerms os.FileMode = 0640

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTemp create new Temp structure
func NewTemp(args ...string) (*Temp, error) {
	tempDir := path.Clean(Dir)

	if len(args) != 0 {
		tempDir = path.Clean(args[0])
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

// MkDir make temporary directory
func (t *Temp) MkDir(args ...string) (string, error) {
	if t == nil {
		return "", fmt.Errorf("Temp struct is nil")
	}

	name := ""

	if len(args) != 0 {
		name = args[0]
	}

	tmpDir := getTempName(t.Dir, name)
	err := os.MkdirAll(tmpDir, t.DirPerms)

	if err != nil {
		return "", err
	}

	t.targets = append(t.targets, tmpDir)

	return tmpDir, err
}

// MkFile make temporary file
func (t *Temp) MkFile(args ...string) (*os.File, string, error) {
	if t == nil {
		return nil, "", fmt.Errorf("Temp struct is nil")
	}

	name := ""

	if len(args) != 0 {
		name = args[0]
	}

	tmpFile := getTempName(t.Dir, name)
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, t.FilePerms)

	if err != nil {
		return nil, "", err
	}

	t.targets = append(t.targets, tmpFile)

	return fd, tmpFile, nil
}

// MkName return name for temporary object
func (t *Temp) MkName(args ...string) string {
	name := ""

	if len(args) != 0 {
		name = args[0]
	}

	tmpObj := getTempName(t.Dir, name)
	t.targets = append(t.targets, tmpObj)

	return tmpObj
}

// Clean remove all temporary targets
func (t *Temp) Clean() {
	if t == nil || t.targets == nil || len(t.targets) == 0 {
		return
	}

	for _, target := range t.targets {
		os.RemoveAll(target)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getTempName return name of temporary file
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
