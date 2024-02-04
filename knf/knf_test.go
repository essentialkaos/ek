package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
    [formatting]
test1:      1
            test2:2

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
  test7: no

[integer]
  test1: 1
  test2: -5
  test3: 10000000
  test4: A
  test5: 0xFF
  test6: 123.4
  test7: 123.456789
  test8: 0xZZYY
  test9: ABCD

[file-mode]
  test1: 644
  test2: 0644
  test3: 0
  test4: ABC
  test5: true

[duration]
	test1: 0
	test2: 60
	test3: ABC
	test4: true

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

[k]
  t: 1
`

const _CONFIG_MALF_DATA = `
  test1: 123
  test2: 111
`

const (
	_CONFIG_FILE_NAME           = "knf-config-test.conf"
	_CONFIG_EMPTY_FILE_NAME     = "knf-config-test-empty.conf"
	_CONFIG_MALFORMED_FILE_NAME = "knf-config-test-malf.conf"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type KNFSuite struct {
	ConfigPath            string
	EmptyConfigPath       string
	MalformedConfigPath   string
	NonReadableConfigPath string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&KNFSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) {
	check.TestingT(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *KNFSuite) SetUpSuite(c *check.C) {
	tmpdir := c.MkDir()

	s.ConfigPath = tmpdir + "/" + _CONFIG_FILE_NAME
	s.EmptyConfigPath = tmpdir + "/" + _CONFIG_EMPTY_FILE_NAME
	s.MalformedConfigPath = tmpdir + "/" + _CONFIG_MALFORMED_FILE_NAME

	if runtime.GOOS == "darwin" {
		s.NonReadableConfigPath = "/etc/master.passwd"
	} else {
		s.NonReadableConfigPath = "/etc/sudoers"
	}

	err := os.WriteFile(s.ConfigPath, []byte(_CONFIG_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	err = os.WriteFile(s.EmptyConfigPath, []byte(""), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	err = os.WriteFile(s.MalformedConfigPath, []byte(_CONFIG_MALF_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *KNFSuite) TestErrors(c *check.C) {
	global = nil

	err := Global("/_not_exists_")
	c.Assert(err, check.NotNil)
	c.Assert(err, check.ErrorMatches, `open /_not_exists_: no such file or directory`)

	err = Global(s.EmptyConfigPath)
	c.Assert(err, check.NotNil)
	c.Assert(err, check.ErrorMatches, `Configuration file doesn't contain any valid data`)

	err = Global(s.NonReadableConfigPath)
	c.Assert(err, check.NotNil)
	c.Assert(err, check.ErrorMatches, `open .*: permission denied`)

	err = Global(s.MalformedConfigPath)

	c.Assert(err, check.NotNil)
	c.Assert(err.Error(), check.Equals, "Error at line 2: Data defined before section")

	updated, err := Reload()

	c.Assert(err, check.NotNil)
	c.Assert(err, check.DeepEquals, ErrNilConfig)
	c.Assert(updated, check.IsNil)

	c.Assert(GetS("test"), check.Equals, "")
	c.Assert(GetI("test"), check.Equals, 0)
	c.Assert(GetI("test"), check.Equals, 0)
	c.Assert(GetU("test"), check.Equals, uint(0))
	c.Assert(GetI64("test"), check.Equals, int64(0))
	c.Assert(GetU64("test"), check.Equals, uint64(0))
	c.Assert(GetF("test"), check.Equals, 0.0)
	c.Assert(GetB("test"), check.Equals, false)
	c.Assert(GetM("test"), check.Equals, os.FileMode(0))
	c.Assert(GetD("test", Second), check.Equals, time.Duration(0))
	c.Assert(Is("test", ""), check.Equals, false)
	c.Assert(HasSection("test"), check.Equals, false)
	c.Assert(HasProp("test"), check.Equals, false)
	c.Assert(Sections(), check.HasLen, 0)
	c.Assert(Props("test"), check.HasLen, 0)
	c.Assert(Validate([]*Validator{}), check.DeepEquals, []error{ErrNilConfig})

	config := &Config{mx: &sync.RWMutex{}}

	c.Assert(config.GetS("test"), check.Equals, "")
	c.Assert(config.GetI("test"), check.Equals, 0)
	c.Assert(config.GetF("test"), check.Equals, 0.0)
	c.Assert(config.GetB("test"), check.Equals, false)
	c.Assert(config.GetM("test"), check.Equals, os.FileMode(0))
	c.Assert(config.GetD("test", Second), check.Equals, time.Duration(0))
	c.Assert(config.Is("test", ""), check.Equals, true)
	c.Assert(config.HasSection("test"), check.Equals, false)
	c.Assert(config.HasProp("test"), check.Equals, false)
	c.Assert(config.Sections(), check.HasLen, 0)
	c.Assert(config.Props("test"), check.HasLen, 0)
	c.Assert(config.Validate([]*Validator{}), check.HasLen, 0)

	updated, err = config.Reload()

	c.Assert(updated, check.IsNil)
	c.Assert(err, check.NotNil)
	c.Assert(err, check.DeepEquals, ErrCantReload)

	config = &Config{file: "/_not_exists_", mx: &sync.RWMutex{}}

	updated, err = config.Reload()

	c.Assert(updated, check.IsNil)
	c.Assert(err, check.NotNil)

	config, err = Parse([]byte(_CONFIG_DATA))

	c.Assert(err, check.IsNil)
	c.Assert(config, check.NotNil)

	updated, err = config.Reload()

	c.Assert(updated, check.IsNil)
	c.Assert(err, check.NotNil)
}

