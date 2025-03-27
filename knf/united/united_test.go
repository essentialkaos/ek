package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"
	"time"

	"github.com/essentialkaos/ek/v13/knf"
	"github.com/essentialkaos/ek/v13/options"

	knfv "github.com/essentialkaos/ek/v13/knf/validators"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
[test]
  string: Test
  integer: 123
  float: 234.5
  boolean: true
  file-mode: 0644
  duration: 24
  size: 3mb
  time-duration: 5m
  timestamp: 1709629048
  timezone: Europe/Zurich
  list: Test1, Test2
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type UnitedSuite struct {
	config *knf.Config
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UnitedSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UnitedSuite) SetUpSuite(c *C) {
	configFile := c.MkDir() + "/config.knf"
	err := os.WriteFile(configFile, []byte(_CONFIG_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	config, err := knf.Read(configFile)

	if err != nil {
		c.Fatal(err.Error())
	}

	s.config = config
}

func (s *UnitedSuite) TestKNFOnly(c *C) {
	err := Combine(nil)

	c.Assert(err, Equals, knf.ErrNilConfig)

	global = nil

	c.Assert(GetS("test:string"), Equals, "")
	c.Assert(GetS("test:string", "TEST"), Equals, "TEST")
	c.Assert(GetI("test:integer"), Equals, 0)
	c.Assert(GetI("test:integer", 1234), Equals, 1234)
	c.Assert(GetI64("test:integer"), Equals, int64(0))
	c.Assert(GetI64("test:integer", 1234), Equals, int64(1234))
	c.Assert(GetU("test:integer"), Equals, uint(0))
	c.Assert(GetU("test:integer", 1234), Equals, uint(1234))
	c.Assert(GetU64("test:integer"), Equals, uint64(0))
	c.Assert(GetU64("test:integer", 1234), Equals, uint64(1234))
	c.Assert(GetF("test:float"), Equals, 0.0)
	c.Assert(GetF("test:float", 1234.5), Equals, 1234.5)
	c.Assert(GetB("test:boolean"), Equals, false)
	c.Assert(GetB("test:boolean", true), Equals, true)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0))
	c.Assert(GetM("test:file-mode", 0600), Equals, os.FileMode(0600))
	c.Assert(GetD("test:duration", MINUTE), Equals, time.Duration(0))
	c.Assert(GetD("test:duration", MINUTE, time.Minute), Equals, time.Minute)
	c.Assert(GetSZ("test:size"), Equals, uint64(0))
	c.Assert(GetSZ("test:size", 1024), Equals, uint64(1024))
	c.Assert(GetTD("test:time-duration"), Equals, time.Duration(0))
	c.Assert(GetTD("test:time-duration", time.Minute), Equals, time.Minute)
	c.Assert(GetTS("test:timestamp").IsZero(), Equals, true)
	c.Assert(GetTS("test:timestamp", time.Unix(1704067200, 0)).IsZero(), Equals, false)
	c.Assert(GetTZ("test:timezone"), IsNil)
	c.Assert(GetTZ("test:timezone", time.Local), NotNil)
	c.Assert(GetL("test:list"), IsNil)
	c.Assert(GetL("test:list", []string{"test"}), NotNil)
	c.Assert(Is("test:string", ""), Equals, false)
	c.Assert(Has("test:string"), Equals, false)
	c.Assert(GetMapping("test:string").Property, Equals, "")

	err = Combine(s.config,
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:size", "test-size", "TEST_SIZE"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	c.Assert(err, IsNil)
	c.Assert(GetS("test:string"), Equals, "Test")
	c.Assert(GetI("test:integer"), Equals, 123)
	c.Assert(GetI64("test:integer"), Equals, int64(123))
	c.Assert(GetU("test:integer"), Equals, uint(123))
	c.Assert(GetU64("test:integer"), Equals, uint64(123))
	c.Assert(GetF("test:float"), Equals, 234.5)
	c.Assert(GetB("test:boolean"), Equals, true)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0644))
	c.Assert(GetD("test:duration", MINUTE), Equals, 24*time.Minute)
	c.Assert(GetSZ("test:size"), Equals, uint64(3*1024*1024))
	c.Assert(GetTD("test:time-duration"), Equals, 5*time.Minute)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1709629048))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Zurich")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1", "Test2"})
	c.Assert(GetS("test:unknown"), Equals, "")
	c.Assert(GetS("test:unknown", "TestABCD"), Equals, "TestABCD")

	c.Assert(Has("test:string"), Equals, true)
	c.Assert(Has("test:abcd"), Equals, false)

	c.Assert(Is("test:string", nil), Equals, false)
	c.Assert(Is("test:string", "Test"), Equals, true)
	c.Assert(Is("test:integer", 123), Equals, true)
	c.Assert(Is("test:integer", int64(123)), Equals, true)
	c.Assert(Is("test:integer", uint(123)), Equals, true)
	c.Assert(Is("test:integer", uint64(123)), Equals, true)
	c.Assert(Is("test:float", 234.5), Equals, true)
	c.Assert(Is("test:file-mode", os.FileMode(0644)), Equals, true)
	c.Assert(Is("test:duration", 24*time.Second), Equals, true)
	c.Assert(Is("test:timestamp", time.Unix(1709629048, 0)), Equals, true)
	c.Assert(Is("test:list", []string{"Test1", "Test2"}), Equals, true)
	c.Assert(Is("test:boolean", true), Equals, true)

	l, _ := time.LoadLocation("Europe/Zurich")
	c.Assert(Is("test:timezone", l), Equals, true)

	mapping := GetMapping("test:string")

	c.Assert(mapping.Option, Equals, "test-string")
	c.Assert(mapping.Variable, Equals, "TEST_STRING")
}

