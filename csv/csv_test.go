package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	. "pkg.re/essentialkaos/check.v1"
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

func (s *CSVSuite) TestRead(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

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

func (s *CSVSuite) TestReadTo(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

	defer fd.Close()

	reader := NewReader(fd)
	reader.Comma = ','

	count := 0

	var rec []string

	for {
		err := reader.ReadTo(rec)

		if rec == nil {
			c.Assert(err, NotNil)
			rec = make([]string, 8)
			continue
		}

		if err == io.EOF {
			break
		}

		switch count {
		case 0:
			c.Assert(rec, DeepEquals, []string{"123", "ABC", "A_C", "A C", "", "", "", ""})
		case 1:
			c.Assert(rec, DeepEquals, []string{"123", "ABC", "", "", "", "", "", ""})
		case 2:
			c.Assert(rec, DeepEquals, []string{"123", "ABC", "A_C", "A C", "123", "ABC", "A_C", "A C"})
		}

		count++
	}
}

func (s *CSVSuite) TestLineParser(c *C) {
	data := make([]string, 2)

	parseAndFill("ABCD", data, ";")
	c.Assert(data, DeepEquals, []string{"ABCD", ""})

	parseAndFill("", data, ";")
	c.Assert(data, DeepEquals, []string{"", ""})

	parseAndFill("A;B;C;D;E", data, ";")
	c.Assert(data, DeepEquals, []string{"A", "B"})
}

func (s *CSVSuite) BenchmarkRead(c *C) {
	fd, _ := os.Open(s.dataFile)

	for i := 0; i < c.N; i++ {
		reader := NewReader(fd)
		reader.Comma = ','

		for {
			_, err := reader.Read()

			if err == io.EOF {
				break
			}
		}
	}

	fd.Close()
}

func (s *CSVSuite) BenchmarkReadTo(c *C) {
	fd, _ := os.Open(s.dataFile)

	for i := 0; i < c.N; i++ {
		reader := NewReader(fd)
		reader.Comma = ','

		k := make([]string, 10)

		for {
			err := reader.ReadTo(k)

			if err == io.EOF {
				break
			}
		}
	}

	fd.Close()
}
