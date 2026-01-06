package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SpinnerSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SpinnerSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SpinnerSuite) TestSpinner(c *C) {
	Done(true) // skipped
	Show("ABCD")
	Show("ABCD") // skipped
	time.Sleep(time.Millisecond * 800)
	Update("ABCD")
	time.Sleep(time.Millisecond * 100)
	Done(true)
	Update("ABCD") // skipped
	Show("ABCD")
	time.Sleep(time.Millisecond * 10)
	Done(false)
	Show("ABCD")
	time.Sleep(time.Millisecond * 10)
	Skip()

	DisableAnimation = true
	Show("ABCD")
	Done(true)

	Skip()
}

func (s *SpinnerSuite) TestAux(c *C) {
	DurationFormat = DURATION_SHORT
	c.Assert(formatDuration(time.Minute), Equals, "1:00")

	DurationFormat = DURATION_MINI
	c.Assert(formatDuration(time.Minute), Equals, "1m")

	DurationFormat = DURATION_SIMPLE
	c.Assert(formatDuration(time.Minute), Equals, "1 minute")
}
