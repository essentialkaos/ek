// Package log provides an improved logger
package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

	"github.com/essentialkaos/ek/v12/fmtc"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DEBUG uint8 = 0  // DEBUG debug messages
	INFO  uint8 = 1  // INFO info messages
	WARN  uint8 = 2  // WARN warning messages
	ERROR uint8 = 3  // ERROR error messages
	CRIT  uint8 = 4  // CRIT critical error messages
	AUX   uint8 = 99 // AUX unskipable messages (separators, headers, etc...)
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ILogger is interface for compatible loggers
type ILogger interface {
	Aux(f string, a ...any) error
	Debug(f string, a ...any) error
	Info(f string, a ...any) error
	Warn(f string, a ...any) error
	Error(f string, a ...any) error
	Crit(f string, a ...any) error
	Print(level uint8, f string, a ...any) error
}

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
	minLevel uint8
	perms    os.FileMode
	useBufIO bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Global is global logger struct
var Global = &Logger{
	PrefixWarn:  true,
	PrefixError: true,
	PrefixCrit:  true,

	minLevel: INFO,
	mu:       &sync.Mutex{},
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
	CRIT:  "{r*}",
}

// TimeFormat contains format string for time in logs
var TimeFormat = "2006/01/02 15:04:05.000"

// ////////////////////////////////////////////////////////////////////////////////// //

// Errors
var (
	// ErrNilLogger is returned by Logger struct methods if struct is nil
	ErrNilLogger = errors.New("Logger is nil")

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

		minLevel: INFO,
		mu:       &sync.Mutex{},
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
func MinLevel(level any) error {
	return Global.MinLevel(level)
}

// Set changes global logger output target
func Set(file string, perms os.FileMode) error {
	return Global.Set(file, perms)
}

// EnableBufIO enables buffered I/O
func EnableBufIO(interval time.Duration) {
	Global.EnableBufIO(interval)
}

// Flush writes buffered data to file
func Flush() error {
	return Global.Flush()
}

// Print writes message to global logger output
func Print(level uint8, f string, a ...any) error {
	return Global.Print(level, f, a...)
}

// Debug writes debug message to global logger output
func Debug(f string, a ...any) error {
	return Global.Debug(f, a...)
}

// Info writes info message to global logger output
func Info(f string, a ...any) error {
	return Global.Info(f, a...)
}

// Warn writes warning message to global logger output
func Warn(f string, a ...any) error {
	return Global.Warn(f, a...)
}

// Error writes error message to global logger output
func Error(f string, a ...any) error {
	return Global.Error(f, a...)
}

// Crit writes critical message to global logger output
func Crit(f string, a ...any) error {
	return Global.Crit(f, a...)
}

// Aux writes unskippable message (for separators/headers)
func Aux(f string, a ...any) error {
	return Global.Aux(f, a...)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reopen closes file descriptor and opens it again
// Useful for log rotation
func (l *Logger) Reopen() error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
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
func (l *Logger) MinLevel(level any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
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

	l.minLevel = levelCode

	return nil
}

// EnableBufIO enables buffered I/O support
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

// Set changes logger output target
func (l *Logger) Set(file string, perms os.FileMode) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
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

// Print writes message to logger output
func (l *Logger) Print(level uint8, f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.minLevel > level {
		return nil
	}

	w := l.getWritter(level)
	showPrefix := l.showPrefix(level)

	if f == "" || f[len(f)-1:] != "\n" {
		f += "\n"
	}

	var err error

	switch {
	case l.UseColors && showPrefix:
		_, err = fmt.Fprintf(w, fmtc.Render("{s-}%s{!} "+Colors[level]+"%s %s{!}"), getTime(), PrefixMap[level], fmt.Sprintf(f, a...))
	case l.UseColors && !showPrefix:
		_, err = fmt.Fprintf(w, fmtc.Render("{s-}%s{!} "+Colors[level]+"%s{!}"), getTime(), fmt.Sprintf(f, a...))
	case !l.UseColors && showPrefix:
		_, err = fmt.Fprintf(w, "%s %s %s", getTime(), PrefixMap[level], fmt.Sprintf(f, a...))
	case !l.UseColors && !showPrefix:
		_, err = fmt.Fprintf(w, "%s %s", getTime(), fmt.Sprintf(f, a...))
	}

	return err
}

// Flush writes buffered data to file
func (l *Logger) Flush() error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
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

// Debug writes debug message to logger output
func (l *Logger) Debug(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(DEBUG, f, a...)
}

// Info writes info message to logger output
func (l *Logger) Info(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(INFO, f, a...)
}

// Warn writes warning message to logger output
func (l *Logger) Warn(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(WARN, f, a...)
}

// Error writes error message to logger output
func (l *Logger) Error(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(ERROR, f, a...)
}

// Crit writes critical message to logger output
func (l *Logger) Crit(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(CRIT, f, a...)
}

// Aux writes unfiltered message (for separators/headers) to logger output
func (l *Logger) Aux(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
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

func (l *Logger) showPrefix(level uint8) bool {
	switch {
	case level == DEBUG && l.PrefixDebug,
		level == INFO && l.PrefixInfo,
		level == WARN && l.PrefixWarn,
		level == ERROR && l.PrefixError,
		level == CRIT && l.PrefixCrit:
		return true
	}

	return false
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

func convertMinLevelValue(level any) (uint8, error) {
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
