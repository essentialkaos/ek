package z7

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/essentialkaos/ek/fsutil"
	"github.com/essentialkaos/ek/mathutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _BINARY = "7za"

const (
	_COMPRESSION_MIN     = 0
	_COMPRESSION_MAX     = 9
	_COMPRESSION_DEFAULT = 4
)

const (
	_COMMAND_ADD       = "a"
	_COMMAND_BENCHMARK = "b"
	_COMMAND_DELETE    = "d"
	_COMMAND_LIST      = "l"
	_COMMAND_TEST      = "t"
	_COMMAND_UPDATE    = "u"
	_COMMAND_EXTRACT   = "x"
)

const (
	_TYPE_7Z   = "7z"
	_TYPE_ZIP  = "zip"
	_TYPE_GZIP = "gzip"
	_TYPE_XZ   = "xz"
	_TYPE_BZIP = "bzip2"
)

const _TEST_OK_VALUE = "Everything is Ok"

// Props contains properties for packing/unpacking data
type Props struct {
	Dir            string
	File           string
	FileList       string
	Exclude        string
	ExcludeFile    string
	Compression    int
	Output         string
	Password       string
	Threads        int
	EncryptHeaders bool
	Recurse        bool
	WorkingDir     string
}

// Info contains info about archive
type Info struct {
	Path         string
	Type         string
	Method       []string
	Solid        bool
	Blocks       int
	PhysicalSize int
	HeadersSize  int
	Files        []*FileInfo
}

