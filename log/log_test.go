package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/essentialkaos/ek.v13/fmtc"
	"github.com/essentialkaos/ek.v13/fsutil"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type LogSuite struct {
	TempDir string
}

type JSONRecord struct {
	Level  string `json:"level"`
	TS     string `json:"ts"`
	Caller string `json:"caller"`
	Msg    string `json:"msg"`
	ID     int    `json:"id"`
	User   string `json:"user"`
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

	err = l.Divider()

	c.Assert(err, NotNil)
	c.Assert(err, DeepEquals, ErrNilLogger)

	c.Assert(l.Is(DEBUG), Equals, false)

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
	c.Assert(l.MinLevel("fatal"), IsNil)
	c.Assert(l.MinLevel(int8(1)), IsNil)
	c.Assert(l.MinLevel(int16(1)), IsNil)
	c.Assert(l.MinLevel(int32(1)), IsNil)
	c.Assert(l.MinLevel(int64(1)), IsNil)
	c.Assert(l.MinLevel(uint(1)), IsNil)
	c.Assert(l.MinLevel(uint8(1)), IsNil)
	c.Assert(l.MinLevel(uint16(1)), IsNil)
	c.Assert(l.MinLevel(uint32(1)), IsNil)
	c.Assert(l.MinLevel(uint64(1)), IsNil)

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

	err = l.Print(INFO, "Info message")

	c.Assert(err, IsNil)

	err = l.Print(ERROR, "Error message")

	c.Assert(err, IsNil)

	fmtc.DisableColors = true

	l.UseColors = true
	l.PrefixError = true

	err = l.Print(INFO, "Info message")

	c.Assert(err, IsNil)

	err = l.Print(ERROR, "Error message")

	c.Assert(err, IsNil)

	fmtc.DisableColors = false

	l.UseJSON = true

	err = l.Print(ERROR, "Error message")
}

func (ls *LogSuite) TestWithoutPrefixes(c *C) {
	logfile := ls.TempDir + "/basic.log"
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

	c.Assert(l.Print(DEBUG, "Test debug %d", DEBUG), IsNil)
	c.Assert(l.Print(INFO, "Test info %d", INFO), IsNil)
	c.Assert(l.Print(WARN, "Test warn %d", WARN), IsNil)
	c.Assert(l.Print(ERROR, "Test error %d", ERROR), IsNil)
	c.Assert(l.Print(CRIT, "Test crit %d", CRIT), IsNil)

	c.Assert(l.Print(DEBUG, "Test debug"), IsNil)
	c.Assert(l.Print(INFO, "Test info"), IsNil)
	c.Assert(l.Print(WARN, "Test warn"), IsNil)
	c.Assert(l.Print(ERROR, "Test error"), IsNil)
	c.Assert(l.Print(CRIT, "Test crit"), IsNil)

	c.Assert(l.Debug("Test debug %d\n", DEBUG), IsNil)
	c.Assert(l.Info("Test info %d\n", INFO), IsNil)
	c.Assert(l.Warn("Test warn %d\n", WARN), IsNil)
	c.Assert(l.Error("Test error %d\n", ERROR), IsNil)
	c.Assert(l.Crit("Test crit %d\n", CRIT), IsNil)

	l.Divider()

	l.Print(DEBUG, "")

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 18)

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

	c.Assert(dataSlice[15][28:], Equals, "--------------------------------------------------------------------------------")

	c.Assert(dataSlice[16][28:], Equals, "")
}

func (ls *LogSuite) TestWithPrefixes(c *C) {
	logfile := ls.TempDir + "/prefixes.log"
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

	c.Assert(Print(DEBUG, "Test debug %d", DEBUG), IsNil)
	c.Assert(Print(INFO, "Test info %d", INFO), IsNil)
	c.Assert(Print(WARN, "Test warn %d", WARN), IsNil)
	c.Assert(Print(ERROR, "Test error %d", ERROR), IsNil)
	c.Assert(Print(CRIT, "Test crit %d", CRIT), IsNil)
	c.Assert(Print(AUX, "Test aux %d", AUX), IsNil)

	c.Assert(Print(DEBUG, "Test debug"), IsNil)
	c.Assert(Print(INFO, "Test info"), IsNil)
	c.Assert(Print(WARN, "Test warn"), IsNil)
	c.Assert(Print(ERROR, "Test error"), IsNil)
	c.Assert(Print(CRIT, "Test crit"), IsNil)
	c.Assert(Print(AUX, "Test aux"), IsNil)

	c.Assert(Debug("Test debug %d", DEBUG), IsNil)
	c.Assert(Info("Test info %d", INFO), IsNil)
	c.Assert(Warn("Test warn %d", WARN), IsNil)
	c.Assert(Error("Test error %d", ERROR), IsNil)
	c.Assert(Crit("Test crit %d", CRIT), IsNil)
	c.Assert(Aux("Test aux %d", AUX), IsNil)

	c.Assert(Divider(), IsNil)

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 20)

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
	c.Assert(dataSlice[18][28:], Equals, "--------------------------------------------------------------------------------")
}