func (s *KNFSuite) TestParsing(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(err, check.IsNil)
	c.Assert(global.File(), check.Equals, s.ConfigPath)

	_, err = Reload()

	c.Assert(err, check.IsNil)
	c.Assert(global.File(), check.Equals, s.ConfigPath)

	config, err := Parse([]byte(_CONFIG_DATA))

	c.Assert(err, check.IsNil)
	c.Assert(config, check.NotNil)
}

func (s *KNFSuite) TestSections(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	sections := Sections()

	c.Assert(sections, check.HasLen, 9)
	c.Assert(
		sections,
		check.DeepEquals,
		[]string{
			"formatting",
			"string",
			"boolean",
			"integer",
			"file-mode",
			"duration",
			"comment",
			"macro",
			"k",
		},
	)
}

func (s *KNFSuite) TestProps(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	props := Props("file-mode")

	c.Assert(props, check.HasLen, 5)
	c.Assert(
		props,
		check.DeepEquals,
		[]string{"test1", "test2", "test3", "test4", "test5"},
	)
}

func (s *KNFSuite) TestCheckers(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(HasSection("string"), check.Equals, true)
	c.Assert(HasSection("strings"), check.Equals, false)

	c.Assert(HasProp("string:test1"), check.Equals, true)
	c.Assert(HasProp("string:test6"), check.Equals, false)
	c.Assert(HasProp("strings:test6"), check.Equals, false)
}

func (s *KNFSuite) TestFormatting(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetI("formatting:test1"), check.Equals, 1)
	c.Assert(GetI("formatting:test2"), check.Equals, 2)
	c.Assert(GetI("formatting:test3"), check.Equals, 3)
}

func (s *KNFSuite) TestStrings(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetS("string:test1"), check.Equals, "test")
	c.Assert(GetS("string:test2"), check.Equals, "true")
	c.Assert(GetS("string:test3"), check.Equals, "4500")
	c.Assert(GetS("string:test4"), check.Equals, "!$%^&")
	c.Assert(GetS("string:test5"), check.Equals, "long long long long text for test")
}

func (s *KNFSuite) TestBoolean(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetB("boolean:test1"), check.Equals, true)
	c.Assert(GetB("boolean:test2"), check.Equals, false)
	c.Assert(GetB("boolean:test3"), check.Equals, false)
	c.Assert(GetB("boolean:test4"), check.Equals, true)
	c.Assert(GetB("boolean:test5"), check.Equals, false)
	c.Assert(GetB("boolean:test6"), check.Equals, true)
	c.Assert(GetB("boolean:test7"), check.Equals, false)
}

