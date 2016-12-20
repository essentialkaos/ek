package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PasswdSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PasswdSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PasswdSuite) TestStrengthCheck(c *C) {
	weakPass1 := "fgiaft"
	weakPass2 := "FgA13"
	weakPass3 := "FgaCCvfaD"
	mediumPass := "AcDr123"
	strongPass := "AbCDEf34%;"

	c.Assert(GetPasswordStrength(weakPass1), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(weakPass2), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(weakPass3), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(mediumPass), Equals, STRENGTH_MEDIUM)
	c.Assert(GetPasswordStrength(strongPass), Equals, STRENGTH_STRONG)
	c.Assert(GetPasswordStrength(""), Equals, STRENGTH_WEAK)
}

func (s *PasswdSuite) TestGenPassword(c *C) {
	c.Assert(GenPassword(0, STRENGTH_WEAK), Equals, "")
	c.Assert(GenPassword(16, STRENGTH_WEAK), HasLen, 16)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_WEAK)), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_MEDIUM)), Equals, STRENGTH_MEDIUM)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_STRONG)), Equals, STRENGTH_STRONG)
	c.Assert(GetPasswordStrength(GenPassword(4, STRENGTH_STRONG)), Equals, STRENGTH_STRONG)

	c.Assert(GetPasswordStrength(GenPassword(16, -100)), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(GenPassword(4, 100)), Equals, STRENGTH_STRONG)
}

func (s *PasswdSuite) TestEncrypt(c *C) {
	hp, err := Encrypt("Test123", "ABCD1234ABCD1234")

	c.Assert(hp, NotNil)
	c.Assert(err, IsNil)

	c.Assert(Check("Test123", "ABCD1234ABCD1234", hp), Equals, true)
	c.Assert(Check("Test123", "ABCD1234ABCD1234", "A1236"), Equals, false)
	c.Assert(Check("Test123", "ABCD1234ABCD1234", "VEVTdA"), Equals, false)
	c.Assert(Check("Test123", "", hp), Equals, false)
	c.Assert(Check("", "ABCD1234ABCD1234", hp), Equals, false)
	c.Assert(Check("", "ABCD1234ABCD1234", hp), Equals, false)
}

func (s *PasswdSuite) TestEncryptErrors(c *C) {
	var err error

	_, err = Encrypt("", "ABCD1234ABCD1234")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Password can't be empty")

	_, err = Encrypt("Test123", "")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Pepper can't be empty")

	_, err = Encrypt("Test123", "ABCD1234ABCD12")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Pepper have invalid size")

	_, ok := unpadData([]byte("-"))

	c.Assert(ok, Equals, false)
}

func (s *PasswdSuite) BenchmarkEncrypt(c *C) {
	for i := 0; i < c.N; i++ {
		Encrypt("Test123", "ABCD1234ABCD1234")
	}
}

func (s *PasswdSuite) BenchmarkCheck(c *C) {
	for i := 0; i < c.N; i++ {
		Check("Test123", "ABCD1234ABCD1234", "jXtzmneskO_ht9VNsuwq68O-jwj3PBxewGrr3YUKf8f7zPqNSlO-Eg7x2KlmoK-wOivvvdaiDpDH_3o5LdWP7ULf6K490KpoNhTZ5XOfaYc")
	}
}
