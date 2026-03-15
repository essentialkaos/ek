// Package log provides an improved logger
package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	DEBUG uint8 = 0  // DEBUG represents the debug log level
	INFO  uint8 = 1  // INFO represents the informational log level
	WARN  uint8 = 2  // WARN represents the warning log level
	ERROR uint8 = 3  // ERROR represents the error log level
	CRIT  uint8 = 4  // CRIT represents the critical (fatal) log level
	AUX   uint8 = 99 // AUX represents an unskippable log level used for separators and headers
)

const (
	// DATE_LAYOUT_TEXT is the datetime format used when rendering plain-text log entries
	DATE_LAYOUT_TEXT = "2006/01/02 15:04:05.000"

	// DATE_LAYOUT_JSON is the datetime format used when rendering JSON log entries
	DATE_LAYOUT_JSON = time.RFC3339
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ILogger defines the interface for loggers compatible with this package
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

// Logger is the core logging struct supporting both plain-text and JSON output
type Logger struct {
	PrefixDebug bool // Show prefix for debug messages
	PrefixInfo  bool // Show prefix for info messages
	PrefixWarn  bool // Show prefix for warning messages
	PrefixError bool // Show prefix for error messages
	PrefixCrit  bool // Show prefix for critical/fatal messages

	TimeLayout         string // Date and time layout used for rendering dates
	UseColors          bool   // Enable ANSI escape codes for colors in output
	UseJSON            bool   // Encode messages to JSON
	WithCaller         bool   // Show caller info
	WithFullCallerPath bool   // Show full path of caller
	DiscardFields      bool   // Don't write fields to log

	file          string
	buf           bytes.Buffer
	fd            *os.File
	w             *bufio.Writer
	mu            *sync.Mutex
	minLevel      uint8
	perms         os.FileMode
	useBufIO      bool
	bufIOStopChan chan struct{}
}

// F is an alias for [Field]
type F = Field

// Field holds a key-value pair attached to a structured log entry
type Field struct {
	Key   string
	Value any
}

// Fields is an ordered collection of Field values for structured logging
type Fields struct {
	data []Field
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Global is the package-level default logger instance
var Global = &Logger{
	PrefixWarn:  true,
	PrefixError: true,
	PrefixCrit:  true,

	minLevel: INFO,
	mu:       &sync.Mutex{},
}

// ////////////////////////////////////////////////////////////////////////////////// //

// PrefixMap maps log levels to their human-readable string prefixes
var PrefixMap = map[uint8]string{
	DEBUG: "[DEBUG]",
	INFO:  "[INFO]",
	WARN:  "[WARNING]",
	ERROR: "[ERROR]",
	CRIT:  "[CRITICAL]",
}

// Colors maps log levels to fmtc color tags used in colorized output
var Colors = map[uint8]string{
	DEBUG: "{s-}",
	INFO:  "",
	WARN:  "{y}",
	ERROR: "{#208}",
	CRIT:  "{#196}{*}",
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilLogger is returned when a method is called on a nil logger
	ErrNilLogger = errors.New("logger is nil")

	// ErrUnexpectedLevel is returned when an unrecognized level value is passed
	// to [MinLevel]
	ErrUnexpectedLevel = errors.New("unexpected level type")

	// ErrOutputNotSet is returned by [Reopen] when no output file has been configured
	ErrOutputNotSet = errors.New("output file is not set")
)

// ////////////////////////////////////////////////////////////////////////////////// //

var logLevelsNames = map[string]uint8{
	"debug":    0,
	"info":     1,
	"":         1, // default
	"warn":     2,
	"warning":  2,
	"error":    3,
	"crit":     4,
	"critical": 4,
	"fatal":    4,
}

var logLevels = []string{
	"",
	"debug",
	"info",
	"warn", "warning",
	"error",
	"crit", "critical", "fatal",
}

// ////////////////////////////////////////////////////////////////////////////////// //

// New creates and returns a new Logger writing to the given file with the
// specified permissions
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

// Reopen closes and reopens the global logger's file descriptor, useful for log
// rotation
func Reopen() error {
	return Global.Reopen()
}

// MinLevel sets the minimum log level for the global logger
func MinLevel(level any) error {
	return Global.MinLevel(level)
}

// Set changes the output target of the global logger to the given file
func Set(file string, perms os.FileMode) error {
	return Global.Set(file, perms)
}

// EnableBufIO enables buffered I/O on the global logger with the given
// flush interval
func EnableBufIO(interval time.Duration) {
	Global.EnableBufIO(interval)
}

// Flush writes any buffered data to the global logger's output file
func Flush() error {
	return Global.Flush()
}

// Print writes a message at the given level to the global logger
func Print(level uint8, f string, a ...any) error {
	return Global.Print(level, f, a...)
}

// Debug writes a debug-level message to the global logger
func Debug(f string, a ...any) error {
	return Global.Debug(f, a...)
}

// Info writes an info-level message to the global logger
func Info(f string, a ...any) error {
	return Global.Info(f, a...)
}

// Warn writes a warning-level message to the global logger
func Warn(f string, a ...any) error {
	return Global.Warn(f, a...)
}

// Error writes an error-level message to the global logger
func Error(f string, a ...any) error {
	return Global.Error(f, a...)
}

// Crit writes a critical-level message to the global logger
func Crit(f string, a ...any) error {
	return Global.Crit(f, a...)
}

// Aux writes an unskippable message (e.g. separator or header) to the
// global logger
func Aux(f string, a ...any) error {
	return Global.Aux(f, a...)
}

// Divider writes an 80-character horizontal rule to the global logger
func Divider() error {
	return Global.Divider()
}

// Is reports whether the global logger's minimum level is at or below
// the given level
func Is(level uint8) bool {
	return Global.Is(level)
}

// Levels returns a slice of all supported log level name strings
func Levels() []string {
	return logLevels
}

// PanicHandler recovers from a panic and logs it at the critical level.
// It must be called via defer, not directly.
func PanicHandler(msg string) {
	r := recover()

	if r == nil {
		return
	}

	if Global == nil {
		fmt.Fprintf(os.Stderr, "%s: %v (%s)\n", msg, r, extractPanicPath(debug.Stack()))
		return
	}

	if Global.UseJSON {
		Global.Crit("%s: %v", msg, r, F{"panic-stack", string(debug.Stack())})
	} else {
		Global.Crit("%s: %v (%s)", msg, r, extractPanicPath(debug.Stack()))
	}

}

// ////////////////////////////////////////////////////////////////////////////////// //

// Reopen closes and reopens the logger's file descriptor, useful for log rotation
func (l *Logger) Reopen() error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	l.mu.Lock()

	if l.fd == nil {
		l.mu.Unlock()
		return ErrOutputNotSet
	}

	file, perms := l.file, l.perms

	l.mu.Unlock()

	return l.Set(file, perms)
}

// MinLevel sets the minimum level below which messages are not written
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

// EnableBufIO enables buffered I/O with the given periodic flush interval
func (l *Logger) EnableBufIO(flushInterval time.Duration) {
	if l == nil || l.mu == nil {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.useBufIO = true

	if l.fd != nil {
		l.w = bufio.NewWriter(l.fd)
	}

	if l.bufIOStopChan != nil {
		close(l.bufIOStopChan)
	}

	l.bufIOStopChan = make(chan struct{})

	go l.flushDaemon(flushInterval, l.bufIOStopChan)
}

// Set changes the logger's output to the given file path and permission mode
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

// Print writes a formatted message at the specified level to the logger output
func (l *Logger) Print(level uint8, f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.minLevel > level {
		return nil
	}

	if l.UseJSON {
		return l.writeJSON(level, f, a...)
	}

	return l.writeText(level, f, a...)
}

// Flush writes any pending buffered data to the underlying file
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

// Debug writes a debug-level message to the logger
func (l *Logger) Debug(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(DEBUG, f, a...)
}

// Info writes an info-level message to the logger
func (l *Logger) Info(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(INFO, f, a...)
}

// Warn writes a warning-level message to the logger
func (l *Logger) Warn(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(WARN, f, a...)
}

// Error writes an error-level message to the logger
func (l *Logger) Error(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(ERROR, f, a...)
}

// Crit writes a critical-level message to the logger
func (l *Logger) Crit(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(CRIT, f, a...)
}

// Aux writes an unskippable message (e.g. separator or header) to the logger
func (l *Logger) Aux(f string, a ...any) error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	return l.Print(AUX, f, a...)
}

// Divider writes an 80-character horizontal rule to the logger output
func (l *Logger) Divider() error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	if l.UseJSON {
		return nil
	}

	return l.Print(AUX, strings.Repeat("-", 80))
}

