package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	"github.com/essentialkaos/ek/v13/knf"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
    [formatting]
test1:      1
            test2:2

    test3: 3 

[string]
  test1: test
  test2: true
  test3: 4500
  test4: !$%^&
  test5: long long long long text for test
  test6: 

[boolean]
  test1: true
  test2: false
  test3: 0
  test4: 1
  test5:
  test6: example for test
  test7: no

[integer]
  test1: 5
  test2: -5
  test3: 10000000
  test4: A
  test5: 0xFF
  test6: 123.4
  test7: 123.456789
  test8: 0xZZYY
  test9: ABCD

[file-mode]
  test1: 644
  test2: 0644
  test3: 0
  test4: ABC
  test5: true

[size]
	test1: 3mb

[comment]
  test1: 100
  # test2: 100

[macro]
  test1: 100
  test2: {macro:test1}.50
  test3: Value is {macro:test2}
  test4: "{macro:test3}"
  test5: {ABC}
  test6: {}

[k]
  t: 1
`

// ////////////////////////////////////////////////////////////////////////////////// //

type ValidatorSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&ValidatorSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) {
	check.TestingT(t)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ValidatorSuite) TestBasicValidators(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	var errs []error

	validators := knf.Validators{
		{"integer:test1", Set, nil},
		{"integer:test1", Greater, 0},
		{"integer:test1", Greater, 0.5},
		{"integer:test1", Less, 10},
		{"integer:test1", Less, 10.1},
		{"integer:test1", NotEquals, 10},
		{"integer:test1", NotEquals, 10.1},
		{"integer:test1", NotEquals, "123"},
	}

	validators = validators.Add(knf.Validators{
		{"string:test3", HasPrefix, "45"},
		{"string:test3", HasSuffix, "00"},
		{"string:test1", LenGreater, 3},
		{"string:test1", LenLess, 10},
		{"string:test1", LenEquals, 4},
	})

	errs = knf.Validate(validators)

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate(knf.Validators{
		{"boolean:test5", Set, nil},
		{"integer:test1", Greater, 10},
		{"integer:test1", Less, 3},
		{"integer:test1", NotEquals, 5},
		{"integer:test1", Greater, "12345"},
		{"integer:test1", SetToAny, []string{"A", "B", "C"}},
		{"integer:test1", SetToAnyIgnoreCase, []string{"A", "B", "C"}},
		{"string:test3", HasPrefix, "AB"},
		{"string:test3", HasSuffix, "CD"},
		{"string:test1", LenLess, 3},
		{"string:test1", LenGreater, 10},
		{"string:test1", LenEquals, 10},
	})

	c.Assert(errs, check.HasLen, 12)

	c.Assert(errs[0].Error(), check.Equals, "Property boolean:test5 must be set")
	c.Assert(errs[1].Error(), check.Equals, "Property integer:test1 can't be less than 10")
	c.Assert(errs[2].Error(), check.Equals, "Property integer:test1 can't be greater than 3")
	c.Assert(errs[3].Error(), check.Equals, "Property integer:test1 can't be equal 5")
	c.Assert(errs[4].Error(), check.Equals, "Validator knf.Greater doesn't support input with type <string> for checking integer:test1 property")
	c.Assert(errs[5].Error(), check.Equals, "Property integer:test1 doesn't contains any valid value")
	c.Assert(errs[6].Error(), check.Equals, "Property integer:test1 doesn't contains any valid value")
	c.Assert(errs[7].Error(), check.Equals, `Property string:test3 must have prefix "AB"`)
	c.Assert(errs[8].Error(), check.Equals, `Property string:test3 must have suffix "CD"`)
	c.Assert(errs[9].Error(), check.Equals, "Property string:test1 value can't be longer than 3 symbols")
	c.Assert(errs[10].Error(), check.Equals, "Property string:test1 value can't be shorter than 10 symbols")
	c.Assert(errs[11].Error(), check.Equals, "Property string:test1 must be 10 symbols long")

	cfgFile := createConfig(c, `
