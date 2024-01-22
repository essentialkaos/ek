package progress

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ProgressSuite struct{}

type DummyReader struct {
	io.Reader
}

type DummyWriter struct {
	io.Writer
}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ProgressSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ProgressSuite) TestBar(c *C) {
	pb := New(-1, "ABCD")

	c.Assert(pb, NotNil)

	pbs := DefaultSettings
	pbs.Width = 60
	pbs.RefreshRate = time.Microsecond

	err := pb.UpdateSettings(pbs)
	c.Assert(err, NotNil)

	pbs.RefreshRate = time.Millisecond

	err = pb.UpdateSettings(pbs)
	c.Assert(err, IsNil)

	pb.SetName("ABC")
	c.Assert(pb.Name(), Equals, "ABC")

	pb.SetTotal(2000)
	c.Assert(pb.Total(), Equals, int64(2000))
	c.Assert(pb.passThruCalc, NotNil)
	pb.SetTotal(2200)
	c.Assert(pb.Total(), Equals, int64(2200))

	pb.SetCurrent(100)
	c.Assert(pb.Current(), Equals, int64(100))

	pb.Add(1)
	c.Assert(pb.Current(), Equals, int64(101))

	pb.Add64(1)
	c.Assert(pb.Current(), Equals, int64(102))

	c.Assert(pb.IsFinished(), Equals, false)
	c.Assert(pb.IsStarted(), Equals, false)

	pb.Reader(&DummyReader{Reader: nil})
	c.Assert(pb.reader, NotNil)

	pb.Writer(&DummyWriter{Writer: nil})
	c.Assert(pb.writer, NotNil)

	pr := pb.Reader(&DummyReader{Reader: nil})
	pw := pb.Writer(&DummyWriter{Writer: nil})

	pr.Read(nil)
	c.Assert(pb.Current(), Equals, int64(202))

	pw.Write(nil)
	c.Assert(pb.Current(), Equals, int64(302))

	pb.Finish() // should be skipped (not started)
	c.Assert(pb.IsFinished(), Equals, false)

	pb.Start()
	pb.renderElements(false)

	c.Assert(pb.IsStarted(), Equals, true)
	pb.Start() // should be skipped (already started)
	pb.Finish()

	time.Sleep(time.Millisecond * 10)

	pb = New(100, "ABCD")
	pb.settings.RefreshRate = time.Millisecond
	pb.Start()
	pb.SetCurrent(1000)
	time.Sleep(time.Millisecond * 100)
	pb.Finish()
}

func (s *ProgressSuite) TestSettingsValidation(c *C) {
	pbs := DefaultSettings

	c.Assert(pbs.Validate(), IsNil)

	pbs.RefreshRate = 2
	c.Assert(pbs.Validate().Error(), Equals, "RefreshRate too small (less than 1ms)")

	pbs.RemainingColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "RemainingColorTag value is not a valid color tag")

	pbs.SpeedColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "SpeedColorTag value is not a valid color tag")

	pbs.ProgressColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "ProgressColorTag value is not a valid color tag")

	pbs.PercentColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "PercentColorTag value is not a valid color tag")

	pbs.BarBgColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "BarBgColorTag value is not a valid color tag")

	pbs.BarFgColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "BarFgColorTag value is not a valid color tag")

	pbs.NameColorTag = "{ABCD}"
	c.Assert(pbs.Validate().Error(), Equals, "NameColorTag value is not a valid color tag")
}

func (s *ProgressSuite) TestBarRender(c *C) {
	pb := New(100, "ABCD")
	c.Assert(pb, NotNil)

	pb.current = 50

	vl := pb.renderBar(80)
	c.Assert(vl, Equals, "{r}————{s-}————{!}")

	vl = pb.renderBar(100)
	c.Assert(vl, Equals, "{r}——{s-}———{!}")

	pb.settings.BarFgColorTag = ""

	vl = pb.renderBar(80)
	c.Assert(vl, Equals, "————    ")

	pb.current = 150
	pb.settings.BarFgColorTag = "{r}"

	vl = pb.renderBar(80)
	c.Assert(vl, Equals, "{r}————————{!}")

	pb.settings.BarFgColorTag = ""

	vl = pb.renderBar(80)
	c.Assert(vl, Equals, "————————")
}

