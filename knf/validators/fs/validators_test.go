package fs

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/essentialkaos/ek/v13/knf"

	. "github.com/essentialkaos/check"
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

	errs := knf.Validate(knf.Validators{
		{"test:test0", Perms, "FR"},
		{"test:test1", Perms, "FR"},
	})

	c.Assert(errs, HasLen, 0)

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate(knf.Validators{
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
		{"test:test1", Perms, ""},
		{"test:test1", Perms, 100},
	})

	c.Assert(errs, HasLen, 13)

	c.Assert(errs[0].Error(), Equals, "Property test:test1 must be path to file")
	c.Assert(errs[1].Error(), Equals, "Property test:test1 must be path to readable file")
	c.Assert(errs[2].Error(), Equals, "Property test:test1 must be path to writable file")
	c.Assert(errs[3].Error(), Equals, "Property test:test1 must be path to executable file")
	c.Assert(errs[4].Error(), Equals, "Property test:test1 must be path to readable/writable file")
	c.Assert(errs[5].Error(), Equals, "Property test:test1 must be path to directory")
	c.Assert(errs[6].Error(), Equals, "Property test:test1 must be path to readable directory")
	c.Assert(errs[7].Error(), Equals, "Property test:test1 must be path to readable directory")
	c.Assert(errs[8].Error(), Equals, "Property test:test1 must be path to writable directory")
	c.Assert(errs[9].Error(), Equals, "Property test:test1 must be path to readable/writable directory")
	c.Assert(errs[10].Error(), Equals, "Property test:test1 must be path to object with given permissions (WX)")
	c.Assert(errs[11].Error(), Equals, "Validator fs.Perms requires non-empty input for checking test:test1 property")
	c.Assert(errs[12].Error(), Equals, "Validator fs.Perms doesn't support input with type <int> for checking test:test1 property")
}

func (s *ValidatorSuite) TestOwnerValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate(knf.Validators{
		{"test:test0", Owner, "root"},
		{"test:test1", Owner, "root"},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"test:test1", Owner, "nobody"},
		{"test:test1", Owner, "somerandomuser"},
		{"test:test1", Owner, ""},
		{"test:test1", Owner, 100},
	})

	c.Assert(errs, HasLen, 4)

	c.Assert(errs[0].Error(), Equals, "User nobody must be owner of /etc/passwd")
	c.Assert(errs[1].Error(), Equals, `Can't find user "somerandomuser" on system`)
	c.Assert(errs[2].Error(), Equals, "Validator fs.Owner requires non-empty input for checking test:test1 property")
	c.Assert(errs[3].Error(), Equals, "Validator fs.Owner doesn't support input with type <int> for checking test:test1 property")

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate(knf.Validators{
		{"test:test1", Owner, "root"},
	})

	c.Assert(errs, HasLen, 1)
	c.Assert(errs[0].Error(), Equals, `Can't get owner for "/etc/__unknown__"`)
}

func (s *ValidatorSuite) TestOwnerGroupValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	var errs []error

	if runtime.GOOS == "darwin" {
		errs = knf.Validate(knf.Validators{
			{"test:test0", OwnerGroup, "wheel"},
			{"test:test1", OwnerGroup, "wheel"},
		})
	} else {
		errs = knf.Validate(knf.Validators{
			{"test:test0", OwnerGroup, "root"},
			{"test:test1", OwnerGroup, "root"},
		})
	}

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"test:test1", OwnerGroup, "daemon"},
		{"test:test1", OwnerGroup, "somerandomgroup"},
		{"test:test1", OwnerGroup, ""},
		{"test:test1", OwnerGroup, 100},
	})

	c.Assert(errs, HasLen, 4)

	c.Assert(errs[0].Error(), Equals, "Group daemon must be owner of /etc/passwd")
	c.Assert(errs[1].Error(), Equals, `Can't find group "somerandomgroup" on system`)
	c.Assert(errs[2].Error(), Equals, "Validator fs.OwnerGroup requires non-empty input for checking test:test1 property")
	c.Assert(errs[3].Error(), Equals, "Validator fs.OwnerGroup doesn't support input with type <int> for checking test:test1 property")

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate(knf.Validators{
		{"test:test1", OwnerGroup, "daemon"},
	})

	c.Assert(errs, HasLen, 1)
	c.Assert(errs[0].Error(), Equals, `Can't get owner group for "/etc/__unknown__"`)
}

func (s *ValidatorSuite) TestFileModeValidator(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate(knf.Validators{
		{"test:test0", FileMode, os.FileMode(0644)},
		{"test:test1", FileMode, os.FileMode(0644)},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"test:test1", FileMode, os.FileMode(0777)},
	})

	c.Assert(errs, HasLen, 1)
	c.Assert(errs[0].Error(), Equals, "/etc/passwd has different mode (644 != 777)")

	configFile = createConfig(c, "/etc/__unknown__")

	err = knf.Global(configFile)
	c.Assert(err, IsNil)

	errs = knf.Validate(knf.Validators{
		{"test:test1", FileMode, os.FileMode(0644)},
		{"test:test1", FileMode, os.FileMode(0)},
		{"test:test1", FileMode, 100},
	})

	c.Assert(errs, HasLen, 3)

	c.Assert(errs[0].Error(), Equals, `Can't get mode for "/etc/__unknown__"`)
	c.Assert(errs[1].Error(), Equals, "Validator fs.FileMode requires non-empty input for checking test:test1 property")
	c.Assert(errs[2].Error(), Equals, "Validator fs.FileMode doesn't support input with type <int> for checking test:test1 property")
}

func (s *ValidatorSuite) TestMatchPattern(c *C) {
	configFile := createConfig(c, "/etc/passwd")

	err := knf.Global(configFile)
	c.Assert(err, IsNil)

	errs := knf.Validate(knf.Validators{
		{"test:test0", MatchPattern, "/etc/*"},
		{"test:test1", MatchPattern, "/etc/*"},
	})

	c.Assert(errs, HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"test:test1", MatchPattern, "/var/*"},
		{"test:test1", MatchPattern, "[]a"},
		{"test:test1", MatchPattern, ""},
		{"test:test1", MatchPattern, 100},
	})

	c.Assert(errs, HasLen, 4)

	c.Assert(errs[0].Error(), Equals, `Property test:test1 must match shell pattern "/var/*"`)
	c.Assert(errs[1].Error(), Equals, "Can't parse shell pattern: syntax error in pattern")
	c.Assert(errs[2].Error(), Equals, "Validator fs.MatchPattern requires non-empty input for checking test:test1 property")
	c.Assert(errs[3].Error(), Equals, "Validator fs.MatchPattern doesn't support input with type <int> for checking test:test1 property")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *C, data string) string {
	configPath := c.MkDir() + "/config.knf"
	configData := fmt.Sprintf(_CONFIG_TEMPLATE, data)

	err := os.WriteFile(configPath, []byte(configData), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
