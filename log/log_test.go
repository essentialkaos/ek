package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"ek/fsutil"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type LogSuite struct {
	TempDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&LogSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ls *LogSuite) SetUpSuite(c *C) {
	ls.TempDir = c.MkDir()
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (ls *LogSuite) TestWithoutPrefixes(c *C) {
	logfile := ls.TempDir + "/file1.log"
	l, err := New(logfile, 0644)

	l.MinLevel(DEBUG)

	c.Assert(l, Not(IsNil))
	c.Assert(err, IsNil)

	l.PrefixDebug = false
	l.PrefixInfo = false
	l.PrefixWarn = false
	l.PrefixError = false
	l.PrefixCrit = false

	c.Assert(fsutil.GetPerm(logfile), Equals, os.FileMode(0644))

	l.Print(DEBUG, "Test debug %d", DEBUG)
	l.Print(INFO, "Test info %d", INFO)
	l.Print(WARN, "Test warn %d", WARN)
	l.Print(ERROR, "Test error %d", ERROR)
	l.Print(CRIT, "Test crit %d", CRIT)

	l.Print(DEBUG, "Test debug")
	l.Print(INFO, "Test info")
	l.Print(WARN, "Test warn")
	l.Print(ERROR, "Test error")
	l.Print(CRIT, "Test crit")

	l.Debug("Test debug %d\n", DEBUG)
	l.Info("Test info %d\n", INFO)
	l.Warn("Test warn %d\n", WARN)
	l.Error("Test error %d\n", ERROR)
	l.Crit("Test crit %d\n", CRIT)

	l.Flush()

	data, err := ioutil.ReadFile(logfile)

	c.Assert(len(data), Not(Equals), 0)
	c.Assert(err, IsNil)

	dataSlice := strings.Split(string(data[:]), "\n")

	c.Assert(len(dataSlice), Equals, 16)

	c.Assert(dataSlice[0][28:], Equals, "Test debug 0")
	c.Assert(dataSlice[1][28:], Equals, "Test info 1")
	c.Assert(dataSlice[2][28:], Equals, "Test warn 2")
	c.Assert(dataSlice[3][28:], Equals, "Test error 3")
	c.Assert(dataSlice[4][28:], Equals, "Test crit 4")

	c.Assert(dataSlice[5][28:], Equals, "Test debug")
	c.Assert(dataSlice[6][28:], Equals, "Test info")
	c.Assert(dataSlice[7][28:], Equals, "Test warn")
	c.Assert(dataSlice[8][28:], Equals, "Test error")
	c.Assert(dataSlice[9][28:], Equals, "Test crit")

	c.Assert(dataSlice[10][28:], Equals, "Test debug 0")
	c.Assert(dataSlice[11][28:], Equals, "Test info 1")
	c.Assert(dataSlice[12][28:], Equals, "Test warn 2")
	c.Assert(dataSlice[13][28:], Equals, "Test error 3")
	c.Assert(dataSlice[14][28:], Equals, "Test crit 4")
}

func (ls *LogSuite) TestWithPrefixes(c *C) {
	logfile := ls.TempDir + "/file2.log"
	l, err := New(logfile, 0644)

	l.MinLevel(DEBUG)

	c.Assert(l, Not(IsNil))
	c.Assert(err, IsNil)

	l.PrefixDebug = true
	l.PrefixInfo = true
	l.PrefixWarn = true
	l.PrefixError = true
	l.PrefixCrit = true

	c.Assert(fsutil.GetPerm(logfile), Equals, os.FileMode(0644))

	l.Print(DEBUG, "Test debug %d", DEBUG)
	l.Print(INFO, "Test info %d", INFO)
	l.Print(WARN, "Test warn %d", WARN)
	l.Print(ERROR, "Test error %d", ERROR)
	l.Print(CRIT, "Test crit %d", CRIT)

	l.Print(DEBUG, "Test debug")
	l.Print(INFO, "Test info")
	l.Print(WARN, "Test warn")
	l.Print(ERROR, "Test error")
	l.Print(CRIT, "Test crit")

	l.Debug("Test debug %d", DEBUG)
	l.Info("Test info %d", INFO)
	l.Warn("Test warn %d", WARN)
	l.Error("Test error %d", ERROR)
	l.Crit("Test crit %d", CRIT)

	l.Flush()

	data, err := ioutil.ReadFile(logfile)

	c.Assert(len(data), Not(Equals), 0)
	c.Assert(err, IsNil)

	dataSlice := strings.Split(string(data[:]), "\n")

	c.Assert(len(dataSlice), Equals, 16)

	c.Assert(dataSlice[0][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[1][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[2][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[3][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[4][28:], Equals, "[CRITICAL] Test crit 4")

	c.Assert(dataSlice[5][28:], Equals, "[DEBUG] Test debug")
	c.Assert(dataSlice[6][28:], Equals, "[INFO] Test info")
	c.Assert(dataSlice[7][28:], Equals, "[WARNING] Test warn")
	c.Assert(dataSlice[8][28:], Equals, "[ERROR] Test error")
	c.Assert(dataSlice[9][28:], Equals, "[CRITICAL] Test crit")

	c.Assert(dataSlice[10][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[11][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[12][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[13][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[14][28:], Equals, "[CRITICAL] Test crit 4")
}
