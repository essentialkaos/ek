package arg

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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

	m1 := Map{"": {}}
	m2 := Map{"t:": {}}

	c.Assert(args.AddMap(m1), Not(HasLen), 0)
	c.Assert(args.AddMap(m2), Not(HasLen), 0)

	m3 := Map{"test": {}}
	m4 := Map{"t:test2": {}}

	c.Assert(args.AddMap(m3), HasLen, 0)
	c.Assert(args.AddMap(m4), HasLen, 0)

	m5 := Map{
		"t:test":  {},
		"T:test1": {},
		"t:test2": {},
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
			"s:string": {},
			"i:int":    {Type: INT},
			"f:float":  {Type: FLOAT},
			"b:bool":   {Type: BOOL},
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

func (s *ArgUtilSuite) TestLimiters(c *C) {
	argline := "--int1 1 --int2 5 --int3 10 --float1 1.0 --float2 5.0 --float3 10.0"

	argsMap := Map{
		"int1":   {Type: INT, Min: 3, Max: 7},
		"int2":   {Type: INT, Min: 3, Max: 7},
		"int3":   {Type: INT, Min: 3, Max: 7},
		"float1": {Type: FLOAT, Min: 3.0, Max: 7.0},
		"float2": {Type: FLOAT, Min: 3.0, Max: 7.0},
		"float3": {Type: FLOAT, Min: 3.0, Max: 7.0},
	}

	args := NewArguments()
	args.Parse(strings.Split(argline, " "), argsMap)

	c.Assert(args.GetI("int1"), Equals, 3)
	c.Assert(args.GetI("int2"), Equals, 5)
	c.Assert(args.GetI("int3"), Equals, 7)
	c.Assert(args.GetF("float1"), Equals, 3.0)
	c.Assert(args.GetF("float2"), Equals, 5.0)
	c.Assert(args.GetF("float3"), Equals, 7.0)
}

func (s *ArgUtilSuite) TestConflicts(c *C) {
	argline := "--test1 abc --test2 123"

	argsMap := Map{
		"test1": {Conflicts: "test2"},
		"test2": {},
	}

	args := NewArguments()
	_, errs := args.Parse(strings.Split(argline, " "), argsMap)

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].(ArgumentError).Type, Equals, ERROR_CONFLICT)
	c.Assert(errs[0].Error(), Equals, "Argument test1 conflicts with argument test2")
}

func (s *ArgUtilSuite) TestBound(c *C) {
	argline := "--test1 abc"

	argsMap := Map{
		"test1": {Bound: "test2"},
		"test2": {},
	}

	args := NewArguments()
	_, errs := args.Parse(strings.Split(argline, " "), argsMap)

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].(ArgumentError).Type, Equals, ERROR_BOUND_NOT_SET)
	c.Assert(errs[0].Error(), Equals, "Argument test2 must be defined with argument test1")
}

func (s *ArgUtilSuite) TestGetters(c *C) {
	argline := "file.mp3 -s STRING --required TEST -i 320 -b -f 1.098765 -S2 100 -f1 5 -f2 1 -ms ABC --merg-string DEF -mi 6 --merg-int 6 -f3 12 -mf 10.1 -mf 10.1 -i1 5"

	argsMap := Map{
		"s:string":          {Type: STRING, Value: "STRING"},
		"S:empty-string":    {Type: STRING},
		"r:required":        {Required: true, Alias: "A:alias"},
		"i:int":             {Type: INT},
		"i1:int-between":    {Type: INT, Min: 1, Max: 3},
		"I:not-set-int":     {Type: INT, Value: 0},
		"b:bool":            {Type: BOOL},
		"B:empty-bool":      {Type: BOOL},
		"B1:not-set-bool":   {Type: BOOL, Value: false},
		"f:float":           {Type: FLOAT},
		"F:not-set-float":   {Type: FLOAT, Value: 0.0},
		"f1:float-max":      {Type: FLOAT, Max: 3.0},
		"f2:float-min":      {Type: FLOAT, Min: 3.0},
		"f3:float-between":  {Type: FLOAT, Min: 3.0, Max: 10.0},
		"ms:merg-string":    {Mergeble: true},
		"mi:merg-int":       {Type: INT, Mergeble: true},
		"mf:merg-float":     {Type: FLOAT, Mergeble: true},
		"S1:not-set-string": {Type: STRING, Value: ""},
		"S2:string-as-num":  {Type: STRING},
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

	_, errs = NewArguments().Parse([]string{"--test", "100"}, Map{"t:test": {Type: 10}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Unsuported argument type 10")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"s:string": {}, "s:trace": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -s defined 2 or more times")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t=100"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	fArgs, errs := NewArguments().Parse([]string{"-", "--"}, Map{"t:test": {}})

	c.Assert(errs, HasLen, 0)
	c.Assert(fArgs, DeepEquals, []string{"-", "--"})

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--asd="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --asd has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-j="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument -j has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"t:test": {Required: true}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Required argument test is not set")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t"}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Non-boolean argument --test is empty")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test=!"}, Map{"t:test": {Type: FLOAT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"-t=!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{"--test", "!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Argument --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"t:test": nil})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Struct for argument --test is nil")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewArguments().Parse([]string{}, Map{"": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Some argument does not have a name")
}

func (s *ArgUtilSuite) TestMerging(c *C) {
	c.Assert(Q(), Equals, "")
	c.Assert(Q("test"), Equals, "test")
	c.Assert(Q("test1", "test2"), Equals, "test1 test2")
}
