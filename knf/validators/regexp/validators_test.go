package regexp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	"github.com/essentialkaos/ek/v12/knf"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
[regexp]
  test1: TEST1
  test2: test
  test3:
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

func (s *ValidatorSuite) TestRegexpValidator(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"regexp:test1", Regexp, `^[A-Z0-9]{4,5}$`},
		{"regexp:test2", Regexp, ``},
		{"regexp:test3", Regexp, `^[A-Z0-9]{4,5}$`},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"regexp:test2", Regexp, `^[A-Z0-9]{4}$`},
		{"regexp:test2", Regexp, `\`},
	})

	c.Assert(errs, check.HasLen, 2)
	c.Assert(errs[0].Error(), check.Equals, "Property regexp:test2 must match regexp pattern ^[A-Z0-9]{4}$")
	c.Assert(errs[1].Error(), check.Equals, "Can't use given regexp pattern: error parsing regexp: trailing backslash at end of expression: ``")
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