[test]
  empty:
  string: test
  integer: 10
  float: 10.0
  size: 3mb
  boolean: false`)

	cfg, err := knf.Read(cfgFile)

	c.Assert(err, check.IsNil)
	c.Assert(cfg, check.NotNil)

	c.Assert(Set(cfg, "test:empty", nil), check.NotNil)
	c.Assert(Set(cfg, "test:string", nil), check.IsNil)

	c.Assert(Less(cfg, "test:integer", 5).Error(), check.Equals, "Property test:integer can't be greater than 5")
	c.Assert(Less(cfg, "test:integer", int64(5)).Error(), check.Equals, "Property test:integer can't be greater than 5")
	c.Assert(Less(cfg, "test:integer", uint(5)).Error(), check.Equals, "Property test:integer can't be greater than 5")
	c.Assert(Less(cfg, "test:integer", uint64(5)).Error(), check.Equals, "Property test:integer can't be greater than 5")
	c.Assert(Less(cfg, "test:integer", 30), check.IsNil)
	c.Assert(Less(cfg, "test:float", 5.1).Error(), check.Equals, "Property test:float can't be greater than 5.1")
	c.Assert(Less(cfg, "test:float", 30.1), check.IsNil)

	c.Assert(Greater(cfg, "test:integer", 30).Error(), check.Equals, "Property test:integer can't be less than 30")
	c.Assert(Greater(cfg, "test:integer", int64(30)).Error(), check.Equals, "Property test:integer can't be less than 30")
	c.Assert(Greater(cfg, "test:integer", uint(30)).Error(), check.Equals, "Property test:integer can't be less than 30")
	c.Assert(Greater(cfg, "test:integer", uint64(30)).Error(), check.Equals, "Property test:integer can't be less than 30")
	c.Assert(Greater(cfg, "test:integer", 5), check.IsNil)
	c.Assert(Greater(cfg, "test:float", 30.1).Error(), check.Equals, "Property test:float can't be less than 30.1")
	c.Assert(Greater(cfg, "test:float", 5.1), check.IsNil)

	c.Assert(SizeGreater(cfg, "test:size", 10*1024*1024).Error(), check.Equals, "Property test:size can't be less than 10485760 bytes")
	c.Assert(SizeGreater(cfg, "test:size", int64(10*1024*1024)).Error(), check.Equals, "Property test:size can't be less than 10485760 bytes")
	c.Assert(SizeGreater(cfg, "test:size", uint(10*1024*1024)).Error(), check.Equals, "Property test:size can't be less than 10485760 bytes")
	c.Assert(SizeGreater(cfg, "test:size", uint64(10*1024*1024)).Error(), check.Equals, "Property test:size can't be less than 10485760 bytes")
	c.Assert(SizeGreater(cfg, "test:size", float64(10*1024*1024)).Error(), check.Equals, "Property test:size can't be less than 10485760 bytes")
	c.Assert(SizeGreater(cfg, "test:size", false).Error(), check.Equals, "Validator knf.SizeGreater doesn't support input with type <bool> for checking test:size property")
	c.Assert(SizeGreater(cfg, "test:size", uint64(1*1024*1024)), check.IsNil)

	c.Assert(SizeLess(cfg, "test:size", 1*1024*1024).Error(), check.Equals, "Property test:size can't be greater than 1048576 bytes")
	c.Assert(SizeLess(cfg, "test:size", int64(1*1024*1024)).Error(), check.Equals, "Property test:size can't be greater than 1048576 bytes")
	c.Assert(SizeLess(cfg, "test:size", uint(1*1024*1024)).Error(), check.Equals, "Property test:size can't be greater than 1048576 bytes")
	c.Assert(SizeLess(cfg, "test:size", uint64(1*1024*1024)).Error(), check.Equals, "Property test:size can't be greater than 1048576 bytes")
	c.Assert(SizeLess(cfg, "test:size", float64(1*1024*1024)).Error(), check.Equals, "Property test:size can't be greater than 1048576 bytes")
	c.Assert(SizeLess(cfg, "test:size", false).Error(), check.Equals, "Validator knf.SizeLess doesn't support input with type <bool> for checking test:size property")
	c.Assert(SizeLess(cfg, "test:size", uint64(10*1024*1024)), check.IsNil)

	c.Assert(InRange(cfg, "test:integer", Range{50, 100}).Error(), check.Equals, "Property test:integer must be in range 50-100")
	c.Assert(InRange(cfg, "test:integer", Range{1, 100}), check.IsNil)
	c.Assert(InRange(cfg, "test:integer", Range{uint(1), uint(100)}), check.IsNil)
	c.Assert(InRange(cfg, "test:integer", Range{50.5, 100.0}).Error(), check.Equals, "Property test:integer must be in range 50.5-100")
	c.Assert(InRange(cfg, "test:integer", Range{1.0, 100.0}), check.IsNil)
	c.Assert(InRange(cfg, "test:integer", false).Error(), check.Equals, "Validator knf.InRange doesn't support input with type <bool> for checking test:integer property")
	c.Assert(InRange(cfg, "test:integer", Range{true, 100.0}).Error(), check.Equals, "Validator knf.InRange doesn't support type <bool> for 'Range.From' value")
	c.Assert(InRange(cfg, "test:integer", Range{1.0, true}).Error(), check.Equals, "Validator knf.InRange doesn't support type <bool> for 'Range.To' value")

	c.Assert(NotEquals(cfg, "test:empty", "").Error(), check.Equals, "Property test:empty can't be equal \"\"")
	c.Assert(NotEquals(cfg, "test:string", "test").Error(), check.Equals, "Property test:string can't be equal \"test\"")
	c.Assert(NotEquals(cfg, "test:integer", 10).Error(), check.Equals, "Property test:integer can't be equal 10")
	c.Assert(NotEquals(cfg, "test:float", 10.0).Error(), check.Equals, "Property test:float can't be equal 10.000000")
	c.Assert(NotEquals(cfg, "test:boolean", false).Error(), check.Equals, "Property test:boolean can't be equal false")

	c.Assert(NotEquals(cfg, "test:empty", "1"), check.IsNil)
	c.Assert(NotEquals(cfg, "test:string", "testtest"), check.IsNil)
	c.Assert(NotEquals(cfg, "test:integer", 15), check.IsNil)
	c.Assert(NotEquals(cfg, "test:float", 130.0), check.IsNil)
	c.Assert(NotEquals(cfg, "test:boolean", true), check.IsNil)

	c.Assert(NotEquals(cfg, "test:boolean", true), check.IsNil)
	c.Assert(NotEquals(cfg, "test:boolean", true), check.IsNil)
	c.Assert(NotEquals(cfg, "test:boolean", true), check.IsNil)

	c.Assert(SetToAny(cfg, "test:string", []string{"A", "B", "test"}), check.IsNil)
	c.Assert(SetToAny(cfg, "test:string", []string{"A", "B"}).Error(), check.Equals, "Property test:string doesn't contains any valid value")

	c.Assert(SetToAnyIgnoreCase(cfg, "test:string", []string{"A", "B", "TEST"}), check.IsNil)
	c.Assert(SetToAnyIgnoreCase(cfg, "test:string", []string{"A", "B"}).Error(), check.Equals, "Property test:string doesn't contains any valid value")

	c.Assert(SetToAny(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(SetToAnyIgnoreCase(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(Less(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(Greater(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(NotEquals(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(HasPrefix(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(HasSuffix(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenLess(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenGreater(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenEquals(cfg, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")

	c.Assert(HasPrefix(cfg, "test:string", ""), check.ErrorMatches, "Validator knf..* requires non-empty input for checking test:string property")
	c.Assert(HasSuffix(cfg, "test:string", ""), check.ErrorMatches, "Validator knf..* requires non-empty input for checking test:string property")
}

func (s *ValidatorSuite) TestTypeValidators(c *check.C) {
	cfgFile := createConfig(c, `
