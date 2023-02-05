package events

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"sync/atomic"
	"testing"
	"time"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type EventsSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&EventsSuite{})

var counter uint32

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *EventsSuite) TestBasicErrors(c *C) {
	var d *Dispatcher

	c.Assert(d.AddHandler("test", basicTestHandler), NotNil)
	c.Assert(d.RemoveHandler("test", basicTestHandler), NotNil)
	c.Assert(d.Dispatch("test", nil), NotNil)
	c.Assert(d.HasHandler("test", nil), Equals, false)
	c.Assert(d.DispatchAndWait("test", nil), NotNil)

	c.Assert(validateArguments(d, "test", basicTestHandler, true), NotNil)

	d = NewDispatcher()

	c.Assert(d.AddHandler("test", basicTestHandler), IsNil)
	c.Assert(d.AddHandler("test", basicTestHandler), NotNil)
	c.Assert(d.RemoveHandler("test", basicTestHandler), IsNil)
	c.Assert(d.RemoveHandler("test", basicTestHandler), NotNil)

	c.Assert(d.Dispatch("unknown", basicTestHandler), NotNil)
	c.Assert(d.DispatchAndWait("unknown", basicTestHandler), NotNil)

	c.Assert(validateArguments(d, "", basicTestHandler, true), NotNil)
	c.Assert(validateArguments(d, "test", nil, true), NotNil)
}

func (s *EventsSuite) TestDispatch(c *C) {
	counter = 0

	d := NewDispatcher()

	c.Assert(d.AddHandler("test", asyncTestHandler1), IsNil)
	c.Assert(d.AddHandler("test", asyncTestHandler2), IsNil)
	c.Assert(d.HasHandler("test", asyncTestHandler1), Equals, true)
	c.Assert(d.HasHandler("test", asyncTestHandler2), Equals, true)

	c.Assert(d.Dispatch("test", uint32(3)), IsNil)

	time.Sleep(15 * time.Millisecond)

	c.Assert(atomic.LoadUint32(&counter), Equals, uint32(7))
}

func (s *EventsSuite) TestDispatchAndWait(c *C) {
	counter = 0

	d := NewDispatcher()

	c.Assert(d.AddHandler("test", asyncTestHandler1), IsNil)
	c.Assert(d.AddHandler("test", asyncTestHandler2), IsNil)
	c.Assert(d.HasHandler("test", asyncTestHandler1), Equals, true)
	c.Assert(d.HasHandler("test", asyncTestHandler2), Equals, true)

	c.Assert(d.DispatchAndWait("test", uint32(5)), IsNil)

	c.Assert(atomic.LoadUint32(&counter), Equals, uint32(11))
}

// ////////////////////////////////////////////////////////////////////////////////// //

func basicTestHandler(payload interface{}) {
	return
}

func asyncTestHandler1(payload interface{}) {
	atomic.AddUint32(&counter, payload.(uint32))
}

func asyncTestHandler2(payload interface{}) {
	atomic.AddUint32(&counter, payload.(uint32)+1)
}
