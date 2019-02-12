package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StdLogger is logger wrapper compatible with stdlib logger
type StdLogger struct {
	Logger *Logger
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	stdExitFunc  = func(code int) { os.Exit(code) }
	stdPanicFunc = func(message string) { panic(message) }
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Fatal is analog of Fatal from stdlib
func (l *StdLogger) Fatal(v ...interface{}) {
	l.Logger.Print(CRIT, fmt.Sprint(v...))
	stdExitFunc(1)
}

// Fatalf is analog of Fatalf from stdlib
func (l *StdLogger) Fatalf(format string, v ...interface{}) {
	l.Logger.Print(CRIT, fmt.Sprintf(format, v...))
	stdExitFunc(1)
}

// Fatalln is analog of Fatalln from stdlib
func (l *StdLogger) Fatalln(v ...interface{}) {
	l.Logger.Print(CRIT, fmt.Sprintln(v...))
	stdExitFunc(1)
}

// Output is analog of Output from stdlib
func (l *StdLogger) Output(calldepth int, s string) error {
	_, err := l.Logger.Print(INFO, s)
	return err
}

// Panic is analog of Panic from stdlib
func (l *StdLogger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Logger.Print(CRIT, s)
	stdPanicFunc(s)
}

// Panicf is analog of Panicf from stdlib
func (l *StdLogger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Logger.Print(CRIT, s)
	stdPanicFunc(s)
}

// Panicln is analog of Panicln from stdlib
func (l *StdLogger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Logger.Print(CRIT, s)
	stdPanicFunc(s)
}

// Print is analog of Print from stdlib
func (l *StdLogger) Print(v ...interface{}) {
	l.Logger.Print(INFO, fmt.Sprint(v...))
}

// Printf is analog of Printf from stdlib
func (l *StdLogger) Printf(format string, v ...interface{}) {
	l.Logger.Print(INFO, fmt.Sprintf(format, v...))
}

// Println is analog of Println from stdlib
func (l *StdLogger) Println(v ...interface{}) {
	l.Logger.Print(INFO, fmt.Sprintln(v...))
}
