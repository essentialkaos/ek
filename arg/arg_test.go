package arg

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ArgUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ArgUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ArgUtilSuite) TestAdd(c *C) {
	args := NewArguments()

	c.Assert(args.Add("", &V{}), NotNil)
	c.Assert(args.Add("t:", &V{}), NotNil)

	c.Assert(args.Add("test", &V{}), IsNil)
	c.Assert(args.Add(":test1", &V{}), IsNil)
	c.Assert(args.Add("t:test2", &V{}), IsNil)

	c.Assert(args.Add("t1:test", &V{}), NotNil)
	c.Assert(args.Add("t:test3", &V{}), NotNil)
}

func (s *ArgUtilSuite) TestAddMap(c *C) {
	args := NewArguments()

	m1 := map[string]*V{"": &V{}}
	m2 := map[string]*V{"t:": &V{}}

	c.Assert(args.AddMap(m1), Not(HasLen), 0)
	c.Assert(args.AddMap(m2), Not(HasLen), 0)

	m3 := map[string]*V{"test": &V{}}
	m4 := map[string]*V{"t:test2": &V{}}

	c.Assert(args.AddMap(m3), HasLen, 0)
	c.Assert(args.AddMap(m4), HasLen, 0)

	m5 := map[string]*V{
		"t:test":  &V{},
		"T:test1": &V{},
		"t:test2": &V{},
	}

	c.Assert(args.AddMap(m5), HasLen, 2)
}

func (s *ArgUtilSuite) TestParsing(c *C) {
	argline := "file.mp3 -f mp3 --bitrate 320 --encode -S 1.52 -S1 5.0 -S2 1.0 --name Dj -n Super --name Star -s 10 -s 5"

	argsMap := map[string]*V{
		"a:author":  &V{Type: STRING, Value: "User"},
		"f:format":  &V{Required: true, Alias: "F:form"},
		"b:bitrate": &V{Type: INT},
		"E:encode":  &V{Type: BOOL},
		"D:decode":  &V{Type: BOOL},
		"S:scale":   &V{Type: FLOAT},
		"S1:scale1": &V{Type: FLOAT, Max: 3.0},
		"S2:scale2": &V{Type: FLOAT, Min: 3.0},
		"n:name":    &V{Mergeble: true},
		"s:summ":    &V{Type: INT, Mergeble: true},
	}

	args := NewArguments()
	args.Parse(strings.Split(argline, " "), argsMap)

	c.Assert(args.GetS("a:author"), Equals, "User")
	c.Assert(args.GetS("author"), Equals, "User")

	c.Assert(args.GetS("f:format"), Equals, "mp3")
	c.Assert(args.GetS("format"), Equals, "mp3")
	c.Assert(args.GetS("F:form"), Equals, "mp3")
	c.Assert(args.GetS("form"), Equals, "mp3")

	c.Assert(args.GetS("bitrate"), Equals, "320")
	c.Assert(args.GetI("bitrate"), Equals, 320)
	c.Assert(args.GetF("bitrate"), Equals, 320.0)

	c.Assert(args.GetS("encode"), Equals, "true")
	c.Assert(args.GetI("encode"), Equals, 1)
	c.Assert(args.GetB("encode"), Equals, true)
	c.Assert(args.GetF("encode"), Equals, 1.0)

	c.Assert(args.GetS("decode"), Equals, "")
	c.Assert(args.GetI("decode"), Equals, -1)
	c.Assert(args.GetB("decode"), Equals, false)
	c.Assert(args.GetF("decode"), Equals, -1.0)

	c.Assert(args.GetS("scale"), Equals, "1.52")
	c.Assert(args.GetI("scale"), Equals, 1)
	c.Assert(args.GetF("scale"), Equals, 1.52)

	c.Assert(args.GetF("scale1"), Equals, 3.0)
	c.Assert(args.GetF("scale2"), Equals, 3.0)

	c.Assert(args.GetS("name"), Equals, "Dj Super Star")
	c.Assert(args.GetI("summ"), Equals, 15)
}
