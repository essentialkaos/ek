package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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

var exitFunc = os.Exit

// ////////////////////////////////////////////////////////////////////////////////// //

// Fatal is analog of Fatal from stdlib
func (l *StdLogger) Fatal(v ...any) {
	l.Logger.Print(CRIT, fmt.Sprint(v...))
	l.Logger.Flush()
	exitFunc(1)
}

// Fatalf is analog of Fatalf from stdlib
func (l *StdLogger) Fatalf(format string, v ...any) {
	l.Logger.Print(CRIT, fmt.Sprintf(format, v...))
	l.Logger.Flush()
	exitFunc(1)
}

// Fatalln is analog of Fatalln from stdlib
func (l *StdLogger) Fatalln(v ...any) {
	l.Logger.Print(CRIT, fmt.Sprintln(v...))
	l.Logger.Flush()
	exitFunc(1)
}

// Output is analog of Output from stdlib
func (l *StdLogger) Output(calldepth int, s string) error {
	return l.Logger.Print(INFO, s)
}

// Panic is analog of Panic from stdlib
func (l *StdLogger) Panic(v ...any) {
	s := fmt.Sprint(v...)
	l.Logger.Print(CRIT, s)
	l.Logger.Flush()
	panic(s)
}

// Panicf is analog of Panicf from stdlib
func (l *StdLogger) Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	l.Logger.Print(CRIT, s)
	l.Logger.Flush()
	panic(s)
}

// Panicln is analog of Panicln from stdlib
func (l *StdLogger) Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	l.Logger.Print(CRIT, s)
	l.Logger.Flush()
	panic(s)
}

// Print is analog of Print from stdlib
func (l *StdLogger) Print(v ...any) {
	l.Logger.Print(INFO, fmt.Sprint(v...))
}

// Printf is analog of Printf from stdlib
func (l *StdLogger) Printf(format string, v ...any) {
	l.Logger.Print(INFO, fmt.Sprintf(format, v...))
}

// Println is analog of Println from stdlib
func (l *StdLogger) Println(v ...any) {
	l.Logger.Print(INFO, fmt.Sprintln(v...))
}
