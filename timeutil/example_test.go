package timeutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePrettyDuration() {
	// you can use int/int64/uint as seconds...
	fmt.Println(PrettyDuration(300))

	// ...or time.Duration types
	fmt.Println(PrettyDuration(123456 * time.Second))

	// Output:
	// 5 minutes
	// 1 day 10 hours 17 minutes and 36 seconds
}

func ExamplePrettyDurationSimple() {
	// you can use int/int64/uint as seconds...
	fmt.Println(PrettyDurationSimple(300))

	// ...or time.Duration types
	fmt.Println(PrettyDurationSimple(50400 * time.Second))

	// Output:
	// 5 minutes
	// 14 hours
}

func ExamplePrettyDurationInDays() {
	fmt.Println(PrettyDurationInDays(2 * time.Hour))
	fmt.Println(PrettyDurationInDays(168 * time.Hour))

	// Output:
	// 1 day
	// 7 days
}

func ExampleParseDuration() {
	d, _ := ParseDuration("2w3d10h20m35s")

	fmt.Println(PrettyDuration(d))

	// Output:
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
	fmt.Println(ShortDuration(3215*time.Millisecond, true))

	// Output:
	// 1:25
	// 0:03.215
}

func ExampleMiniDuration() {
	fmt.Println(MiniDuration(36 * time.Hour))
	fmt.Println(MiniDuration(18 * time.Second))
	fmt.Println(MiniDuration(time.Second / 125))

	// You can remove or change separator
	fmt.Println(MiniDuration(time.Second/2000, ""))

	// Output:
	// 2 d
	// 18 s
	// 8 ms
	// 500μs
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

func ExampleStartOfHour() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfHour(d))

	// Output:
	// 2021-06-15 12:00:00 +0000 UTC
}

func ExampleStartOfDay() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfDay(d))

	// Output:
	// 2021-06-15 00:00:00 +0000 UTC
}

func ExampleStartOfWeek() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfWeek(d, time.Monday))

	// Output:
	// 2021-06-14 00:00:00 +0000 UTC
}

func ExampleStartOfMonth() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfMonth(d))

	// Output:
	// 2021-06-01 00:00:00 +0000 UTC
}

func ExampleStartOfYear() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfYear(d))

	// Output:
	// 2021-01-01 00:00:00 +0000 UTC
}

func ExamplePrevDay() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevDay(d))

	// Output:
	// 2021-05-31 12:30:15 +0000 UTC
}

func ExamplePrevMonth() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevMonth(d))

	// Output:
	// 2021-05-01 12:30:15 +0000 UTC
}

func ExamplePrevYear() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevYear(d))

	// Output:
	// 2020-06-01 12:30:15 +0000 UTC
}

func ExampleNextDay() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextDay(d))

	// Output:
	// 2021-06-02 12:30:15 +0000 UTC
}

func ExampleNextMonth() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextMonth(d))

	// Output:
	// 2021-07-01 12:30:15 +0000 UTC
}

func ExampleNextYear() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextYear(d))

	// Output:
	// 2022-06-01 12:30:15 +0000 UTC
}

func ExamplePrevWorkday() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevWorkday(d))

	// Output:
	// 2021-06-04 12:30:15 +0000 UTC
}

func ExamplePrevWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevWeekend(d))

	// Output:
	// 2021-06-05 12:30:15 +0000 UTC
}

func ExampleNextWorkday() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextWorkday(d))

	// Output:
	// 2021-06-07 12:30:15 +0000 UTC
}

func ExampleNextWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextWeekend(d))

	// Output:
	// 2021-06-12 12:30:15 +0000 UTC
}

func ExampleIsWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(IsWeekend(d))

	// Output:
	// true
}

func ExampleFromISOWeek() {
	t := FromISOWeek(25, 2021, time.UTC)

	fmt.Println(t)

	// Output:
	// 2021-06-18 00:00:00 +0000 UTC
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
