package crypto

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type CryptoSuite struct {
	TmpDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CryptoSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CryptoSuite) SetUpSuite(c *C) {
	s.TmpDir = c.MkDir()
}

func (s *CryptoSuite) TestJCHash(c *C) {
	c.Assert(JCHash(1, 1), Equals, 0)
	c.Assert(JCHash(36, 49), Equals, 8)
	c.Assert(JCHash(0xDEAD10CC, 1), Equals, 0)
	c.Assert(JCHash(0xDEAD10CC, 1000), Equals, 361)
	c.Assert(JCHash(128, 1024), Equals, 267)
}

func (s *CryptoSuite) TestJCHashNegative(c *C) {
	c.Assert(JCHash(0, -10), Equals, 0)
	c.Assert(JCHash(0xDEAD10CC, -1000), Equals, 0)
}

func (s *CryptoSuite) TestStrengthCheck(c *C) {
	weakPass1 := "fgiaft"
	weakPass2 := "FgA13"
	mediumPass := "AcDr123"
	strongPass := "AbCDEf34%;"

	c.Assert(GetPasswordStrength(weakPass1), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(weakPass2), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(weakPass2), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(mediumPass), Equals, STRENGTH_MEDIUM)
	c.Assert(GetPasswordStrength(strongPass), Equals, STRENGTH_STRONG)
	c.Assert(GetPasswordStrength(""), Equals, STRENGTH_WEAK)
}

func (s *CryptoSuite) TestGenPassword(c *C) {
	c.Assert(GenPassword(0, STRENGTH_WEAK), Equals, "")
	c.Assert(GenPassword(16, STRENGTH_WEAK), HasLen, 16)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_WEAK)), Equals, STRENGTH_WEAK)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_MEDIUM)), Equals, STRENGTH_MEDIUM)
	c.Assert(GetPasswordStrength(GenPassword(16, STRENGTH_STRONG)), Equals, STRENGTH_STRONG)
	c.Assert(GetPasswordStrength(GenPassword(4, STRENGTH_STRONG)), Equals, STRENGTH_STRONG)
}

func (s *CryptoSuite) TestGenAuth(c *C) {
	ad1 := GenAuth(16, STRENGTH_WEAK)

	c.Assert(ad1.Password, HasLen, 16)
	c.Assert(ad1.Salt, HasLen, 16)
	c.Assert(ad1.Hash, HasLen, 64)
	c.Assert(GetPasswordStrength(ad1.Password), Equals, STRENGTH_WEAK)
}

func (s *CryptoSuite) TestGenHash(c *C) {
	pass := "test1234test"
	salt := "saLT"
	hash := "070dd77d9bf913db7681b2271b2328154b53aeb4b70c3742880c08ed32188456"

	v := GenHash(pass, salt)

	c.Assert(v, HasLen, 64)
	c.Assert(v, Equals, hash)
}

func (s *CryptoSuite) TestGenUUID(c *C) {
	c.Assert(GenUUID(), HasLen, 36)
}

func (s *CryptoSuite) TestFileHash(c *C) {
	tempFile := s.TmpDir + "/test.log"

	err := ioutil.WriteFile(tempFile, []byte("ABCDEF12345\n\n"), 0644)

	c.Assert(err, IsNil)
	c.Assert(FileHash(tempFile), Equals, "2d7ec20906125cd23fee7b628b98463d554b1105b141b2d39a19bac5f3274dec")
	c.Assert(FileHash(s.TmpDir+"/not-exist.log"), Equals, "")
}
