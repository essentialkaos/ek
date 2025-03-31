package reutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"regexp"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ReUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ReUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ReUtilSuite) TestReplace(c *C) {
	re := regexp.MustCompile(`[a-z0-9]{2,4}`)

	_, err := Replace(nil, "", nil)
	c.Assert(err, Equals, ErrNilRegex)

	_, err = Replace(re, "", nil)
	c.Assert(err, Equals, ErrNilFunc)

	str, err := Replace(re, "", func(_ string, _ []string) string { return "" })
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "")

	str, err = Replace(re, "#####", func(_ string, _ []string) string { return "" })
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "#####")

	str, err = Replace(re, "test.AF-12345678", func(found string, submatch []string) string { return "[" + found + "]" })
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "[test].AF-[1234][5678]")

	re = regexp.MustCompile(`^([a-z0-9]+).AF\-([\d]+)$`)

	str, err = Replace(re, "test.AF-12345678", func(found string, submatch []string) string { return "[" + found + "]" })
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "[test.AF-12345678]")

	re = regexp.MustCompile(`^test$`)

	str, err = Replace(re, "test", func(found string, submatch []string) string { return "12345678" })
	c.Assert(err, IsNil)
	c.Assert(str, Equals, "12345678")
}
