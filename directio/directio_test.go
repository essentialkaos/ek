package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type DirectIOSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&DirectIOSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *DirectIOSuite) TestReading(c *C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/tmp_data"
	payload := []byte(strings.Repeat("DATA1", 123))

	err := ioutil.WriteFile(tmpFile, payload, 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	data, err := ReadFile(tmpFile)

	c.Assert(err, IsNil)
	c.Assert(data, NotNil)

	c.Assert(string(data), Equals, strings.Repeat("DATA1", 123))

	data, err = ReadFile(tmpDir + "/not_exist")

	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `open .*: no such file or directory`)
	c.Assert(data, IsNil)
}

func (s *DirectIOSuite) TestWriting(c *C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/tmp_data"
	payload := []byte(strings.Repeat("DATA", 123))

	err := WriteFile(tmpFile, payload, 0644)

	c.Assert(err, IsNil)

	data, err := ioutil.ReadFile(tmpFile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Equals, len(payload))

	err = WriteFile("/not_exist", payload, 0644)

	c.Assert(err, NotNil)

	c.Assert(err, ErrorMatchesOS, map[string]string{
		"darwin": `open /not_exist: read-only file system`,
		"linux":  `open /not_exist: permission denied`,
	})
}
