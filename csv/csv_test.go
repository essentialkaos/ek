package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io"
	"os"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _DATA = `ID,FIRST NAME,LAST NAME,BALANCE
1,John,Doe,0.34
2,Fiammetta,Miriana,30
3,Mathew,Timothy,34.19371,1,2
4,Lou
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

	err := os.WriteFile(s.dataFile, []byte(_DATA), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func (s *CSVSuite) TestRead(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

	defer fd.Close()

	line := 0
	reader := NewReader(fd).WithComma(',').WithHeaderSkip(true)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		switch line {
		case 0:
			c.Assert(row, HasLen, 4)
			c.Assert(row, DeepEquals, Row{"1", "John", "Doe", "0.34"})
		case 1:
			c.Assert(row, HasLen, 4)
			c.Assert(row, DeepEquals, Row{"2", "Fiammetta", "Miriana", "30"})
		case 2:
			c.Assert(row, HasLen, 6)
			c.Assert(row, DeepEquals, Row{"3", "Mathew", "Timothy", "34.19371", "1", "2"})
		}

		line++
	}
}

func (s *CSVSuite) TestReadTo(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

	defer fd.Close()

	line := 0
	reader := NewReader(fd).WithComma(',').WithHeaderSkip(true)

	var row Row

	for {
		err := reader.ReadTo(row)

		if row == nil {
			c.Assert(err, NotNil)
			row = make(Row, 4)
			continue
		}

		if err == io.EOF {
			break
		}

		switch line {
		case 0:
			c.Assert(row, DeepEquals, Row{"1", "John", "Doe", "0.34"})
		case 1:
			c.Assert(row, DeepEquals, Row{"2", "Fiammetta", "Miriana", "30"})
		case 2:
			c.Assert(row, DeepEquals, Row{"3", "Mathew", "Timothy", "34.19371"})
		case 3:
			c.Assert(row, DeepEquals, Row{"4", "Lou", "", ""})
		}

		line++
	}
}

func (s *CSVSuite) TestLineParser(c *C) {
	data := make(Row, 2)

	parseAndFill("ABCD", data, ";")
	c.Assert(data, DeepEquals, Row{"ABCD", ""})

	parseAndFill("", data, ";")
	c.Assert(data, DeepEquals, Row{"", ""})

	parseAndFill("A;B;C;D;E", data, ";")
	c.Assert(data, DeepEquals, Row{"A", "B"})
}

func (s *CSVSuite) TestNil(c *C) {
	var r *Reader
	var b = Row{""}

	_, err := r.Read()

	c.Assert(err, DeepEquals, ErrNilReader)
	c.Assert(r.ReadTo(b), DeepEquals, ErrNilReader)
	c.Assert(r.WithComma('X'), IsNil)
	c.Assert(r.WithHeaderSkip(false), IsNil)
}

func (s *CSVSuite) TestRow(c *C) {
	r := Row{"test", "", "1234", "12.34", "Yes"}
	c.Assert(r.Size(), Equals, 5)
	c.Assert(r.Cells(), Equals, 4)
	c.Assert(r.IsEmpty(), Equals, false)
	c.Assert(r.Has(0), Equals, true)
	c.Assert(r.Has(1), Equals, false)
	c.Assert(r.Get(0), Equals, "test")
	c.Assert(r.Get(10), Equals, "")

	ci, err := r.GetI(2)
	c.Assert(ci, Equals, 1234)
	c.Assert(err, IsNil)

	cf, err := r.GetF(3)
	c.Assert(cf, Equals, 12.34)
	c.Assert(err, IsNil)

	cu, err := r.GetU(2)
	c.Assert(cu, Equals, uint64(1234))
	c.Assert(err, IsNil)

	c.Assert(r.GetB(1), Equals, false)
	c.Assert(r.GetB(4), Equals, true)

	dummyFn := func(i int, v string) error { return nil }
	c.Assert(r.ForEach(dummyFn), IsNil)

	dummyFn = func(i int, v string) error { return errors.New("1") }
	c.Assert(r.ForEach(dummyFn), NotNil)

	c.Assert(r.ToString(';'), Equals, "test;;1234;12.34;Yes")
	c.Assert(string(r.ToBytes(';')), Equals, "test;;1234;12.34;Yes")
}

// ////////////////////////////////////////////////////////////////////////////////// //

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
