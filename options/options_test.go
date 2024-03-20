package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type OptUtilSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&OptUtilSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *OptUtilSuite) TestAdd(c *C) {
	opts := &Options{}

	c.Assert(opts.Add("", &V{}), Equals, ErrEmptyName)
	c.Assert(opts.Add("t:", &V{}), Equals, ErrEmptyName)

	c.Assert(opts.Add("test", &V{}), IsNil)
	c.Assert(opts.Add(":test1", &V{}), IsNil)
	c.Assert(opts.Add("t:test2", &V{}), IsNil)

	c.Assert(opts.Add("t1:test", &V{}), ErrorMatches, `Option --test defined 2 or more times`)
	c.Assert(opts.Add("t:test3", &V{}), ErrorMatches, `Option -t defined 2 or more times`)
	c.Assert(opts.Add("t:test3", nil), ErrorMatches, `Struct for option --test3 is nil`)
}

func (s *OptUtilSuite) TestMap(c *C) {
	var m Map

	c.Assert(m.Set("test", &V{}), Equals, ErrNilMap)
	c.Assert(m.Delete("test"), Equals, false)

	m = Map{}

	c.Assert(m.Set("", &V{}), Equals, ErrEmptyName)
	c.Assert(m.Set("test", nil), ErrorMatches, `Struct for option --test is nil`)
	c.Assert(m.Delete("_unknown_"), Equals, false)

	c.Assert(m.Set("test", &V{}), IsNil)
	c.Assert(m.Delete("test"), Equals, true)
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

	c.Assert(opts.AddMap(nil), HasLen, 1)
	c.Assert(opts.AddMap(nil)[0], Equals, ErrNilMap)

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
	c.Assert(Split("s:string"), IsNil)
	c.Assert(Is("s:string", ""), Equals, false)
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
	c.Assert(Is("s:string", "Test"), Equals, true)
	c.Assert(Is("string1", "Test"), Equals, false)
	c.Assert(Split("s:string"), DeepEquals, []string{"Test"})
	c.Assert(Split("s:string1"), IsNil)
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
	c.Assert(errs[0], ErrorMatches, "Option test1 conflicts with option test2")

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
	c.Assert(errs[0], ErrorMatches, "Option test2 must be defined with option test1")

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
	argline := "file.mp3 -SAT -s STRING --required TEST -i 320 -b -f 1.098765 -S2 100 -f1 5 -f2 1 -ms ABC --merg-string DEF -mi 6 --merg-int 6 -f3 12 -mf 10.1 -mf 10.1 -i1 5"

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
	args, errs := opts.Parse(strings.Split(argline, " "), optMap)

	c.Assert(errs, HasLen, 0)
	c.Assert(args, HasLen, 2)

	c.Assert(args.Get(0).String(), Equals, "file.mp3")
	c.Assert(args.Get(1).String(), Equals, "-SAT")

	c.Assert(opts.GetS("_not_exist_"), Equals, "")
	c.Assert(opts.GetI("_not_exist_"), Equals, 0)
	c.Assert(opts.GetB("_not_exist_"), Equals, false)
	c.Assert(opts.GetF("_not_exist_"), Equals, 0.0)

	c.Assert(opts.Is("_not_exist_", ""), Equals, false)

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
	c.Assert(opts.Split("merg-string"), DeepEquals, []string{"ABC", "DEF"})
	c.Assert(opts.GetI("merg-int"), Equals, 12)
	c.Assert(opts.GetF("merg-float"), Equals, 20.2)

	c.Assert(opts.Is("s:string", "STRING"), Equals, true)
	c.Assert(opts.Is("int", 320), Equals, true)
	c.Assert(opts.Is("float", 1.098765), Equals, true)
	c.Assert(opts.Is("b:bool", true), Equals, true)
	c.Assert(opts.Is("s:string", uint(1)), Equals, false)

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

func (s *OptUtilSuite) TestValueConversion(c *C) {
	optMap := Map{"t:test": {Value: 1}}
	opts := NewOptions()
	opts.Parse([]string{}, optMap)

	c.Assert(opts.GetS("t:test"), Equals, "1")

	optMap = Map{"t:test": {Value: true}}
	opts = NewOptions()
	opts.Parse([]string{}, optMap)

	c.Assert(opts.GetS("t:test"), Equals, "true")

	optMap = Map{"t:test": {Value: 158}}
	opts = NewOptions()
	opts.Parse([]string{}, optMap)

	c.Assert(opts.GetS("t:test"), Equals, "158")
}

func (s *OptUtilSuite) TestParsing(c *C) {
	_, errs := NewOptions().Parse([]string{}, Map{})

	c.Assert(errs, HasLen, 0)

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test", "100"}, Map{"t:test": {Type: 10}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test has unsupported type")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"s:string": {}, "s:trace": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option -s defined 2 or more times")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t=100"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option -t is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test=abcd"}, Map{"s:string": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test is not supported")

	// //////////////////////////////////////////////////////////////////////////////// //

	fArgs, errs := NewOptions().Parse([]string{"-", "--"}, Map{"t:test": {}})

	c.Assert(errs, HasLen, 0)
	c.Assert(fArgs, DeepEquals, Arguments{"-", "--"})

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--asd="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --asd has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-j="}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option -j has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": {Required: true}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Required option test is not set")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t"}, Map{"t:test": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Non-boolean option --test is empty")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test=!"}, Map{"t:test": {Type: FLOAT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"-t=!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{"--test", "!"}, Map{"t:test": {Type: INT}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option --test has wrong format")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": nil})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Struct for option --test is nil")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"t:test": {Value: []string{}}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], ErrorMatches, "Option test contains unsupported default value")

	// //////////////////////////////////////////////////////////////////////////////// //

	_, errs = NewOptions().Parse([]string{}, Map{"": {}})

	c.Assert(errs, Not(HasLen), 0)
	c.Assert(errs[0], Equals, ErrEmptyName)
}

