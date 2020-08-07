package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type OptUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&OptUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *OptUtilSuite) TestAdd(c *C) {
	opts := &Options{}

	c.Assert(opts.Add("", &V{}), NotNil)
	c.Assert(opts.Add("t:", &V{}), NotNil)

	c.Assert(opts.Add("test", &V{}), IsNil)
	c.Assert(opts.Add(":test1", &V{}), IsNil)
	c.Assert(opts.Add("t:test2", &V{}), IsNil)

	c.Assert(opts.Add("t1:test", &V{}), NotNil)
	c.Assert(opts.Add("t:test3", &V{}), NotNil)
	c.Assert(opts.Add("t:test3", nil), NotNil)
}

func (s *OptUtilSuite) TestAddMap(c *C) {
	opts := NewOptions()

	m1 := Map{"": {}}
	m2 := Map{"t:": {}}

	c.Assert(opts.AddMap(m1), Not(HasLen), 0)
	c.Assert(opts.AddMap(m2), Not(HasLen), 0)

	m3 := Map{"test": {}}
	m4 := Map{"t:test2": {}}

	c.Assert(opts.AddMap(m3), HasLen, 0)
	c.Assert(opts.AddMap(m4), HasLen, 0)

	m5 := Map{
		"t:test":  {},
		"T:test1": {},
		"t:test2": {},
	}

	c.Assert(opts.AddMap(m5), HasLen, 2)
}

func (s *OptUtilSuite) TestGlobal(c *C) {
	long, short := ParseOptionName("t:test")

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

	global = NewOptions()

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

func (s *OptUtilSuite) TestLimiters(c *C) {
	argline := "--int1 1 --int2 5 --int3 10 --float1 1.0 --float2 5.0 --float3 10.0"

	optMap := Map{
		"int1":   {Type: INT, Min: 3, Max: 7},
		"int2":   {Type: INT, Min: 3, Max: 7},
		"int3":   {Type: INT, Min: 3, Max: 7},
		"float1": {Type: FLOAT, Min: 3.0, Max: 7.0},
		"float2": {Type: FLOAT, Min: 3.0, Max: 7.0},
		"float3": {Type: FLOAT, Min: 3.0, Max: 7.0},
	}

	opts := NewOptions()
	opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(opts.GetI("int1"), Equals, 3)
	c.Assert(opts.GetI("int2"), Equals, 5)
	c.Assert(opts.GetI("int3"), Equals, 7)
	c.Assert(opts.GetF("float1"), Equals, 3.0)
	c.Assert(opts.GetF("float2"), Equals, 5.0)
	c.Assert(opts.GetF("float3"), Equals, 7.0)
}

func (s *OptUtilSuite) TestConflicts(c *C) {
	argline := "--test1 abc --test2 123"

	optMap := Map{
		"test1": {Conflicts: "test2"},
		"test2": {},
	}

	opts := NewOptions()
	_, errs := opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].(OptionError).Type, Equals, ERROR_CONFLICT)
	c.Assert(errs[0].Error(), Equals, "Option test1 conflicts with option test2")

	argline = "--test0 xyz"

	optMap = Map{
		"test0": {},
		"test1": {Conflicts: "test2"},
		"test2": {},
	}

	opts = NewOptions()
	_, errs = opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(errs, HasLen, 0)
}

func (s *OptUtilSuite) TestBound(c *C) {
	argline := "--test1 abc"

	optMap := Map{
		"test1": {Bound: "test2"},
		"test2": {},
	}

	opts := NewOptions()
	_, errs := opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].(OptionError).Type, Equals, ERROR_BOUND_NOT_SET)
	c.Assert(errs[0].Error(), Equals, "Option test2 must be defined with option test1")

	argline = "--test0 xyz"

	optMap = Map{
		"test0": {},
		"test1": {Bound: "test2"},
		"test2": {},
	}

	opts = NewOptions()
	_, errs = opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(errs, HasLen, 0)
}