func (ls *LogSuite) TestJSON(c *C) {
	logfile := ls.TempDir + "/json.log"
	l, err := New(logfile, 0644)

	l.MinLevel(DEBUG)

	c.Assert(err, IsNil)
	c.Assert(l, Not(IsNil))

	l.UseJSON = true
	l.WithCaller = true
	l.TimeLayout = time.RFC822

	c.Assert(fsutil.GetMode(logfile), Equals, os.FileMode(0644))

	c.Assert(l.Print(DEBUG, "Test debug %d (%s)", DEBUG, F{"id", 101}, "test1", F{"user", "john"}), IsNil)
	c.Assert(l.Print(INFO, "Test info %d", INFO, F{"id", 102}, F{"user", "bob"}), IsNil)
	c.Assert(l.Print(WARN, "Test warn %d", WARN, F{"id", 103}, F{"user", "frida"}), IsNil)
	c.Assert(l.Print(ERROR, "Test error %d", ERROR, F{"id", 104}, F{"user", "alisa"}), IsNil)
	c.Assert(l.Print(CRIT, "Test crit %d", CRIT, F{"id", 105}, F{"user", "simon"}), IsNil)

	c.Assert(l.Info("Test message"), IsNil)

	l.TimeLayout = ""

	c.Assert(l.Info("Test message"), IsNil)

	l.Print(DEBUG, "")
	l.Aux("Test")

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 8)

	records := parseJSONRecords(dataSlice)

	c.Assert(len(records), Equals, 7)

	c.Assert(records[0].Level, Equals, "debug")
	c.Assert(records[0].TS, Not(Equals), "")
	c.Assert(records[0].Caller, Not(Equals), "")
	c.Assert(records[0].Msg, Equals, "Test debug 0 (test1)")
	c.Assert(records[0].ID, Equals, 101)
	c.Assert(records[0].User, Equals, "john")

	c.Assert(records[1].Level, Equals, "info")
	c.Assert(records[1].TS, Not(Equals), "")
	c.Assert(records[1].Caller, Not(Equals), "")
	c.Assert(records[1].Msg, Equals, "Test info 1")
	c.Assert(records[1].ID, Equals, 102)
	c.Assert(records[1].User, Equals, "bob")

	c.Assert(records[2].Level, Equals, "warn")
	c.Assert(records[2].TS, Not(Equals), "")
	c.Assert(records[2].Caller, Not(Equals), "")
	c.Assert(records[2].Msg, Equals, "Test warn 2")
	c.Assert(records[2].ID, Equals, 103)
	c.Assert(records[2].User, Equals, "frida")

	c.Assert(records[3].Level, Equals, "error")
	c.Assert(records[3].TS, Not(Equals), "")
	c.Assert(records[3].Caller, Not(Equals), "")
	c.Assert(records[3].Msg, Equals, "Test error 3")
	c.Assert(records[3].ID, Equals, 104)
	c.Assert(records[3].User, Equals, "alisa")

	c.Assert(records[4].Level, Equals, "fatal")
	c.Assert(records[4].TS, Not(Equals), "")
	c.Assert(records[4].Caller, Not(Equals), "")
	c.Assert(records[4].Msg, Equals, "Test crit 4")
	c.Assert(records[4].ID, Equals, 105)
	c.Assert(records[4].User, Equals, "simon")

	c.Assert(records[5].Level, Equals, "info")
	c.Assert(records[5].TS, Not(Equals), "")
	c.Assert(records[5].Caller, Not(Equals), "")
	c.Assert(records[5].Msg, Equals, "Test message")
	c.Assert(records[5].ID, Equals, 0)
	c.Assert(records[5].User, Equals, "")
}

func (ls *LogSuite) TestWithCaller(c *C) {
	logfile := ls.TempDir + "/caller.log"
	l, err := New(logfile, 0644)

	c.Assert(err, IsNil)

	l.WithCaller = true

	c.Assert(l.Print(INFO, "Test info 1"), IsNil)

	l.UseColors = true
	fmtc.DisableColors = true

	c.Assert(l.Print(INFO, "Test info 2"), IsNil)

	fmtc.DisableColors = false

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 3)

	c.Assert(dataSlice[0][28:], Equals, "(log/log_test.go:470) Test info 1")
	c.Assert(dataSlice[1][28:], Equals, "(log/log_test.go:475) Test info 2")
}

func (ls *LogSuite) TestWithFields(c *C) {
	logfile := ls.TempDir + "/fields.log"
	l, err := New(logfile, 0644)

	c.Assert(err, IsNil)

	c.Assert(l.Info("Test info %d %s", 1, F{"name", "john"}, F{"id", 1}, F{"", 99}, "test"), IsNil)

	l.UseColors = true
	fmtc.DisableColors = true

	c.Assert(l.Info("Test info %d %s", 2, F{"name", "john"}, F{"id", 1}, F{"", 99}, "test"), IsNil)

	fmtc.DisableColors = false

	data, err := os.ReadFile(logfile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Not(Equals), 0)

	dataSlice := strings.Split(string(data), "\n")

	c.Assert(len(dataSlice), Equals, 3)

	c.Assert(dataSlice[0][28:], Equals, "Test info 1 test {id: 1 | name: john}")
	c.Assert(dataSlice[1][28:], Equals, "Test info 2 test {id: 1 | name: john}")
}

