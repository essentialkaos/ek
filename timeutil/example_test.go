package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePrettyDuration() {
	// you can use int ...
	fmt.Println(PrettyDuration(300))

	// ... and time.Duration types
	fmt.Println(PrettyDuration(123456 * time.Second))

	// Output:
	// 5 minutes
	// 1 day 10 hours 17 minutes and 36 seconds
}

func ExampleParseDuration() {
	fmt.Println(ParseDuration("2w3d10h20m35s"))
	fmt.Println(PrettyDuration(ParseDuration("2w3d10h20m35s")))

	// Output:
	// 1506035
	// 2 weeks 3 days 10 hours 20 minutes and 35 seconds
}

func ExampleFormat() {
	date := time.Date(2010, 6, 15, 15, 30, 45, 1234, time.Local)

	fmt.Println(Format(date, "%A %d/%b/%Y %H:%M:%S.%N"))

	// Output:
	// Tuesday 15/Jun/2010 15:30:45.000001234
}

func ExampleDurationToSeconds() {
	fmt.Println(DurationToSeconds(time.Minute))

	// Output:
	// 60
}
