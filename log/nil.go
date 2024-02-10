package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// NilLogger is Logger (ILogger) compatible logger that doesn't print anything
type NilLogger struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

// Aux do nothing
func (l *NilLogger) Aux(f string, a ...any) error {
	return nil
}

// Debug do nothing
func (l *NilLogger) Debug(f string, a ...any) error {
	return nil
}

// Info do nothing
func (l *NilLogger) Info(f string, a ...any) error {
	return nil
}

// Warn do nothing
func (l *NilLogger) Warn(f string, a ...any) error {
	return nil
}

// Error do nothing
func (l *NilLogger) Error(f string, a ...any) error {
	return nil
}

// Crit do nothing
func (l *NilLogger) Crit(f string, a ...any) error {
	return nil
}

// Print do nothing
func (l *NilLogger) Print(level uint8, f string, a ...any) error {
	return nil
}
