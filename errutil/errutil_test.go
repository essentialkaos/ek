package errutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ErrSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ErrSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ErrSuite) TestPositive(c *C) {
	var nilErrs *Errors

	errs := NewErrors()

	errs.Add()
	errs.Add(nil)
	errs.Add(nilErrs)
	errs.Add(errors.New("1"))
	errs.Add(errors.New("2"))
	errs.Add(errors.New("3"))
	errs.Add(errors.New("4"))
	errs.Add(errors.New("5"))

	c.Assert(errs.Num(), Equals, 5)
	c.Assert(errs.All(), HasLen, 5)
	c.Assert(errs.HasErrors(), Equals, true)
	c.Assert(errs.Last(), DeepEquals, errors.New("5"))
	c.Assert(errs.All(), DeepEquals,
		[]error{
			errors.New("1"),
			errors.New("2"),
			errors.New("3"),
			errors.New("4"),
			errors.New("5"),
		},
	)
	c.Assert(errs.Add(nil), NotNil)
	c.Assert(errs.Error(), Equals, "  1\n  2\n  3\n  4\n  5\n")
}

func (s *ErrSuite) TestSizeLimit(c *C) {
	errs := NewErrors(3)

	errs.Add(errors.New("1"))
	errs.Add(errors.New("2"))

	c.Assert(errs.HasErrors(), Equals, true)
	c.Assert(errs.Num(), Equals, 2)
	c.Assert(errs.All(), HasLen, 2)

	errs.Add(errors.New("3"))
	errs.Add(errors.New("4"))
	errs.Add(errors.New("5"))
	errs.Add(errors.New("6"))

	c.Assert(errs.HasErrors(), Equals, true)
	c.Assert(errs.Num(), Equals, 3)
	c.Assert(errs.All(), HasLen, 3)

	errList := errs.All()

	c.Assert(errList[0].Error(), Equals, "4")
	c.Assert(errList[2].Error(), Equals, "6")

	errs = NewErrors(-10)

	c.Assert(errs.maxSize, Equals, 0)
}

func (s *ErrSuite) TestAdd(c *C) {
	errs1 := NewErrors()
	errs2 := NewErrors()

	errs1.Add(errors.New("1"))
	errs1.Add(errors.New("2"))

	errs2.Add(errors.New("3"))
	errs2.Add(errors.New("4"))

	errs1.Add(errs2)
	errs1.Add([]error{errors.New("5"), errors.New("6")})

	c.Assert(errs1.HasErrors(), Equals, true)
	c.Assert(errs1.Num(), Equals, 6)
	c.Assert(errs1.All(), HasLen, 6)
}

func (s *ErrSuite) TestNegative(c *C) {
	errs := NewErrors()

	c.Assert(errs.All(), HasLen, 0)
	c.Assert(errs.HasErrors(), Equals, false)
	c.Assert(errs.Last(), IsNil)
	c.Assert(errs.Error(), Equals, "")
}

func (s *ErrSuite) TestNil(c *C) {
	var errs *Errors

	c.Assert(errs.Num(), Equals, 0)
	c.Assert(errs.All(), HasLen, 0)
	c.Assert(errs.HasErrors(), Equals, false)
	c.Assert(errs.Last(), IsNil)
	c.Assert(errs.Error(), Equals, "")
}

func (s *ErrSuite) TestChain(c *C) {
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
