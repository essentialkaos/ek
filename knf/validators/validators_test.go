package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"testing"

	"github.com/essentialkaos/ek/v12/knf"

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
  test1: 1
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

	errs = knf.Validate([]*knf.Validator{
		{"integer:test1", Set, nil},
		{"integer:test1", Less, 0},
		{"integer:test1", Less, 0.5},
		{"integer:test1", Greater, 10},
		{"integer:test1", Greater, 10.1},
		{"integer:test1", Equals, 10},
		{"integer:test1", Equals, 10.1},
		{"integer:test1", Equals, "123"},
		{"string:test3", NotPrefix, "45"},
		{"string:test3", NotSuffix, "00"},
		{"string:test1", LenLess, 3},
		{"string:test1", LenGreater, 10},
		{"string:test1", LenNotEquals, 4},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"boolean:test5", Set, nil},
		{"integer:test1", Less, 10},
		{"integer:test1", Greater, 0},
		{"integer:test1", Equals, 1},
		{"integer:test1", Greater, "12345"},
		{"integer:test1", SetToAny, []string{"A", "B", "C"}},
		{"integer:test1", SetToAnyIgnoreCase, []string{"A", "B", "C"}},
		{"string:test3", NotPrefix, "AB"},
		{"string:test3", NotSuffix, "CD"},
		{"string:test1", LenLess, 10},
		{"string:test1", LenGreater, 3},
		{"string:test1", LenNotEquals, 10},
	})

	c.Assert(errs, check.HasLen, 12)

	c.Assert(errs[0].Error(), check.Equals, "Property boolean:test5 must be set")
	c.Assert(errs[1].Error(), check.Equals, "Property integer:test1 can't be less than 10")
	c.Assert(errs[2].Error(), check.Equals, "Property integer:test1 can't be greater than 0")
	c.Assert(errs[3].Error(), check.Equals, "Property integer:test1 can't be equal 1")
	c.Assert(errs[4].Error(), check.Equals, "Validator knf.Greater doesn't support input with type <string> for checking integer:test1 property")
	c.Assert(errs[5].Error(), check.Equals, "Property integer:test1 doesn't contains any valid value")
	c.Assert(errs[6].Error(), check.Equals, "Property integer:test1 doesn't contains any valid value")
	c.Assert(errs[7].Error(), check.Equals, `Property string:test3 must have prefix "AB"`)
	c.Assert(errs[8].Error(), check.Equals, `Property string:test3 must have suffix "CD"`)
	c.Assert(errs[9].Error(), check.Equals, "Property string:test1 value can't be shorter than 10 symbols")
	c.Assert(errs[10].Error(), check.Equals, "Property string:test1 value can't be longer than 3 symbols")
	c.Assert(errs[11].Error(), check.Equals, "Property string:test1 must be 10 symbols long")

	fakeConfigFile := createConfig(c, `
[test]
  empty:
  string: test
  integer: 10
  float: 10.0
  boolean: false`)

	fakeConfig, err := knf.Read(fakeConfigFile)

	c.Assert(err, check.IsNil)
	c.Assert(fakeConfig, check.NotNil)

	c.Assert(Set(fakeConfig, "test:empty", nil), check.NotNil)
	c.Assert(Set(fakeConfig, "test:string", nil), check.IsNil)

	c.Assert(Less(fakeConfig, "test:integer", 30).Error(), check.Equals, "Property test:integer can't be less than 30")
	c.Assert(Less(fakeConfig, "test:integer", 5), check.IsNil)
	c.Assert(Less(fakeConfig, "test:float", 30.0).Error(), check.Equals, "Property test:float can't be less than 30")
	c.Assert(Less(fakeConfig, "test:float", 5.0), check.IsNil)

	c.Assert(Greater(fakeConfig, "test:integer", 5).Error(), check.Equals, "Property test:integer can't be greater than 5")
	c.Assert(Greater(fakeConfig, "test:integer", 30), check.IsNil)
	c.Assert(Greater(fakeConfig, "test:float", 5.0).Error(), check.Equals, "Property test:float can't be greater than 5")
	c.Assert(Greater(fakeConfig, "test:float", 30.0), check.IsNil)

	c.Assert(Equals(fakeConfig, "test:empty", "").Error(), check.Equals, "Property test:empty can't be equal \"\"")
	c.Assert(Equals(fakeConfig, "test:string", "test").Error(), check.Equals, "Property test:string can't be equal \"test\"")
	c.Assert(Equals(fakeConfig, "test:integer", 10).Error(), check.Equals, "Property test:integer can't be equal 10")
	c.Assert(Equals(fakeConfig, "test:float", 10.0).Error(), check.Equals, "Property test:float can't be equal 10.000000")
	c.Assert(Equals(fakeConfig, "test:boolean", false).Error(), check.Equals, "Property test:boolean can't be equal false")

	c.Assert(Equals(fakeConfig, "test:empty", "1"), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:string", "testtest"), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:integer", 15), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:float", 130.0), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:boolean", true), check.IsNil)

	c.Assert(Equals(fakeConfig, "test:boolean", true), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:boolean", true), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:boolean", true), check.IsNil)

	c.Assert(SetToAny(fakeConfig, "test:string", []string{"A", "B", "test"}), check.IsNil)
	c.Assert(SetToAny(fakeConfig, "test:string", []string{"A", "B"}).Error(), check.Equals, "Property test:string doesn't contains any valid value")

	c.Assert(SetToAnyIgnoreCase(fakeConfig, "test:string", []string{"A", "B", "TEST"}), check.IsNil)
	c.Assert(SetToAnyIgnoreCase(fakeConfig, "test:string", []string{"A", "B"}).Error(), check.Equals, "Property test:string doesn't contains any valid value")

	c.Assert(SetToAny(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(SetToAnyIgnoreCase(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(Less(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(Greater(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(Equals(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(NotLen(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(NotPrefix(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(NotSuffix(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenLess(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenGreater(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
	c.Assert(LenNotEquals(fakeConfig, "test:string", float32(1.1)), check.ErrorMatches, "Validator knf..* doesn't support input with type <float32> for checking test:string property")
}

func (s *ValidatorSuite) TestTypeValidators(c *check.C) {
	fakeConfigFile := createConfig(c, `
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
  test5: ABCD`)

	fakeConfig, err := knf.Read(fakeConfigFile)

	c.Assert(err, check.IsNil)
	c.Assert(fakeConfig, check.NotNil)

	c.Assert(TypeBool(fakeConfig, "boolean:test1", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test2", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test3", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test4", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test5", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test6", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test7", nil), check.IsNil)
	c.Assert(TypeBool(fakeConfig, "boolean:test8", nil).Error(), check.Equals, "Property boolean:test8 contains unsupported boolean value (disabled)")

	c.Assert(TypeNum(fakeConfig, "num:test1", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test2", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test3", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test4", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test5", nil), check.NotNil)
	c.Assert(TypeNum(fakeConfig, "float:test3", nil).Error(), check.Equals, "Property float:test3 contains unsupported numeric value (0.6)")

	c.Assert(TypeFloat(fakeConfig, "float:test1", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test2", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test3", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test4", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test5", nil).Error(), check.Equals, "Property float:test5 contains unsupported float value (ABCD)")
}

func (s *ValidatorSuite) TestDeprecated(c *check.C) {
	var err error

	configFile := createConfig(c, _CONFIG_DATA)

	err = knf.Global(configFile)
	c.Assert(err, check.IsNil)

	var errs []error

	errs = knf.Validate([]*knf.Validator{
		{"boolean:test5", Empty, nil},
		{"integer:test1", NotContains, []string{"A", "B", "C"}},
		{"string:test3", NotLen, 8},
		{"string:test3", NotLen, 4},
	})

	c.Assert(errs, check.HasLen, 3)
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
