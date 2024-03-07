package united

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"
	"time"

	"github.com/essentialkaos/ek/v12/knf"
	"github.com/essentialkaos/ek/v12/options"

	knfv "github.com/essentialkaos/ek/v12/knf/validators"

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
  time-duration: 5m
  timestamp: 1709629048
  timezone: Europe/Zurich
  list: Test1, Test2
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type UnitedSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&UnitedSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *UnitedSuite) SetUpSuite(c *C) {
	configFile := c.MkDir() + "/config.knf"
	err := os.WriteFile(configFile, []byte(_CONFIG_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	err = knf.Global(configFile)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *UnitedSuite) TestKNFOnly(c *C) {
	global = nil

	c.Assert(GetS("test:string"), Equals, "Test")
	c.Assert(GetI("test:integer"), Equals, 123)
	c.Assert(GetI64("test:integer"), Equals, int64(123))
	c.Assert(GetU("test:integer"), Equals, uint(123))
	c.Assert(GetU64("test:integer"), Equals, uint64(123))
	c.Assert(GetF("test:float"), Equals, 234.5)
	c.Assert(GetB("test:boolean"), Equals, true)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0644))
	c.Assert(GetD("test:duration", knf.Minute), Equals, 24*time.Minute)
	c.Assert(GetTD("test:time-duration"), Equals, 5*time.Minute)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1709629048))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Zurich")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1", "Test2"})

	err := Combine(
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
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
	c.Assert(GetD("test:duration", knf.Minute), Equals, 24*time.Minute)
	c.Assert(GetTD("test:time-duration"), Equals, 5*time.Minute)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1709629048))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Zurich")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1", "Test2"})
	c.Assert(GetS("test:unknown"), Equals, "")
	c.Assert(GetS("test:unknown", "TestABCD"), Equals, "TestABCD")
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
			"test-time-duration": {},
			"test-timestamp":     {},
			"test-timezone":      {},
			"test-list":          {},
		},
	)

	c.Assert(errs, HasLen, 0)

	err := Combine(
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	c.Assert(err, IsNil)

	optionHasFunc = opts.Has
	optionGetFunc = opts.GetS

	c.Assert(GetS("test:string"), Equals, "TestOpt")
	c.Assert(GetI("test:integer"), Equals, 456)
	c.Assert(GetI64("test:integer"), Equals, int64(456))
	c.Assert(GetU("test:integer"), Equals, uint(456))
	c.Assert(GetU64("test:integer"), Equals, uint64(456))
	c.Assert(GetF("test:float"), Equals, 567.8)
	c.Assert(GetB("test:boolean"), Equals, false)
	c.Assert(GetM("test:file-mode"), Equals, os.FileMode(0640))
	c.Assert(GetD("test:duration", knf.Hour), Equals, 35*time.Hour)
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
	os.Setenv("TEST_TIME_DURATION", "19m")
	os.Setenv("TEST_TIMESTAMP", "1591014600")
	os.Setenv("TEST_TIMEZONE", "Europe/Berlin")
	os.Setenv("TEST_LIST", "Test1Env,Test2Env")

	err := Combine(
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
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
	c.Assert(GetD("test:duration", knf.Hour), Equals, 17*time.Hour)
	c.Assert(GetTD("test:time-duration"), Equals, 19*time.Minute)
	c.Assert(GetTS("test:timestamp").Unix(), Equals, int64(1591014600))
	c.Assert(GetTZ("test:timezone").String(), Equals, "Europe/Berlin")
	c.Assert(GetL("test:list"), DeepEquals, []string{"Test1Env", "Test2Env"})
}

func (s *UnitedSuite) TestValidation(c *C) {
	global = nil

	errs := Validate([]*knf.Validator{
		{"test:string", knfv.Empty, nil},
	})

	c.Assert(errs, HasLen, 1)

	err := Combine(
		Mapping{"test:string", "test-string", "TEST_STRING"},
		Mapping{"test:integer", "test-integer", "TEST_INTEGER"},
		Mapping{"test:float", "test-float", "TEST_FLOAT"},
		Mapping{"test:boolean", "test-boolean", "TEST_BOOLEAN"},
		Mapping{"test:file-mode", "test-file-mode", "TEST_FILE_MODE"},
		Mapping{"test:duration", "test-duration", "TEST_DURATION"},
		Mapping{"test:time-duration", "test-time-duration", "TEST_TIME_DURATION"},
		Mapping{"test:timestamp", "test-timestamp", "TEST_TIMESTAMP"},
		Mapping{"test:timezone", "test-timezone", "TEST_TIMEZONE"},
		Mapping{"test:list", "test-list", "TEST_LIST"},
	)

	c.Assert(err, IsNil)

	errs = Validate([]*knf.Validator{
		{"test:string", knfv.Empty, nil},
		{"test:integer", knfv.Greater, 100},
	})

	c.Assert(errs, HasLen, 1)
}
