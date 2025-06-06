// Package log provides an improved logger
package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/strutil"
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

const (
	DATE_LAYOUT_TEXT = "2006/01/02 15:04:05.000" // Datetime layout for text logs
	DATE_LAYOUT_JSON = time.RFC3339              // Datetime layout for JSON logs
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

	file     string
	buf      bytes.Buffer
	fd       *os.File
	w        *bufio.Writer
	mu       *sync.Mutex
	minLevel uint8
	perms    os.FileMode
	useBufIO bool
}

// F is a shortcut for Field struct
type F = Field

// Field contains key and value for JSON log
type Field struct {
	Key   string
	Value any
}

// Fields is a field collection
type Fields struct {
	data []Field
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
	ERROR: "{#208}",
	CRIT:  "{#196}{*}",
}

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

// Divider writes simple divider to logger output
func Divider() error {
	return Global.Divider()
}

// Is returns true if current minimal logging level is equal or greater than
// given
func Is(level uint8) bool {
	return Global.Is(level)
}

// Levels returns slice with all supported log levels
func Levels() []string {
	return logLevels
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

	go l.flushDaemon(flushInterval)
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

	if l.UseJSON {
		return l.writeJSON(level, f, a...)
	}

	return l.writeText(level, f, a...)
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

// Divider writes simple divider to logger output
func (l *Logger) Divider() error {
	if l == nil || l.mu == nil {
		return ErrNilLogger
	}

	if l.UseJSON {
		return nil
	}

	return l.Print(AUX, strings.Repeat("-", 80))
}

// Is returns true if current minimal logging level is equal or greater than
// given
func (l *Logger) Is(level uint8) bool {
	return l != nil && level >= l.minLevel
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of field
func (f Field) String() string {
	return fmt.Sprintf("%s:%v", f.Key, f.Value)
}

// Mask masks part of the field value
func (f Field) Mask() Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Mask(v, len(v)/3, 999999, '*')
	return f
}

// Head trims the last part of field value
func (f Field) Head(size int) Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Head(v, size)
	return f
}

// Tail trims the first part of field value
func (f Field) Tail(size int) Field {
	v := fmt.Sprintf("%v", f.Value)
	f.Value = strutil.Tail(v, size)
	return f
}

// Compact shrinks field value to given size
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

// NewFields creates new fields collection
func NewFields(fields ...Field) *Fields {
	result := &Fields{}
	result.Add(fields...)
	return result
}

// Add adds given fields to collection
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

// AddF creates and adds field to collection
func (f *Fields) AddF(key string, value any) *Fields {
	if f == nil || key == "" {
		return f
	}

	f.data = append(f.data, F{key, value})

	return f
}

// Reset removes all fields from collection
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
			fmt.Fprintf(&l.buf, PrefixMap[level]+" ")
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
	var w io.Writer

	if l.fd == nil {
		if l.UseJSON {
			w = os.Stdout
		} else {
			switch level {
			case ERROR, CRIT:
				w = os.Stderr
			default:
				w = os.Stdout
			}
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
func (l *Logger) flushDaemon(interval time.Duration) {
	for range time.NewTicker(interval).C {
		l.Flush()
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

	for i, f := range fields {
		switch t := f.(type) {
		case Field:
			l.writeJSONField(t)
			if i+1 != len(fields) {
				l.buf.WriteRune(',')
			}
		}
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
		switch t := f.(type) {
		case Field:
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
