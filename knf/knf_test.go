package knf

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

	"github.com/essentialkaos/ek/v13/errors"

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

[time-duration]
	test1: 0
	test2: 15
	test3: 45s
	test4: 12M
	test5: ABCD

[timestamp]
	test1: 0
	test2: 1709629048
	test3: ABCD

[timezone]
	test1: Europe/Zurich
	test2: ABCD

[list]
	test1:
	test2: Test1, Test2

[size]
	test1: 100
	test2: 1mb
	test3: 2.3 gb

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

const _MERGE_DATA = `
[string]
  test1: test-new

[boolean]
  test1: false

[extra]
	test1: extra-data
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

	c.Assert(GetS("test:test"), check.Equals, "")
	c.Assert(GetI("test:test"), check.Equals, 0)
	c.Assert(GetI("test:test"), check.Equals, 0)
	c.Assert(GetU("test:test"), check.Equals, uint(0))
	c.Assert(GetI64("test:test"), check.Equals, int64(0))
	c.Assert(GetU64("test:test"), check.Equals, uint64(0))
	c.Assert(GetF("test:test"), check.Equals, 0.0)
	c.Assert(GetB("test:test"), check.Equals, false)
	c.Assert(GetM("test:test"), check.Equals, os.FileMode(0))
	c.Assert(GetD("test:test", SECOND), check.Equals, time.Duration(0))
	c.Assert(GetSZ("test:test"), check.Equals, uint64(0))
	c.Assert(GetTD("test:test"), check.Equals, time.Duration(0))
	c.Assert(GetTS("test:test").IsZero(), check.Equals, true)
	c.Assert(GetTZ("test:test"), check.IsNil)
	c.Assert(GetL("test:test"), check.IsNil)
	c.Assert(Is("test:test", ""), check.Equals, false)
	c.Assert(HasSection("test"), check.Equals, false)
	c.Assert(Has("test:test"), check.Equals, false)
	c.Assert(Sections(), check.HasLen, 0)
	c.Assert(Props("test"), check.HasLen, 0)
	c.Assert(Validate(Validators{}), check.DeepEquals, errors.Errors{ErrNilConfig})
	c.Assert(Alias("test:test", "test:test"), check.NotNil)
	c.Assert(global.Merge(nil), check.NotNil)

	config := &Config{mx: &sync.RWMutex{}}

	c.Assert(config.GetS("test:test"), check.Equals, "")
	c.Assert(config.GetI("test:test"), check.Equals, 0)
	c.Assert(config.GetF("test:test"), check.Equals, 0.0)
	c.Assert(config.GetB("test:test"), check.Equals, false)
	c.Assert(config.GetM("test:test"), check.Equals, os.FileMode(0))
	c.Assert(config.GetD("test:test", SECOND), check.Equals, time.Duration(0))
	c.Assert(config.GetSZ("test:test"), check.Equals, uint64(0))
	c.Assert(config.GetTD("test:test"), check.Equals, time.Duration(0))
	c.Assert(config.GetTS("test:test").IsZero(), check.Equals, true)
	c.Assert(config.GetTZ("test:test"), check.IsNil)
	c.Assert(config.GetL("test:test"), check.IsNil)
	c.Assert(config.Is("test:test", ""), check.Equals, true)
	c.Assert(config.HasSection("test"), check.Equals, false)
	c.Assert(config.Has("test:test"), check.Equals, false)
	c.Assert(config.Sections(), check.HasLen, 0)
	c.Assert(config.Props("test"), check.HasLen, 0)
	c.Assert(config.Validate(Validators{}), check.HasLen, 0)
	c.Assert(config.Merge(nil), check.NotNil)

	c.Assert(config.Alias("", ""), check.NotNil)
	c.Assert(config.Alias("test", ""), check.NotNil)
	c.Assert(config.Alias("test", "test"), check.NotNil)
	c.Assert(config.Alias("test:test", "test"), check.NotNil)

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

func (s *KNFSuite) TestMerging(c *check.C) {
	c1, err := Parse([]byte(_CONFIG_DATA))

	c.Assert(err, check.IsNil)
	c.Assert(c1, check.NotNil)

	c2, err := Parse([]byte(_MERGE_DATA))

	c.Assert(err, check.IsNil)
	c.Assert(c2, check.NotNil)

	err = c1.Merge(c2)

	c.Assert(err, check.IsNil)

	c.Assert(c1.GetS("string:test1"), check.Equals, "test-new")
	c.Assert(c1.GetB("boolean:test1"), check.Equals, false)
	c.Assert(c1.GetS("extra:test1"), check.Equals, "extra-data")
}

func (s *KNFSuite) TestAlias(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(Alias("string:test1", "string:testX"), check.IsNil)

	c.Assert(GetS("string:testX"), check.Equals, "test")
}

func (s *KNFSuite) TestSections(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	sections := Sections()

	c.Assert(sections, check.HasLen, 14)
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
			"time-duration",
			"timestamp",
			"timezone",
			"list",
			"size",
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

	c.Assert(Has("string:test1"), check.Equals, true)
	c.Assert(Has("string:test6"), check.Equals, false)
	c.Assert(Has("strings:test6"), check.Equals, false)
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

	c.Assert(GetD("duration:test1", SECOND), check.Equals, time.Duration(0))
	c.Assert(GetD("duration:test2", SECOND), check.Equals, time.Minute)
	c.Assert(GetD("duration:test3", SECOND), check.Equals, time.Duration(0))
	c.Assert(GetD("duration:test4", SECOND), check.Equals, time.Duration(0))
}

func (s *KNFSuite) TestTimeDuration(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetTD("time-duration:test1"), check.Equals, time.Duration(0))
	c.Assert(GetTD("time-duration:test2"), check.Equals, 15*time.Second)
	c.Assert(GetTD("time-duration:test3"), check.Equals, 45*time.Second)
	c.Assert(GetTD("time-duration:test4"), check.Equals, 12*time.Minute)
	c.Assert(GetTD("time-duration:test5"), check.Equals, time.Duration(0))
}

func (s *KNFSuite) TestTimestamp(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetTS("timestamp:test1"), check.DeepEquals, time.Unix(0, 0))
	c.Assert(GetTS("timestamp:test2"), check.DeepEquals, time.Unix(1709629048, 0))
	c.Assert(GetTS("timestamp:test3"), check.DeepEquals, time.Time{})
}

func (s *KNFSuite) TestTimezone(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	l, _ := time.LoadLocation("Europe/Zurich")

	c.Assert(GetTZ("timezone:test1"), check.DeepEquals, l)
	c.Assert(GetTZ("timezone:test2"), check.IsNil)
}

func (s *KNFSuite) TestList(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetL("list:test1"), check.HasLen, 0)
	c.Assert(GetL("list:test2"), check.HasLen, 2)
}

func (s *KNFSuite) TestSize(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetSZ("size:test1"), check.Equals, uint64(100))
	c.Assert(GetSZ("size:test2"), check.Equals, uint64(1024*1024))
	c.Assert(GetSZ("size:test3"), check.Equals, uint64(2469606195))
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
	c.Assert(Is("timestamp:test2", time.Unix(1709629048, 0)), check.Equals, true)
	c.Assert(Is("list:test2", []string{"Test1", "Test2"}), check.Equals, true)

	l, _ := time.LoadLocation("Europe/Zurich")
	c.Assert(Is("timezone:test1", l), check.Equals, true)

	c.Assert(Is("integer:test1", nil), check.Equals, false)
}

func (s *KNFSuite) TestComments(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	c.Assert(GetI("comment:test1"), check.Equals, 100)
	c.Assert(GetI("comment:test2"), check.Not(check.Equals), 100)
	c.Assert(Has("comment:test2"), check.Equals, false)
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

	c.Assert(nilConf.getValue("formatting:test1"), check.Equals, "")

	c.Assert(nilConf.GetS("formatting:test1"), check.Equals, "")
	c.Assert(nilConf.GetI("formatting:test1"), check.Equals, 0)
	c.Assert(nilConf.GetI64("formatting:test1"), check.Equals, int64(0))
	c.Assert(nilConf.GetU("formatting:test1"), check.Equals, uint(0))
	c.Assert(nilConf.GetU64("formatting:test1"), check.Equals, uint64(0))
	c.Assert(nilConf.GetF("formatting:test1"), check.Equals, 0.0)
	c.Assert(nilConf.GetB("formatting:test1"), check.Equals, false)
	c.Assert(nilConf.GetM("formatting:test1"), check.Equals, os.FileMode(0))
	c.Assert(nilConf.GetD("formatting:test1", SECOND), check.Equals, time.Duration(0))
	c.Assert(nilConf.GetSZ("size:test100"), check.Equals, uint64(0))
	c.Assert(nilConf.GetTD("formatting:test1"), check.Equals, time.Duration(0))
	c.Assert(nilConf.GetTS("formatting:test1").IsZero(), check.Equals, true)
	c.Assert(nilConf.GetTZ("formatting:test1"), check.IsNil)
	c.Assert(nilConf.GetL("formatting:test1"), check.IsNil)
	c.Assert(nilConf.Is("formatting:test1", ""), check.Equals, false)
	c.Assert(nilConf.HasSection("formatting"), check.Equals, false)
	c.Assert(nilConf.Has("formatting:test1"), check.Equals, false)
	c.Assert(nilConf.Sections(), check.HasLen, 0)
	c.Assert(nilConf.Props("formatting"), check.HasLen, 0)
	c.Assert(nilConf.File(), check.Equals, "")
	c.Assert(nilConf.Alias("test:test", "test:test"), check.NotNil)

	_, err := nilConf.Reload()

	c.Assert(err, check.NotNil)
	c.Assert(err, check.DeepEquals, ErrNilConfig)

	errs := nilConf.Validate(Validators{})

	c.Assert(errs, check.Not(check.HasLen), 0)
	c.Assert(errs, check.DeepEquals, errors.Errors{ErrNilConfig})
}

func (s *KNFSuite) TestDefault(c *check.C) {
	global = nil

	l, _ := time.LoadLocation("Asia/Yerevan")

	c.Assert(GetS("string:test100", "fail"), check.Equals, "fail")
	c.Assert(GetB("boolean:test100", true), check.Equals, true)
	c.Assert(GetI("integer:test100", 9999), check.Equals, 9999)
	c.Assert(GetU("integer:test100", 9999), check.Equals, uint(9999))
	c.Assert(GetI64("integer:test100", 9999), check.Equals, int64(9999))
	c.Assert(GetU64("integer:test100", 9999), check.Equals, uint64(9999))
	c.Assert(GetF("integer:test100", 123.45), check.Equals, 123.45)
	c.Assert(GetM("file-mode:test100", 0755), check.Equals, os.FileMode(0755))
	c.Assert(GetD("duration:test100", SECOND, time.Minute), check.Equals, time.Minute)
	c.Assert(GetSZ("size:test100", 1024), check.Equals, uint64(1024))
	c.Assert(GetTD("duration:test100", time.Minute), check.Equals, time.Minute)
	c.Assert(GetTS("duration:test100", time.Now()).IsZero(), check.Equals, false)
	c.Assert(GetTZ("duration:test100", l), check.Equals, l)
	c.Assert(GetL("duration:test100", []string{"A", "B"}), check.DeepEquals, []string{"A", "B"})
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
	c.Assert(GetD("duration:test100", SECOND, time.Minute), check.Equals, time.Minute)
	c.Assert(GetSZ("size:test100", 1024), check.Equals, uint64(1024))
	c.Assert(GetTD("duration:test100", time.Minute), check.Equals, time.Minute)
	c.Assert(GetTS("duration:test100", time.Now()).IsZero(), check.Equals, false)
	c.Assert(GetTZ("duration:test100", l), check.Equals, l)
	c.Assert(GetL("duration:test100", []string{"A", "B"}), check.DeepEquals, []string{"A", "B"})
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
	c.Assert(nc.GetD("duration:test100", SECOND, time.Minute), check.Equals, time.Minute)
	c.Assert(nc.GetSZ("size:test100", 1024), check.Equals, uint64(1024))
	c.Assert(nc.GetTD("duration:test100", time.Minute), check.Equals, time.Minute)
	c.Assert(nc.GetTS("duration:test100", time.Now()).IsZero(), check.Equals, false)
	c.Assert(nc.GetTZ("duration:test100", l), check.Equals, l)
	c.Assert(nc.GetL("duration:test100", []string{"A", "B"}), check.DeepEquals, []string{"A", "B"})
	c.Assert(nc.GetS("string:test6", "fail"), check.Equals, "fail")
}

func (s *KNFSuite) TestSimpleValidator(c *check.C) {
	err := Global(s.ConfigPath)

	c.Assert(global, check.NotNil)
	c.Assert(err, check.IsNil)

	var simpleValidator = func(config IConfig, prop string, value any) error {
		if prop == "string:test2" {
			return fmt.Errorf("ERROR")
		}

		return nil
	}

	validators := Validators{
		{"string:test2", simpleValidator, nil},
	}

	validators = validators.AddIf(false, Validators{
		{"string:test2", simpleValidator, nil},
	})

	validators = validators.AddIf(true, Validators{
		{"string:test2", simpleValidator, nil},
	})

	errs := Validate(validators)

	c.Assert(errs, check.HasLen, 2)
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

func (s *KNFSuite) TestHelpers(c *check.C) {
	c.Assert(Q("section", "prop"), check.Equals, "section:prop")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *KNFSuite) BenchmarkBasic(c *check.C) {
	Global(s.ConfigPath)

	for range c.N {
		GetS("string:test1")
	}
}