func (s *ProgressSuite) TestBarPlaceholderRender(c *C) {
	pb := New(-1, "ABCD")
	c.Assert(pb, NotNil)

	vl := pb.renderBar(80)
	c.Assert(vl, Equals, "{r}—{s-}—{s-}—{r}—{s-}—{s-}—{r}—{s-}—{!}")

	pb.settings.BarFgColorTag = ""

	vl = pb.renderBar(80)
	c.Assert(vl, Equals, " —  —  —")
	vl = pb.renderBar(80)
	c.Assert(vl, Equals, "  —  —  ")
}

func (s *ProgressSuite) TestNameRender(c *C) {
	pb := New(1000, "ABCD")
	c.Assert(pb, NotNil)

	vl, sz := pb.renderName()

	c.Assert(vl, Equals, "{b}ABCD{!}")
	c.Assert(sz, Equals, 4)

	pb.settings.NameSize = 10

	vl, sz = pb.renderName()

	c.Assert(vl, Equals, "{b}      ABCD{!}")
	c.Assert(sz, Equals, 10)

	pb.settings.NameColorTag = ""

	vl, sz = pb.renderName()

	c.Assert(vl, Equals, "      ABCD")
	c.Assert(sz, Equals, 10)
}

func (s *ProgressSuite) TestPercRender(c *C) {
	pb := New(1000, "ABCD")
	c.Assert(pb, NotNil)

	pb.current = 123

	vl, sz := pb.renderPercentage()

	c.Assert(vl, Equals, "{m} 12.3%{!}")
	c.Assert(sz, Equals, 6)

	pb.total = -1

	vl, sz = pb.renderPercentage()

	c.Assert(vl, Equals, "{m}  0.0%{!}")
	c.Assert(sz, Equals, 6)

	pb.total = 100
	pb.current = 1000

	vl, sz = pb.renderPercentage()

	c.Assert(vl, Equals, "{m}100%{!}")
	c.Assert(sz, Equals, 4)

	pb.settings.PercentColorTag = ""

	vl, sz = pb.renderPercentage()

	c.Assert(vl, Equals, "100%")
	c.Assert(sz, Equals, 4)
}

func (s *ProgressSuite) TestProgressRender(c *C) {
	pb := New(10000000, "ABCD")
	c.Assert(pb, NotNil)

	pb.current = 123456

	vl, sz := pb.renderProgress()

	c.Assert(vl, Equals, "{g}0.1/9.5 MB{!}")
	c.Assert(sz, Equals, 10)

	pb.settings.IsSize = false

	vl, sz = pb.renderProgress()

	c.Assert(vl, Equals, "{g} 0.1/10.0M{!}")
	c.Assert(sz, Equals, 10)

	pb.settings.ProgressColorTag = ""

	vl, sz = pb.renderProgress()

	c.Assert(vl, Equals, " 0.1/10.0M")
	c.Assert(sz, Equals, 10)

	pb.current = 1
	pb.total = 10

	vl, sz = pb.renderProgress()

	c.Assert(vl, Equals, " 1/10")
	c.Assert(sz, Equals, 5)
}

func (s *ProgressSuite) TestSpeedRender(c *C) {
	pb := New(10000000, "ABCD")
	c.Assert(pb, NotNil)

	vl, sz := pb.renderSpeed(123456.0)

	c.Assert(vl, Equals, "{r} 120.6 KB/s{!}")
	c.Assert(sz, Equals, 11)

	pb.settings.IsSize = false

	vl, sz = pb.renderSpeed(123456.0)

	c.Assert(vl, Equals, "{r} 123.5K/s{!}")
	c.Assert(sz, Equals, 9)

	pb.settings.SpeedColorTag = ""

	vl, sz = pb.renderSpeed(123456.0)

	c.Assert(vl, Equals, " 123.5K/s")
	c.Assert(sz, Equals, 9)
}

