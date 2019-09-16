package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	. "pkg.re/check.v1"

	"pkg.re/essentialkaos/ek.v11/fsutil"
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

type ErrReaderWriter struct{}

func (e *ErrReaderWriter) Read(p []byte) (n int, err error) {
	return 0, errors.New("ERROR")
}

func (e *ErrReaderWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *JSONSuite) SetUpSuite(c *C) {
	s.TmpDir = c.MkDir()
}

func (s *JSONSuite) TestDecoding(c *C) {
	jsonFile := s.TmpDir + "/file1.json"

	err := ioutil.WriteFile(jsonFile, []byte(_JSON_DATA), 0644)

	c.Assert(err, IsNil)

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

	err := EncodeToFile(jsonFile, testStruct, 0640)

	c.Assert(err, IsNil)
	c.Assert(fsutil.GetMode(jsonFile), Equals, os.FileMode(0640))

	data, err := ioutil.ReadFile(jsonFile)

	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, _JSON_DATA)
	c.Assert(string(data), Equals, _JSON_DATA)

	err = EncodeToFile("/test.json", testStruct)

	c.Assert(err, NotNil)

	err = EncodeToFile(jsonFile, map[float64]int{3.14: 123})

	c.Assert(err, NotNil)
}

func (s *JSONSuite) TestCompression(c *C) {
	jsonFile := s.TmpDir + "/file3.gz"

	testStruct := &TestStruct{
		String:  "test",
		Integer: 912,
		Boolean: true,
	}

	err := WriteGz(jsonFile, testStruct)

	c.Assert(err, IsNil)
	c.Assert(fsutil.IsNonEmpty(jsonFile), Equals, true)

	testStructDec := &TestStruct{}

	err = ReadGz(jsonFile, testStructDec)

	c.Assert(err, IsNil)
	c.Assert(testStruct.String, Equals, testStructDec.String)
	c.Assert(testStruct.Integer, Equals, testStructDec.Integer)
	c.Assert(testStruct.Boolean, Equals, testStructDec.Boolean)
}

func (s *JSONSuite) TestAux(c *C) {
	erw := &ErrReaderWriter{}

	err := readData(erw, nil, true)
	c.Assert(err, NotNil)

	GzipLevel = -15
	err = writeData(erw, nil, true)
	c.Assert(err, NotNil)
	GzipLevel = 1
}
