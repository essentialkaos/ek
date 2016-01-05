package arg

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ArgUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ArgUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ArgUtilSuite) TestAdd(c *C) {
	args := &Arguments{}

	c.Assert(args.Add("", &V{}), NotNil)
	c.Assert(args.Add("t:", &V{}), NotNil)

	c.Assert(args.Add("test", &V{}), IsNil)
	c.Assert(args.Add(":test1", &V{}), IsNil)
	c.Assert(args.Add("t:test2", &V{}), IsNil)

	c.Assert(args.Add("t1:test", &V{}), NotNil)
	c.Assert(args.Add("t:test3", &V{}), NotNil)
	c.Assert(args.Add("t:test3", nil), NotNil)
}

func (s *ArgUtilSuite) TestAddMap(c *C) {
	args := NewArguments()

	m1 := Map{"": &V{}}
	m2 := Map{"t:": &V{}}

	c.Assert(args.AddMap(m1), Not(HasLen), 0)
	c.Assert(args.AddMap(m2), Not(HasLen), 0)

	m3 := Map{"test": &V{}}
	m4 := Map{"t:test2": &V{}}

	c.Assert(args.AddMap(m3), HasLen, 0)
	c.Assert(args.AddMap(m4), HasLen, 0)

	m5 := Map{
		"t:test":  &V{},
		"T:test1": &V{},
		"t:test2": &V{},
	}

	c.Assert(args.AddMap(m5), HasLen, 2)
}

func (s *ArgUtilSuite) TestGlobal(c *C) {
	long, short := ParseArgName("t:test")

	c.Assert(short, Equals, "t")
	c.Assert(long, Equals, "test")

	c.Assert(GetS("s:string"), Equals, "")
	c.Assert(GetI("i:int"), Equals, 0)
	c.Assert(GetF("f:float"), Equals, 0.0)
	c.Assert(GetB("b:bool"), Equals, false)
	c.Assert(Has("s:string"), Equals, false)

	c.Assert(Add("t:test", &V{}), IsNil)

	global = nil

	c.Assert(AddMap(Map{}), IsNil)

	global = nil

	Parse(Map{})

	global = NewArguments()

	_, errs := global.Parse(
		strings.Split("-s Test -i 123 -f 100.5 -b", " "),
		Map{
			"s:string": &V{},
			"i:int":    &V{Type: INT},
			"f:float":  &V{Type: FLOAT},
			"b:bool":   &V{Type: BOOL},
		},
	)

	c.Assert(errs, HasLen, 0)

	c.Assert(Has("s:string"), Equals, true)
	c.Assert(Has("string1"), Equals, false)
	c.Assert(GetS("s:string"), Equals, "Test")
	c.Assert(GetI("i:int"), Equals, 123)
	c.Assert(GetF("f:float"), Equals, 100.5)
	c.Assert(GetB("b:bool"), Equals, true)
}

