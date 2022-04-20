package network

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	"github.com/essentialkaos/ek/v12/knf"

	. "github.com/essentialkaos/check"
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
	c.Assert(errs[0].Error(), Equals, "300.0.400.5 is not a valid IP address")
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
	c.Assert(errs[0].Error(), Equals, "ABCD is not a valid port number")
	c.Assert(errs[1].Error(), Equals, "78361 is not a valid port number")
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
	c.Assert(errs[0].Error(), Equals, "ABCD is not a valid MAC address: address ABCD: invalid MAC address")
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
	c.Assert(errs[0].Error(), Equals, "127.0.0.1/200 is not a valid CIDR address: invalid CIDR address: 127.0.0.1/200")
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
	c.Assert(errs[0].Error(), Equals, "google.com/abcd.php is not a valid URL address: parse \"google.com/abcd.php\": invalid URI for request")
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
