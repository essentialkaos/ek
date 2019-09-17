package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"pkg.re/essentialkaos/ek.v11/knf"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_TEMPLATE = `
[test]
	test0:
	test1: %s
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

func (s *ValidatorSuite) TestPermsValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"test:test0", Perms, "FR"},
		{"test:test1", Perms, "FR"},
	})

	c.Assert(errs, HasLen, 0)

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", Perms, "F"},
		{"test:test1", Perms, "FR"},
		{"test:test1", Perms, "FW"},
		{"test:test1", Perms, "FX"},
		{"test:test1", Perms, "FRW"},
		{"test:test1", Perms, "DX"},
		{"test:test1", Perms, "DRX"},
		{"test:test1", Perms, "DRX"},
		{"test:test1", Perms, "DWX"},
		{"test:test1", Perms, "DRWX"},
		{"test:test1", Perms, "WX"},
	})

	c.Assert(errs, HasLen, 11)
}

func (s *ValidatorSuite) TestOwnerValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"test:test0", Owner, "root"},
		{"test:test1", Owner, "root"},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", Owner, "ftp"},
		{"test:test1", Owner, "somerandomuser"},
	})

	c.Assert(errs, HasLen, 2)

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", Owner, "root"},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestOwnerGroupValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"test:test0", OwnerGroup, "root"},
		{"test:test1", OwnerGroup, "root"},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", OwnerGroup, "ftp"},
		{"test:test1", OwnerGroup, "somerandomuser"},
	})

	c.Assert(errs, HasLen, 2)

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", OwnerGroup, "root"},
	})

	c.Assert(errs, HasLen, 1)
}

func (s *ValidatorSuite) TestFileModeValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate([]*knf.Validator{
		{"test:test0", FileMode, os.FileMode(0644)},
		{"test:test1", FileMode, os.FileMode(0644)},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", FileMode, os.FileMode(0777)},
	})

	c.Assert(errs, HasLen, 1)

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate([]*knf.Validator{
		{"test:test1", FileMode, os.FileMode(0644)},
	})

	c.Assert(errs, HasLen, 1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *C, data string) string {
	configPath := c.MkDir() + "/config.knf"
	configData := fmt.Sprintf(_CONFIG_TEMPLATE, data)

	err := ioutil.WriteFile(configPath, []byte(configData), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