func (s *ArgUtilSuite) TestGetters(c *C) {
	argline := "file.mp3 -s STRING --required TEST -i 320 -b -f 1.098765 -S2 100 -f1 5 -f2 1 -ms ABC --merg-string DEF -mi 6 --merg-int 6 -f3 12 -mf 10.1 -mf 10.1 -i1 5"

	argsMap := Map{
		"s:string":          &V{Type: STRING, Value: "STRING"},
		"S:empty-string":    &V{Type: STRING},
		"r:required":        &V{Required: true, Alias: "A:alias"},
		"i:int":             &V{Type: INT},
		"i1:int-between":    &V{Type: INT, Min: 1, Max: 3},
		"I:not-set-int":     &V{Type: INT, Value: 0},
		"b:bool":            &V{Type: BOOL},
		"B:empty-bool":      &V{Type: BOOL},
		"B1:not-set-bool":   &V{Type: BOOL, Value: false},
		"f:float":           &V{Type: FLOAT},
		"F:not-set-float":   &V{Type: FLOAT, Value: 0.0},
		"f1:float-max":      &V{Type: FLOAT, Max: 3.0},
		"f2:float-min":      &V{Type: FLOAT, Min: 3.0},
		"f3:float-between":  &V{Type: FLOAT, Min: 3.0, Max: 10.0},
		"ms:merg-string":    &V{Mergeble: true},
		"mi:merg-int":       &V{Type: INT, Mergeble: true},
		"mf:merg-float":     &V{Type: FLOAT, Mergeble: true},
		"S1:not-set-string": &V{Type: STRING, Value: ""},
		"S2:string-as-num":  &V{Type: STRING},
	}

	args := NewArguments()
	args.Parse(strings.Split(argline, " "), argsMap)

	c.Assert(args.GetS("_not_exist_"), Equals, "")
	c.Assert(args.GetI("_not_exist_"), Equals, 0)
	c.Assert(args.GetB("_not_exist_"), Equals, false)
	c.Assert(args.GetF("_not_exist_"), Equals, 0.0)

	c.Assert(args.GetS("s:string"), Equals, "STRING")
	c.Assert(args.GetS("string"), Equals, "STRING")
	c.Assert(args.GetS("S:empty-string"), Equals, "")
	c.Assert(args.GetS("empty-string"), Equals, "")
	c.Assert(args.GetB("string"), Equals, true)
	c.Assert(args.GetI("string"), Equals, 0)
	c.Assert(args.GetF("string"), Equals, 0.0)
	c.Assert(args.GetB("empty-string"), Equals, false)
	c.Assert(args.GetI("empty-string"), Equals, 0)
	c.Assert(args.GetF("empty-string"), Equals, 0.0)
	c.Assert(args.GetB("not-set-string"), Equals, false)
	c.Assert(args.GetI("not-set-string"), Equals, 0)
	c.Assert(args.GetF("not-set-string"), Equals, 0.0)
	c.Assert(args.GetS("S2:string-as-num"), Equals, "100")
	c.Assert(args.GetB("S2:string-as-num"), Equals, true)
	c.Assert(args.GetI("S2:string-as-num"), Equals, 100)
	c.Assert(args.GetF("S2:string-as-num"), Equals, 100.0)

	c.Assert(args.GetS("r:required"), Equals, "TEST")
	c.Assert(args.GetS("required"), Equals, "TEST")
	c.Assert(args.GetS("A:alias"), Equals, "TEST")
	c.Assert(args.GetS("alias"), Equals, "TEST")

	c.Assert(args.GetS("int"), Equals, "320")
	c.Assert(args.GetB("int"), Equals, true)
	c.Assert(args.GetI("int"), Equals, 320)
	c.Assert(args.GetF("int"), Equals, 320.0)
	c.Assert(args.GetI("int-between"), Equals, 3)
	c.Assert(args.GetS("not-set-int"), Equals, "0")
	c.Assert(args.GetB("not-set-int"), Equals, false)
	c.Assert(args.GetI("not-set-int"), Equals, 0)
	c.Assert(args.GetF("not-set-int"), Equals, 0.0)

	c.Assert(args.GetS("b:bool"), Equals, "true")
	c.Assert(args.GetI("b:bool"), Equals, 1)
	c.Assert(args.GetB("b:bool"), Equals, true)
	c.Assert(args.GetF("b:bool"), Equals, 1.0)
	c.Assert(args.GetS("empty-bool"), Equals, "")
	c.Assert(args.GetI("empty-bool"), Equals, 0)
	c.Assert(args.GetB("empty-bool"), Equals, false)
	c.Assert(args.GetF("empty-bool"), Equals, 0.0)
	c.Assert(args.GetS("not-set-bool"), Equals, "false")
	c.Assert(args.GetI("not-set-bool"), Equals, 0)
	c.Assert(args.GetB("not-set-bool"), Equals, false)
	c.Assert(args.GetF("not-set-bool"), Equals, 0.0)

	c.Assert(args.GetS("float"), Equals, "1.098765")
	c.Assert(args.GetB("float"), Equals, true)
	c.Assert(args.GetI("float"), Equals, 1)
	c.Assert(args.GetF("float"), Equals, 1.098765)
	c.Assert(args.GetS("not-set-float"), Equals, "0")
	c.Assert(args.GetB("not-set-float"), Equals, false)
	c.Assert(args.GetI("not-set-float"), Equals, 0)
	c.Assert(args.GetF("not-set-float"), Equals, 0.0)

	c.Assert(args.GetF("float-max"), Equals, 3.0)
	c.Assert(args.GetF("float-min"), Equals, 3.0)
	c.Assert(args.GetF("float-between"), Equals, 10.0)

	c.Assert(args.GetS("merg-string"), Equals, "ABC DEF")
	c.Assert(args.GetI("merg-int"), Equals, 12)
	c.Assert(args.GetF("merg-float"), Equals, 20.2)

	c.Assert(args.Has("_not_exist_"), Equals, false)
	c.Assert(args.Has("empty-string"), Equals, false)
	c.Assert(args.Has("s:string"), Equals, true)
}

func (s *ArgUtilSuite) TestParsing(c *C) {
	_, errs := NewArguments().Parse([]string{}, Map{})

	c.Assert(errs, HasLen, 0)

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"s:string": &V{}, "s:trace": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -s defined 2 or more times")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t"}, Map{"s:string": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t=100"}, Map{"s:string": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test"}, Map{"s:string": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	fArgs, errs := NewArguments().Parse([]string{"-", "--"}, Map{"t:test": &V{}})

	c.Assert(errs, HasLen, 0)
	c.Assert(fArgs, DeepEquals, []string{"-", "--"})

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--asd="}, Map{"t:test": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --asd has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-j="}, Map{"t:test": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -j has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"t:test": &V{Required: true}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Required argument test is not set")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t"}, Map{"t:test": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Non-boolean argument --test is empty")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test=!"}, Map{"t:test": &V{Type: FLOAT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t=!"}, Map{"t:test": &V{Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test", "!"}, Map{"t:test": &V{Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"t:test": nil})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Struct for argument --test is nil")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"": &V{}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Some argument does not have a name")
}
