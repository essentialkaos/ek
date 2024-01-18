package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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

type ContainerSuite struct {
	DataDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ContainerSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ContainerSuite) SetUpSuite(c *C) {
	s.DataDir = c.MkDir()

	os.WriteFile(s.DataDir+"/default", []byte(`sysfs /sys sysfs rw,nosuid,nodev,noexec,relatime 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
devtmpfs /dev devtmpfs rw,nosuid,size=930048k,nr_inodes=232512,mode=755 0 0
securityfs /sys/kernel/security securityfs rw,nosuid,nodev,noexec,relatime 0 0`), 0644)

	os.WriteFile(s.DataDir+"/lxc", []byte(`none /sys/fs/cgroup cgroup2 rw,nosuid,nodev,noexec,relatime 0 0
lxcfs /proc/cpuinfo fuse.lxcfs rw,nosuid,nodev,relatime,user_id=0,group_id=0,allow_other 0 0`), 0644)

	os.WriteFile(s.DataDir+"/docker", []byte(`overlay / overlay rw,seclabel,relatime,lowerdir=/var/lib/docker/overlay2/l/ONS52X3BOCT7XPZRIOTDXVOTI5:/var/lib/docker/overlay2/l/RII7KWRJQAKYT6PQDZWIH3LQPY:/var/lib/docker/overlay2/l/K7NWZSOOPD6IQA3ZBMAFNCV2UK,upperdir=/var/lib/docker/overlay2/b912553379d74e8dc8b13c8bc97a1478324255fc249121bc3140c77edf10000c/diff,workdir=/var/lib/docker/overlay2/b912553379d74e8dc8b13c8bc97a1478324255fc249121bc3140c77edf10000c/work 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,seclabel,nosuid,size=65536k,mode=755 0 0`), 0644)

	os.WriteFile(s.DataDir+"/podman", []byte(`overlay / overlay rw,context="system_u:object_r:container_file_t:s0:c858,c956",relatime,lowerdir=/var/lib/containers/storage/overlay/l/4WA73D64E37PLK3SAORPPUISJK,upperdir=/var/lib/containers/storage/overlay/5d57db59db9567665efd0e17756445580b53be42396198b35a94f10ff30416be/diff,workdir=/var/lib/containers/storage/overlay/5d57db59db9567665efd0e17756445580b53be42396198b35a94f10ff30416be/work,metacopy=on,volatile 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,context="system_u:object_r:container_file_t:s0:c858,c956",nosuid,noexec,size=65536k,mode=755,inode64 0 0`), 0644)
}

func (s *ContainerSuite) TestGetEngine(c *C) {
	mountsFile = "/__unknown__"
	c.Assert(GetEngine(), Equals, "")

	mountsFile = s.DataDir + "/default"
	c.Assert(GetEngine(), Equals, "")

	mountsFile = s.DataDir + "/lxc"
	c.Assert(GetEngine(), Equals, LXC)

	mountsFile = s.DataDir + "/docker"
	c.Assert(GetEngine(), Equals, DOCKER)

	mountsFile = s.DataDir + "/podman"
	c.Assert(GetEngine(), Equals, PODMAN)
}
