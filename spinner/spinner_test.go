package spinner

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"
	"time"

	. "pkg.re/essentialkaos/check.v1"
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
	time.Sleep(time.Millisecond * 100)
	Update("ABCD")
	time.Sleep(time.Millisecond * 100)
	Done(true)
	Update("ABCD") // skipped
	Show("ABCD")
	time.Sleep(time.Millisecond * 10)
	Done(false)

	DisableAnimation = true
	Show("ABCD")
	Done(true)
}
