package time

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	"github.com/essentialkaos/ek/v13/knf"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
[format]
  test1:
  test2: %D
  test3: %Y/%m/%d %H:%M:%S %Z
  test4: %H:%M:%i

[timezone]
	test1:
	test2: Local
	test3: UTC
	test4: Etc/GMT+12
	test5: Europe/Dublin
	test6: Europe/Guernsey
	test7: Europe/Hogwarts
`

// ////////////////////////////////////////////////////////////////////////////////// //

type ValidatorSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&ValidatorSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) {
	check.TestingT(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ValidatorSuite) TestFormatValidator(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	errs := knf.Validate(knf.Validators{
		{"format:test1", Format, nil},
		{"format:test2", Format, nil},
		{"format:test3", Format, nil},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"format:test4", Format, nil},
	})

	c.Assert(errs, check.HasLen, 1)

	c.Assert(errs[0].Error(), check.Equals, `Property format:test4 contains invalid time format: Invalid control sequence "%i"`)
}

func (s *ValidatorSuite) TestTimezoneValidator(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	errs := knf.Validate(knf.Validators{
		{"timezone:test1", Timezone, nil},
		{"timezone:test2", Timezone, nil},
		{"timezone:test3", Timezone, nil},
		{"timezone:test4", Timezone, nil},
		{"timezone:test5", Timezone, nil},
		{"timezone:test6", Timezone, nil},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"timezone:test7", Timezone, nil},
	})

	c.Assert(errs, check.HasLen, 1)

	c.Assert(errs[0].Error(), check.Equals, `Property timezone:test7 contains invalid time zone name: unknown time zone Europe/Hogwarts`)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *check.C, data string) string {
	configPath := c.MkDir() + "/config.knf"

	err := os.WriteFile(configPath, []byte(data), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
