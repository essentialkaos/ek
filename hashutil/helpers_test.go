package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha1"
	"crypto/sha256"
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type HashUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&HashUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *HashUtilSuite) TestFileHash(c *C) {
	testFile := c.MkDir() + "/test.log"

	err := os.WriteFile(testFile, []byte("ABCDEF12345\n\n"), 0644)
	c.Assert(err, IsNil)

	c.Assert(File("_unknown_", sha1.New()), Equals, "")
	c.Assert(File(testFile, nil), Equals, "")
	c.Assert(File(testFile, sha1.New()), Equals, "9267257cafff1df7a8c0dea354d71c7221d17eda")
	c.Assert(File(testFile, sha256.New()), Equals, "2d7ec20906125cd23fee7b628b98463d554b1105b141b2d39a19bac5f3274dec")
}

func (s *HashUtilSuite) TestDataHash(c *C) {
	c.Assert(Bytes([]byte(""), sha1.New()), Equals, "")
	c.Assert(Bytes([]byte("TEST1234!"), nil), Equals, "")
	c.Assert(Bytes([]byte("TEST1234!"), sha1.New()), Equals, "54b9cb418d8426c4fcf9c91a5923375fd19e6df7")
}

func (s *HashUtilSuite) TestString(c *C) {
	c.Assert(String("", sha1.New()), Equals, "")
	c.Assert(String("TEST1234!", nil), Equals, "")
	c.Assert(String("TEST1234!", sha1.New()), Equals, "54b9cb418d8426c4fcf9c91a5923375fd19e6df7")
}

func (s *HashUtilSuite) TestSum(c *C) {
	c.Assert(Sum(nil), Equals, "")
}
