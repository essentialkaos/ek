package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"sort"
	"testing"

	check "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type FSSuite struct {
	TempDir string
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { check.TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&FSSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (fs *FSSuite) SetUpSuite(c *check.C) {
	fs.TempDir = c.MkDir()
}

func (fs *FSSuite) TestList(c *check.C) {
	os.Mkdir(fs.TempDir+"/.dir0", 0755)

	os.Create(fs.TempDir + "/.file0")
	os.Create(fs.TempDir + "/file1.mp3")
	os.Create(fs.TempDir + "/file2.jpg")

	os.Mkdir(fs.TempDir+"/dir1", 0755)
	os.Mkdir(fs.TempDir+"/dir2", 0755)

	os.Create(fs.TempDir + "/dir1/file3.mp3")
	os.Create(fs.TempDir + "/dir2/file4.wav")

	os.Mkdir(fs.TempDir+"/dir1/dir3", 0755)

	listing1 := List(fs.TempDir, false)
	listing2 := List(fs.TempDir, true)
	listing3 := ListAll(fs.TempDir, false)
	listing4 := ListAll(fs.TempDir, true)
	listing5 := ListAllDirs(fs.TempDir, false)
	listing6 := ListAllDirs(fs.TempDir, true)
	listing7 := ListAllFiles(fs.TempDir, false)
	listing8 := ListAllFiles(fs.TempDir, true)
	listing9 := ListAllFiles(fs.TempDir, true, &ListingFilter{MatchPatterns: []string{"*.mp3", "*.wav"}})
	listing10 := ListAllFiles(fs.TempDir, true, &ListingFilter{NotMatchPatterns: []string{"*.mp3"}})

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
}

func (fs *FSSuite) TestProperPath(c *check.C) {
	paths := []string{"/root", "/bin", "/tmp"}

	c.Assert(ProperPath("FR", paths), check.Equals, "")
	c.Assert(ProperPath("DR", paths), check.Equals, "/bin")
	c.Assert(ProperPath("DRW", paths), check.Equals, "/tmp")
	c.Assert(ProperPath("D", paths), check.Equals, "/root")
	c.Assert(ProperPath("DX", paths), check.Equals, "/bin")
}
