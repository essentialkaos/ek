package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"strings"
	"testing"

	. "pkg.re/check.v1"
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
	payload := []byte(strings.Repeat("DATA1", 6543))

	err := ioutil.WriteFile(tmpFile, payload, 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	data, err := ReadFile(tmpFile)

	c.Assert(err, IsNil)
	c.Assert(data, NotNil)

	c.Assert(string(data), Equals, strings.Repeat("DATA1", 6543))

	data, err = ReadFile(tmpDir + "/not_exist")

	c.Assert(err, NotNil)
	c.Assert(data, IsNil)
}

func (s *DirectIOSuite) TestWriting(c *C) {
	tmpDir := c.MkDir()
	tmpFile := tmpDir + "/tmp_data"
	payload := []byte(strings.Repeat("DATA", 6543))

	err := WriteFile(tmpFile, payload, 0644)

	c.Assert(err, IsNil)

	data, err := ioutil.ReadFile(tmpFile)

	c.Assert(err, IsNil)
	c.Assert(len(data), Equals, len(payload))

	err = WriteFile("/not_exist", payload, 0644)

	c.Assert(err, NotNil)
}

func (s *DirectIOSuite) BenchmarkAllocation(c *C) {
	for i := 0; i < c.N; i++ {
		block := allocateBlock()
		freeBlock(block)
	}
}