func (s *ProgressSuite) TestRemainingRender(c *C) {
	pb := New(10000000, "ABCD")
	c.Assert(pb, NotNil)

	vl, sz := pb.renderRemaining(8 * time.Second)

	c.Assert(vl, Equals, "{c} 0:08{!}")
	c.Assert(sz, Equals, 5)

	vl, sz = pb.renderRemaining(148 * time.Second)

	c.Assert(vl, Equals, "{c} 2:28{!}")
	c.Assert(sz, Equals, 5)

	pb.settings.RemainingColorTag = ""

	vl, sz = pb.renderRemaining(148 * time.Second)

	c.Assert(vl, Equals, " 2:28")
	c.Assert(sz, Equals, 5)
}

func (s *ProgressSuite) TestAux(c *C) {
	ct, tt, lt := getPrettyCTSize(1, 15)

	c.Assert(ct, Equals, "1.0")
	c.Assert(tt, Equals, "15.0")
	c.Assert(lt, Equals, " B")

	ct, tt, lt = getPrettyCTSize(123, 15*1024)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "15.0")
	c.Assert(lt, Equals, " KB")

	ct, tt, lt = getPrettyCTSize(123123, 16*1024*1024)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "16.0")
	c.Assert(lt, Equals, " MB")

	ct, tt, lt = getPrettyCTSize(123123123, 17*1024*1024*1024)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "17.0")
	c.Assert(lt, Equals, " GB")

	ct, tt, lt = getPrettyCTNum(1, 15)

	c.Assert(ct, Equals, "1")
	c.Assert(tt, Equals, "15")
	c.Assert(lt, Equals, "")

	ct, tt, lt = getPrettyCTNum(123, 15*1000)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "15.0")
	c.Assert(lt, Equals, "K")

	ct, tt, lt = getPrettyCTNum(123123, 16*1000*1000)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "16.0")
	c.Assert(lt, Equals, "M")

	ct, tt, lt = getPrettyCTNum(123123123, 17*1000*1000*1000)

	c.Assert(ct, Equals, "0.1")
	c.Assert(tt, Equals, "17.0")
	c.Assert(lt, Equals, "B")

	c.Assert(formatSpeedNum(123.0), Equals, "   123/s")
	c.Assert(formatSpeedNum(123.0*1000.0), Equals, "   123K/s")
	c.Assert(formatSpeedNum(123.0*1000.0*1000.0), Equals, "   123M/s")
	c.Assert(formatSpeedNum(123.0*1000.0*1000.0*1000.0), Equals, "   123B/s")
}

func (s *ProgressSuite) TestNil(c *C) {
	var pb *Bar

	c.Assert(func() { pb.Start() }, NotPanics)
	c.Assert(func() { pb.Finish() }, NotPanics)
	c.Assert(func() { pb.UpdateSettings(Settings{}) }, NotPanics)
	c.Assert(func() { pb.SetName("test") }, NotPanics)
	c.Assert(func() { pb.SetTotal(100.0) }, NotPanics)
	c.Assert(func() { pb.SetCurrent(100.0) }, NotPanics)
	c.Assert(func() { pb.Add(1) }, NotPanics)
	c.Assert(func() { pb.Add64(1) }, NotPanics)

	c.Assert(pb.Name(), Equals, "")
	c.Assert(pb.Total(), Equals, int64(0))
	c.Assert(pb.Current(), Equals, int64(0))
	c.Assert(pb.IsFinished(), Equals, false)
	c.Assert(pb.IsStarted(), Equals, false)
	c.Assert(pb.Reader(nil), Equals, nil)
	c.Assert(pb.Writer(nil), Equals, nil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (r *DummyReader) Read(p []byte) (int, error) {
	return 100, nil
}

func (r *DummyReader) Close() error {
	return nil
}

func (w *DummyWriter) Write(p []byte) (int, error) {
	return 100, nil
}

func (r *DummyWriter) Close() error {
	return nil
}
