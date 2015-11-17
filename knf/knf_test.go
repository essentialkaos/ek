package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	check "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
    [formating]
test1: 1
            test2: 2

      test3: 3

[string]
  test1: test
  test2: true
  test3: 4500
  test4: !$%^&
  test5: long long long long text for test
  test6: 

[boolean]
  test1: true
  test2: false
  test3: 0
  test4: 1
  test5:
  test6: example for test

[integer]
  test1: 1
  test2: -5
  test3: 10000000
  test4: A
  test5: 0xFF
  test6: 123.4
  test7: 123.456789

[file-mode]
  test1: 644
  test2: 0644
  test3: 0
  test4: ABC
  test5: true

[comment]
  test1: 100
  # test2: 100

[macro]
  test1: 100
  test2: {macro:test1}.50
  test3: Value is {macro:test2}
  test4: "{macro:test3}"
  test5: {ABC}
  test6: {}
`

const _CONFIG_FILE_NAME = "knf-config-test.conf"

// ////////////////////////////////////////////////////////////////////////////////// //

type KNFSuite struct {
	Config     *Config
	ConfigPath string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&KNFSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) {
	check.TestingT(t)
}

func (s *KNFSuite) SetUpSuite(c *check.C) {
	tmpdir := c.MkDir()

	s.ConfigPath = tmpdir + "/" + _CONFIG_FILE_NAME

	err := ioutil.WriteFile(s.ConfigPath, []byte(_CONFIG_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *KNFSuite) SetUpTest(c *check.C) {
	cf, err := Read(s.ConfigPath + "123")

	c.Assert(err, check.NotNil)

	cf, err = Read(s.ConfigPath)

	if err != nil {
		c.Fatal(err.Error())
	}

	if cf == nil {
		c.Fatal("Config struct is nil")
	}

	s.Config = cf
}

func (s *KNFSuite) TestSections(c *check.C) {
	sections := s.Config.Sections()

	c.Assert(len(sections), check.Equals, 7)
	c.Assert(
		sections,
		check.DeepEquals,
		[]string{"formating", "string", "boolean", "integer", "file-mode", "comment", "macro"},
	)
}

func (s *KNFSuite) TestProps(c *check.C) {
	props := s.Config.Props("integer")

	c.Assert(len(props), check.Equals, 7)
	c.Assert(
		props,
		check.DeepEquals,
		[]string{"test1", "test2", "test3", "test4", "test5", "test6", "test7"},
	)
}

func (s *KNFSuite) TestCheckers(c *check.C) {
	c.Assert(s.Config.HasSection("string"), check.Equals, true)
	c.Assert(s.Config.HasSection("strings"), check.Equals, false)

	c.Assert(s.Config.HasProp("string:test1"), check.Equals, true)
	c.Assert(s.Config.HasProp("string:test9"), check.Equals, false)
	c.Assert(s.Config.HasProp("strings:test9"), check.Equals, false)
}

func (s *KNFSuite) TestFormating(c *check.C) {
	c.Assert(s.Config.GetI("formating:test1"), check.Equals, 1)
	c.Assert(s.Config.GetI("formating:test2"), check.Equals, 2)
	c.Assert(s.Config.GetI("formating:test3"), check.Equals, 3)
}

func (s *KNFSuite) TestStrings(c *check.C) {
	c.Assert(s.Config.GetS("string:test1"), check.Equals, "test")
	c.Assert(s.Config.GetS("string:test2"), check.Equals, "true")
	c.Assert(s.Config.GetS("string:test3"), check.Equals, "4500")
	c.Assert(s.Config.GetS("string:test4"), check.Equals, "!$%^&")
	c.Assert(s.Config.GetS("string:test5"), check.Equals, "long long long long text for test")
}

func (s *KNFSuite) TestBoolean(c *check.C) {
	c.Assert(s.Config.GetB("boolean:test1"), check.Equals, true)
	c.Assert(s.Config.GetB("boolean:test2"), check.Equals, false)
	c.Assert(s.Config.GetB("boolean:test3"), check.Equals, false)
	c.Assert(s.Config.GetB("boolean:test4"), check.Equals, true)
	c.Assert(s.Config.GetB("boolean:test5"), check.Equals, false)
	c.Assert(s.Config.GetB("boolean:test6"), check.Equals, true)
}

func (s *KNFSuite) TestInteger(c *check.C) {
	c.Assert(s.Config.GetI("integer:test1"), check.Equals, 1)
	c.Assert(s.Config.GetI("integer:test2"), check.Equals, -5)
	c.Assert(s.Config.GetI("integer:test3"), check.Equals, 10000000)
	c.Assert(s.Config.GetI("integer:test4"), check.Equals, 0)
	c.Assert(s.Config.GetI("integer:test5"), check.Equals, 0xFF)
	c.Assert(s.Config.GetF("integer:test6"), check.Equals, 123.4)
	c.Assert(s.Config.GetF("integer:test7"), check.Equals, 123.456789)
}

func (s *KNFSuite) TestFileMode(c *check.C) {
	c.Assert(s.Config.GetM("file-mode:test1"), check.Equals, os.FileMode(0644))
	c.Assert(s.Config.GetM("file-mode:test2"), check.Equals, os.FileMode(0644))
	c.Assert(s.Config.GetM("file-mode:test3"), check.Equals, os.FileMode(0))
	c.Assert(s.Config.GetM("file-mode:test4"), check.Equals, os.FileMode(0))
	c.Assert(s.Config.GetM("file-mode:test5"), check.Equals, os.FileMode(0))
}

func (s *KNFSuite) TestComments(c *check.C) {
	c.Assert(s.Config.GetI("comment:test1"), check.Equals, 100)
	c.Assert(s.Config.GetI("comment:test2"), check.Not(check.Equals), 100)
	c.Assert(s.Config.HasProp("comment:test2"), check.Equals, false)
}

func (s *KNFSuite) TestMacro(c *check.C) {
	c.Assert(s.Config.GetS("macro:test1"), check.Equals, "100")
	c.Assert(s.Config.GetI("macro:test1"), check.Equals, 100)
	c.Assert(s.Config.GetS("macro:test2"), check.Equals, "100.50")
	c.Assert(s.Config.GetS("macro:test3"), check.Equals, "Value is 100.50")
	c.Assert(s.Config.GetS("macro:test4"), check.Equals, "\"Value is 100.50\"")
	c.Assert(s.Config.GetS("macro:test5"), check.Equals, "{ABC}")
	c.Assert(s.Config.GetS("macro:test6"), check.Equals, "{}")
}

func (s *KNFSuite) TestNil(c *check.C) {
	var nilConf *Config

	c.Assert(nilConf.GetS("formating:test1"), check.Equals, "")
	c.Assert(nilConf.GetI("formating:test1"), check.Equals, 0)
	c.Assert(nilConf.GetF("formating:test1"), check.Equals, 0.0)
	c.Assert(nilConf.GetB("formating:test1"), check.Equals, false)
	c.Assert(nilConf.HasSection("formating"), check.Equals, false)
	c.Assert(nilConf.HasProp("formating:test1"), check.Equals, false)
	c.Assert(len(nilConf.Sections()), check.Equals, 0)
}

func (s *KNFSuite) TestDefault(c *check.C) {
	c.Assert(s.Config.GetS("string:test100", "fail"), check.Equals, "fail")
	c.Assert(s.Config.GetB("boolean:test100", true), check.Equals, true)
	c.Assert(s.Config.GetI("integer:test100", 9999), check.Equals, 9999)
	c.Assert(s.Config.GetF("integer:test100", 123.45), check.Equals, 123.45)
	c.Assert(s.Config.GetM("file-mode:test100", 0755), check.Equals, os.FileMode(0755))
	c.Assert(s.Config.GetS("string:test6", "fail"), check.Equals, "fail")
}

func (s *KNFSuite) TestValidation(c *check.C) {
	var e []error

	e = s.Config.Validate([]*Validator{
		&Validator{"integer:test1", Empty, nil},
		&Validator{"integer:test1", Less, 0},
		&Validator{"integer:test1", Less, 0.5},
		&Validator{"integer:test1", Greater, 10},
		&Validator{"integer:test1", Greater, 10.1},
		&Validator{"integer:test1", Equals, 10},
		&Validator{"integer:test1", Equals, 10.1},
		&Validator{"integer:test1", Equals, "123"},
	})

	c.Assert(len(e), check.Equals, 0)

	e = s.Config.Validate([]*Validator{
		&Validator{"boolean:test5", Empty, nil},
		&Validator{"integer:test1", Less, 10},
		&Validator{"integer:test1", Greater, 0},
		&Validator{"integer:test1", Equals, 1},
		&Validator{"integer:test1", Greater, "12345"},
	})

	c.Assert(len(e), check.Equals, 5)

	c.Assert(e[0].Error(), check.Equals, "Property boolean:test5 can't be empty")
	c.Assert(e[1].Error(), check.Equals, "Property integer:test1 can't be less than 10")
	c.Assert(e[2].Error(), check.Equals, "Property integer:test1 can't be greater than 0")
	c.Assert(e[3].Error(), check.Equals, "Property integer:test1 can't be equal 1")
	c.Assert(e[4].Error(), check.Equals, "Wrong validator for property integer:test1")
}
