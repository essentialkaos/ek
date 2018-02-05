package log

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Example_logger() {
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

	// AUX message it's unskippable message which will be printed to log file with
	// any minimum level
	logger.Aux("This is aux message")

	// For log rotation we provide method Reopen
	logger.Reopen()

	// If buffered IO is used, you should flush data before exit
	logger.Flush()
}
