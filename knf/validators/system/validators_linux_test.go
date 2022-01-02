//go:build linux
// +build linux

package system

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"pkg.re/essentialkaos/ek.v12/knf"

	. "pkg.re/essentialkaos/check.v1"
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
}
