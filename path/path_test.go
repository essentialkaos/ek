package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type PathUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PathUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PathUtilSuite) TestSanitize(c *C) {
	c.Assert(Sanitize(""), Equals, "")
	c.Assert(Sanitize(`file:name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file*name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file?name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file"name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file<name>.txt`), Equals, "file_name_.txt")
	c.Assert(Sanitize(`file|name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file\name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize(`file/name.txt`), Equals, "file_name.txt")
	c.Assert(Sanitize("file\x00name.txt"), Equals, "file_name.txt")
	c.Assert(Sanitize("file\x1fname.txt"), Equals, "filename.txt")
	c.Assert(Sanitize(".filename.txt"), Equals, "filename.txt")
	c.Assert(Sanitize("filename.txt."), Equals, "filename.txt")
	c.Assert(Sanitize(" filename.txt "), Equals, "filename.txt")
	c.Assert(Sanitize("фото.jpg"), Equals, "фото.jpg")
	c.Assert(Sanitize("фото?.jpg"), Equals, "фото_.jpg")
	c.Assert(Sanitize(string(make([]byte, 300))), HasLen, 255)
}
