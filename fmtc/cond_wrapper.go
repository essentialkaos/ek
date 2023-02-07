package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "io"

// ////////////////////////////////////////////////////////////////////////////////// //

type CondWrapper struct {
	match bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

// If returns wrapper for printing messages if condition is true
func If(cond bool) CondWrapper {
	return CondWrapper{cond}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Print formats using the default formats for its operands and writes to standard
// output.
func (cw CondWrapper) Print(a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Print(a...)
}

// Println formats using the default formats for its operands and writes to standard
// output. Spaces are always added between operands and a newline is appended. It
// returns the number of bytes written and any write error encountered.
func (cw CondWrapper) Println(a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Println(a...)
}

// Printf formats according to a format specifier and writes to standard output. It
// returns the number of bytes written and any write error encountered.
func (cw CondWrapper) Printf(f string, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Printf(f, a...)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string. It returns the
// number of bytes written and any write error encountered.
func (cw CondWrapper) Fprint(w io.Writer, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Fprint(w, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It returns
// the number of bytes written and any write error encountered.
func (cw CondWrapper) Fprintln(w io.Writer, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Fprintln(w, a...)
}

// Fprintf formats according to a format specifier and writes to w. It returns
// the number of bytes written and any write error encountered.
func (cw CondWrapper) Fprintf(w io.Writer, f string, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return Fprintf(w, f, a...)
}

// Sprint formats using the default formats for its operands and returns the
// resulting string. Spaces are added between operands when neither is a string.
func (cw CondWrapper) Sprint(a ...any) string {
	if cw.match == false {
		return ""
	}

	return Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting
// string.
func (cw CondWrapper) Sprintf(f string, a ...any) string {
	if cw.match == false {
		return ""
	}

	return Sprintf(f, a...)
}

// Sprintln formats using the default formats for its operands and returns the
// resulting string. Spaces are always added between operands and a newline is
// appended.
func (cw CondWrapper) Sprintln(a ...any) string {
	if cw.match == false {
		return ""
	}

	return Sprintln(a...)
}

// TPrint removes all content on the current line and prints the new message
func (cw CondWrapper) TPrint(a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TPrint(a...)
}

// TPrintf removes all content on the current line and prints the new message
func (cw CondWrapper) TPrintf(f string, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TPrintf(f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// with a new line symbol at the end
func (cw CondWrapper) TPrintln(a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TPrintln(a...)
}

// LPrint formats using the default formats for its operands and writes to standard
// output limited by the text size
func (cw CondWrapper) LPrint(maxSize int, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return LPrint(maxSize, a...)
}

// LPrintf formats according to a format specifier and writes to standard output
// limited by the text size
func (cw CondWrapper) LPrintf(maxSize int, f string, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return LPrintf(maxSize, f, a...)
}

// LPrintln formats using the default formats for its operands and writes to standard
// output limited by the text size
func (cw CondWrapper) LPrintln(maxSize int, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return LPrintln(maxSize, a...)
}

// TLPrint removes all content on the current line and prints the new message
// limited by the text size
func (cw CondWrapper) TLPrint(maxSize int, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TLPrint(maxSize, a...)
}

// TLPrintf removes all content on the current line and prints the new message
// limited by the text size
func (cw CondWrapper) TLPrintf(maxSize int, f string, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TLPrintf(maxSize, f, a...)
}

// TPrintln removes all content on the current line and prints the new message
// limited by the text size with a new line symbol at the end
func (cw CondWrapper) TLPrintln(maxSize int, a ...any) (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return TLPrintln(maxSize, a...)
}

// NewLine prints a newline to standard output
func (cw CondWrapper) NewLine() (int, error) {
	if cw.match == false {
		return 0, nil
	}

	return NewLine()
}

// Bell prints alert (bell) symbol
func (cw CondWrapper) Bell() {
	if cw.match == false {
		return
	}

	Bell()
}
