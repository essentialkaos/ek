package container

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"math"
	"os"
	"strconv"
	"sync"
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

	os.WriteFile(s.DataDir+"/.dockerenv", []byte(``), 0644)

	os.WriteFile(s.DataDir+"/default", []byte(`sysfs /sys sysfs rw,nosuid,nodev,noexec,relatime 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
devtmpfs /dev devtmpfs rw,nosuid,size=930048k,nr_inodes=232512,mode=755 0 0
securityfs /sys/kernel/security securityfs rw,nosuid,nodev,noexec,relatime 0 0`), 0644)

	os.WriteFile(s.DataDir+"/yandex", []byte(`overlay-container /function/code/rootfs overlay rw,relatime,lowerdir=/function/code/3c854c8cbf469fda815b8f6183300c07cfa2fbb5703859ca79aff93ae934961b,upperdir=/tmp/.overlay/tmp/diff,workdir=/tmp/.overlay/tmp/work 0 0
devtmpfs /function/code/rootfs/dev devtmpfs rw,relatime,size=18464k,nr_inodes=4616,mode=755 0 0
/var/run /function/code/rootfs/etc/resolv.conf tmpfs rw,relatime,size=20400k,nr_inodes=5100 0 0`), 0644)

	os.WriteFile(s.DataDir+"/lxc", []byte(`none /sys/fs/cgroup cgroup2 rw,nosuid,nodev,noexec,relatime 0 0
lxcfs /proc/cpuinfo fuse.lxcfs rw,nosuid,nodev,relatime,user_id=0,group_id=0,allow_other 0 0`), 0644)

	os.WriteFile(s.DataDir+"/docker", []byte(`overlay / overlay rw,seclabel,relatime,lowerdir=/var/lib/docker/overlay2/l/ONS52X3BOCT7XPZRIOTDXVOTI5:/var/lib/docker/overlay2/l/RII7KWRJQAKYT6PQDZWIH3LQPY:/var/lib/docker/overlay2/l/K7NWZSOOPD6IQA3ZBMAFNCV2UK,upperdir=/var/lib/docker/overlay2/b912553379d74e8dc8b13c8bc97a1478324255fc249121bc3140c77edf10000c/diff,workdir=/var/lib/docker/overlay2/b912553379d74e8dc8b13c8bc97a1478324255fc249121bc3140c77edf10000c/work 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,seclabel,nosuid,size=65536k,mode=755 0 0`), 0644)

	os.WriteFile(s.DataDir+"/podman", []byte(`overlay / overlay rw,context="system_u:object_r:container_file_t:s0:c858,c956",relatime,lowerdir=/var/lib/containers/storage/overlay/l/4WA73D64E37PLK3SAORPPUISJK,upperdir=/var/lib/containers/storage/overlay/5d57db59db9567665efd0e17756445580b53be42396198b35a94f10ff30416be/diff,workdir=/var/lib/containers/storage/overlay/5d57db59db9567665efd0e17756445580b53be42396198b35a94f10ff30416be/work,metacopy=on,volatile 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,context="system_u:object_r:container_file_t:s0:c858,c956",nosuid,noexec,size=65536k,mode=755,inode64 0 0`), 0644)

	os.WriteFile(s.DataDir+"/docker-runsc", []byte(`none /etc/hosts 9p rw,trans=fd,rfdno=7,wfdno=7,aname=/,dfltuid=4294967294,dfltgid=4294967294,dcache=1000,cache=remote_revalidating,disable_fifo_open,directfs 0 0
none /etc/hostname 9p rw,trans=fd,rfdno=6,wfdno=6,aname=/,dfltuid=4294967294,dfltgid=4294967294,dcache=1000,cache=remote_revalidating,disable_fifo_open,directfs 0 0`), 0644)

	os.WriteFile(s.DataDir+"/containerd", []byte(`overlay / overlay rw,relatime,lowerdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/111/fs,upperdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/144/fs,workdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/144/work,xino=off 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,nosuid,size=65536k,mode=755 0 0
tmpfs /run/secrets/kubernetes.io/serviceaccount tmpfs ro,relatime,size=2760104k 0 0`), 0644)

	os.WriteFile(s.DataDir+"/host", []byte(`overlay / overlay rw,relatime,lowerdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/111/fs,upperdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/144/fs,workdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/144/work,xino=off 0 0
systemd-1 /proc/sys/fs/binfmt_misc autofs rw,relatime,fd=29,pgrp=1,timeout=0,minproto=5,maxproto=5,direct,pipe_ino=23748 0 0
proc /proc proc rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev tmpfs rw,nosuid,size=65536k,mode=755 0 0
tmpfs /run/secrets/kubernetes.io/serviceaccount tmpfs ro,relatime,size=2760104k 0 0`), 0644)
}