func (s *OptUtilSuite) TestGetters(c *C) {
	argline := "file.mp3 -s STRING --required TEST -i 320 -b -f 1.098765 -S2 100 -f1 5 -f2 1 -ms ABC --merg-string DEF -mi 6 --merg-int 6 -f3 12 -mf 10.1 -mf 10.1 -i1 5"

	optMap := Map{
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
		"M:mixed":           {Type: MIXED},
	}

	opts := NewOptions()
	opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(opts.GetS("_not_exist_"), Equals, "")
	c.Assert(opts.GetI("_not_exist_"), Equals, 0)
	c.Assert(opts.GetB("_not_exist_"), Equals, false)
	c.Assert(opts.GetF("_not_exist_"), Equals, 0.0)

	c.Assert(opts.GetS("s:string"), Equals, "STRING")
	c.Assert(opts.GetS("string"), Equals, "STRING")
	c.Assert(opts.GetS("S:empty-string"), Equals, "")
	c.Assert(opts.GetS("empty-string"), Equals, "")
	c.Assert(opts.GetB("string"), Equals, true)
	c.Assert(opts.GetI("string"), Equals, 0)
	c.Assert(opts.GetF("string"), Equals, 0.0)
	c.Assert(opts.GetB("empty-string"), Equals, false)
	c.Assert(opts.GetI("empty-string"), Equals, 0)
	c.Assert(opts.GetF("empty-string"), Equals, 0.0)
	c.Assert(opts.GetB("not-set-string"), Equals, false)
	c.Assert(opts.GetI("not-set-string"), Equals, 0)
	c.Assert(opts.GetF("not-set-string"), Equals, 0.0)
	c.Assert(opts.GetS("S2:string-as-num"), Equals, "100")
	c.Assert(opts.GetB("S2:string-as-num"), Equals, true)
	c.Assert(opts.GetI("S2:string-as-num"), Equals, 100)
	c.Assert(opts.GetF("S2:string-as-num"), Equals, 100.0)

	c.Assert(opts.GetS("r:required"), Equals, "TEST")
	c.Assert(opts.GetS("required"), Equals, "TEST")
	c.Assert(opts.GetS("A:alias"), Equals, "TEST")
	c.Assert(opts.GetS("alias"), Equals, "TEST")

	c.Assert(opts.GetS("int"), Equals, "320")
	c.Assert(opts.GetB("int"), Equals, true)
	c.Assert(opts.GetI("int"), Equals, 320)
	c.Assert(opts.GetF("int"), Equals, 320.0)
	c.Assert(opts.GetI("int-between"), Equals, 3)
	c.Assert(opts.GetS("not-set-int"), Equals, "0")
	c.Assert(opts.GetB("not-set-int"), Equals, false)
	c.Assert(opts.GetI("not-set-int"), Equals, 0)
	c.Assert(opts.GetF("not-set-int"), Equals, 0.0)

	c.Assert(opts.GetS("b:bool"), Equals, "true")
	c.Assert(opts.GetI("b:bool"), Equals, 1)
	c.Assert(opts.GetB("b:bool"), Equals, true)
	c.Assert(opts.GetF("b:bool"), Equals, 1.0)
	c.Assert(opts.GetS("empty-bool"), Equals, "")
	c.Assert(opts.GetI("empty-bool"), Equals, 0)
	c.Assert(opts.GetB("empty-bool"), Equals, false)
	c.Assert(opts.GetF("empty-bool"), Equals, 0.0)
	c.Assert(opts.GetS("not-set-bool"), Equals, "false")
	c.Assert(opts.GetI("not-set-bool"), Equals, 0)
	c.Assert(opts.GetB("not-set-bool"), Equals, false)
	c.Assert(opts.GetF("not-set-bool"), Equals, 0.0)

	c.Assert(opts.GetS("float"), Equals, "1.098765")
	c.Assert(opts.GetB("float"), Equals, true)
	c.Assert(opts.GetI("float"), Equals, 1)
	c.Assert(opts.GetF("float"), Equals, 1.098765)
	c.Assert(opts.GetS("not-set-float"), Equals, "0")
	c.Assert(opts.GetB("not-set-float"), Equals, false)
	c.Assert(opts.GetI("not-set-float"), Equals, 0)
	c.Assert(opts.GetF("not-set-float"), Equals, 0.0)

	c.Assert(opts.GetF("float-max"), Equals, 3.0)
	c.Assert(opts.GetF("float-min"), Equals, 3.0)
	c.Assert(opts.GetF("float-between"), Equals, 10.0)

	c.Assert(opts.GetS("merg-string"), Equals, "ABC DEF")
	c.Assert(opts.GetI("merg-int"), Equals, 12)
	c.Assert(opts.GetF("merg-float"), Equals, 20.2)

	c.Assert(opts.Has("_not_exist_"), Equals, false)
	c.Assert(opts.Has("empty-string"), Equals, false)
	c.Assert(opts.Has("s:string"), Equals, true)
}

func (s *OptUtilSuite) TestMixed(c *C) {
	optMap := Map{
		"M:mixed": {Type: MIXED},
		"t:test":  {},
	}

	argline := "-M -t 123"
	opts := NewOptions()
	opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(opts.Has("M:mixed"), Equals, true)
	c.Assert(opts.GetS("M:mixed"), Equals, "true")
	c.Assert(opts.GetS("t:test"), Equals, "123")

	argline = "-M TEST123 --test 123"
	opts = NewOptions()
	opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(opts.Has("M:mixed"), Equals, true)
	c.Assert(opts.GetS("M:mixed"), Equals, "TEST123")
	c.Assert(opts.GetS("t:test"), Equals, "123")

	argline = "--test 123 -M"
	opts = NewOptions()
	opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(opts.Has("M:mixed"), Equals, true)
	c.Assert(opts.GetS("M:mixed"), Equals, "true")
	c.Assert(opts.GetS("t:test"), Equals, "123")
}

func (s *OptUtilSuite) TestParsing(c *C) {
	_, errs := NewOptions().Parse([]string{}, Map{})

	c.Assert(errs, HasLen, 0)

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test", "100"}, Map{"t:test": {Type: 10}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --test has unsupported type")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"s:string": {}, "s:trace": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option -s defined 2 or more times")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t=100"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --test is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	fArgs, errs := NewOptions().Parse([]string{"-", "--"}, Map{"t:test": {}})

	c.Assert(errs, HasLen, 0)
	c.Assert(fArgs, DeepEquals, []string{"-", "--"})

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--asd="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --asd has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-j="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option -j has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": {Required: true}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Required option test is not set")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t"}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Non-boolean option --test is empty")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test=!"}, Map{"t:test": {Type: FLOAT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t=!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test", "!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": nil})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Struct for option --test is nil")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": {Value: []string{}}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Option test contains unsupported default value")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0].Error(), Equals, "Some option does not have a name")
}

func (s *OptUtilSuite) TestMerging(c *C) {
	c.Assert(Q(), Equals, "")
	c.Assert(Q("test"), Equals, "test")
	c.Assert(Q("test1", "test2"), Equals, "test1 test2")
}

func (s *OptUtilSuite) TestGuessType(c *C) {
	c.Assert(guessType(nil), Equals, STRING)
	c.Assert(guessType("test"), Equals, STRING)
	c.Assert(guessType(4), Equals, INT)
	c.Assert(guessType(true), Equals, BOOL)
	c.Assert(guessType(3.3), Equals, FLOAT)
}
