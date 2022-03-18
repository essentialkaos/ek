package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	"github.com/essentialkaos/ek/knf"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
[user]
  test0:
  test1: root
  test2: somerandomuser

[group]
  test0:
  test1: daemon
  test2: somerandomgroup

[interface]
  test0:
  test1: lo
  test2: abc
`

// ////////////////////////////////////////////////////////////////////////////////// //

type ValidatorSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ValidatorSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) {
	TestingT(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ValidatorSuite) TestUserValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"user:test0", User, nil},
		{"user:test1", User, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"user:test2", User, nil},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestGroupValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"group:test0", Group, nil},
		{"group:test1", Group, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"group:test2", Group, nil},
	})

	c.Assert(errs, HasLen, 1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *C, data string) string {
	configPath := c.MkDir() + "/config.knf"

	err := ioutil.WriteFile(configPath, []byte(data), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
