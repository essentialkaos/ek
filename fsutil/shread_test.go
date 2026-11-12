package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestShred(c *check.C) {
	tmpDir := c.MkDir()
	bufferSize := 1024 * 1024

	for _, zeroFill := range []bool{false, true} {
		for _, size := range []int{0, 1, bufferSize - 1, bufferSize, bufferSize + 17} {
			name := fmt.Sprintf("file-%t-%d", zeroFill, size)
			path := filepath.Join(tmpDir, name)
			comment := check.Commentf("zeroFill=%t, size=%d", zeroFill, size)

			c.Assert(
				os.WriteFile(path, bytes.Repeat([]byte{0xAA}, size), 0600),
				check.IsNil,
				comment,
			)

			c.Assert(Shred(path, zeroFill), check.IsNil, comment)

			_, err := os.Lstat(path)
			c.Assert(errors.Is(err, os.ErrNotExist), check.Equals, true, comment)
		}
	}
}

func (s *FSSuite) TestShredMissingFile(c *check.C) {
	path := filepath.Join(c.MkDir(), "missing")

	err := Shred(path, true)

	c.Assert(err, check.NotNil)
	c.Assert(errors.Is(err, os.ErrNotExist), check.Equals, true)
}

func (s *FSSuite) TestShredDirectory(c *check.C) {
	path := c.MkDir()

	c.Assert(Shred(path, true), check.ErrorMatches, `.*not a regular file`)

	info, err := os.Stat(path)
	c.Assert(err, check.IsNil)
	c.Assert(info.IsDir(), check.Equals, true)
}

func (s *FSSuite) TestShredSymlink(c *check.C) {
	tmpDir := c.MkDir()
	target := filepath.Join(tmpDir, "target")
	link := filepath.Join(tmpDir, "link")
	content := []byte("must remain unchanged")

	c.Assert(os.WriteFile(target, content, 0600), check.IsNil)

	if err := os.Symlink(target, link); err != nil {
		c.Skip("Cannot create symbolic links: " + err.Error())
	}

	c.Assert(Shred(link, true), check.ErrorMatches, `.*not a regular file`)

	got, err := os.ReadFile(target)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, content)

	info, err := os.Lstat(link)
	c.Assert(err, check.IsNil)
	c.Assert(info.Mode()&os.ModeSymlink != 0, check.Equals, true)
}

func (s *FSSuite) TestShredReadOnlyFile(c *check.C) {
	path := filepath.Join(c.MkDir(), "readonly")

	c.Assert(os.WriteFile(path, []byte("sensitive data"), 0600), check.IsNil)
	c.Assert(os.Chmod(path, 0400), check.IsNil)

	c.Assert(Shred(path, false), check.IsNil)

	_, err := os.Lstat(path)
	c.Assert(errors.Is(err, os.ErrNotExist), check.Equals, true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestOverwriteFileZeroFill(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")
	content := bytes.Repeat([]byte{0xAA}, (1024*1024)+17)

	c.Assert(os.WriteFile(path, content, 0600), check.IsNil)

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	c.Assert(err, check.IsNil)

	overwriteErr := overwriteFile(file, true)
	closeErr := file.Close()

	c.Assert(overwriteErr, check.IsNil)
	c.Assert(closeErr, check.IsNil)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, make([]byte, len(content)))
}

func (s *FSSuite) TestOverwriteFileRandomFill(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")
	content := bytes.Repeat([]byte{0xAA}, 4096)

	c.Assert(os.WriteFile(path, content, 0600), check.IsNil)

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	c.Assert(err, check.IsNil)

	overwriteErr := overwriteFile(file, false)
	closeErr := file.Close()

	c.Assert(overwriteErr, check.IsNil)
	c.Assert(closeErr, check.IsNil)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.HasLen, len(content))
	c.Assert(bytes.Equal(got, content), check.Equals, false)
	c.Assert(bytes.Equal(got, make([]byte, len(content))), check.Equals, false)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestOverwritePassZeroFill(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")
	content := bytes.Repeat([]byte{0xAA}, 37)

	c.Assert(os.WriteFile(path, content, 0600), check.IsNil)

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	c.Assert(err, check.IsNil)

	// Ensure the pass seeks back to offset zero.
	_, seekErr := file.Seek(11, 0)

	// A nonzero buffer verifies that the zero-fill pass clears reused data.
	buffer := bytes.Repeat([]byte{0xFF}, 8)
	overwriteErr := overwritePass(file, 29, buffer, true)
	closeErr := file.Close()

	c.Assert(seekErr, check.IsNil)
	c.Assert(overwriteErr, check.IsNil)
	c.Assert(closeErr, check.IsNil)

	expected := append(make([]byte, 29), content[29:]...)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, expected)
}

func (s *FSSuite) TestOverwritePassEmptyBuffer(c *check.C) {
	file, err := os.Create(filepath.Join(c.MkDir(), "data"))
	c.Assert(err, check.IsNil)

	overwriteErr := overwritePass(file, 1, nil, false)
	closeErr := file.Close()

	c.Assert(overwriteErr, check.ErrorMatches, "empty overwrite buffer")
	c.Assert(closeErr, check.IsNil)
}

func (s *FSSuite) TestOverwritePassReadOnlyFile(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")
	content := []byte("must remain unchanged")

	c.Assert(os.WriteFile(path, content, 0600), check.IsNil)

	file, err := os.Open(path)
	c.Assert(err, check.IsNil)

	overwriteErr := overwritePass(file, int64(len(content)), make([]byte, 8), true)
	closeErr := file.Close()

	c.Assert(overwriteErr, check.ErrorMatches, `write data: .*`)
	c.Assert(closeErr, check.IsNil)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, content)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FSSuite) TestOverwriteAndClose(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")

	c.Assert(os.WriteFile(path, []byte("sensitive data"), 0600), check.IsNil)

	file, err := os.OpenFile(path, os.O_RDWR, 0)
	c.Assert(err, check.IsNil)

	c.Assert(overwriteAndClose(file, true), check.IsNil)

	_, err = file.Stat()
	c.Assert(errors.Is(err, os.ErrClosed), check.Equals, true)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, make([]byte, len("sensitive data")))
}

func (s *FSSuite) TestOverwriteAndCloseOnFailure(c *check.C) {
	path := filepath.Join(c.MkDir(), "data")
	content := []byte("must remain unchanged")

	c.Assert(os.WriteFile(path, content, 0600), check.IsNil)

	// A read-only handle causes overwriting to fail even when running as root.
	file, err := os.Open(path)
	c.Assert(err, check.IsNil)

	c.Assert(overwriteAndClose(file, true), check.ErrorMatches, `random pass 1: .*`)

	_, err = file.Stat()
	c.Assert(errors.Is(err, os.ErrClosed), check.Equals, true)

	got, err := os.ReadFile(path)
	c.Assert(err, check.IsNil)
	c.Assert(got, check.DeepEquals, content)
}

func (s *FSSuite) TestOverwriteAndCloseClosedFile(c *check.C) {
	file, err := os.Create(filepath.Join(c.MkDir(), "data"))
	c.Assert(err, check.IsNil)
	c.Assert(file.Close(), check.IsNil)

	err = overwriteAndClose(file, false)

	c.Assert(err, check.NotNil)
	c.Assert(errors.Is(err, os.ErrClosed), check.Equals, true)

	// Both inspecting the closed handle and closing it again must be reported.
	c.Assert(err.Error(), check.Matches, `(?s)inspect opened file: .*\nclose file: .*`)
}
