package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fsutil"
	"ek/rand"
	"os"
	"path"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Temp is basic temp struct
type Temp struct {
	Dir string

	targets []string
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Dir is default temporary directory
var Dir = "/tmp"

// ////////////////////////////////////////////////////////////////////////////////// //

// NewTemp create new Temp structure
func NewTemp(args ...string) *Temp {
	if len(args) == 0 {
		return &Temp{Dir: Dir}
	}

	return &Temp{Dir: path.Clean(args[0])}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// MkDir make temporary directory
func (t *Temp) MkDir(args ...string) (string, error) {
	name := ""

	if len(args) != 0 {
		name = args[0]
	}

	tmpDir := getTempName(t.Dir, name)
	err := os.MkdirAll(tmpDir, 0750)

	if err == nil {
		t.targets = append(t.targets, tmpDir)
	}

	return tmpDir, err
}

// MkFile make temporary file
func (t *Temp) MkFile(args ...string) (*os.File, string, error) {
	name := ""

	if len(args) != 0 {
		name = args[0]
	}

	tmpFile := getTempName(t.Dir, name)
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE, 0640)

	if err != nil {
		return fd, "", err
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
	if t.targets == nil || len(t.targets) == 0 {
		return
	}

	for _, target := range t.targets {
		os.RemoveAll(target)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getTempName(dir, name string) string {
	var result string

	for {
		if name != "" {
			result = path.Join(dir, "_"+name+"_"+rand.String(12))
		} else {
			result = path.Join(dir, "_tmp_"+rand.String(12))
		}

		if !fsutil.IsExist(result) {
			break
		}
	}

	return result
}
