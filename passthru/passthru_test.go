package passthru

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"io"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type PassthruSuite struct{}

type DummyReader struct {
	io.Reader
	setError bool
}

type DummyWriter struct {
	io.Writer
	setError bool
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&PassthruSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *PassthruSuite) TestNil(c *C) {
	var r *Reader

	_, err := r.Read([]byte{})
	c.Assert(err, Equals, ErrNilReader)
	c.Assert(r.Current(), Equals, int64(0))
	c.Assert(r.Total(), Equals, int64(0))
	c.Assert(r.Progress(), Equals, 0.0)

	rs, rr := r.Speed()
	c.Assert(rs, Equals, 0.0)
	c.Assert(rr, Equals, time.Duration(0))

	c.Assert(func() { r.SetTotal(1) }, NotPanics)

	var w *Writer

	_, err = w.Write([]byte{})
	c.Assert(err, Equals, ErrNilWriter)
	c.Assert(w.Current(), Equals, int64(0))
	c.Assert(w.Total(), Equals, int64(0))
	c.Assert(w.Progress(), Equals, 0.0)

	ws, wr := w.Speed()
	c.Assert(ws, Equals, 0.0)
	c.Assert(wr, Equals, time.Duration(0))

	c.Assert(func() { w.SetTotal(1) }, NotPanics)

	var cl *Calculator

	cs, cr := cl.Calculate(1)
	c.Assert(cs, Equals, 0.0)
	c.Assert(cr, Equals, time.Duration(0))

	c.Assert(func() { cl.SetTotal(1) }, NotPanics)
}

func (s *PassthruSuite) TestReader(c *C) {
	r := NewReader(&DummyReader{}, 1000)

	c.Assert(r, NotNil)

	r.Update = func(_ int) {}

	n, err := r.Read([]byte{})
	c.Assert(n, Equals, int(100))
	c.Assert(err, Not(Equals), ErrNilReader)

	c.Assert(r.Current(), Equals, int64(100))
	c.Assert(r.Total(), Equals, int64(1000))
	c.Assert(r.Progress(), Equals, 10.0)

	r.SetTotal(2000)
	c.Assert(r.Total(), Equals, int64(2000))

	rs, _ := r.Speed()

	c.Assert(rs, Not(Equals), 0.0)

	r = NewReader(&DummyReader{setError: true}, 1000)

	_, err = r.Read([]byte{})
	c.Assert(err, NotNil)
}

func (s *PassthruSuite) TestWriter(c *C) {
	w := NewWriter(&DummyWriter{}, 1000)

	c.Assert(w, NotNil)

	w.Update = func(_ int) {}

	n, err := w.Write([]byte{})
	c.Assert(n, Equals, int(100))
	c.Assert(err, Not(Equals), ErrNilWriter)

	c.Assert(w.Current(), Equals, int64(100))
	c.Assert(w.Total(), Equals, int64(1000))
	c.Assert(w.Progress(), Equals, 10.0)

	w.SetTotal(2000)
	c.Assert(w.Total(), Equals, int64(2000))

	ws, _ := w.Speed()

	c.Assert(ws, Not(Equals), 0.0)

	w = NewWriter(&DummyWriter{setError: true}, 1000)

	_, err = w.Write([]byte{})
	c.Assert(err, NotNil)
}

func (s *PassthruSuite) TestCalculator(c *C) {
	cl := NewCalculator(0, 1.0)
	c.Assert(cl, NotNil)

	cs, cr := cl.Calculate(1)
	c.Assert(cs, Equals, 0.0)
	c.Assert(cr, Equals, time.Duration(0))

	cl = NewCalculator(1000, 1.0)
	c.Assert(cl, NotNil)

	c.Assert(cl, NotNil)
	c.Assert(func() { cl.SetTotal(10000) }, NotPanics)
	c.Assert(cl.total, Equals, 10000.0)

	cl.Calculate(1)

	time.Sleep(500 * time.Millisecond)

	cl.Calculate(2)

	cs, cr = cl.Calculate(1)
	c.Assert(cs, Not(Equals), 0.0)
	c.Assert(cr, Not(Equals), time.Duration(0))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (r *DummyReader) Read(p []byte) (int, error) {
	if r.setError {
		return 100, errors.New("ERROR")
	}

	return 100, nil
}

func (r *DummyReader) Close() error {
	return nil
}

func (w *DummyWriter) Write(p []byte) (int, error) {
	if w.setError {
		return 100, errors.New("ERROR")
	}

	return 100, nil
}

func (w *DummyWriter) Close() error {
	return nil
}
