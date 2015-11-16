// Package with improved logger
package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// DEBUG debug messages
// INFO info messages
// WARN warning messages
// ERROR error messages
// CRIT critical error messages
// AUX unskipable messages (separators, headers, etc...)
const (
	DEBUG = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
	CRIT  = 4
	AUX   = 99
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Logger is a basic logger struct
type Logger struct {
	PrefixDebug bool // Prefix for debug messages
	PrefixInfo  bool // Prefix for info messages
	PrefixWarn  bool // Prefix for warning messages
	PrefixError bool // Prefix for error messages
	PrefixCrit  bool // Prefix for critical error messages

	file  string
	perms os.FileMode
	fd    *os.File
	w     *bufio.Writer
	level int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var globalLogger = &Logger{
	PrefixWarn:  true,
	PrefixError: true,
	PrefixCrit:  true,

	level: INFO,
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Map with messages prefixes
var PrefixMap = map[int]string{
	DEBUG: "[DEBUG]",
	INFO:  "[INFO]",
	WARN:  "[WARNING]",
	ERROR: "[ERROR]",
	CRIT:  "[CRITICAL]",
}

// TimeFormat contains format string for time in logs
var TimeFormat = "2006/01/02 15:04:05.000"

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates new logger struct
func New(file string, perms os.FileMode) (*Logger, error) {
	logger := &Logger{
		PrefixWarn:  true,
		PrefixCrit:  true,
		PrefixError: true,

		level: INFO,
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
	return globalLogger.Reopen()
}

// MinLevel defines minimal logging level for global logger (1 by default)
func MinLevel(level int) {
	globalLogger.MinLevel(level)
}

// Set change global logger output target
func Set(file string, perms os.FileMode) error {
	return globalLogger.Set(file, perms)
}

// Flush write buffered data to file
func Flush() error {
	return globalLogger.Flush()
}

// Print write message to global logger output
func Print(level int, f string, a ...interface{}) (int, error) {
	return globalLogger.Print(level, f, a...)
}

// Debug write debug message to global logger output
func Debug(f string, a ...interface{}) (int, error) {
	return globalLogger.Debug(f, a...)
}

// Info write info message to global logger output
func Info(f string, a ...interface{}) (int, error) {
	return globalLogger.Info(f, a...)
}

// Warn write warning message to global logger output
func Warn(f string, a ...interface{}) (int, error) {
	return globalLogger.Warn(f, a...)
}

// Error write error message to global logger output
func Error(f string, a ...interface{}) (int, error) {
	return globalLogger.Error(f, a...)
}

// Crit write critical message to global logger output
func Crit(f string, a ...interface{}) (int, error) {
	return globalLogger.Crit(f, a...)
}

// Aux write unskipable message (for separators/headers)
func Aux(f string, a ...interface{}) (int, error) {
	return globalLogger.Aux(f, a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reopen close file descriptor and open again
// Useful for log rotation
func (l *Logger) Reopen() error {
	if l.fd == nil {
		return errors.New("Output file is not set")
	}

	l.w.Flush()
	l.fd.Close()

	return l.Set(l.file, l.perms)
}

// MinLevel defines minimal logging level for logger (1 by default)
func (l *Logger) MinLevel(level int) {
	switch {
	case level < DEBUG:
		level = DEBUG
	case level > CRIT:
		level = CRIT
	}

	l.level = level
}

// Set change logger output target
func (l *Logger) Set(file string, perms os.FileMode) error {
	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perms)

	if err != nil {
		return err
	}

	// Flush data if writter exist
	if l.w != nil {
		l.w.Flush()
	}

	if l.fd != nil {
		l.fd.Close()
	}

	l.fd, l.file, l.perms = fd, file, perms
	l.w = bufio.NewWriter(l.fd)

	return nil
}

// Print write message to logger output
func (l *Logger) Print(level int, f string, a ...interface{}) (int, error) {
	if l.level > level {
		return 0, nil
	}

	var w io.Writer

	if l.fd == nil {
		switch level {
		case ERROR, CRIT:
			w = os.Stderr
		default:
			w = os.Stdout
		}
	} else {
		w = l.w
	}

	var showPrefixes bool

	switch {
	case level == DEBUG && l.PrefixDebug,
		level == INFO && l.PrefixInfo,
		level == WARN && l.PrefixWarn,
		level == ERROR && l.PrefixError,
		level == CRIT && l.PrefixCrit:
		showPrefixes = true
	}

	if f[len(f)-1:] != "\n" {
		f += "\n"
	}

	if showPrefixes {
		return fmt.Fprintf(w, "%s %s %s", getTime(), PrefixMap[level], fmt.Sprintf(f, a...))
	}

	return fmt.Fprintf(w, "%s %s", getTime(), fmt.Sprintf(f, a...))
}

// Flush write buffered data to file
func (l *Logger) Flush() error {
	if l.w == nil {
		return nil
	}

	return l.w.Flush()
}

// Debug write debug message to logger output
func (l *Logger) Debug(f string, a ...interface{}) (int, error) {
	return l.Print(DEBUG, f, a...)
}

// Info write info message to logger output
func (l *Logger) Info(f string, a ...interface{}) (int, error) {
	return l.Print(INFO, f, a...)
}

// Warn write warning message to logger output
func (l *Logger) Warn(f string, a ...interface{}) (int, error) {
	return l.Print(WARN, f, a...)
}

// Error write error message to logger output
func (l *Logger) Error(f string, a ...interface{}) (int, error) {
	return l.Print(ERROR, f, a...)
}

// Crit write critical message to logger output
func (l *Logger) Crit(f string, a ...interface{}) (int, error) {
	return l.Print(CRIT, f, a...)
}

// Aux write unskipable message (for separators/headers)
func (l *Logger) Aux(f string, a ...interface{}) (int, error) {
	return l.Print(AUX, f, a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func getTime() string {
	return "[ " + time.Now().Format(TimeFormat) + " ]"
}
