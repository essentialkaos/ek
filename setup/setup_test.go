package setup

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SetupSuite struct {
	App App
	Bin binaryInfo
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SetupSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SetupSuite) SetUpSuite(c *C) {
	s.App = App{
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
	}

	s.Bin = binaryInfo{File: "/usr/bin/echo", Name: "test1"}
}

func (s *SetupSuite) TestInstall(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()

	err := installFiles(s.App, s.Bin)
	c.Assert(err, IsNil)

	isGen, _ := isGeneratedUnit(s.Bin.ServiceUnitPath())
	c.Assert(isGen, Equals, true)
	_, err = isGeneratedUnit("/_unknown_")
	c.Assert(err, NotNil)

	app := s.App
	app.Options = nil

	err = installFiles(app, s.Bin)
	c.Assert(err, IsNil)

	c.Assert(s.Bin.IsBinInstalled(), Equals, true)
	c.Assert(s.Bin.IsServiceInstalled(), Equals, true)
}

func (s *SetupSuite) TestInstallErrors(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()

	logDir = "/_unknown_"
	err := installFiles(s.App, s.Bin)
	c.Assert(err, NotNil)

	serviceDir = "/_unknown_"
	err = installFiles(s.App, s.Bin)
	c.Assert(err, NotNil)

	binaryDir = "/_unknown_"
	err = installFiles(s.App, s.Bin)
	c.Assert(err, NotNil)
}

func (s *SetupSuite) TestUninstall(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()

	err := installFiles(s.App, s.Bin)
	c.Assert(err, IsNil)

	err = uninstallFiles(s.App, s.Bin, true)
	c.Assert(err, IsNil)
}

func (s *SetupSuite) TestUninstallErrors(c *C) {
	serviceDir = c.MkDir()
	binaryDir = c.MkDir()
	logDir = c.MkDir()

	err := installFiles(s.App, s.Bin)
	c.Assert(err, IsNil)

	binaryDir = "/_unknown_"
	err = uninstallFiles(s.App, s.Bin, true)
	c.Assert(err, NotNil)

	serviceDir = "/_unknown_"
	err = uninstallFiles(s.App, s.Bin, true)
	c.Assert(err, NotNil)
}

func (s *SetupSuite) TestAux(c *C) {
	bin := getBinaryInfo()

	c.Assert(bin.File, Not(Equals), "")
	c.Assert(bin.Name, Not(Equals), "")

	c.Assert(s.Bin.checkForInstall(), NotNil)
	c.Assert(s.Bin.checkForUninstall(), NotNil)
}
