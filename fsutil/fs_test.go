package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"os"
	"sort"
	"testing"

	check "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type FSSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { check.TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&FSSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestList(c *check.C) {
	tmpDir := c.MkDir()

	os.Mkdir(tmpDir+"/.dir0", 0755)

	os.Create(tmpDir + "/.file0")

	ioutil.WriteFile(tmpDir+"/file1.mp3", []byte("TESTDATA12345678"), 644)
	ioutil.WriteFile(tmpDir+"/file2.jpg", []byte("TESTDATA"), 644)

	os.Mkdir(tmpDir+"/dir1", 0755)
	os.Mkdir(tmpDir+"/dir2", 0755)

	os.Create(tmpDir + "/dir1/file3.mp3")
	os.Create(tmpDir + "/dir2/file4.wav")

	os.Mkdir(tmpDir+"/dir1/dir3", 0755)

	listing1 := List(tmpDir, false)
	listing2 := List(tmpDir, true)
	listing3 := ListAll(tmpDir, false)
	listing4 := ListAll(tmpDir, true)
	listing5 := ListAllDirs(tmpDir, false)
	listing6 := ListAllDirs(tmpDir, true)
	listing7 := ListAllFiles(tmpDir, false)
	listing8 := ListAllFiles(tmpDir, true)
	listing9 := ListAllFiles(tmpDir, true, ListingFilter{MatchPatterns: []string{"*.mp3", "*.wav"}})
	listing10 := ListAllFiles(tmpDir, true, ListingFilter{NotMatchPatterns: []string{"*.mp3"}})
	listing11 := List(tmpDir, true, ListingFilter{Perms: "DR"})
	listing12 := List(tmpDir, true, ListingFilter{NotPerms: "DR"})
	listing13 := ListAllFiles(tmpDir, true, ListingFilter{NotMatchPatterns: []string{"*.mp3"}, SizeZero: true})
	listing14 := ListAllFiles(tmpDir, false, ListingFilter{SizeEqual: 16})
	listing15 := ListAllFiles(tmpDir, false, ListingFilter{SizeLess: 12, SizeGreater: 5})
	listing16 := ListAllFiles(tmpDir, false, ListingFilter{SizeGreater: 12})

	sort.Strings(listing1)
	sort.Strings(listing2)
	sort.Strings(listing3)
	sort.Strings(listing4)
	sort.Strings(listing5)
	sort.Strings(listing6)
	sort.Strings(listing7)
	sort.Strings(listing8)
	sort.Strings(listing9)
	sort.Strings(listing10)
	sort.Strings(listing11)
	sort.Strings(listing12)
	sort.Strings(listing13)
	sort.Strings(listing14)
	sort.Strings(listing15)
	sort.Strings(listing16)

	c.Assert(
		listing1,
		check.DeepEquals,
		[]string{".dir0", ".file0", "dir1", "dir2", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing2,
		check.DeepEquals,
		[]string{"dir1", "dir2", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing3,
		check.DeepEquals,
		[]string{".dir0", ".file0", "dir1", "dir1/dir3", "dir1/file3.mp3", "dir2", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing4,
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir1/file3.mp3", "dir2", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing5,
		check.DeepEquals,
		[]string{".dir0", "dir1", "dir1/dir3", "dir2"},
	)

	c.Assert(
		listing6,
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir2"},
	)

	c.Assert(
		listing7,
		check.DeepEquals,
		[]string{".file0", "dir1/file3.mp3", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing8,
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "dir2/file4.wav", "file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing9,
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(
		listing10,
		check.DeepEquals,
		[]string{"dir2/file4.wav", "file2.jpg"},
	)

	c.Assert(
		listing11,
		check.DeepEquals,
		[]string{"dir1", "dir2"},
	)

	c.Assert(
		listing12,
		check.DeepEquals,
		[]string{"file1.mp3", "file2.jpg"},
	)

	c.Assert(
		listing13,
		check.DeepEquals,
		[]string{"dir2/file4.wav"},
	)

	c.Assert(
		listing14,
		check.DeepEquals,
		[]string{"file1.mp3"},
	)

	c.Assert(
		listing15,
		check.DeepEquals,
		[]string{"file2.jpg"},
	)

	c.Assert(
		listing16,
		check.DeepEquals,
		[]string{"file1.mp3"},
	)
}

func (s *FSSuite) TestProperPath(c *check.C) {
	tmpFile := c.MkDir() + "/test.txt"

	os.OpenFile(tmpFile, os.O_CREATE, 0644)

	paths := []string{"/etc/sudoers", "/etc/passwd", tmpFile}

	c.Assert(ProperPath("DR", paths), check.Equals, "")
	c.Assert(ProperPath("FR", paths), check.Equals, "/etc/passwd")
	c.Assert(ProperPath("FRW", paths), check.Equals, tmpFile)
	c.Assert(ProperPath("FRWS", paths), check.Equals, "")
	c.Assert(ProperPath("F", paths), check.Equals, "/etc/sudoers")

	os.Remove(tmpFile)
}

func (s *FSSuite) TestWalker(c *check.C) {
	tmpDir := c.MkDir()

	os.Chdir(tmpDir)

	tmpDir, _ = os.Getwd()

	os.MkdirAll(tmpDir+"/dir1/dir2/dir3/dir4", 0755)
	os.Chdir(tmpDir)

	c.Assert(Current(), check.Equals, tmpDir)
	c.Assert(Pop(), check.Equals, tmpDir)

	dirStack = nil

	c.Assert(Push("dir1"), check.Equals, tmpDir+"/dir1")
	c.Assert(Push("dir9"), check.Equals, "")
	c.Assert(Push("dir2/dir3"), check.Equals, tmpDir+"/dir1/dir2/dir3")
	c.Assert(Push("dir4"), check.Equals, tmpDir+"/dir1/dir2/dir3/dir4")
	c.Assert(Push("dir9"), check.Equals, "")
	c.Assert(Pop(), check.Equals, tmpDir+"/dir1/dir2/dir3")
	c.Assert(Pop(), check.Equals, tmpDir+"/dir1")
	c.Assert(Pop(), check.Equals, tmpDir)
	c.Assert(Pop(), check.Equals, tmpDir)
}
