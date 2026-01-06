package errors

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ErrorsSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ErrorsSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ErrorsSuite) TestStd(c *C) {
	c.Assert(New("test"), NotNil)
	c.Assert(Is(nil, nil), Equals, true)
	c.Assert(As(nil, nil), Equals, false)
	c.Assert(Join(New("1"), New("2")), NotNil)
	c.Assert(Unwrap(New("1")), IsNil)
}

func (s *ErrorsSuite) TestPositive(c *C) {
	var nilBundle *Bundle
	var nilErrs Errors

	errs := NewBundle()

	errs.Add()
	errs.Add(nil)
	errs.Add(nilBundle)
	errs.Add(nilErrs)

	errs.Add(errors.New("1"))
	errs.Add(errors.New("2"))
	errs.Add(errors.New("3"))
	errs.Add(errors.New("4"))
	errs.Add(errors.New("5"))

	c.Assert(errs.All(), HasLen, 5)
	c.Assert(errs.Num(), Equals, 5)
	c.Assert(errs.IsEmpty(), Equals, false)
	c.Assert(errs.First(), DeepEquals, errors.New("1"))
	c.Assert(errs.Last(), DeepEquals, errors.New("5"))
	c.Assert(errs.Get(0), DeepEquals, errors.New("1"))
	c.Assert(errs.Get(4), DeepEquals, errors.New("5"))
	c.Assert(errs.Get(100), IsNil)
	c.Assert(errs.Get(-100), IsNil)
	c.Assert(errs.All(), DeepEquals,
		Errors{
			errors.New("1"),
			errors.New("2"),
			errors.New("3"),
			errors.New("4"),
			errors.New("5"),
		},
	)
	c.Assert(errs.Add(nil), NotNil)
	c.Assert(errs.Error("  "), Equals, "  1\n  2\n  3\n  4\n  5")

	errs.Reset()

	c.Assert(errs.Num(), Equals, 0)
}

func (s *ErrorsSuite) TestSizeLimit(c *C) {
	errs := NewBundle(3)

	errs.Add(errors.New("1"))
	errs.Add(errors.New("2"))

	c.Assert(errs.IsEmpty(), Equals, false)
	c.Assert(errs.Num(), Equals, 2)
	c.Assert(errs.All(), HasLen, 2)

	errs.Add(errors.New("3"))
	errs.Add(errors.New("4"))
	errs.Add(errors.New("5"))
	errs.Add(errors.New("6"))

	c.Assert(errs.IsEmpty(), Equals, false)
	c.Assert(errs.Num(), Equals, 3)
	c.Assert(errs.Cap(), Equals, 3)
	c.Assert(errs.All(), HasLen, 3)

	errList := errs.All()

	c.Assert(errList[0].Error(), Equals, "4")
	c.Assert(errList[2].Error(), Equals, "6")

	errs = NewBundle(-10)

	c.Assert(errs.capacity, Equals, 0)
}

func (s *ErrorsSuite) TestAdd(c *C) {
	errs1 := NewBundle()
	errs2 := NewBundle()

	var errs3 Bundle

	errs1.Add(errors.New("1"))
	errs1.Add(errors.New("2"))

	errs2.Add(errors.New("3"))
	errs2.Add(errors.New("4"))

	errs3.Add(errors.New("5"))
	errs3.Add(errors.New("6"))

	errs1.Add(errs2)
	errs1.Add(errs3)

	errs1.Add([]error{errors.New("7"), errors.New("8")})
	errs1.Add([]string{"9", "10"})
	errs1.Add("11")
	errs1.Add(Errors{New("12")})
	errs1.Addf("Test %s %d", "error", 100)

	c.Assert(errs1.IsEmpty(), Equals, false)
	c.Assert(errs1.Num(), Equals, 13)
	c.Assert(errs1.All(), HasLen, 13)
}

func (s *ErrorsSuite) TestNegative(c *C) {
	errs := NewBundle()

	c.Assert(errs.All(), HasLen, 0)
	c.Assert(errs.IsEmpty(), Equals, true)
	c.Assert(errs.Last(), IsNil)
	c.Assert(errs.Error(""), Equals, "")
}

func (s *ErrorsSuite) TestNil(c *C) {
	var errs *Bundle

	c.Assert(errs.Num(), Equals, 0)
	c.Assert(errs.Cap(), Equals, 0)
	c.Assert(errs.All(), HasLen, 0)
	c.Assert(errs.IsEmpty(), Equals, true)
	c.Assert(errs.First(), IsNil)
	c.Assert(errs.Last(), IsNil)
	c.Assert(errs.Error(""), Equals, "")
	c.Assert(errs.Get(10), IsNil)
}

func (s *ErrorsSuite) TestNoInit(c *C) {
	var errs Bundle

	c.Assert(errs.Num(), Equals, 0)
	c.Assert(errs.Cap(), Equals, 0)
	c.Assert(errs.All(), HasLen, 0)
	c.Assert(errs.IsEmpty(), Equals, true)
	c.Assert(errs.First(), IsNil)
	c.Assert(errs.Last(), IsNil)
	c.Assert(errs.Error(""), Equals, "")

	c.Assert(errs.Add(errors.New("1")), NotNil)
	c.Assert(errs.Last(), DeepEquals, errors.New("1"))
}

func (s *ErrorsSuite) TestChain(c *C) {
	f1 := func() error {
		return nil
	}

	f2 := func() error {
		return errors.New("Error #2")
	}

	f3 := func() error {
		return nil
	}

	c.Assert(Chain(f1, f2, f3), NotNil)
	c.Assert(Chain(f1, f3), IsNil)
}

func (s *ErrorsSuite) TestToBundle(c *C) {
	errs := ToBundle([]error{
		errors.New("Error 1"),
		errors.New("Error 2"),
	})

	c.Assert(errs.Num(), Equals, 2)
	c.Assert(errs.Last(), DeepEquals, errors.New("Error 2"))
}