func (ls *LogSuite) TestBufIODaemon(c *C) {
	logfile := ls.TempDir + "/bufio-daemon.log"
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
	logfile := ls.TempDir + "/bufio.log"
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
	logfile := ls.TempDir + "/stdlogger.log"
	l := &Logger{mu: &sync.Mutex{}}

	l.Set(logfile, 0644)

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

	data, err := os.ReadFile(logfile)

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

func (ls *LogSuite) TestTimeLayout(c *C) {
	l := &Logger{mu: &sync.Mutex{}}
	t := time.Unix(1704067200, 0)

	jf := l.formatDateTime(t, true)
	tf := l.formatDateTime(t, false)

	l.TimeLayout = time.Kitchen
	cf := l.formatDateTime(t, false)

	c.Assert(jf, Not(Equals), "")
	c.Assert(tf, Not(Equals), "")
	c.Assert(cf, Not(Equals), "")

	c.Assert(jf != tf, Equals, true)
	c.Assert(tf != cf, Equals, true)
	c.Assert(jf != cf, Equals, true)
}

func (ls *LogSuite) TestFieldEncoding(c *C) {
	f := F{"test", 123}
	c.Assert(f.String(), Equals, "test:123")

	l := &Logger{mu: &sync.Mutex{}}
	l.writeJSONField(F{"test", "abcd"})
	c.Assert(l.buf.String(), Equals, "\"test\":\"abcd\"")
	l.buf.Reset()

	l.writeJSONField(F{"test", false})
	c.Assert(l.buf.String(), Equals, "\"test\":false")
	l.buf.Reset()

	l.writeJSONField(F{"test", 123})
	c.Assert(l.buf.String(), Equals, "\"test\":123")
	l.buf.Reset()

	l.writeJSONField(F{"test", int8(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", int8(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", int16(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", int32(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", int64(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint(123)})
	c.Assert(l.buf.String(), Equals, "\"test\":123")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint8(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint8(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint16(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint32(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", uint64(33)})
	c.Assert(l.buf.String(), Equals, "\"test\":33")
	l.buf.Reset()

	l.writeJSONField(F{"test", float32(1.23)})
	c.Assert(l.buf.String(), Equals, "\"test\":1.23")
	l.buf.Reset()

	l.writeJSONField(F{"test", float64(1.23)})
	c.Assert(l.buf.String(), Equals, "\"test\":1.23")
	l.buf.Reset()

	l.writeJSONField(F{"test", time.Minute - (150 * time.Millisecond)})
	c.Assert(l.buf.String(), Equals, "\"test\":59.85")
	l.buf.Reset()

	l.writeJSONField(F{"test", time.Now()})
	c.Assert(l.buf.String(), Not(Equals), "")
	l.buf.Reset()

	l.writeJSONField(F{"test", []string{"A"}})
	c.Assert(l.buf.String(), Equals, "\"test\":\"[A]\"")
	l.buf.Reset()
}

func (ls *LogSuite) TestFields(c *C) {
	var fp *Fields

	c.Assert(fp.Add(), IsNil)
	c.Assert(fp.Reset(), IsNil)

	var fv Fields

	c.Assert(fv.Add(F{}, F{"testF1", 1}, F{"testF2", true}), NotNil)
	c.Assert(fv.data, HasLen, 2)
	c.Assert(fv.Reset(), NotNil)
	c.Assert(fv.data, HasLen, 0)
}

func (ls *LogSuite) TestPayloadSplitter(c *C) {
	var fv Fields
	fv.Add(F{}, F{"testF1", 1})

	fp := &Fields{}
	fp.Add(F{}, F{"testF2", true})

	p := []any{1, F{"name", "john"}, F{"id", 1}, F{"", 99}, "test", fp, fv}
	p1, p2 := splitPayload(p)

	c.Assert(p1, HasLen, 2)
	c.Assert(p1, DeepEquals, []any{1, "test"})
	c.Assert(p2, HasLen, 4)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *LogSuite) BenchmarkJSONWrite(c *C) {
	l, err := New(os.DevNull, 0644)

	c.Assert(err, IsNil)

	l.UseJSON = true

	f1, f2 := F{"test1", 1}, F{"test2", false}

	for i := 0; i < c.N; i++ {
		l.Info("Test %s %s", "test", f1, "abcd", f2)
	}
}

func (s *LogSuite) BenchmarkTextWrite(c *C) {
	l, err := New(os.DevNull, 0644)

	c.Assert(err, IsNil)

	f1, f2 := F{"test1", 1}, F{"test2", false}

	for i := 0; i < c.N; i++ {
		l.Info("Test %s %s", "test", f1, "abcd", f2)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func parseJSONRecords(data []string) []*JSONRecord {
	var result []*JSONRecord

	for _, l := range data {
		if l == "" {
			continue
		}

		r := &JSONRecord{}
		json.Unmarshal([]byte(l), r)
		result = append(result, r)
	}

	return result
}