func (s *KNFSuite) TestInteger(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetI("integer:test1"), check.Equals, 1)
	c.Assert(GetI("integer:test2"), check.Equals, -5)
	c.Assert(GetI("integer:test3"), check.Equals, 10000000)
	c.Assert(GetI("integer:test4"), check.Equals, 0)
	c.Assert(GetI("integer:test5"), check.Equals, 0xFF)
	c.Assert(GetF("integer:test6"), check.Equals, 123.4)
	c.Assert(GetF("integer:test7"), check.Equals, 123.456789)
	c.Assert(GetF("integer:test8"), check.Equals, 0.0)
	c.Assert(GetI("integer:test8"), check.Equals, 0)
	c.Assert(GetF("integer:test9"), check.Equals, 0.0)

	c.Assert(GetU("integer:test1"), check.Equals, uint(1))
	c.Assert(GetI64("integer:test1"), check.Equals, int64(1))
	c.Assert(GetU64("integer:test1"), check.Equals, uint64(1))
}

func (s *KNFSuite) TestFileMode(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetM("file-mode:test1"), check.Equals, os.FileMode(0644))
	c.Assert(GetM("file-mode:test2"), check.Equals, os.FileMode(0644))
	c.Assert(GetM("file-mode:test3"), check.Equals, os.FileMode(0))
	c.Assert(GetM("file-mode:test4"), check.Equals, os.FileMode(0))
	c.Assert(GetM("file-mode:test5"), check.Equals, os.FileMode(0))
}

func (s *KNFSuite) TestDuration(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetD("duration:test1", Second), check.Equals, time.Duration(0))
	c.Assert(GetD("duration:test2", Second), check.Equals, time.Minute)
	c.Assert(GetD("duration:test3", Second), check.Equals, time.Duration(0))
	c.Assert(GetD("duration:test4", Second), check.Equals, time.Duration(0))
}

func (s *KNFSuite) TestIs(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(Is("string:test1", "test"), check.Equals, true)
	c.Assert(Is("string:test6", ""), check.Equals, true)
	c.Assert(Is("boolean:test1", true), check.Equals, true)
	c.Assert(Is("integer:test1", 1), check.Equals, true)
	c.Assert(Is("integer:test6", 123.4), check.Equals, true)
	c.Assert(Is("integer:test1", uint(1)), check.Equals, true)
	c.Assert(Is("integer:test1", uint64(1)), check.Equals, true)
	c.Assert(Is("integer:test1", int64(1)), check.Equals, true)
	c.Assert(Is("file-mode:test1", os.FileMode(0644)), check.Equals, true)
	c.Assert(Is("duration:test2", time.Minute), check.Equals, true)

	c.Assert(Is("integer:test1", nil), check.Equals, false)
}

func (s *KNFSuite) TestComments(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetI("comment:test1"), check.Equals, 100)
	c.Assert(GetI("comment:test2"), check.Not(check.Equals), 100)
	c.Assert(HasProp("comment:test2"), check.Equals, false)
}

func (s *KNFSuite) TestMacro(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetS("macro:test1"), check.Equals, "100")
	c.Assert(GetI("macro:test1"), check.Equals, 100)
	c.Assert(GetS("macro:test2"), check.Equals, "100.50")
	c.Assert(GetS("macro:test3"), check.Equals, "Value is 100.50")
	c.Assert(GetS("macro:test4"), check.Equals, "\"Value is 100.50\"")
	c.Assert(GetS("macro:test5"), check.Equals, "{ABC}")
	c.Assert(GetS("macro:test6"), check.Equals, "{}")
}

func (s *KNFSuite) TestNil(c *check.C) {
	var nilConf *Config

	c.Assert(nilConf.GetS("formatting:test1"), check.Equals, "")
	c.Assert(nilConf.GetI("formatting:test1"), check.Equals, 0)
	c.Assert(nilConf.GetF("formatting:test1"), check.Equals, 0.0)
	c.Assert(nilConf.GetB("formatting:test1"), check.Equals, false)
	c.Assert(nilConf.GetM("formatting:test1"), check.Equals, os.FileMode(0))
	c.Assert(nilConf.GetD("formatting:test1", Second), check.Equals, time.Duration(0))
	c.Assert(nilConf.Is("formatting:test1", ""), check.Equals, false)
	c.Assert(nilConf.HasSection("formatting"), check.Equals, false)
	c.Assert(nilConf.HasProp("formatting:test1"), check.Equals, false)
	c.Assert(nilConf.Sections(), check.HasLen, 0)
	c.Assert(nilConf.Props("formatting"), check.HasLen, 0)
	c.Assert(nilConf.File(), check.Equals, "")

	_, err := nilConf.Reload()

	c.Assert(err, check.NotNil)
	c.Assert(err, check.DeepEquals, ErrNilConfig)

	errs := nilConf.Validate([]*Validator{})

	c.Assert(errs, check.Not(check.HasLen), 0)
	c.Assert(errs, check.DeepEquals, []error{ErrNilConfig})
}