func (s *ContainerSuite) TestGetEngine(c *C) {
	mountsFile = "/__unknown__"
	c.Assert(GetEngine(), Equals, "")

	once = sync.Once{}

	mountsFile = s.DataDir + "/default"
	c.Assert(GetEngine(), Equals, "")
	c.Assert(IsContainer(), Equals, false)

	once = sync.Once{}

	mountsFile = s.DataDir + "/host"
	c.Assert(GetEngine(), Equals, "")
	c.Assert(IsK8s(), Equals, false)

	once = sync.Once{}

	mountsFile = s.DataDir + "/yandex"
	c.Assert(GetEngine(), Equals, YANDEX)
	c.Assert(IsK8s(), Equals, false)

	once = sync.Once{}

	mountsFile = s.DataDir + "/lxc"
	c.Assert(GetEngine(), Equals, LXC)

	once = sync.Once{}

	mountsFile = s.DataDir + "/containerd"
	c.Assert(GetEngine(), Equals, CONTAINERD)
	c.Assert(IsK8s(), Equals, true)

	dockerEnv = s.DataDir + "/.dockerenv"

	isK8s = false
	once = sync.Once{}

	mountsFile = s.DataDir + "/docker"
	c.Assert(GetEngine(), Equals, DOCKER)

	once = sync.Once{}

	mountsFile = s.DataDir + "/docker-runsc"
	c.Assert(GetEngine(), Equals, DOCKER_RUNSC)

	dockerEnv = "/.dockerenv"
	once = sync.Once{}

	mountsFile = s.DataDir + "/podman"
	c.Assert(GetEngine(), Equals, PODMAN)

	// Check cached info
	mountsFile = s.DataDir + "/podman"
	c.Assert(GetEngine(), Equals, PODMAN)
	c.Assert(IsContainer(), Equals, true)
}

func (s *ContainerSuite) TestGetCPUQuota(c *C) {
	cgroupDir = c.MkDir()

	quota, unlimited, err := GetCPUQuota()
	c.Assert(os.IsNotExist(err), Equals, true)
	c.Assert(quota, Equals, 0.0)
	c.Assert(unlimited, Equals, false)

	c.Assert(os.WriteFile(cgroupDir+"/cpu.max", []byte("50000 100000\n"), 0644), IsNil)

	quota, unlimited, err = GetCPUQuota()
	c.Assert(err, IsNil)
	c.Assert(quota, Equals, 0.5)
	c.Assert(unlimited, Equals, false)

	c.Assert(os.WriteFile(cgroupDir+"/cpu.max", []byte("max 100000\n"), 0644), IsNil)

	quota, unlimited, err = GetCPUQuota()
	c.Assert(err, IsNil)
	c.Assert(quota, Equals, 0.0)
	c.Assert(unlimited, Equals, true)

	for _, value := range []string{"50000", "invalid 100000", "50000 0", "50000 invalid"} {
		c.Assert(os.WriteFile(cgroupDir+"/cpu.max", []byte(value), 0644), IsNil)

		quota, unlimited, err = GetCPUQuota()
		c.Assert(err, NotNil, Commentf("cpu.max = %q", value))
		c.Assert(quota, Equals, 0.0)
		c.Assert(unlimited, Equals, false)
	}
}

func (s *ContainerSuite) TestGetMemLimit(c *C) {
	cgroupDir = c.MkDir()

	limit, unlimited, err := GetMemLimit()
	c.Assert(os.IsNotExist(err), Equals, true)
	c.Assert(limit, Equals, uint64(0))
	c.Assert(unlimited, Equals, false)

	c.Assert(os.WriteFile(cgroupDir+"/memory.max", []byte("536870912\n"), 0644), IsNil)

	limit, unlimited, err = GetMemLimit()
	c.Assert(err, IsNil)
	c.Assert(limit, Equals, uint64(536870912))
	c.Assert(unlimited, Equals, false)

	c.Assert(os.WriteFile(cgroupDir+"/memory.max", []byte("max\n"), 0644), IsNil)

	limit, unlimited, err = GetMemLimit()
	c.Assert(err, IsNil)
	c.Assert(limit, Equals, uint64(0))
	c.Assert(unlimited, Equals, true)

	c.Assert(os.WriteFile(cgroupDir+"/memory.max", []byte("-1\n"), 0644), IsNil)

	limit, unlimited, err = GetMemLimit()
	c.Assert(err, NotNil)
	c.Assert(limit, Equals, uint64(0))
	c.Assert(unlimited, Equals, false)
}

func (s *ContainerSuite) TestGetCPUSet(c *C) {
	cgroupDir = c.MkDir()

	cpus, err := GetCPUSet()
	c.Assert(os.IsNotExist(err), Equals, true)
	c.Assert(cpus, IsNil)

	c.Assert(os.WriteFile(cgroupDir+"/cpuset.cpus.effective", []byte("1-2,5\n"), 0644), IsNil)

	cpus, err = GetCPUSet()
	c.Assert(err, IsNil)
	c.Assert(cpus, DeepEquals, []bool{false, true, true, false, false, true})

	c.Assert(os.WriteFile(cgroupDir+"/cpuset.cpus.effective", []byte("\n"), 0644), IsNil)

	cpus, err = GetCPUSet()
	c.Assert(err, IsNil)
	c.Assert(cpus, DeepEquals, []bool{})

	for _, value := range []string{"3-1", "0,,2", "invalid"} {
		c.Assert(os.WriteFile(cgroupDir+"/cpuset.cpus.effective", []byte(value), 0644), IsNil)

		cpus, err = GetCPUSet()
		c.Assert(err, NotNil, Commentf("cpuset.cpus.effective = %q", value))
		c.Assert(cpus, IsNil)
	}

	id := strconv.Itoa(math.MaxInt)

	c.Assert(os.WriteFile(cgroupDir+"/cpuset.cpus.effective", []byte(id+"\n"), 0644), IsNil)

	cpus, err = GetCPUSet()
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "CPU ID too large for slice indexing: "+id)
	c.Assert(cpus, IsNil)
}