[boolean]
  test1: 
  test2: 0
  test3: 1
  test4: True
  test5: false
  test6: Yes
  test7: no
  test8: disabled

[num]
  test1: 
  test2: 0
  test3: -100
  test4: 657
  test5: ABCD

[float]
  test1: 
  test2: 0
  test3: 0.6
  test4: -0.45
  test5: ABCD

[size]
	test1:
	test2: 3mb
	test3: 12.33km`)

	cfg, err := knf.Read(cfgFile)

	c.Assert(err, check.IsNil)
	c.Assert(cfg, check.NotNil)

	c.Assert(TypeBool(cfg, "boolean:test1", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test2", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test3", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test4", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test5", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test6", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test7", nil), check.IsNil)
	c.Assert(TypeBool(cfg, "boolean:test8", nil).Error(), check.Equals, "Property boolean:test8 contains unsupported boolean value (disabled)")

	c.Assert(TypeNum(cfg, "num:test1", nil), check.IsNil)
	c.Assert(TypeNum(cfg, "num:test2", nil), check.IsNil)
	c.Assert(TypeNum(cfg, "num:test3", nil), check.IsNil)
	c.Assert(TypeNum(cfg, "num:test4", nil), check.IsNil)
	c.Assert(TypeNum(cfg, "num:test5", nil), check.NotNil)
	c.Assert(TypeNum(cfg, "float:test3", nil).Error(), check.Equals, "Property float:test3 contains unsupported numeric value (0.6)")

	c.Assert(TypeFloat(cfg, "float:test1", nil), check.IsNil)
	c.Assert(TypeFloat(cfg, "float:test2", nil), check.IsNil)
	c.Assert(TypeFloat(cfg, "float:test3", nil), check.IsNil)
	c.Assert(TypeFloat(cfg, "float:test4", nil), check.IsNil)
	c.Assert(TypeFloat(cfg, "float:test5", nil).Error(), check.Equals, "Property float:test5 contains unsupported float value (ABCD)")

	c.Assert(TypeSize(cfg, "size:test1", nil), check.IsNil)
	c.Assert(TypeSize(cfg, "size:test2", nil), check.IsNil)
	c.Assert(TypeSize(cfg, "size:test3", nil).Error(), check.Equals, "Property size:test3 contains unsupported size value (12.33km)")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *check.C, data string) string {
	configPath := c.MkDir() + "/config.knf"

	err := os.WriteFile(configPath, []byte(data), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
