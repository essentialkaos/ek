package input

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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
	_, err := validatorNotEmpty("")
	c.Assert(err, Equals, ErrIsEmpty)
	_, err = validatorNotEmpty("    ")
	c.Assert(err, Equals, ErrIsEmpty)
	_, err = validatorNotEmpty("test")
	c.Assert(err, IsNil)

	_, err = validatorIsNumber("ABCD")
	c.Assert(err, Equals, ErrInvalidNumber)
	_, err = validatorIsNumber("1234")
	c.Assert(err, IsNil)
	_, err = validatorIsNumber("-1234")
	c.Assert(err, IsNil)
	_, err = validatorIsNumber("")
	c.Assert(err, IsNil)

	_, err = validatorIsFloat("ABCD")
	c.Assert(err, Equals, ErrInvalidFloat)
	_, err = validatorIsFloat("1234.56")
	c.Assert(err, IsNil)
	_, err = validatorIsFloat("-1234.56")
	c.Assert(err, IsNil)
	_, err = validatorIsFloat("")
	c.Assert(err, IsNil)

	_, err = validatorIsEmail("ABCD")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = validatorIsEmail("@test")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = validatorIsEmail("abcd@")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = validatorIsEmail("abcd@test")
	c.Assert(err, Equals, ErrInvalidEmail)
	_, err = validatorIsEmail("")
	c.Assert(err, IsNil)
	_, err = validatorIsEmail("test@domain.com")
	c.Assert(err, IsNil)

	_, err = validatorIsURL("abcd")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = validatorIsURL("abcd.com")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = validatorIsURL("https://abcd")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = validatorIsURL("test://abcd.com")
	c.Assert(err, Equals, ErrInvalidURL)
	_, err = validatorIsURL("")
	c.Assert(err, IsNil)
	_, err = validatorIsURL("https://domain.com")
	c.Assert(err, IsNil)
}
