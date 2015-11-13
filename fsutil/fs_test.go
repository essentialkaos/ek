package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	check "gopkg.in/check.v1"
	"os"
	"testing"
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

	c.Assert(
		List(fs.TempDir, false),
		check.DeepEquals,
		[]string{"dir1", ".file0", "file2.jpg", "dir2", "file1.mp3", ".dir0"},
	)

	c.Assert(
		List(fs.TempDir, true),
		check.DeepEquals,
		[]string{"dir1", "file2.jpg", "dir2", "file1.mp3"},
	)

	c.Assert(ListAll(fs.TempDir, false),
		check.DeepEquals,
		[]string{"dir1", "dir1/file3.mp3", "dir1/dir3", ".file0", "file2.jpg", "dir2", "dir2/file4.wav", "file1.mp3", ".dir0"},
	)

	c.Assert(ListAll(fs.TempDir, true),
		check.DeepEquals,
		[]string{"dir1", "dir1/file3.mp3", "dir1/dir3", "file2.jpg", "dir2", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(ListAllDirs(fs.TempDir, false),
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir2", ".dir0"},
	)

	c.Assert(ListAllDirs(fs.TempDir, true),
		check.DeepEquals,
		[]string{"dir1", "dir1/dir3", "dir2"},
	)

	c.Assert(ListAllFiles(fs.TempDir, false),
		check.DeepEquals,
		[]string{"dir1/file3.mp3", ".file0", "file2.jpg", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(ListAllFiles(fs.TempDir, true),
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "file2.jpg", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(
		ListAllFiles(fs.TempDir, true, &ListingFilter{MatchPatterns: []string{"*.mp3", "*.wav"}}),
		check.DeepEquals,
		[]string{"dir1/file3.mp3", "dir2/file4.wav", "file1.mp3"},
	)

	c.Assert(
		ListAllFiles(fs.TempDir, true, &ListingFilter{NotMatchPatterns: []string{"*.mp3"}}),
		check.DeepEquals,
		[]string{"file2.jpg", "dir2/file4.wav"},
	)
}
