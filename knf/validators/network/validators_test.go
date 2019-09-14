package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	"pkg.re/essentialkaos/ek.v11/knf"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
[ip]
	test0:
	test1: 127.0.0.1
	test2: 300.0.400.5

[port]
	test0:
	test1: 1045
	test2: ABCD
	test3: 78361

[mac]
	test0:
	test1: 00:00:5e:00:53:01
	test2: ABCD
	
[cidr]
	test0:
	test1: 192.0.2.1/24
	test2: 127.0.0.1/200

[url]
	test0:
	test1: https://google.com
	test2: google.com/abcd.php
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

func (s *ValidatorSuite) TestIPValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"ip:test0", IP, nil},
		{"ip:test1", IP, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"ip:test2", IP, nil},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestPortValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"port:test0", Port, nil},
		{"port:test1", Port, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"port:test2", Port, nil},
		{"port:test3", Port, nil},
	})

	c.Assert(errs, HasLen, 2)
}

func (s *ValidatorSuite) TestMACValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"mac:test0", MAC, nil},
		{"mac:test1", MAC, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"mac:test2", MAC, nil},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestCIDRValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"cidr:test0", CIDR, nil},
		{"cidr:test1", CIDR, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"cidr:test2", CIDR, nil},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestURLValidator(c *C) {
	configFile := createConfig(c, _CONFIG_DATA)

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"url:test0", URL, nil},
		{"url:test1", URL, nil},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"url:test2", URL, nil},
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
