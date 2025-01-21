package setup

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SetupSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SetupSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SetupSuite) TestInstall(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()
	configDir = c.MkDir()

	app, bin := s.generateApp(), s.generateBin()

	err := installFiles(app, bin)
	c.Assert(err, IsNil)

	isGen, _ := isGeneratedUnit(bin.ServiceUnitPath())
	c.Assert(isGen, Equals, true)
	_, err = isGeneratedUnit("/_unknown_")
	c.Assert(err, NotNil)

	app = s.generateApp()
	app.Options = nil

	err = installFiles(app, bin)
	c.Assert(err, IsNil)

	c.Assert(bin.IsBinInstalled(), Equals, true)
	c.Assert(bin.IsServiceInstalled(), Equals, true)
}

func (s *SetupSuite) TestInstallErrors(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()

	app, bin := s.generateApp(), s.generateBin()

	logDir = "/_unknown_"
	err := installFiles(app, bin)
	c.Assert(err, NotNil)

	serviceDir = "/_unknown_"
	err = installFiles(app, bin)
	c.Assert(err, NotNil)

	binaryDir = "/_unknown_"
	err = installFiles(app, bin)
	c.Assert(err, NotNil)

	app.Configs = []Config{
		{"abcd/test1.knf", []byte("[main]\n  test: 1\n\n"), 0},
	}

	err = installFiles(app, bin)
	c.Assert(err, NotNil)
}

func (s *SetupSuite) TestUninstall(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()
	configDir = c.MkDir()

	app, bin := s.generateApp(), s.generateBin()

	err := installFiles(app, bin)
	c.Assert(err, IsNil)

	app.Configs = append(app.Configs, Config{Name: "testX.knf"})

	err = uninstallFiles(app, bin, true)
	c.Assert(err, IsNil)

	app.Configs = nil

	err = installFiles(app, bin)
	c.Assert(err, IsNil)

	err = uninstallFiles(app, bin, true)
	c.Assert(err, IsNil)

	c.Assert(uninstallConfigurationFiles(app), IsNil)
}

func (s *SetupSuite) TestUninstallErrors(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()

	app, bin := s.generateApp(), s.generateBin()

	err := installFiles(app, bin)
	c.Assert(err, IsNil)

	binaryDir = "/_unknown_"
	err = uninstallFiles(app, bin, true)
	c.Assert(err, NotNil)

	serviceDir = "/_unknown_"
	err = uninstallFiles(app, bin, true)
	c.Assert(err, NotNil)
}

func (s *SetupSuite) TestAux(c *C) {
	bin := getBinaryInfo()

	app, bin := s.generateApp(), s.generateBin()

	app.Install()
	app.Uninstall(false)

	c.Assert(bin.File, Not(Equals), "")
	c.Assert(bin.Name, Not(Equals), "")

	c.Assert(checkForInstall(app, bin), NotNil)
	c.Assert(checkForUninstall(app, bin), NotNil)

	app = s.generateApp()
	app.Configs = []Config{
		{"abcd/test1.knf", []byte("[main]\n  test: 1\n\n"), 0},
	}

	c.Assert(checkForInstall(app, bin), DeepEquals, errors.New("Configuration file name \"abcd/test1.knf\" is invalid"))
	c.Assert(checkForUninstall(app, bin), DeepEquals, errors.New("Configuration file name \"abcd/test1.knf\" is invalid"))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SetupSuite) generateApp() App {
	return App{
		Name:               "Test",
		Options:            []string{"--config", "/etc/test.knf"},
		DocsURL:            "https://domain.com",
		User:               "nobody",
		Identifier:         "TEST",
		WorkingDir:         "/srv/test",
		StopSignal:         "TERM",
		ReloadSignal:       "HUP",
		WithLog:            true,
		WithoutPrivateTemp: false,
		Configs: []Config{
			{"test1.knf", []byte("[main]\n  test: 1\n\n"), 0},
			{"test2.knf", []byte("[main]\n  test: 2\n\n"), 0644},
		},
	}
}

func (s *SetupSuite) generateBin() binaryInfo {
	return binaryInfo{File: "/usr/bin/echo", Name: "test1"}
}
