package csv

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
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

	c.Assert(reader.Line(), Equals, 5)
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

	c.Assert(reader.Line(), Equals, 5)
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
	c.Assert(r.Line(), Equals, 0)
}

func (s *CSVSuite) TestRow(c *C) {
	r := Row{"test", "", "33", "12.34", "Yes"}
	c.Assert(r.Size(), Equals, 5)
	c.Assert(r.Cells(), Equals, 4)
	c.Assert(r.IsEmpty(), Equals, false)
	c.Assert(r.Has(0), Equals, true)
	c.Assert(r.Has(1), Equals, false)
	c.Assert(r.Get(0), Equals, "test")
	c.Assert(r.Get(10), Equals, "")

	i, err := r.GetI(2)
	c.Assert(i, Equals, 33)
	c.Assert(err, IsNil)
	i8, err := r.GetI8(2)
	c.Assert(i8, Equals, int8(33))
	c.Assert(err, IsNil)
	i16, err := r.GetI16(2)
	c.Assert(i16, Equals, int16(33))
	c.Assert(err, IsNil)
	i32, err := r.GetI32(2)
	c.Assert(i32, Equals, int32(33))
	c.Assert(err, IsNil)
	i64, err := r.GetI64(2)
	c.Assert(i64, Equals, int64(33))
	c.Assert(err, IsNil)

	f64, err := r.GetF(3)
	c.Assert(f64, Equals, 12.34)
	c.Assert(err, IsNil)
	f32, err := r.GetF32(3)
	c.Assert(f32, Equals, float32(12.34))
	c.Assert(err, IsNil)

	u, err := r.GetU(2)
	c.Assert(u, Equals, uint(33))
	c.Assert(err, IsNil)
	u8, err := r.GetU8(2)
	c.Assert(u8, Equals, uint8(33))
	c.Assert(err, IsNil)
	u16, err := r.GetU16(2)
	c.Assert(u16, Equals, uint16(33))
	c.Assert(err, IsNil)
	u32, err := r.GetU32(2)
	c.Assert(u32, Equals, uint32(33))
	c.Assert(err, IsNil)
	u64, err := r.GetU64(2)
	c.Assert(u64, Equals, uint64(33))
	c.Assert(err, IsNil)

	c.Assert(r.GetB(1), Equals, false)
	c.Assert(r.GetB(4), Equals, true)

	dummyFn := func(i int, v string) error { return nil }
	c.Assert(r.ForEach(dummyFn), IsNil)

	dummyFn = func(i int, v string) error { return errors.New("1") }
	c.Assert(r.ForEach(dummyFn), NotNil)

	c.Assert(r.ToString(';'), Equals, "test;;33;12.34;Yes")
	c.Assert(string(r.ToBytes(';')), Equals, "test;;33;12.34;Yes")

	_, err = r.GetI8(0)
	c.Assert(err, NotNil)
	_, err = r.GetI16(0)
	c.Assert(err, NotNil)
	_, err = r.GetI32(0)
	c.Assert(err, NotNil)
	_, err = r.GetI64(0)
	c.Assert(err, NotNil)
	_, err = r.GetU(0)
	c.Assert(err, NotNil)
	_, err = r.GetU8(0)
	c.Assert(err, NotNil)
	_, err = r.GetU16(0)
	c.Assert(err, NotNil)
	_, err = r.GetU32(0)
	c.Assert(err, NotNil)
	_, err = r.GetU64(0)
	c.Assert(err, NotNil)
	_, err = r.GetF32(0)
	c.Assert(err, NotNil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *CSVSuite) BenchmarkRead(c *C) {
	fd, _ := os.Open(s.dataFile)

	for range c.N {
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

	for range c.N {
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
