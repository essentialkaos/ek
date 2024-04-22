package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	"github.com/essentialkaos/ek/v12/fmtc"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type SupportSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&SupportSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

var gomodData = []byte(`module github.com/essentialkaos/ek/v12

go 1.18

require (
  github.com/essentialkaos/check v1.4.0
  github.com/essentialkaos/go-linenoise/v3 v3.4.0
  golang.org/x/crypto v0.21.0
  golang.org/x/sys v0.18.0
)

require (
  github.com/kr/pretty v0.3.1 // indirect
  github.com/kr/text v0.2.0 // indirect
  github.com/rogpeppe/go-internal v1.11.0 // indirect
)
`)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *SupportSuite) TestCollect(c *C) {
	i := Collect("test", "1.2.3")

	os.Setenv("SUPPORT_VAR", "123")

	chks := []Check{
		Check{CHECK_OK, "Test", ""},
		Check{CHECK_OK, "Test", "Test message"},
		Check{CHECK_ERROR, "Test", "Test message"},
		Check{CHECK_WARN, "Test", "Test message"},
		Check{CHECK_SKIP, "Test", "Test message"},
		Check{CHECK_WARN, "", "Test message"},
	}

	deps := []Dep{
		Dep{"test", "1.0", ""},
		Dep{"test", "1.0", "0000"},
	}

	pkgs := []Pkg{
		Pkg{"test", "1.2.3"},
		Pkg{"test", ""},
		Pkg{},
	}

	apps := []App{
		App{"test", "1.2.3"},
		App{"test", ""},
		App{},
	}

	c.Assert(i, NotNil)
	c.Assert(i.WithDeps(deps), NotNil)
	c.Assert(i.WithRevision(""), NotNil)
	c.Assert(i.WithRevision("1234567"), NotNil)
	c.Assert(i.WithPackages(pkgs), NotNil)
	c.Assert(i.WithApps(apps...), NotNil)
	c.Assert(i.WithChecks(chks...), NotNil)
	c.Assert(i.WithEnvVars("", "SUPPORT_VAR", "TERM", "CI"), NotNil)
	c.Assert(i.WithNetwork(&NetworkInfo{PublicIP: "192.168.1.1", Hostname: "test.loc"}), NotNil)
	c.Assert(i.WithFS([]FSInfo{FSInfo{}, FSInfo{"/", "/dev/vda1", "ext4", 1000, 10000}}), NotNil)

	i.Print()

	i.Build = nil
	i.OS = nil
	i.Env = nil
	i.Pkgs = nil
	i.Apps = nil
	i.Checks = nil
	i.Network = nil
	i.FS = nil
	i.Deps = nil

	i.Print()

	i.System.ContainerEngine = "docker"
	i.printOSInfo()
	i.System.ContainerEngine = "podman"
	i.printOSInfo()
	i.System.ContainerEngine = "lxc"
	i.printOSInfo()
	i.System.ContainerEngine = "yandex"
	i.printOSInfo()
}

func (s *SupportSuite) TestNil(c *C) {
	var i *Info

	c.Assert(i.WithDeps(nil), IsNil)
	c.Assert(i.WithRevision(""), IsNil)
	c.Assert(i.WithPackages(nil), IsNil)
	c.Assert(i.WithApps(App{}), IsNil)
	c.Assert(i.WithChecks(Check{}), IsNil)
	c.Assert(i.WithEnvVars(""), IsNil)
	c.Assert(i.WithNetwork(nil), IsNil)
	c.Assert(i.WithFS(nil), IsNil)

	c.Assert(func() { i.Print() }, NotPanics)
}

func (s *SupportSuite) TestColorBulletGen(c *C) {
	c.Assert(getHashColorBullet(""), Equals, "")

	fmtc.DisableColors = true
	c.Assert(getHashColorBullet("1a2b3c4"), Equals, "")
	fmtc.DisableColors = false

	getHashColorBullet("1a2b3c4")
}

func (s *SupportSuite) TestSizeCalc(c *C) {
	c.Assert(getMaxKeySize([]EnvVar{
		EnvVar{"Test", "1"}, EnvVar{"TestABCD1234", "1"}, EnvVar{"Te", "1"},
	}), Equals, 12)

	c.Assert(getMaxAppNameSize([]App{
		App{"test", "1"}, App{"testABCD1234", "1"}, App{"t", "1"},
	}), Equals, 12)

	c.Assert(getMaxDeviceNameSize([]FSInfo{
		FSInfo{Device: "/dev/sda1"}, FSInfo{Device: "/dev/test/test"}, FSInfo{Device: "/dev"},
	}), Equals, 14)
}
