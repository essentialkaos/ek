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
	r := NewReader(fd, ',').WithHeader(true)
	m := map[string]string{}

	for {
		rr, err := r.Read()

		if err == io.EOF {
			break
		}

		switch line {
		case 0:
			c.Assert(rr, HasLen, 4)
			c.Assert(rr, DeepEquals, Row{"1", "John", "Doe", "0.34"})
			c.Assert(r.Header, DeepEquals, Header{"ID", "FIRST NAME", "LAST NAME", "BALANCE"})
			c.Assert(r.Header.Map(m, rr), IsNil)
			c.Assert(m["ID"], Equals, "1")
			c.Assert(m["FIRST NAME"], Equals, "John")
			c.Assert(m["LAST NAME"], Equals, "Doe")
			c.Assert(m["BALANCE"], Equals, "0.34")
		case 1:
			c.Assert(rr, HasLen, 4)
			c.Assert(rr, DeepEquals, Row{"2", "Fiammetta", "Miriana", "30"})
		case 2:
			c.Assert(rr, HasLen, 6)
			c.Assert(rr, DeepEquals, Row{"3", "Mathew", "Timothy", "34.19371", "1", "2"})
		}

		line++
	}

	c.Assert(r.Line(), Equals, 5)
}

func (s *CSVSuite) TestReadTo(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

	defer fd.Close()

	line := 0
	r := NewReader(fd, ',').WithHeader(true)

	var rr Row

	for {
		err := r.ReadTo(rr)

		if rr == nil {
			c.Assert(err, NotNil)
			rr = make(Row, 4)
			continue
		}

		if err == io.EOF {
			break
		}

		switch line {
		case 0:
			c.Assert(rr, DeepEquals, Row{"1", "John", "Doe", "0.34"})
		case 1:
			c.Assert(rr, DeepEquals, Row{"2", "Fiammetta", "Miriana", "30"})
		case 2:
			c.Assert(rr, DeepEquals, Row{"3", "Mathew", "Timothy", "34.19371"})
		case 3:
			c.Assert(rr, DeepEquals, Row{"4", "Lou", "", ""})
		}

		line++
	}

	c.Assert(r.Line(), Equals, 5)
}

func (s *CSVSuite) TestSeq(c *C) {
	fd, err := os.Open(s.dataFile)

	c.Assert(err, IsNil)
	c.Assert(fd, NotNil)

	defer fd.Close()

	r := NewReader(fd, ',').WithHeader(true)

	for line, rr := range r.Seq {
		switch line {
		case 2:
			c.Assert(rr, DeepEquals, Row{"1", "John", "Doe", "0.34"})
		case 3:
			c.Assert(rr, DeepEquals, Row{"2", "Fiammetta", "Miriana", "30"})
		case 4:
			c.Assert(rr, DeepEquals, Row{"3", "Mathew", "Timothy", "34.19371", "1", "2"})
		case 5:
			c.Assert(rr, DeepEquals, Row{"4", "Lou"})
		}
	}

	c.Assert(r.Error(), IsNil)
}

func (s *CSVSuite) TestReadErrors(c *C) {
	var fd *os.File
	r := NewReader(fd, ',').WithHeader(true)

	_, err := r.Read()
	c.Assert(err, NotNil)

	row := make(Row, 4)

	err = r.ReadTo(row)
	c.Assert(err, NotNil)

	var h Header
	var m map[string]string

	rr := Row{"1", "John", "Doe", "0.34"}

	c.Assert(h.Map(m, rr), Equals, ErrEmptyHeader)

	h = Header{"ID", "FIRST NAME", "LAST NAME", "BALANCE"}

	c.Assert(h.Map(m, rr), Equals, ErrNilMap)
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
	c.Assert(r.WithHeader(false), IsNil)
	c.Assert(r.Line(), Equals, 0)
	c.Assert(r.Error(), IsNil)

	r.Seq(nil)
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
		r := NewReader(fd, ',')

		for {
			_, err := r.Read()

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
		r := NewReader(fd, ',')
		k := make([]string, 10)

		for {
			err := r.ReadTo(k)

			if err == io.EOF {
				break
			}
		}
	}

	fd.Close()
}