func (s *UnitedSuite) TestWithOptions(c *C) {
	opts := options.NewOptions()

	_, errs := opts.Parse(
		[]string{
			"--test-string", "TestOpt",
			"--test-integer", "456",
			"--test-float", "567.8",
			"--test-boolean", "false",
			"--test-file-mode", "0640",
			"--test-duration", "35",
			"--test-size", "5MB",
			"--test-time-duration", "3h",
			"--test-timestamp", "1704067200",
			"--test-timezone", "Europe/Prague",
			"--test-list", "Test1Opt,Test2Opt",
		}, options.Map{
			"test-string":        {},
			"test-integer":       {Type: options.INT},
			"test-float":         {Type: options.FLOAT},
			"test-boolean":       {Type: options.MIXED},
			"test-file-mode":     {},
			"test-duration":      {Type: options.INT},
			"test-size":          {},
			"test-time-duration": {},
			"test-timestamp":     {},
			"test-timezone":      {},
			"test-list":          {},
		},
	)

	c.Assert(errs, HasLen, 0)

	err := Combine(s.config,
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:size", "test-size", "TEST_SIZE"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	optionHasFunc = opts.Has
	optionGetFunc = opts.GetS

	c.Assert(err, IsNil)
	c.Assert(GetS("test:string"), Equals, "TestOpt")
	c.Assert(GetI("test:integer"), Equals, 456)
	c.Assert(GetI64("test:integer"), Equals, int64(456))
	c.Assert(GetU("test:integer"), Equals, uint(456))
	c.Assert(GetU64("test:integer"), Equals, uint64(456))
	c.Assert(GetF("test:float"), Equals, 567.8)
	c.Assert(GetB("test:boolean"), Equals, false)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0640))
	c.Assert(GetD("test:duration", HOUR), Equals, 35*time.Hour)
	c.Assert(GetSZ("test:size"), Equals, uint64(5*1024*1024))
	c.Assert(GetTD("test:time-duration"), Equals, 3*time.Hour)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1704067200))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Prague")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1Opt", "Test2Opt"})

	optionHasFunc = options.Has
	optionGetFunc = options.GetS
}

func (s *UnitedSuite) TestWithEnv(c *C) {
	os.Setenv("TEST_STRING", "TestEnv")
	os.Setenv("TEST_INTEGER", "789")
	os.Setenv("TEST_FLOAT", "678.9")
	os.Setenv("TEST_BOOLEAN", "no")
	os.Setenv("TEST_FILE_MODE", "0600")
	os.Setenv("TEST_DURATION", "17")
	os.Setenv("TEST_SIZE", "30kB")
	os.Setenv("TEST_TIME_DURATION", "19m")
	os.Setenv("TEST_TIMESTAMP", "1591014600")
	os.Setenv("TEST_TIMEZONE", "Europe/Berlin")
	os.Setenv("TEST_LIST", "Test1Env,Test2Env")

	err := Combine(s.config,
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:size", "test-size", "TEST_SIZE"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	c.Assert(err, IsNil)
	c.Assert(GetS("test:string"), Equals, "TestEnv")
	c.Assert(GetI("test:integer"), Equals, 789)
	c.Assert(GetI64("test:integer"), Equals, int64(789))
	c.Assert(GetU("test:integer"), Equals, uint(789))
	c.Assert(GetU64("test:integer"), Equals, uint64(789))
	c.Assert(GetF("test:float"), Equals, 678.9)
	c.Assert(GetB("test:boolean"), Equals, false)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0600))
	c.Assert(GetD("test:duration", HOUR), Equals, 17*time.Hour)
	c.Assert(GetSZ("test:size"), Equals, uint64(30*1024))
	c.Assert(GetTD("test:time-duration"), Equals, 19*time.Minute)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1591014600))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Berlin")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1Env", "Test2Env"})
}

func (s *UnitedSuite) TestCombineSimple(c *C) {
	err := CombineSimple(nil, "test:string")
	c.Assert(err, NotNil)

	err = CombineSimple(s.config, "test:string")
	c.Assert(err, IsNil)
}

func (s *UnitedSuite) TestValidation(c *C) {
	global = nil

	errs := Validate([]*knf.Validator{
		{"test:string", knfv.Set, nil},
	})

	c.Assert(errs, HasLen, 1)

	err := Combine(s.config,
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:size", "test-size", "TEST_SIZE"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	c.Assert(err, IsNil)

	errs = Validate([]*knf.Validator{
		{"test:string", knfv.Set, nil},
		{"test:integer", knfv.Less, 100},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *UnitedSuite) TestHelpers(c *C) {
	c.Assert(O(""), Equals, "")
	c.Assert(O("section:prop-one"), Equals, "section-prop-one")
	c.Assert(O("Section:Prop-Two"), Equals, "section-prop-two")

	c.Assert(E(""), Equals, "")
	c.Assert(E("section:prop-one"), Equals, "SECTION_PROP_ONE")
	c.Assert(E("Section:Prop-Two"), Equals, "SECTION_PROP_TWO")

	m := Simple("test:option-one")

	c.Assert(m.Property, Equals, "test:option-one")
	c.Assert(m.Option, Equals, "test-option-one")
	c.Assert(m.Variable, Equals, "TEST_OPTION_ONE")

	var op options.Map

	c.Assert(func() { AddOptions(op, "test") }, NotPanics)

	op = options.Map{}

	AddOptions(op, "test")

	c.Assert(op, HasLen, 1)
}