// Is reports whether the given level meets or exceeds the logger's minimum level
func (l *Logger) Is(level uint8) bool {
	return l != nil && level >= l.minLevel
}

// PanicHandler recovers from a panic and logs it at the critical level.
// It must be called via defer, not directly.
func (l *Logger) PanicHandler(msg string) {
	if l == nil || l.mu == nil {
		return
	}

	r := recover()

	if r != nil {
		stack := debug.Stack()

		if l.UseJSON {
			l.Crit("%s: %v", msg, r, F{"panic-stack", string(stack)})
		} else {
			l.Crit("%s: %v (%s)", msg, r, extractPanicPath(stack))
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns the field formatted as "key:value"
func (f Field) String() string {
	return fmt.Sprintf("%s:%v", f.Key, f.Value)
}

// Mask returns a copy of the field with the middle portion of its value obscured
func (f Field) Mask() Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Mask(v, len(v)/3, 999999, '*')
	return f
}

// Head returns a copy of the field with its value truncated to the first size bytes
func (f Field) Head(size int) Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Head(v, size)
	return f
}

// Tail returns a copy of the field with its value truncated to the last size bytes
func (f Field) Tail(size int) Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Tail(v, size)
	return f
}

// Compact returns a copy of the field with its value shortened to size bytes,
// using an ellipsis in the middle
func (f Field) Compact(size int) Field {
	v := fmt.Sprintf("%v", f.Value)

	if len(v) < size {
		return f
	}

	if size < 3 {
		f.Value = strutil.Head(v, size)
		return f
	}

	size--

	f.Value = strutil.Head(v, size/2) + "…" + strutil.Tail(v, size/2)

	return f
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewFields creates a new Fields collection pre-populated with the given fields
func NewFields(fields ...Field) *Fields {
	result := &Fields{}
	result.Add(fields...)
	return result
}

// Add appends one or more fields to the collection, skipping any with empty keys
func (f *Fields) Add(fields ...Field) *Fields {
	if f == nil || len(fields) == 0 {
		return f
	}

	for _, ff := range fields {
		if ff.Key != "" {
			f.data = append(f.data, ff)
		}
	}

	return f
}

// AddF creates a new field from the given key and value and appends it to the
// collection
func (f *Fields) AddF(key string, value any) *Fields {
	if f == nil || key == "" {
		return f
	}

	f.data = append(f.data, F{key, value})

	return f
}

// Reset removes all fields from the collection
func (f *Fields) Reset() *Fields {
	if f == nil {
		return f
	}

	f.data = nil

	return f
}

// ////////////////////////////////////////////////////////////////////////////////// //

// writeText writes text message into log
func (l *Logger) writeText(level uint8, f string, a ...any) error {
	var color string

	w := l.getWriter(level)

	if l.UseColors {
		color = strutil.B(fmtc.IsTag(Colors[level]), Colors[level], "")
	}

	var err error

	l.buf.Reset()

	if l.UseColors {
		fmtc.Fprintf(&l.buf, "{s}[ %s ]{!} ", l.formatDateTime(time.Now(), false))
	} else {
		l.buf.WriteString("[ " + l.formatDateTime(time.Now(), false) + " ] ")
	}

	if l.WithCaller {
		if l.UseColors {
			fmtc.Fprintf(&l.buf, "{s-}(%s){!} ", getCallerFromStack(l.WithFullCallerPath))
		} else {
			l.buf.WriteString("(" + getCallerFromStack(l.WithFullCallerPath) + ") ")
		}
	}

	if l.isPrefixRequired(level) {
		if l.UseColors {
			fmtc.Fprintf(&l.buf, color+"{@}%s{!} ", PrefixMap[level])
		} else {
			fmt.Fprint(&l.buf, PrefixMap[level]+" ")
		}
	}

	operands, fields := splitPayload(a)

	if l.UseColors {
		fmtc.Fprintf(&l.buf, color+f+"{!}", operands...)
	} else {
		fmt.Fprintf(&l.buf, f, operands...)
	}

	if len(fields) > 0 && !l.DiscardFields {
		l.buf.WriteRune(' ')
		if l.UseColors {
			fmtc.Fprint(&l.buf, strutil.B(level == DEBUG, Colors[DEBUG], "{b}")+fieldsToText(fields)+"{!}")
		} else {
			l.buf.WriteString(fieldsToText(fields))
		}
	}

	if f == "" || strutil.Tail(f, 1) != "\n" {
		l.buf.WriteRune('\n')
	}

	_, err = l.buf.WriteTo(w)

	l.buf.Reset()

	return err
}

// writeJSON writes JSON encoded message into log
func (l *Logger) writeJSON(level uint8, msg string, a ...any) error {
	// Aux in JSON is info
	if level == AUX {
		level = INFO
	}

	if msg == "" && len(a) == 0 {
		return nil
	}

	l.buf.Reset()
	l.buf.WriteRune('{')

	l.writeJSONLevel(level)
	l.writeJSONTimestamp()

	if l.WithCaller {
		l.buf.WriteString(`"caller":"` + getCallerFromStack(l.WithFullCallerPath) + `",`)
	}

	operands, fields := splitPayload(a)

	if msg != "" {
		if len(operands) > 0 {
			l.buf.WriteString(`"msg":` + strconv.Quote(fmt.Sprintf(msg, operands...)))
		} else {
			l.buf.WriteString(`"msg":` + strconv.Quote(msg))
		}
	}

	if len(fields) != 0 {
		if msg != "" {
			l.buf.WriteRune(',')
		}
		l.writeJSONFields(fields)
	}

	l.buf.WriteRune('}')
	l.buf.WriteRune('\n')

	_, err := l.buf.WriteTo(l.getWriter(level))

	l.buf.Reset()

	return err
}

// getWriter returns writer based on logger configuration
func (l *Logger) getWriter(level uint8) io.Writer {
	if l.fd != nil {
		if l.w != nil {
			return l.w
		}

		return l.fd
	}

	if l.UseJSON || (level != ERROR && level != CRIT) {
		return os.Stdout
	}

	return os.Stderr
}

// formatDateTime applies logger datetime layout for given date
func (l *Logger) formatDateTime(t time.Time, isJSON bool) string {
	switch {
	case l.TimeLayout == "" && isJSON:
		return t.Format(DATE_LAYOUT_JSON)
	case l.TimeLayout == "" && !isJSON:
		return t.Format(DATE_LAYOUT_TEXT)
	}

	return t.Format(l.TimeLayout)
}

// isPrefixRequired returns true if prefix must be shown
func (l *Logger) isPrefixRequired(level uint8) bool {
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

// flushDaemon periodically flashes buffered data
func (l *Logger) flushDaemon(interval time.Duration, stopChan <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			l.Flush()
		case <-stopChan:
			return
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// writeJSONLevel writes level JSON into buffer
func (l *Logger) writeJSONLevel(level uint8) {
	switch level {
	case DEBUG:
		l.buf.WriteString(`"level":"debug",`)
	case INFO:
		l.buf.WriteString(`"level":"info",`)
	case WARN:
		l.buf.WriteString(`"level":"warn",`)
	case ERROR:
		l.buf.WriteString(`"level":"error",`)
	case CRIT:
		l.buf.WriteString(`"level":"fatal",`)
	}
}

// writeJSONTimestamp writes timestamp JSON into buffer
func (l *Logger) writeJSONTimestamp() {
	l.buf.WriteString(`"ts":`)

	if l.TimeLayout == "" {
		l.buf.WriteString(strconv.FormatFloat(float64(time.Now().UnixMicro())/1_000_000, 'f', -1, 64))
	} else {
		l.buf.WriteRune('"')
		l.buf.WriteString(l.formatDateTime(time.Now(), true))
		l.buf.WriteRune('"')
	}

	l.buf.WriteRune(',')
}

// writeJSONFields writes fields JSON into buffer
func (l *Logger) writeJSONFields(fields []any) {
	if l.DiscardFields {
		return
	}

	first := true

	for _, f := range fields {
		t, ok := f.(Field)

		if !ok {
			continue
		}

		if !first {
			l.buf.WriteRune(',')
		}

		l.writeJSONField(t)
		first = false
	}
}

// writeJSONField writes field JSON into buffer
func (l *Logger) writeJSONField(field Field) {
	l.buf.WriteString(strconv.Quote(field.Key) + ":")

	switch t := field.Value.(type) {
	case string:
		l.buf.WriteString(strconv.Quote(t))

	case bool:
		l.buf.WriteString(strconv.FormatBool(t))

	case int:
		l.buf.WriteString(strconv.Itoa(t))
	case int8:
		l.buf.WriteString(strconv.FormatInt(int64(t), 10))
	case int16:
		l.buf.WriteString(strconv.FormatInt(int64(t), 10))
	case int32:
		l.buf.WriteString(strconv.FormatInt(int64(t), 10))
	case int64:
		l.buf.WriteString(strconv.FormatInt(t, 10))

	case uint:
		l.buf.WriteString(strconv.FormatUint(uint64(t), 10))
	case uint8:
		l.buf.WriteString(strconv.FormatUint(uint64(t), 10))
	case uint16:
		l.buf.WriteString(strconv.FormatUint(uint64(t), 10))
	case uint32:
		l.buf.WriteString(strconv.FormatUint(uint64(t), 10))
	case uint64:
		l.buf.WriteString(strconv.FormatUint(t, 10))

	case float32:
		l.buf.WriteString(strconv.FormatFloat(float64(t), 'f', -1, 32))
	case float64:
		l.buf.WriteString(strconv.FormatFloat(t, 'f', -1, 64))

	case time.Duration:
		l.buf.WriteString(strconv.FormatFloat(t.Seconds(), 'f', -1, 64))
	case time.Time:
		l.buf.WriteString(strconv.Quote(l.formatDateTime(t, true)))

	default:
		l.buf.WriteString(strconv.Quote(fmt.Sprintf("%v", field.Value)))
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// convertMinLevelValue converts any supported format of minimal level to uint8 used
// for default levels
func convertMinLevelValue(level any) (uint8, error) {
	switch t := level.(type) {
	case int:
		return uint8(t), nil
	case int8:
		return uint8(t), nil
	case int16:
		return uint8(t), nil
	case int32:
		return uint8(t), nil
	case int64:
		return uint8(t), nil
	case uint:
		return uint8(t), nil
	case uint8:
		return uint8(t), nil
	case uint16:
		return uint8(t), nil
	case uint32:
		return uint8(t), nil
	case uint64:
		return uint8(t), nil
	case string:
		code, ok := logLevelsNames[strings.ToLower(t)]

		if !ok {
			return 255, errors.New("Unknown level " + t)
		}

		return code, nil
	}

	return 255, ErrUnexpectedLevel
}

// fieldsToText converts fields slice to string
func fieldsToText(fields []any) string {
	var buf bytes.Buffer

	buf.WriteRune('{')

	for i, f := range fields {
		t, ok := f.(Field)

		if ok {
			v := fmt.Sprintf("%v", t.Value)
			fmt.Fprintf(&buf, "%s: %s", t.Key, strutil.Q(v, `—`))

			if i+1 != len(fields) {
				buf.WriteString(" | ")
			}
		}
	}

	buf.WriteRune('}')

	return buf.String()
}

// splitPayload split mixed payload to format string operands and fields
func splitPayload(payload []any) ([]any, []any) {
	firstField := -1

	// Expand Fields
	for i := range len(payload) {
		switch t := payload[i].(type) {
		case *Fields:
			for _, ff := range t.data {
				payload = append(payload, ff)
			}
		case Fields:
			for _, ff := range t.data {
				payload = append(payload, ff)
			}
		}
	}

	lastField := len(payload)

	// Remove all fields without key
	for i := 0; i < lastField; i++ {
		switch t := payload[i].(type) {
		case Field:
			if t.Key == "" {
				payload[i], payload[lastField-1] = payload[lastField-1], payload[i]
				lastField--
			}
		case Fields, *Fields:
			payload[i], payload[lastField-1] = payload[lastField-1], payload[i]
			lastField--
		}
	}

	payload = payload[:lastField]

	// Move non-field values to the beginning of the slice
	for i, p := range payload {
		switch p.(type) {
		case Field:
			if firstField < 0 {
				firstField = i
			}
		default:
			if firstField > 0 {
				payload[firstField], payload[i] = payload[i], payload[firstField]
				firstField++
			}
		}
	}

	if firstField == -1 {
		return payload, nil
	}

	return payload[:firstField], payload[firstField:]
}

// getCallerFromStack returns caller function and line from stack
func getCallerFromStack(full bool) string {
	pcs := make([]uintptr, 64)
	n := runtime.Callers(2, pcs)

	file := ""
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()

		if !more {
			break
		}

		if file == "" {
			file = frame.File
		}

		if file == frame.File {
			continue
		}

		return extractCallerFromFrame(frame, full)
	}

	return "unknown"
}

// extractCallerFromFrame extracts caller info from frame
func extractCallerFromFrame(f runtime.Frame, full bool) string {
	if full {
		return f.File + ":" + strconv.Itoa(f.Line)
	}

	index := strutil.IndexByteSkip(f.File, '/', -1)
	return f.File[index+1:] + ":" + strconv.Itoa(f.Line)
}

// extractPanicPath tries to extract path to the line with panic
func extractPanicPath(stackData []byte) string {
	stack := string(stackData)
	rs := strings.Index(stack, "panic(")

	if rs == -1 {
		return "unknown position"
	}

	stack = stack[rs:]

	rs = strutil.IndexByteSkip(stack, '\n', 2)
	re := strutil.IndexByteSkip(stack, '\n', 3)

	if rs <= 0 || re <= 0 {
		return "unknown position"
	}

	stack = stack[rs+1 : re]

	fp, _, _ := strings.Cut(strings.Trim(stack, "\t\n"), " ")

	return fp
}