// FileInfo contains info about file inside archive
type FileInfo struct {
	Path       string
	Folder     string
	Size       int
	PackedSize int
	Modified   time.Time
	Created    time.Time
	Accessed   time.Time
	Attributes string
	CRC        int
	Encrypted  bool
	Method     []string
	Block      int
	Comment    string
	HostOS     string
	Version    int
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add add files to archive
func Add(arch interface{}, files ...string) (string, error) {
	return AddList(arch, files)
}

// AddList add files as string slice
func AddList(arch interface{}, files ...[]string) (string, error) {
	props, err := procProps(arch)

	if err != nil {
		return "", err
	}

	var (
		cwd  string
		file string
	)

	file, _ = filepath.Abs(props.File)

	if props.Dir != "" {
		cwd, _ = os.Getwd()
		os.Chdir(props.Dir)
	}

	args := propsToArgs(props, _COMMAND_ADD)

	if len(files) != 0 {
		args = append(args, files[0]...)
	}

	out, err := execBinary(file, _COMMAND_ADD, args)

	if props.Dir != "" {
		os.Chdir(cwd)
	}

	return out, err
}

// Extract extract arhive
func Extract(arch interface{}) (string, error) {
	props, err := procProps(arch)

	if err != nil {
		return "", err
	}

	if !fsutil.IsExist(props.File) {
		return "", errors.New("File " + props.File + " is not exist")
	}

	if !fsutil.IsReadable(props.File) {
		return "", errors.New("File " + props.File + " is not readable")
	}

	if props.Output != "" && !fsutil.IsWritable(props.Output) {
		return "", errors.New("Directory " + props.Output + " is not writable")
	}

	var (
		cwd  string
		file string
	)

	file, _ = filepath.Abs(props.File)

	if props.Dir != "" {
		cwd, _ = os.Getwd()
		os.Chdir(props.Dir)
	}

	args := propsToArgs(props, _COMMAND_EXTRACT)
	out, err := execBinary(file, _COMMAND_EXTRACT, args)

	if props.Dir != "" {
		os.Chdir(cwd)
	}

	return out, err
}

// List return info about archive
func List(arch interface{}) (*Info, error) {
	props, err := procProps(arch)

	if err != nil {
		return &Info{}, err
	}

	if !fsutil.IsExist(props.File) {
		return nil, errors.New("File " + props.File + " is not exist")
	}

	if !fsutil.IsReadable(props.File) {
		return nil, errors.New("File " + props.File + " is not readable")
	}

	args := propsToArgs(props, _COMMAND_LIST)
	out, err := execBinary(props.File, _COMMAND_LIST, args)

	if err != nil {
		return nil, errors.New(out)
	}

	return parseInfoString(out), nil
}

// Test test archive
func Test(arch interface{}) (string, bool) {
	props, err := procProps(arch)

	if err != nil {
		return "", false
	}

	if !fsutil.IsExist(props.File) {
		return "", false
	}

	if !fsutil.IsReadable(props.File) {
		return "", false
	}

	args := propsToArgs(props, _COMMAND_TEST)
	out, err := execBinary(props.File, _COMMAND_TEST, args)

	for _, line := range strings.Split(out, "\n") {
		if line == _TEST_OK_VALUE {
			return out, true
		}
	}

	return out, false
}

// Delete remove files from archive
func Delete(arch interface{}, files ...string) (string, error) {
	return DeleteList(arch, files)
}

// DeleteList remove files provided as string slice from archive
func DeleteList(arch interface{}, files ...[]string) (string, error) {
	props, err := procProps(arch)

	if err != nil {
		return "", err
	}

	args := propsToArgs(props, _COMMAND_DELETE)

	if len(files) != 0 {
		args = append(args, files[0]...)
	}

	out, err := execBinary(props.File, _COMMAND_DELETE, args)

	return out, err
}

// ////////////////////////////////////////////////////////////////////////////////// //

func execBinary(arch string, command string, args []string) (string, error) {
	var cmd = exec.Command(_BINARY)

	cmd.Args = append(cmd.Args, command)
	cmd.Args = append(cmd.Args, arch)
	cmd.Args = append(cmd.Args, args...)

	out, err := cmd.Output()

	return string(out[:]), err
}

func procProps(p interface{}) (*Props, error) {
	switch p.(type) {
	case *Props:
		return p.(*Props), nil
	case string:
		return &Props{File: p.(string), Compression: _COMPRESSION_DEFAULT}, nil
	default:
		return nil, errors.New("Unknown properties type")
	}
}

func propsToArgs(props *Props, command string) []string {
	var args = []string{"", "-y"}

	if command == _COMMAND_ADD {
		compLvl := strconv.Itoa(mathutil.Between(props.Compression, _COMPRESSION_MIN, _COMPRESSION_MAX))

		args = append(args, "-mx="+compLvl)

		switch {
		case props.Threads == -1:
			args = append(args, "-mt=off")
		case props.Threads > 0:
			args = append(args, "-mt="+strconv.Itoa(mathutil.Between(props.Threads, 1, 128)))
		}

		if props.EncryptHeaders {
			args = append(args, "-mhe")
		}

		if props.Exclude != "" {
			args = append(args, "-x"+props.Exclude)
		} else if props.ExcludeFile != "" {
			args = append(args, "-x@"+props.ExcludeFile)
		}

		if props.FileList != "" {
			args = append(args, "-i@"+props.FileList)
		}

	} else if command == _COMMAND_EXTRACT {
		if props.Output != "" {
			args = append(args, "-o"+props.Output)
		}
	} else if command == _COMMAND_LIST {
		args = append(args, "-slt")
	}

	if props.Password != "" {
		args = append(args, "-p"+props.Password)
	}

	if props.Recurse {
		args = append(args, "-r")
	}

	if props.WorkingDir != "" {
		args = append(args, "-w"+props.WorkingDir)
	}

	return args
}

func parseInfoString(infoData string) *Info {
	var data = strings.Split(infoData, "\n")
	var info = &Info{}

	var fStart, fCount, fRecs int
	var rList map[string]string

	info.Path = getValue(data[7])
	info.Type = getValue(data[8])

	if info.Type == _TYPE_7Z {
		fStart = 16
		fRecs = 10

		rList = parseRecordList(data[7:15])

		info.Method = strings.Split(rList["Method"], " ")
		info.Solid = rList["Solid"] == "+"

		info.Blocks, _ = strconv.Atoi(rList["Blocks"])
		info.PhysicalSize, _ = strconv.Atoi(rList["Physical Size"])
		info.HeadersSize, _ = strconv.Atoi(rList["Headers Size"])
	} else if info.Type == _TYPE_ZIP {
		fStart = 12
		fRecs = 14

		rList = parseRecordList(data[7:11])

		info.PhysicalSize, _ = strconv.Atoi(rList["Physical Size"])
	} else if info.Type == _TYPE_GZIP {
		fStart = 11
		fRecs = 7
	} else if info.Type == _TYPE_XZ {
		fStart = 12
		fRecs = 4

		rList = parseRecordList(data[7:11])

		info.Method = strings.Split(rList["Method"], " ")
	} else if info.Type == _TYPE_BZIP {
		fStart = 11
		fRecs = 2
	} else {
		return info
	}

	fCount = (len(data) - fStart - 1) / fRecs

	for i := 0; i < fCount; i++ {
		start := (i * fRecs) + fStart
		end := start + fRecs

		info.Files = append(info.Files, parseFileInfoString(data[start:end]))
	}

	return info
}

func parseFileInfoString(data []string) *FileInfo {
	var info = &FileInfo{}
	var rList = parseRecordList(data)

	crc, _ := strconv.ParseInt(rList["CRC"], 16, 0)

	info.Path = rList["Path"]
	info.Folder = rList["Folder"]
	info.Size, _ = strconv.Atoi(rList["Size"])
	info.PackedSize, _ = strconv.Atoi(rList["Packed Size"])
	info.Modified = parseDateString(rList["Modified"])
	info.Created = parseDateString(rList["Created"])
	info.Accessed = parseDateString(rList["Accessed"])
	info.Attributes = rList["Attributes"]
	info.CRC = int(crc)
	info.Comment = rList["Comment"]
	info.Encrypted = rList["Encrypted"] == "+"
	info.Method = strings.Split(rList["Method"], " ")
	info.Block, _ = strconv.Atoi(rList["Block"])
	info.HostOS = rList["Host OS"]
	info.Version, _ = strconv.Atoi(rList["Version"])

	return info
}

func parseRecordList(data []string) map[string]string {
	var result = make(map[string]string)

	for _, rec := range data {
		if rec != "" {
			revVal := strings.Split(rec, " = ")
			result[revVal[0]] = revVal[1]
		}
	}

	return result
}

func parseDateString(data string) time.Time {
	if data == "" {
		return time.Time{}
	}

	year, _ := strconv.Atoi(data[0:4])
	month, _ := strconv.Atoi(data[5:7])
	day, _ := strconv.Atoi(data[8:10])
	hour, _ := strconv.Atoi(data[11:13])
	min, _ := strconv.Atoi(data[14:16])
	sec, _ := strconv.Atoi(data[17:19])

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
}

func getValue(s string) string {
	return strings.Split(s, " = ")[1]
}
