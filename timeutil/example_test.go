package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
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
	d, _ := ParseDuration("2w3d10h20m35s")

	fmt.Println(d)
	fmt.Println(PrettyDuration(d))

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

func ExampleShortDuration() {
	fmt.Println(ShortDuration(time.Second * 85))

	// Output:
	// 1:25
}

func ExampleDurationToSeconds() {
	fmt.Println(DurationToSeconds(2500 * time.Millisecond))

	// Output:
	// 2.5
}

func ExampleSecondsToDuration() {
	fmt.Println(SecondsToDuration(3600))

	// Output:
	// 1h0m0s
}

func ExampleDate() {
	StartDate = 1577836800

	t := int64(1583020800)
	d := TinyDate(t)

	fmt.Println(t)
	fmt.Println(d)

	// Output:
	// 1583020800
	// 5184000
}

func ExampleDate_Unix() {
	StartDate = 1577836800

	d := TinyDate(1583020800)

	fmt.Println(d.Unix())

	// Output:
	// 1583020800
}

func ExampleDate_Time() {
	StartDate = 1577836800

	d := TinyDate(1583020800)

	fmt.Println(d.Time().In(time.UTC))

	// Output:
	// 2020-03-01 00:00:00 +0000 UTC
}
