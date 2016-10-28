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

func (s *PasswdSuite) TestGenAuth(c *C) {
	ad1 := GenAuth(16, STRENGTH_WEAK)

	c.Assert(ad1.Password, HasLen, 16)
	c.Assert(ad1.Salt, HasLen, 16)
	c.Assert(ad1.Hash, HasLen, 64)
	c.Assert(GetPasswordStrength(ad1.Password), Equals, STRENGTH_WEAK)
}

func (s *PasswdSuite) TestGenHash(c *C) {
	pass := "test1234test"
	salt := "saLT"
	hash := "070dd77d9bf913db7681b2271b2328154b53aeb4b70c3742880c08ed32188456"

	v := GenHash(pass, salt)

	c.Assert(v, HasLen, 64)
	c.Assert(v, Equals, hash)
}

func (s *PasswdSuite) BenchmarkGenHash(c *C) {
	for i := 0; i < c.N; i++ {
		GenHash("TEST1234TEST", "12345678")
	}
}
