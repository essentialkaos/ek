package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _DATA = `123,ABC,A_C,A C,
123,ABC,
123,ABC,A_C,A C,123,ABC,A_C,A C
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type CSVSuite struct {
	dataFile string
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&CSVSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CSVSuite) SetUpSuite(c *C) {
	tmpDir := c.MkDir()

	s.dataFile = tmpDir + "/data.csv"

	err := ioutil.WriteFile(s.dataFile, []byte(_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *CSVSuite) TestParsing(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(fd, NotNil)
	c.Assert(err, IsNil)

	defer fd.Close()

	reader := NewReader(fd)
	reader.Comma = ','

	count := 0

	for {
		rec, err := reader.Read()

		if err == io.EOF {
			break
		}

		switch count {
		case 0:
			c.Assert(rec, HasLen, 5)
			c.Assert(rec, DeepEquals, []string{"123", "ABC", "A_C", "A C", ""})
		case 1:
			c.Assert(rec, HasLen, 3)
			c.Assert(rec, DeepEquals, []string{"123", "ABC", ""})
		case 2:
			c.Assert(rec, HasLen, 8)
			c.Assert(rec, DeepEquals, []string{"123", "ABC", "A_C", "A C", "123", "ABC", "A_C", "A C"})
		}

		count++
	}
}
