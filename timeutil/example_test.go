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
	fmt.Println(PrettyDurationInDays(168 * HOUR))

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
	fmt.Println(MiniDuration(36 * HOUR))
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

func ExampleEndOfHour() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(EndOfHour(d))

	// Output:
	// 2021-06-15 12:59:59.999999999 +0000 UTC
}

func ExampleStartOfDay() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfDay(d))

	// Output:
	// 2021-06-15 00:00:00 +0000 UTC
}

func ExampleEndOfDay() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(EndOfDay(d))

	// Output:
	// 2021-06-15 23:59:59.999999999 +0000 UTC
}

func ExampleStartOfWeek() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfWeek(d, time.Monday))

	// Output:
	// 2021-06-14 00:00:00 +0000 UTC
}

func ExampleEndOfWeek() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(EndOfWeek(d, time.Monday))

	// Output:
	// 2021-06-20 23:59:59.999999999 +0000 UTC
}

func ExampleStartOfMonth() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfMonth(d))

	// Output:
	// 2021-06-01 00:00:00 +0000 UTC
}

func ExampleEndOfMonth() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(EndOfMonth(d))

	// Output:
	// 2021-06-30 23:59:59.999999999 +0000 UTC
}

func ExampleStartOfYear() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(StartOfYear(d))

	// Output:
	// 2021-01-01 00:00:00 +0000 UTC
}

func ExampleEndOfYear() {
	d := time.Date(2021, 6, 15, 12, 30, 15, 0, time.UTC)

	fmt.Println(EndOfYear(d))

	// Output:
	// 2021-12-31 23:59:59.999999999 +0000 UTC
}

func ExamplePrevDay() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevDay(d))

	// Output:
	// 2021-05-31 00:00:00 +0000 UTC
}

func ExamplePrevWeek() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevWeek(d, time.Monday))

	// Output:
	// 2021-05-24 00:00:00 +0000 UTC
}

func ExamplePrevMonth() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevMonth(d))

	// Output:
	// 2021-05-01 00:00:00 +0000 UTC
}

func ExamplePrevYear() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevYear(d))

	// Output:
	// 2020-01-01 00:00:00 +0000 UTC
}

func ExampleNextDay() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextDay(d))

	// Output:
	// 2021-06-02 00:00:00 +0000 UTC
}

func ExampleNextWeek() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextWeek(d, time.Monday))

	// Output:
	// 2021-06-07 00:00:00 +0000 UTC
}

func ExampleNextMonth() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextMonth(d))

	// Output:
	// 2021-07-01 00:00:00 +0000 UTC
}

func ExampleNextYear() {
	d := time.Date(2021, 6, 1, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextYear(d))

	// Output:
	// 2022-01-01 00:00:00 +0000 UTC
}

func ExamplePrevWorkday() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevWorkday(d))

	// Output:
	// 2021-06-04 00:00:00 +0000 UTC
}

func ExamplePrevWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(PrevWeekend(d))

	// Output:
	// 2021-06-05 00:00:00 +0000 UTC
}

func ExampleNextWorkday() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextWorkday(d))

	// Output:
	// 2021-06-07 00:00:00 +0000 UTC
}

func ExampleNextWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(NextWeekend(d))

	// Output:
	// 2021-06-12 00:00:00 +0000 UTC
}

func ExampleIsWeekend() {
	d := time.Date(2021, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(IsWeekend(d))

	// Output:
	// true
}

func ExampleUntil() {
	d := time.Date(2030, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(Until(d, DAY))
}

func ExampleSince() {
	d := time.Date(2012, 6, 6, 12, 30, 15, 0, time.UTC)

	fmt.Println(Since(d, DAY))
}

func ExampleDurationAs() {
	d1 := time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local)
	d2 := time.Date(2024, 6, 15, 14, 15, 45, 0, time.Local)

	d := d2.Sub(d1)

	fmt.Printf("Days: %d\n", DurationAs(d, DAY))
	// Output:
	// Days: 1261
}

func ExampleFromISOWeek() {
	t := FromISOWeek(25, 2021, time.UTC)

	fmt.Println(t)

	// Output:
	// 2021-06-18 00:00:00 +0000 UTC
}

// ////////////////////////////////////////////////////////////////////////////////// //

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

// ////////////////////////////////////////////////////////////////////////////////// //

func ExamplePeriod_Contains() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2023, 6, 15, 18, 45, 30, 0, time.Local),
	}

	d := time.Date(2021, 1, 1, 12, 30, 16, 0, time.Local)

	fmt.Printf("Period contains date: %t\n", p.Contains(d))

	// Output:
	// Period contains date: true
}

func ExamplePeriod_Duration() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 1, 1, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %s\n", p.Duration())

	// Output:
	// Period duration: 1h45m30s
}

func ExamplePeriod_DurationIn() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration (hours): %d\n", p.DurationIn(HOUR))
	fmt.Printf("Period duration (days): %d\n", p.DurationIn(DAY))
	fmt.Printf("Period duration (weeks): %d\n", p.DurationIn(WEEK))

	// Output:
	// Period duration (hours): 3961
	// Period duration (days): 165
	// Period duration (weeks): 23
}

func ExamplePeriod_Seconds() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Seconds())

	// Output:
	// Period duration: 14262330
}

func ExamplePeriod_Minutes() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Minutes())

	// Output:
	// Period duration: 237705
}

func ExamplePeriod_Hours() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Hours())

	// Output:
	// Period duration: 3961
}

func ExamplePeriod_Days() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Days())

	// Output:
	// Period duration: 165
}

func ExamplePeriod_Weeks() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Weeks())

	// Output:
	// Period duration: 23
}

func ExamplePeriod_Years() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2024, 6, 15, 14, 15, 45, 0, time.Local),
	}

	fmt.Printf("Period duration: %d\n", p.Years())

	// Output:
	// Period duration: 3
}

func ExamplePeriod_String() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 1, 1, 14, 15, 45, 0, time.Local),
	}

	fmt.Println(p)

	// Output:
	// 2021/01/01 12:30:15 → 2021/01/01 14:15:45
}

func ExamplePeriod_Stringf() {
	p := Period{
		time.Date(2021, 1, 1, 12, 30, 15, 0, time.Local),
		time.Date(2021, 1, 1, 14, 15, 45, 0, time.Local),
	}

	fmt.Println(p.Stringf(`%H:%M`))

	// Output:
	// 12:30 → 14:15
}
