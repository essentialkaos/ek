package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type InputSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&InputSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *InputSuite) TestValidators(c *C) {
	_, err := NotEmpty.Validate("")
	c.Assert(err, Equals, ErrIsEmpty)
	_, err = NotEmpty.Validate("    ")
	c.Assert(err, Equals, ErrIsEmpty)
	_, err = NotEmpty.Validate("test")
	c.Assert(err, IsNil)

	_, err = IsNumber.Validate("ABCD")
	c.Assert(err, Equals, ErrInvalidNumber)
	_, err = IsNumber.Validate("1234")
	c.Assert(err, IsNil)
	_, err = IsNumber.Validate("-1234")
	c.Assert(err, IsNil)
	_, err = IsNumber.Validate("")
	c.Assert(err, IsNil)

	_, err = IsFloat.Validate("ABCD")
	c.Assert(err, Equals, ErrInvalidFloat)
	_, err = IsFloat.Validate("1234.56")
	c.Assert(err, IsNil)
	_, err = IsFloat.Validate("-1234.56")
	c.Assert(err, IsNil)
	_, err = IsFloat.Validate("")
	c.Assert(err, IsNil)

	_, err = IsEmail.Validate("ABCD")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = IsEmail.Validate("@test")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = IsEmail.Validate("abcd@")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = IsEmail.Validate("abcd@test")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = IsEmail.Validate("")
	c.Assert(err, IsNil)
	_, err = IsEmail.Validate("test@domain.com")
	c.Assert(err, IsNil)

	_, err = IsURL.Validate("abcd")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = IsURL.Validate("abcd.com")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = IsURL.Validate("https://abcd")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = IsURL.Validate("test://abcd.com")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = IsURL.Validate("")
	c.Assert(err, IsNil)
	_, err = IsURL.Validate("https://domain.com")
	c.Assert(err, IsNil)
}
