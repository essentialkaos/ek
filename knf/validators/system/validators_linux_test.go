package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/knf"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ValidatorSuite) TestInterfaceValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"interface:test0", Interface, nil},
		{"interface:test1", Interface, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"interface:test2", Interface, nil},
	})

	c.Assert(errs, HasLen, 1)
	c.Assert(errs[0].Error(), Equals, `Interface "abc" is not present on the system`)
}
