package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ENVSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ENVSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ENVSuite) TestEnv(c *C) {
	if os.Getenv("EK_TEST_PORT") == "" {
		os.Setenv("EK_TEST_PORT", "8080")
	}

	envs := Get()

	c.Assert(envs["EK_TEST_PORT"], Equals, "8080")

	c.Assert(envs.GetS("EK_TEST_PORT"), Equals, "8080")
	c.Assert(envs.GetI("EK_TEST_PORT"), Equals, 8080)
	c.Assert(envs.GetF("EK_TEST_PORT"), Equals, 8080.0)

	c.Assert(envs.GetS("UNKNOWN_VARIABLE"), Equals, "")
	c.Assert(envs.GetI("UNKNOWN_VARIABLE"), Equals, -1)
	c.Assert(envs.GetF("UNKNOWN_VARIABLE"), Equals, -1.0)

	c.Assert(Which("cat"), Not(Equals), "")
	c.Assert(Which("catABCD1234"), Equals, "")

	var v *Variable

	c.Assert(v.Get(), Equals, "")
	c.Assert(v.String(), Equals, "")
	c.Assert(v.Reset(), IsNil)

	v = Var("EK_TEST_PORT")

	c.Assert(v.Get(), Equals, "8080")
	c.Assert(v.String(), Equals, "8080")
	c.Assert(v.Is("8080"), Equals, true)
	c.Assert(v.Reset(), NotNil)
}
