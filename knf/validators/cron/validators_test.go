package cron

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
[cron]
  test1:
  test2: 5 0 * 8 *
  test3: 0 0,12 1 */2 sun
  test4: @daily
  test5: 0 0,456 1 */2 sun
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

func (s *ValidatorSuite) TestCronValidator(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"cron:test1", Expression, nil},
		{"cron:test2", Expression, nil},
		{"cron:test3", Expression, nil},
		{"cron:test4", Expression, nil},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"cron:test5", Expression, nil},
	})

	c.Assert(errs, check.HasLen, 1)

	c.Assert(errs[0].Error(), check.Equals, `Property cron:test5 contains ivalid cron expression: Can't parse token "0,456": strconv.ParseUint: parsing "456": value out of range`)
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