func (s *OptUtilSuite) TestFormat(c *C) {
	c.Assert(Format(""), Equals, "")
	c.Assert(Format("test"), Equals, "--test")
	c.Assert(Format("t:test"), Equals, "-t/--test")
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

func (s *OptUtilSuite) TestArguments(c *C) {
	a := Arguments{"A.txt", "b.png", "c.txt", "d.jpg", "e.txte"}

	c.Assert(a.Has(0), Equals, true)
	c.Assert(a.Last().String(), Equals, "e.txte")
	c.Assert(a.Get(0).ToLower().String(), Equals, "a.txt")
	c.Assert(a.Get(0).ToUpper().String(), Equals, "A.TXT")
	c.Assert(a.Filter("*.txt"), DeepEquals, Arguments{"A.txt", "c.txt"})
	c.Assert(a.Strings(), DeepEquals, []string{"A.txt", "b.png", "c.txt", "d.jpg", "e.txte"})
	c.Assert(a.Flatten(), Equals, "A.txt b.png c.txt d.jpg e.txte")

	a = Arguments{"2", "3"}
	a = a.Append("4", "5")
	a = a.Unshift("0", "1")

	c.Assert(a, DeepEquals, Arguments{"0", "1", "2", "3", "4", "5"})

	a = Arguments{"*.txt", "*.jpg"}
	c.Assert(a.Filter("*.txt"), HasLen, 0)
}

func (s *OptUtilSuite) TestArgumentsConversion(c *C) {
	a := Arguments{"test", "6", "2.67", "true"}

	c.Assert(a.Has(1), Equals, true)
	c.Assert(a.Has(9), Equals, false)

	c.Assert(a.Get(0).String(), Equals, "test")
	c.Assert(a.Get(1).String(), Equals, "6")
	c.Assert(a.Get(2).String(), Equals, "2.67")
	c.Assert(a.Get(3).String(), Equals, "true")
	c.Assert(a.Get(4).String(), Equals, "")

	c.Assert(a.Get(0).Is("test"), Equals, true)
	c.Assert(a.Get(0).Is("abcd"), Equals, false)
	c.Assert(a.Get(1).Is(6), Equals, true)
	c.Assert(a.Get(1).Is(12), Equals, false)
	c.Assert(a.Get(1).Is(int64(6)), Equals, true)
	c.Assert(a.Get(1).Is(int64(12)), Equals, false)
	c.Assert(a.Get(1).Is(uint64(6)), Equals, true)
	c.Assert(a.Get(1).Is(uint64(12)), Equals, false)
	c.Assert(a.Get(2).Is(2.67), Equals, true)
	c.Assert(a.Get(2).Is(3.14), Equals, false)
	c.Assert(a.Get(3).Is(true), Equals, true)
	c.Assert(a.Get(3).Is(false), Equals, false)

	c.Assert(a.Get(0).Is([]string{}), Equals, false)

	vi, err := a.Get(0).Int()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.Atoi: parsing "test": invalid syntax`)
	c.Assert(vi, Equals, 0)
	vi, err = a.Get(1).Int()
	c.Assert(err, IsNil)
	c.Assert(vi, Equals, 6)

	vi64, err := a.Get(0).Int64()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.ParseInt: parsing "test": invalid syntax`)
	c.Assert(vi64, Equals, int64(0))
	vi64, err = a.Get(1).Int64()
	c.Assert(err, IsNil)
	c.Assert(vi64, Equals, int64(6))

	vu, err := a.Get(0).Uint()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.ParseUint: parsing "test": invalid syntax`)
	c.Assert(vu, Equals, uint64(0))
	vu, err = a.Get(1).Uint()
	c.Assert(err, IsNil)
	c.Assert(vu, Equals, uint64(6))

	fv, err := a.Get(0).Float()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `strconv.ParseFloat: parsing "test": invalid syntax`)
	c.Assert(fv, Equals, 0.0)
	fv, err = a.Get(2).Float()
	c.Assert(err, IsNil)
	c.Assert(fv, Equals, 2.67)

	a = Arguments{"true", "yes", "y", "1", "false", "no", "n", "0", "", "TEST"}

	bv, err := a.Get(0).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, true)
	bv, err = a.Get(1).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, true)
	bv, err = a.Get(2).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, true)
	bv, err = a.Get(3).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, true)
	bv, err = a.Get(4).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, false)
	bv, err = a.Get(5).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, false)
	bv, err = a.Get(6).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, false)
	bv, err = a.Get(7).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, false)
	bv, err = a.Get(8).Bool()
	c.Assert(err, IsNil)
	c.Assert(bv, Equals, false)
	bv, err = a.Get(9).Bool()
	c.Assert(err, NotNil)
	c.Assert(err, ErrorMatches, `Unsupported boolean value "TEST"`)
	c.Assert(bv, Equals, false)
}

func (s *OptUtilSuite) TestArgumentsPathShorthand(c *C) {
	a := Arguments{"/my/test/path/to/file.jpg", "//dir////file.jpg"}

	c.Assert(a.Get(0).Base().String(), Equals, "file.jpg")
	c.Assert(a.Get(1).Clean().String(), Equals, "/dir/file.jpg")
	c.Assert(a.Get(0).Dir().String(), Equals, "/my/test/path/to")
	c.Assert(a.Get(0).Ext().String(), Equals, ".jpg")
	c.Assert(a.Get(0).IsAbs(), Equals, true)

	m, err := a.Get(0).Match("*.txt")
	c.Assert(err, IsNil)
	c.Assert(m, Equals, false)

	m, err = a.Get(0).Match("/my/test/path/to/*.jpg")
	c.Assert(err, IsNil)
	c.Assert(m, Equals, true)
}

func (s *OptUtilSuite) TestNilOptions(c *C) {
	var opts *Options

	c.Assert(opts.Add("", &V{}), Equals, ErrNilOptions)
	c.Assert(opts.AddMap(Map{"t:": {}}), DeepEquals, []error{ErrNilOptions})
	c.Assert(opts.GetS("test"), Equals, "")
	c.Assert(opts.GetI("test"), Equals, 0)
	c.Assert(opts.GetB("test"), Equals, false)
	c.Assert(opts.GetF("test"), Equals, 0.0)
	c.Assert(opts.Is("test", ""), Equals, false)
	c.Assert(opts.Has("test"), Equals, false)

	_, errs := opts.Parse([]string{}, Map{"t:": {}})

	c.Assert(errs, DeepEquals, []error{ErrNilOptions})
}

func (s *OptUtilSuite) TestNilArguments(c *C) {
	var a Arguments

	c.Assert(a.Flatten(), Equals, "")
	c.Assert(a.Has(0), Equals, false)
	c.Assert(a.Get(0).String(), Equals, "")
	c.Assert(a.Last().String(), Equals, "")
	_, err := a.Get(0).Int()
	c.Assert(err, NotNil)
	_, err = a.Get(0).Float()
	c.Assert(err, NotNil)
	v, _ := a.Get(0).Bool()
	c.Assert(v, Equals, false)
}
