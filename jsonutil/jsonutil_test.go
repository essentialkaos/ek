package jsonutil

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

	"pkg.re/essentialkaos/ek.v3/crypto"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _JSON_DATA = `{
  "string": "test",
  "integer": 912,
  "boolean": true
}
`

// ////////////////////////////////////////////////////////////////////////////////// //

type TestStruct struct {
	String  string `json:"string"`
	Integer int    `json:"integer"`
	Boolean bool   `json:"boolean"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type JSONSuite struct {
	TmpDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&JSONSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *JSONSuite) SetUpSuite(c *C) {
	s.TmpDir = c.MkDir()
}

func (s *JSONSuite) TestDecoding(c *C) {
	jsonFile := s.TmpDir + "/file1.json"

	err := ioutil.WriteFile(jsonFile, []byte(_JSON_DATA), 0644)

	c.Assert(err, IsNil)
	c.Assert(crypto.FileHash(jsonFile), Equals,
		"b88184cc9c6c517e572a21acae7118d58c485b051c33f3f13d057b43461c4eec")

	testStruct := &TestStruct{}

	err = DecodeFile(s.TmpDir+"/file-not-exists.json", &TestStruct{})

	c.Assert(err, NotNil)

	err = DecodeFile(s.TmpDir+"/file1.json", testStruct)

	c.Assert(err, IsNil)
	c.Assert(testStruct.String, Equals, "test")
	c.Assert(testStruct.Integer, Equals, 912)
	c.Assert(testStruct.Boolean, Equals, true)
}

func (s *JSONSuite) TestEncoding(c *C) {
	jsonFile := s.TmpDir + "/file2.json"

	testStruct := &TestStruct{
		String:  "test",
		Integer: 912,
		Boolean: true,
	}

	err := EncodeToFile(jsonFile, testStruct)

	c.Assert(err, IsNil)
	c.Assert(crypto.FileHash(jsonFile), Equals,
		"b88184cc9c6c517e572a21acae7118d58c485b051c33f3f13d057b43461c4eec")

	err = EncodeToFile("/test.json", testStruct)

	c.Assert(err, NotNil)
}