func (s *KNFSuite) TestDefault(c *check.C) {
	global = nil

	c.Assert(GetS("string:test100", "fail"), check.Equals, "fail")
	c.Assert(GetB("boolean:test100", true), check.Equals, true)
	c.Assert(GetI("integer:test100", 9999), check.Equals, 9999)
	c.Assert(GetU("integer:test100", 9999), check.Equals, uint(9999))
	c.Assert(GetI64("integer:test100", 9999), check.Equals, int64(9999))
	c.Assert(GetU64("integer:test100", 9999), check.Equals, uint64(9999))
	c.Assert(GetF("integer:test100", 123.45), check.Equals, 123.45)
	c.Assert(GetM("file-mode:test100", 0755), check.Equals, os.FileMode(0755))
	c.Assert(GetD("duration:test100", Second, time.Minute), check.Equals, time.Minute)
	c.Assert(GetS("string:test6", "fail"), check.Equals, "fail")

	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetS("string:test100", "fail"), check.Equals, "fail")
	c.Assert(GetB("boolean:test100", true), check.Equals, true)
	c.Assert(GetI("integer:test100", 9999), check.Equals, 9999)
	c.Assert(GetU("integer:test100", 9999), check.Equals, uint(9999))
	c.Assert(GetI64("integer:test100", 9999), check.Equals, int64(9999))
	c.Assert(GetU64("integer:test100", 9999), check.Equals, uint64(9999))
	c.Assert(GetF("integer:test100", 123.45), check.Equals, 123.45)
	c.Assert(GetM("file-mode:test100", 0755), check.Equals, os.FileMode(0755))
	c.Assert(GetD("duration:test100", Second, time.Minute), check.Equals, time.Minute)
	c.Assert(GetS("string:test6", "fail"), check.Equals, "fail")

	var nc *Config

	c.Assert(nc.GetS("string:test100", "fail"), check.Equals, "fail")
	c.Assert(nc.GetB("boolean:test100", true), check.Equals, true)
	c.Assert(nc.GetI("integer:test100", 9999), check.Equals, 9999)
	c.Assert(nc.GetU("integer:test100", 9999), check.Equals, uint(9999))
	c.Assert(nc.GetI64("integer:test100", 9999), check.Equals, int64(9999))
	c.Assert(nc.GetU64("integer:test100", 9999), check.Equals, uint64(9999))
	c.Assert(nc.GetF("integer:test100", 123.45), check.Equals, 123.45)
	c.Assert(nc.GetM("file-mode:test100", 0755), check.Equals, os.FileMode(0755))
	c.Assert(nc.GetD("duration:test100", Second, time.Minute), check.Equals, time.Minute)
	c.Assert(nc.GetS("string:test6", "fail"), check.Equals, "fail")
}

func (s *KNFSuite) TestSimpleValidator(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	var simpleValidator = func(config *Config, prop string, value any) error {
		if prop == "string:test2" {
			return fmt.Errorf("ERROR")
		}

		return nil
	}

	errs := Validate([]*Validator{
		{"string:test2", simpleValidator, nil},
	})

	c.Assert(errs, check.HasLen, 1)
}

func (s *KNFSuite) TestKNFParserExceptions(c *check.C) {
	r := strings.NewReader(`
		[section]
		ABCD
	`)

	_, err := readData(r)
	c.Assert(err.Error(), check.Equals, `Error at line 3: Property must have ":" as a delimiter`)

	r = strings.NewReader(`
		[section]
		A: 1
		A: 2
	`)

	_, err = readData(r)
	c.Assert(err.Error(), check.Equals, `Error at line 4: Property "A" defined more than once`)

	r = strings.NewReader(`
		[section]
		A: {abcd:test}
	`)

	_, err = readData(r)
	c.Assert(err.Error(), check.Equals, "Error at line 3: Unknown property {abcd:test}")
}
