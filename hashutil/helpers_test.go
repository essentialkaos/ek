package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type HashUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&HashUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *HashUtilSuite) TestCopy(c *C) {
	src := &bytes.Buffer{}
	dst := &bytes.Buffer{}

	src.WriteString("TEST1234!")

	_, _, err := Copy(nil, src, sha256.New())
	c.Assert(err, Equals, ErrNilDest)

	_, _, err = Copy(dst, nil, sha256.New())
	c.Assert(err, Equals, ErrNilSource)

	_, _, err = Copy(dst, src, nil)
	c.Assert(err, Equals, ErrNilHasher)

	n, hash, err := Copy(dst, src, sha256.New())
	c.Assert(err, IsNil)
	c.Assert(n, Equals, int64(9))
	c.Assert(hash, Equals, "ffa3f7c87408925c3acee34a21a4b6b5141f3a3948f144abd3c4489f086b6cb4")
}

func (s *HashUtilSuite) TestFileHash(c *C) {
	testFile := c.MkDir() + "/test.log"

	err := os.WriteFile(testFile, []byte("ABCDEF12345\n\n"), 0644)
	c.Assert(err, IsNil)

	c.Assert(File("_unknown_", sha1.New()), Equals, "")
	c.Assert(File(testFile, nil), Equals, "")
	c.Assert(File(testFile, sha1.New()), Equals, "9267257cafff1df7a8c0dea354d71c7221d17eda")
	c.Assert(File(testFile, sha256.New()), Equals, "2d7ec20906125cd23fee7b628b98463d554b1105b141b2d39a19bac5f3274dec")
}

func (s *HashUtilSuite) TestDataHash(c *C) {
	c.Assert(Bytes([]byte(""), sha1.New()), Equals, "")
	c.Assert(Bytes([]byte("TEST1234!"), nil), Equals, "")
	c.Assert(Bytes([]byte("TEST1234!"), sha256.New()), Equals, "ffa3f7c87408925c3acee34a21a4b6b5141f3a3948f144abd3c4489f086b6cb4")
}

func (s *HashUtilSuite) TestString(c *C) {
	c.Assert(String("", sha1.New()), Equals, "")
	c.Assert(String("TEST1234!", nil), Equals, "")
	c.Assert(String("TEST1234!", sha256.New()), Equals, "ffa3f7c87408925c3acee34a21a4b6b5141f3a3948f144abd3c4489f086b6cb4")
}

func (s *HashUtilSuite) TestSum(c *C) {
	c.Assert(Sum(nil), Equals, "")
}

func (s *HashUtilSuite) TestReader(c *C) {
	buf := bytes.NewBufferString("TEST1234!")

	r, err := NewReader(buf, sha256.New())
	c.Assert(err, IsNil)

	data, err := io.ReadAll(r)
	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, "TEST1234!")
	c.Assert(r.Sum(), Equals, "ffa3f7c87408925c3acee34a21a4b6b5141f3a3948f144abd3c4489f086b6cb4")
}

func (s *HashUtilSuite) TestReaderErrors(c *C) {
	buf := &bytes.Buffer{}

	_, err := NewReader(nil, sha256.New())
	c.Assert(err, Equals, ErrNilReader)

	_, err = NewReader(buf, nil)
	c.Assert(err, Equals, ErrNilHasher)

	var r1 *Reader
	r2 := &Reader{}
	r3 := &Reader{r: buf}

	_, err = r1.Read([]byte{})
	c.Assert(err, Equals, ErrNilReader)

	_, err = r2.Read([]byte{})
	c.Assert(err, Equals, ErrNilReader)

	_, err = r3.Read([]byte{})
	c.Assert(err, Equals, ErrNilHasher)

	c.Assert(r1.Sum(), Equals, "")
	c.Assert(r2.Sum(), Equals, "")
	c.Assert(r3.Sum(), Equals, "")
}

func (s *HashUtilSuite) TestWriter(c *C) {
	buf := &bytes.Buffer{}

	w, err := NewWriter(buf, sha256.New())
	c.Assert(err, IsNil)

	_, err = w.Write([]byte("TEST1234!"))
	c.Assert(err, IsNil)

	c.Assert(buf.String(), Equals, "TEST1234!")
	c.Assert(w.Sum(), Equals, "ffa3f7c87408925c3acee34a21a4b6b5141f3a3948f144abd3c4489f086b6cb4")
}

func (s *HashUtilSuite) TestWriterErrors(c *C) {
	buf := &bytes.Buffer{}

	_, err := NewWriter(nil, sha256.New())
	c.Assert(err, Equals, ErrNilWriter)

	_, err = NewWriter(buf, nil)
	c.Assert(err, Equals, ErrNilHasher)

	var w1 *Writer
	w2 := &Writer{}
	w3 := &Writer{w: buf}

	_, err = w1.Write([]byte("TEST1234!"))
	c.Assert(err, Equals, ErrNilWriter)

	_, err = w2.Write([]byte("TEST1234!"))
	c.Assert(err, Equals, ErrNilWriter)

	_, err = w3.Write([]byte("TEST1234!"))
	c.Assert(err, Equals, ErrNilHasher)

	c.Assert(w1.Sum(), Equals, "")
	c.Assert(w2.Sum(), Equals, "")
	c.Assert(w3.Sum(), Equals, "")
}
