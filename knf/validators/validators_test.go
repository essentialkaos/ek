package validators

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	"pkg.re/essentialkaos/ek.v11/knf"

	check "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _CONFIG_DATA = `
    [formating]
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
		{"integer:test1", Empty, nil},
		{"integer:test1", Less, 0},
		{"integer:test1", Less, 0.5},
		{"integer:test1", Greater, 10},
		{"integer:test1", Greater, 10.1},
		{"integer:test1", Equals, 10},
		{"integer:test1", Equals, 10.1},
		{"integer:test1", Equals, "123"},
		{"string:test3", NotLen, 4},
		{"string:test3", NotPrefix, "45"},
		{"string:test3", NotSuffix, "00"},
	})

	c.Assert(errs, check.HasLen, 0)

	errs = knf.Validate([]*knf.Validator{
		{"boolean:test5", Empty, nil},
		{"integer:test1", Less, 10},
		{"integer:test1", Greater, 0},
		{"integer:test1", Equals, 1},
		{"integer:test1", Greater, "12345"},
		{"integer:test1", NotContains, []string{"A", "B", "C"}},
		{"string:test3", NotLen, 8},
		{"string:test3", NotPrefix, "AB"},
		{"string:test3", NotSuffix, "CD"},
	})

	c.Assert(errs, check.HasLen, 9)

	c.Assert(errs[0].Error(), check.Equals, "Property boolean:test5 can't be empty")
	c.Assert(errs[1].Error(), check.Equals, "Property integer:test1 can't be less than 10")
	c.Assert(errs[2].Error(), check.Equals, "Property integer:test1 can't be greater than 0")
	c.Assert(errs[3].Error(), check.Equals, "Property integer:test1 can't be equal 1")
	c.Assert(errs[4].Error(), check.Equals, "Wrong validator for property integer:test1")
	c.Assert(errs[5].Error(), check.Equals, "Property integer:test1 doesn't contains any valid value")
	c.Assert(errs[6].Error(), check.Equals, "Property string:test3 must be 8 symbols long")
	c.Assert(errs[7].Error(), check.Equals, "Property string:test3 must have prefix \"AB\"")
	c.Assert(errs[8].Error(), check.Equals, "Property string:test3 must have suffix \"CD\"")

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

	c.Assert(Empty(fakeConfig, "test:empty", nil), check.NotNil)
	c.Assert(Empty(fakeConfig, "test:string", nil), check.IsNil)

	c.Assert(Less(fakeConfig, "test:integer", 30), check.NotNil)
	c.Assert(Less(fakeConfig, "test:integer", 5), check.IsNil)
	c.Assert(Less(fakeConfig, "test:float", 30.0), check.NotNil)
	c.Assert(Less(fakeConfig, "test:float", 5.0), check.IsNil)
	c.Assert(Less(fakeConfig, "test:string", "30"), check.NotNil)

	c.Assert(Greater(fakeConfig, "test:integer", 5), check.NotNil)
	c.Assert(Greater(fakeConfig, "test:integer", 30), check.IsNil)
	c.Assert(Greater(fakeConfig, "test:float", 5.0), check.NotNil)
	c.Assert(Greater(fakeConfig, "test:float", 30.0), check.IsNil)
	c.Assert(Greater(fakeConfig, "test:string", "30"), check.NotNil)

	c.Assert(Equals(fakeConfig, "test:empty", ""), check.NotNil)
	c.Assert(Equals(fakeConfig, "test:string", "test"), check.NotNil)
	c.Assert(Equals(fakeConfig, "test:integer", 10), check.NotNil)
	c.Assert(Equals(fakeConfig, "test:float", 10.0), check.NotNil)
	c.Assert(Equals(fakeConfig, "test:boolean", false), check.NotNil)

	c.Assert(Equals(fakeConfig, "test:empty", []string{}), check.NotNil)
	c.Assert(Equals(fakeConfig, "test:empty", "1"), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:string", "testtest"), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:integer", 15), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:float", 130.0), check.IsNil)
	c.Assert(Equals(fakeConfig, "test:boolean", true), check.IsNil)

	c.Assert(NotContains(fakeConfig, "test:string", []string{"A", "B", "test"}), check.IsNil)
	c.Assert(NotContains(fakeConfig, "test:string", []string{"A", "B"}), check.NotNil)
	c.Assert(NotContains(fakeConfig, "test:string", 0), check.NotNil)
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
	c.Assert(TypeBool(fakeConfig, "boolean:test8", nil), check.NotNil)

	c.Assert(TypeNum(fakeConfig, "num:test1", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test2", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test3", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test4", nil), check.IsNil)
	c.Assert(TypeNum(fakeConfig, "num:test5", nil), check.NotNil)
	c.Assert(TypeNum(fakeConfig, "float:test3", nil), check.NotNil)

	c.Assert(TypeFloat(fakeConfig, "float:test1", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test2", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test3", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test4", nil), check.IsNil)
	c.Assert(TypeFloat(fakeConfig, "float:test5", nil), check.NotNil)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func createConfig(c *check.C, data string) string {
	configPath := c.MkDir() + "/config.knf"

	err := ioutil.WriteFile(configPath, []byte(data), 0644)

	if err != nil {
		c.Fatal(err.Error())
	}

	return configPath
}
