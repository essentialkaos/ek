package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fsutil"

	. "github.com/essentialkaos/check"
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

func (ls *LogSuite) SetUpTest(c *C) {
	Global = &Logger{
		PrefixWarn:  true,
		PrefixError: true,
		PrefixCrit:  true,

		minLevel: INFO,
		mu:       &sync.Mutex{},
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (ls *LogSuite) TestErrors(c *C) {
	var l *Logger

	l.MinLevel(DEBUG)

	err := l.Flush()

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Reopen()

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Print(DEBUG, "test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Debug("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Info("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Warn("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Error("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Crit("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Aux("test")

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	err = l.Set("", 0)

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	_, err = New("/_not_exist_", 0644)

	c.Assert(err, NotNil)

	err = Reopen()

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrOutputNotSet)

	l.EnableBufIO(time.Second)
}

func (ls *LogSuite) TestLevel(c *C) {
	l := &Logger{minLevel: WARN, mu: &sync.Mutex{}}

	c.Assert(l.Is(DEBUG), Equals, false)
	c.Assert(l.Is(WARN), Equals, true)
	c.Assert(l.Is(ERROR), Equals, true)

	c.Assert(l.MinLevel(-1), IsNil)
	c.Assert(l.MinLevel(6), IsNil)
	c.Assert(l.MinLevel("debug"), IsNil)
	c.Assert(l.MinLevel("info"), IsNil)
	c.Assert(l.MinLevel("warn"), IsNil)
	c.Assert(l.MinLevel("warning"), IsNil)
	c.Assert(l.MinLevel("error"), IsNil)
	c.Assert(l.MinLevel("crit"), IsNil)
	c.Assert(l.MinLevel("critical"), IsNil)
	c.Assert(l.MinLevel(int8(1)), IsNil)
	c.Assert(l.MinLevel(int16(1)), IsNil)
	c.Assert(l.MinLevel(int32(1)), IsNil)
	c.Assert(l.MinLevel(int64(1)), IsNil)
	c.Assert(l.MinLevel(uint(1)), IsNil)
	c.Assert(l.MinLevel(uint8(1)), IsNil)
	c.Assert(l.MinLevel(uint16(1)), IsNil)
	c.Assert(l.MinLevel(uint32(1)), IsNil)
	c.Assert(l.MinLevel(uint64(1)), IsNil)
	c.Assert(l.MinLevel(float32(1)), IsNil)
	c.Assert(l.MinLevel(float64(1)), IsNil)

	c.Assert(l.MinLevel("abcd"), NotNil)
	c.Assert(l.MinLevel(time.Now()), NotNil)

	l.MinLevel("crit")

	err := l.Print(ERROR, "error")

	c.Assert(err, IsNil)
}

func (ls *LogSuite) TestFlush(c *C) {
	l := &Logger{mu: &sync.Mutex{}}

	err := l.Flush()

	c.Assert(err, IsNil)
}

func (ls *LogSuite) TestReopenAndSet(c *C) {
	l, err := New(ls.TempDir+"/set-test.log", 0644)

	c.Assert(err, IsNil)

	err = l.Set(ls.TempDir+"/set-test-2.log", 0644)

	c.Assert(err, IsNil)

	err = l.Reopen()

	c.Assert(err, IsNil)
}

func (ls *LogSuite) TestStdOutput(c *C) {
	var err error

	l := &Logger{mu: &sync.Mutex{}}

	err = l.Print(INFO, "info")

	c.Assert(err, IsNil)

	err = l.Print(ERROR, "error")

	c.Assert(err, IsNil)

	fmtc.DisableColors = true

	l.UseColors = true
	l.PrefixError = true

	err = l.Print(INFO, "info")

	c.Assert(err, IsNil)

	err = l.Print(ERROR, "error")

	c.Assert(err, IsNil)

	fmtc.DisableColors = false
}

func (ls *LogSuite) TestWithoutPrefixes(c *C) {
	logfile := ls.TempDir + "/file1.log"
	l, err := New(logfile, 0644)

	l.MinLevel(DEBUG)

	c.Assert(err, IsNil)
	c.Assert(l, Not(IsNil))

	l.PrefixDebug = false
	l.PrefixInfo = false
	l.PrefixWarn = false
	l.PrefixError = false
	l.PrefixCrit = false

	c.Assert(fsutil.GetMode(logfile), Equals, os.FileMode(0644))

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

	l.Print(DEBUG, "")

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 17)

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

	c.Assert(dataSlice[15][28:], Equals, "")
}

func (ls *LogSuite) TestWithPrefixes(c *C) {
	logfile := ls.TempDir + "/file2.log"
	err := Set(logfile, 0644)

	MinLevel(DEBUG)

	c.Assert(err, IsNil)
	c.Assert(Global, Not(IsNil))
	c.Assert(Is(DEBUG), Equals, true)

	Global.PrefixDebug = true
	Global.PrefixInfo = true
	Global.PrefixWarn = true
	Global.PrefixError = true
	Global.PrefixCrit = true

	c.Assert(fsutil.GetMode(logfile), Equals, os.FileMode(0644))

	Print(DEBUG, "Test debug %d", DEBUG)
	Print(INFO, "Test info %d", INFO)
	Print(WARN, "Test warn %d", WARN)
	Print(ERROR, "Test error %d", ERROR)
	Print(CRIT, "Test crit %d", CRIT)
	Print(AUX, "Test aux %d", AUX)

	Print(DEBUG, "Test debug")
	Print(INFO, "Test info")
	Print(WARN, "Test warn")
	Print(ERROR, "Test error")
	Print(CRIT, "Test crit")
	Print(AUX, "Test aux")

	Debug("Test debug %d", DEBUG)
	Info("Test info %d", INFO)
	Warn("Test warn %d", WARN)
	Error("Test error %d", ERROR)
	Crit("Test crit %d", CRIT)
	Aux("Test aux %d", AUX)

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 19)

	c.Assert(dataSlice[0][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[1][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[2][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[3][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[4][28:], Equals, "[CRITICAL] Test crit 4")
	c.Assert(dataSlice[5][28:], Equals, "Test aux 99")

	c.Assert(dataSlice[6][28:], Equals, "[DEBUG] Test debug")
	c.Assert(dataSlice[7][28:], Equals, "[INFO] Test info")
	c.Assert(dataSlice[8][28:], Equals, "[WARNING] Test warn")
	c.Assert(dataSlice[9][28:], Equals, "[ERROR] Test error")
	c.Assert(dataSlice[10][28:], Equals, "[CRITICAL] Test crit")
	c.Assert(dataSlice[11][28:], Equals, "Test aux")

	c.Assert(dataSlice[12][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[13][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[14][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[15][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[16][28:], Equals, "[CRITICAL] Test crit 4")
	c.Assert(dataSlice[17][28:], Equals, "Test aux 99")
}

func (ls *LogSuite) TestBufIODaemon(c *C) {
	logfile := ls.TempDir + "/file3.log"
	err := Set(logfile, 0644)

	MinLevel(DEBUG)

	c.Assert(err, IsNil)
	c.Assert(Global, Not(IsNil))

	Global.PrefixDebug = true
	Global.PrefixInfo = true
	Global.PrefixWarn = true
	Global.PrefixError = true
	Global.PrefixCrit = true

	c.Assert(fsutil.GetMode(logfile), Equals, os.FileMode(0644))

	EnableBufIO(250 * time.Millisecond)

	Print(DEBUG, "Test debug %d", DEBUG)
	Print(INFO, "Test info %d", INFO)
	Print(WARN, "Test warn %d", WARN)
	Print(ERROR, "Test error %d", ERROR)
	Print(CRIT, "Test crit %d", CRIT)
	Print(AUX, "Test aux %d", AUX)

	Print(DEBUG, "Test debug")
	Print(INFO, "Test info")
	Print(WARN, "Test warn")
	Print(ERROR, "Test error")
	Print(CRIT, "Test crit")
	Print(AUX, "Test aux")

	Debug("Test debug %d", DEBUG)
	Info("Test info %d", INFO)
	Warn("Test warn %d", WARN)
	Error("Test error %d", ERROR)
	Crit("Test crit %d", CRIT)
	Aux("Test aux %d", AUX)

	c.Assert(fsutil.GetSize(logfile), Equals, int64(0))

	time.Sleep(2 * time.Second)

	c.Assert(fsutil.GetSize(logfile), Not(Equals), int64(0))

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 19)

	c.Assert(dataSlice[0][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[1][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[2][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[3][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[4][28:], Equals, "[CRITICAL] Test crit 4")
	c.Assert(dataSlice[5][28:], Equals, "Test aux 99")

	c.Assert(dataSlice[6][28:], Equals, "[DEBUG] Test debug")
	c.Assert(dataSlice[7][28:], Equals, "[INFO] Test info")
	c.Assert(dataSlice[8][28:], Equals, "[WARNING] Test warn")
	c.Assert(dataSlice[9][28:], Equals, "[ERROR] Test error")
	c.Assert(dataSlice[10][28:], Equals, "[CRITICAL] Test crit")
	c.Assert(dataSlice[11][28:], Equals, "Test aux")

	c.Assert(dataSlice[12][28:], Equals, "[DEBUG] Test debug 0")
	c.Assert(dataSlice[13][28:], Equals, "[INFO] Test info 1")
	c.Assert(dataSlice[14][28:], Equals, "[WARNING] Test warn 2")
	c.Assert(dataSlice[15][28:], Equals, "[ERROR] Test error 3")
	c.Assert(dataSlice[16][28:], Equals, "[CRITICAL] Test crit 4")
	c.Assert(dataSlice[17][28:], Equals, "Test aux 99")
}

func (ls *LogSuite) TestBufIO(c *C) {
	logfile := ls.TempDir + "/file4.log"
	err := Set(logfile, 0644)

	c.Assert(err, IsNil)
	c.Assert(Global, Not(IsNil))

	c.Assert(fsutil.GetMode(logfile), Equals, os.FileMode(0644))

	EnableBufIO(time.Minute)

	Aux("Test aux %d", AUX)

	fileSize := fsutil.GetSize(logfile)

	c.Assert(fileSize, Equals, int64(0))

	Reopen()

	fileSize = fsutil.GetSize(logfile)

	c.Assert(fileSize, Not(Equals), int64(0))

	Aux("Test aux %d", AUX)

	c.Assert(fsutil.GetSize(logfile), Equals, fileSize)

	Flush()

	c.Assert(fsutil.GetSize(logfile), Not(Equals), fileSize)
}

func (ls *LogSuite) TestLoggerIsNil(c *C) {
	var l *Logger

	c.Assert(l.Reopen(), Equals, ErrNilLogger)
	c.Assert(l.MinLevel(1), Equals, ErrNilLogger)
	c.Assert(l.Set("", 0644), Equals, ErrNilLogger)
	c.Assert(l.Print(CRIT, ""), Equals, ErrNilLogger)
	c.Assert(l.Flush(), Equals, ErrNilLogger)
	c.Assert(l.Debug(""), Equals, ErrNilLogger)
	c.Assert(l.Info(""), Equals, ErrNilLogger)
	c.Assert(l.Warn(""), Equals, ErrNilLogger)
	c.Assert(l.Error(""), Equals, ErrNilLogger)
	c.Assert(l.Crit(""), Equals, ErrNilLogger)
	c.Assert(l.Aux(""), Equals, ErrNilLogger)

	c.Assert(func() { l.EnableBufIO(time.Second) }, NotPanics)
}

func (ls *LogSuite) TestStdLogger(c *C) {
	l := &Logger{mu: &sync.Mutex{}}
	l.Set(ls.TempDir+"/file5.log", 0644)

	std := &StdLogger{l}

	exitFunc = func(code int) {}

	c.Assert(std.Output(2, "1"), IsNil)

	std.Fatal("2")
	std.Fatalf("%s", "3")
	std.Fatalln("4")
	std.Print("5")
	std.Printf("%s", "6")
	std.Println("7")

	c.Assert(func() { std.Panic("testPanic") }, PanicMatches, "testPanic")
	c.Assert(func() { std.Panicf("%s", "testPanic") }, PanicMatches, "testPanic")
	c.Assert(func() { std.Panicln("testPanic") }, PanicMatches, "testPanic\n")

	data, err := os.ReadFile(ls.TempDir + "/file5.log")

	if err != nil {
		c.Fatal(err)
	}

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 11)
}

func (ls *LogSuite) TestNilLogger(c *C) {
	var l *NilLogger

	c.Assert(func() { l.Aux("test") }, NotPanics)
	c.Assert(func() { l.Debug("test") }, NotPanics)
	c.Assert(func() { l.Info("test") }, NotPanics)
	c.Assert(func() { l.Warn("test") }, NotPanics)
	c.Assert(func() { l.Error("test") }, NotPanics)
	c.Assert(func() { l.Crit("test") }, NotPanics)
	c.Assert(func() { l.Print(0, "test") }, NotPanics)
}
