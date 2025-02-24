package time

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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
[time]
  test1:
  test2: %D
  test3: %Y/%m/%d %H:%M:%S %Z
  test4: %H:%M:%i
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

	errs := knf.Validate([]*knf.Validator{
		{"time:test1", Format, nil},
		{"time:test2", Format, nil},
		{"time:test3", Format, nil},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"time:test4", Format, nil},
	})

	c.Assert(errs, check.HasLen, 1)

	c.Assert(errs[0].Error(), check.Equals, `Property time:test4 contains invalid time format: Invalid control sequence "%i"`)
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
