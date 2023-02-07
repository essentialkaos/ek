package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ContainerSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ContainerSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ContainerSuite) TestGetEngine(c *C) {
	testEnv := c.MkDir() + "/test"

	dockerEnv = "/_unknown_"
	podmanEnv = "/_unknown_"

	c.Assert(GetEngine(), Equals, "")

	os.WriteFile(testEnv, []byte("TEST"), 0644)

	dockerEnv = testEnv
	podmanEnv = "/_unknown_"

	c.Assert(GetEngine(), Equals, DOCKER)

	dockerEnv = "/_unknown_"
	podmanEnv = testEnv

	c.Assert(GetEngine(), Equals, PODMAN)
}
