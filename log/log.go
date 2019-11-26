// Package log provides an improved logger
package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"pkg.re/essentialkaos/ek.v11/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DEBUG uint8 = 0  // DEBUG debug messages
	INFO        = 1  // INFO info messages
	WARN        = 2  // WARN warning messages
	ERROR       = 3  // ERROR error messages
	CRIT        = 4  // CRIT critical error messages
	AUX         = 99 // AUX unskipable messages (separators, headers, etc...)
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Logger is a basic logger struct
type Logger struct {
	PrefixDebug bool // Prefix for debug messages
	PrefixInfo  bool // Prefix for info messages
	PrefixWarn  bool // Prefix for warning messages
	PrefixError bool // Prefix for error messages
	PrefixCrit  bool // Prefix for critical error messages

	UseColors bool // Enable ANSI escape codes for colors in output

	file     string
	fd       *os.File
	w        *bufio.Writer
	mu       *sync.Mutex
	level    uint8
	perms    os.FileMode
	useBufIO bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Global is global logger struct
var Global = &Logger{
	PrefixWarn:  true,
	PrefixError: true,
	PrefixCrit:  true,

	level: INFO,
	mu:    &sync.Mutex{},
}

// ////////////////////////////////////////////////////////////////////////////////// //

// PrefixMap is map with messages prefixes
var PrefixMap = map[uint8]string{
	DEBUG: "[DEBUG]",
	INFO:  "[INFO]",
	WARN:  "[WARNING]",
	ERROR: "[ERROR]",
	CRIT:  "[CRITICAL]",
}

// Colors colors is map with fmtc color tags for every level
var Colors = map[uint8]string{
	DEBUG: "{s-}",
	INFO:  "",
	WARN:  "{y}",
	ERROR: "{r}",
	CRIT:  "{m}",
}

// TimeFormat contains format string for time in logs
var TimeFormat = "2006/01/02 15:04:05.000"

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors
var (
	// ErrLoggerIsNil is returned by Logger struct methods if struct is nil
	ErrLoggerIsNil = errors.New("Logger is nil or not created properly")

	// ErrUnexpectedLevel is returned by the MinLevel method if given level is unknown
	ErrUnexpectedLevel = errors.New("Unexpected level type")

	// ErrOutputNotSet is returned by the Reopen method if output file is not set
	ErrOutputNotSet = errors.New("Output file is not set")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var logLevelsNames = map[string]uint8{
	"debug":    0,
	"info":     1,
	"warn":     2,
	"warning":  2,
	"error":    3,
	"crit":     4,
	"critical": 4,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new logger struct
func New(file string, perms os.FileMode) (*Logger, error) {
	logger := &Logger{
		PrefixWarn:  true,
		PrefixCrit:  true,
		PrefixError: true,

		level: INFO,
		mu:    &sync.Mutex{},
	}

	err := logger.Set(file, perms)

	if err != nil {
		return nil, err
	}

	return logger, nil
}

// Reopen close file descriptor for global logger and open it again
// Useful for log rotation
func Reopen() error {
	return Global.Reopen()
}

// MinLevel defines minimal logging level
func MinLevel(level interface{}) error {
	return Global.MinLevel(level)
}

// Set change global logger output target
func Set(file string, perms os.FileMode) error {
	return Global.Set(file, perms)
}

// EnableBufIO enable buffered I/O
func EnableBufIO(interval time.Duration) {
	Global.EnableBufIO(interval)
}

// Flush write buffered data to file
func Flush() error {
	return Global.Flush()
}

// Print write message to global logger output
func Print(level uint8, f string, a ...interface{}) (int, error) {
	return Global.Print(level, f, a...)
}

// Debug write debug message to global logger output
func Debug(f string, a ...interface{}) (int, error) {
	return Global.Debug(f, a...)
}

// Info write info message to global logger output
func Info(f string, a ...interface{}) (int, error) {
	return Global.Info(f, a...)
}

// Warn write warning message to global logger output
func Warn(f string, a ...interface{}) (int, error) {
	return Global.Warn(f, a...)
}

// Error write error message to global logger output
func Error(f string, a ...interface{}) (int, error) {
	return Global.Error(f, a...)
}

// Crit write critical message to global logger output
func Crit(f string, a ...interface{}) (int, error) {
	return Global.Crit(f, a...)
}

// Aux write unskippable message (for separators/headers)
func Aux(f string, a ...interface{}) (int, error) {
	return Global.Aux(f, a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reopen close file descriptor and open again
// Useful for log rotation
func (l *Logger) Reopen() error {
	if l == nil || l.mu == nil {
		return ErrLoggerIsNil
	}

	l.mu.Lock()

	if l.fd == nil {
		l.mu.Unlock()
		return ErrOutputNotSet
	}

	if l.w != nil {
		l.w.Flush()
	}

	l.fd.Close()
	l.mu.Unlock()

	return l.Set(l.file, l.perms)
}

// MinLevel defines minimal logging level
func (l *Logger) MinLevel(level interface{}) error {
	if l == nil || l.mu == nil {
		return ErrLoggerIsNil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	levelCode, err := convertMinLevelValue(level)

	if err != nil {
		return err
	}

	if levelCode > CRIT {
		levelCode = CRIT
	}

	l.level = levelCode

	return nil
}

// EnableBufIO enable buffered I/O support
func (l *Logger) EnableBufIO(interval time.Duration) {
	if l == nil || l.mu == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.useBufIO = true

	if l.fd != nil {
		l.w = bufio.NewWriter(l.fd)
	}

	go l.flushDaemon(interval)
}

// Set change logger output target
func (l *Logger) Set(file string, perms os.FileMode) error {
	if l == nil || l.mu == nil {
		return ErrLoggerIsNil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perms)

	if err != nil {
		return err
	}

	// Flush data if writer exist
	if l.w != nil {
		l.w.Flush()
		l.w = nil
	}

	if l.fd != nil {
		l.fd.Close()
		l.fd = nil
	}

	l.fd, l.file, l.perms = fd, file, perms

	if l.useBufIO {
		l.w = bufio.NewWriter(l.fd)
	}

	return nil
}

// Print write message to logger output
func (l *Logger) Print(level uint8, f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level > level {
		return 0, nil
	}

	w := l.getWritter(level)

	var showPrefixes bool

	switch {
	case level == DEBUG && l.PrefixDebug,
		level == INFO && l.PrefixInfo,
		level == WARN && l.PrefixWarn,
		level == ERROR && l.PrefixError,
		level == CRIT && l.PrefixCrit:
		showPrefixes = true
	}

	if f == "" || f[len(f)-1:] != "\n" {
		f += "\n"
	}

	var err error
	var n int

	if l.UseColors {
		c := Colors[level]
		if showPrefixes {
			n, err = fmtc.Fprintf(w, "{s-}%s{!} "+c+"%s %s{!}", getTime(), PrefixMap[level], fmt.Sprintf(f, a...))
		} else {
			n, err = fmtc.Fprintf(w, "{s-}%s{!} "+c+"%s{!}", getTime(), fmt.Sprintf(f, a...))
		}
	} else {
		if showPrefixes {
			n, err = fmt.Fprintf(w, "%s %s %s", getTime(), PrefixMap[level], fmt.Sprintf(f, a...))
		} else {
			n, err = fmt.Fprintf(w, "%s %s", getTime(), fmt.Sprintf(f, a...))
		}
	}

	return n, err
}

// Flush write buffered data to file
func (l *Logger) Flush() error {
	if l == nil || l.mu == nil {
		return ErrLoggerIsNil
	}

	l.mu.Lock()

	if l.w == nil {
		l.mu.Unlock()
		return nil
	}

	err := l.w.Flush()

	l.mu.Unlock()

	return err
}

// Debug write debug message to logger output
func (l *Logger) Debug(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(DEBUG, f, a...)
}

// Info write info message to logger output
func (l *Logger) Info(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(INFO, f, a...)
}

// Warn write warning message to logger output
func (l *Logger) Warn(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(WARN, f, a...)
}

// Error write error message to logger output
func (l *Logger) Error(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(ERROR, f, a...)
}

// Crit write critical message to logger output
func (l *Logger) Crit(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(CRIT, f, a...)
}

// Aux write unfiltered message (for separators/headers) to logger output
func (l *Logger) Aux(f string, a ...interface{}) (int, error) {
	if l == nil || l.mu == nil {
		return -1, ErrLoggerIsNil
	}

	return l.Print(AUX, f, a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (l *Logger) getWritter(level uint8) io.Writer {
	var w io.Writer

	if l.fd == nil {
		switch level {
		case ERROR, CRIT:
			w = os.Stderr
		default:
			w = os.Stdout
		}
	} else {
		if l.w != nil {
			w = l.w
		} else {
			w = l.fd
		}
	}

	return w
}

func (l *Logger) flushDaemon(interval time.Duration) {
	for range time.NewTicker(interval).C {
		l.Flush()
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getTime() string {
	return "[ " + time.Now().Format(TimeFormat) + " ]"
}

func convertMinLevelValue(level interface{}) (uint8, error) {
	switch u := level.(type) {

	case int:
		return uint8(u), nil

	case int8:
		return uint8(u), nil

	case int16:
		return uint8(u), nil

	case int32:
		return uint8(u), nil

	case int64:
		return uint8(u), nil

	case uint:
		return uint8(u), nil

	case uint8:
		return uint8(u), nil

	case uint16:
		return uint8(u), nil

	case uint32:
		return uint8(u), nil

	case uint64:
		return uint8(u), nil

	case float32:
		return uint8(u), nil

	case float64:
		return uint8(u), nil

	case string:
		code, ok := logLevelsNames[strings.ToLower(level.(string))]

		if !ok {
			return 255, errors.New("Unknown level " + level.(string))
		}

		return code, nil
	}

	return 255, ErrUnexpectedLevel
}
