package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Example() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Enable buffered IO with 1-second flush interval
	logger.EnableBufIO(time.Second)

	// Set minimal log level
	logger.MinLevel(INFO)

	// Print message to log
	logger.Print(DEBUG, "This is %s message", "debug")

	// Package provides different methods for each level
	logger.Debug("This is %d %s message", 2, "debug")
	logger.Info("This is info message")
	logger.Warn("This is warning message")
	logger.Error("This is error message")
	logger.Crit("This is critical message")

	// Enable colors for output
	logger.UseColors = true

	// Encode messages to JSON
	logger.UseJSON = true

	// Print caller info
	logger.WithCaller = true

	// Use custom date & time layout
	logger.TimeLayout = time.RFC3339

	// Add fields to message
	logger.Debug("This is %d %s message", 2, "debug", F{"user", "bob"}, F{"id", 200})

	// AUX message it's unskippable message which will be printed to log file with
	// any minimum level
	//
	// Note that all AUX messages are dropped when using JSON format
	logger.Aux("This is aux message")

	// Print simple divider
	logger.Divider()

	// For log rotation we provide method Reopen
	logger.Reopen()

	// If buffered IO is used, you should flush data before exit
	logger.Flush()
}

func ExampleNew() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Package provides different methods for each level
	logger.Debug("This is %d %s message", 2, "debug")
	logger.Info("This is info message")
	logger.Warn("This is warning message")
	logger.Error("This is error message")
	logger.Crit("This is critical message")

	// If buffered IO is used, you should flush data before exit
	logger.Flush()
}

func ExampleReopen() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// For log rotation we provide method Reopen
	Reopen()
}

func ExampleMinLevel() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Set minimal log level to error
	MinLevel(ERROR)

	Info("This message is not displayed")
	Error("This message is displayed")
}

func ExampleSet() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = Set("/path/to/file2.log", 0640)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Info("Message will go to file2.log")
}

func ExampleEnableBufIO() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Enable buffered IO with 1-second flush interval
	EnableBufIO(time.Second)

	Info("Info message")

	Flush()
}

func ExampleFlush() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Enable buffered IO with 1-second flush interval
	EnableBufIO(time.Second)

	Info("Info message")

	Flush()
}

func ExamplePrint() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Print(INFO, "Info message")
	Print(ERROR, "Error message")
}

func ExampleDebug() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Debug("Debug message")
}

func ExampleInfo() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Info("Info message")
}

func ExampleWarn() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Warn("Warning message")
}

func ExampleError() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Error("Error message")
}

func ExampleCrit() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Crit("Critical error message")
}

func ExampleAux() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Aux("Auxiliary message")
}

func ExampleDivider() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	Divider()
}

func ExampleIs() {
	err := Set("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if Is(INFO) {
		Info("Info message")
	}
}

func ExampleLogger_Reopen() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// For log rotation we provide method Reopen
	logger.Reopen()
}

func ExampleLogger_MinLevel() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Set minimal log level to error
	logger.MinLevel(ERROR)

	logger.Info("This message is not displayed")
	logger.Error("This message is displayed")
}

func ExampleLogger_Set() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = logger.Set("/path/to/file2.log", 0640)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Info("Message will go to file2.log")
}

func ExampleLogger_EnableBufIO() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Enable buffered IO with 1-second flush interval
	logger.EnableBufIO(time.Second)

	logger.Info("Info message")

	logger.Flush()
}

func ExampleLogger_Flush() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Enable buffered IO with 1-second flush interval
	logger.EnableBufIO(time.Second)

	logger.Info("Info message")

	logger.Flush()
}

func ExampleLogger_Print() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Print(INFO, "Info message")
	logger.Print(ERROR, "Error message")
}

func ExampleLogger_Debug() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Debug("Debug message")
}

func ExampleLogger_Info() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Info("Info message")
}

func ExampleLogger_Warn() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Warn("Warning message")
}

func ExampleLogger_Error() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Error("Error message")
}

func ExampleLogger_Crit() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Crit("Critical error message")
}

func ExampleLogger_Aux() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Aux("Auxiliary message")
}

func ExampleLogger_Divider() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	logger.Divider()
}

func ExampleLogger_Is() {
	logger, err := New("/path/to/file.log", 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if logger.Is(INFO) {
		logger.Info("Info message")
	}
}
