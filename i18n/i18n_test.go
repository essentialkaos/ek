package i18n

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type I18NSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

type testBundle struct {
	GREETING      Text
	EXIT_QUESTION Text
	ERRORS        *testErrors
}

type testErrors struct {
	UNKNOWN_USER Text
	UNKNOWN_ID   Text
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&I18NSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *I18NSuite) TestFallback(c *C) {
	en := &testBundle{
		GREETING: "Hello!",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	ru := &testBundle{
		GREETING: "Привет!",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Неизвестный пользователь {{.Username}}",
		},
	}

	kz := &testBundle{
		GREETING: "Сәлеметсіз бе!",
	}

	_, err := Fallback()
	c.Assert(err, NotNil)

	_, err = Fallback(nil)
	c.Assert(err, NotNil)

	_, err = Fallback(en, nil)
	c.Assert(err, NotNil)

	_, err = Fallback(en, en.ERRORS)
	c.Assert(err, NotNil)

	loc, err := Fallback(en)
	c.Assert(err, IsNil)
	c.Assert(loc, Equals, en)

	loc, err = Fallback(en, ru, kz)
	c.Assert(err, IsNil)

	l := loc.(*testBundle)

	data := Data{
		"Username": "johndoe",
		"ID":       183,
	}

	c.Assert(l.GREETING.String(), Equals, "Сәлеметсіз бе!")
	c.Assert(l.ERRORS.UNKNOWN_USER.With(data), Equals, `Неизвестный пользователь johndoe`)
	c.Assert(l.ERRORS.UNKNOWN_ID.With(data), Equals, `Unknown ID 183`)
}

func (s *I18NSuite) TestIsComplete(c *C) {
	en := &testBundle{
		GREETING:      "Hello!",
		EXIT_QUESTION: "Do you really want to exit?",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	ru := &testBundle{
		GREETING: "Привет!",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Неизвестный пользователь {{.Username}}",
		},
	}

	ic, f := IsComplete(nil)
	c.Assert(ic, Equals, false)
	c.Assert(f, HasLen, 0)

	ic, f = IsComplete(en)
	c.Assert(ic, Equals, true)
	c.Assert(f, HasLen, 0)

	ic, f = IsComplete(ru)
	c.Assert(ic, Equals, false)
	c.Assert(f, DeepEquals, []string{"EXIT_QUESTION", "ERRORS.UNKNOWN_ID"})

	en = &testBundle{
		GREETING:      "Hello!",
		EXIT_QUESTION: "Do you really want to exit?",
	}

	ic, f = IsComplete(en)
	c.Assert(ic, Equals, false)
	c.Assert(f, DeepEquals, []string{"ERRORS"})
}

func (s *I18NSuite) TestValidateBundle(c *C) {
	en := &testBundle{
		GREETING:      "Hello!",
		EXIT_QUESTION: "Do you really want to exit?",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Unknown user {{.Username}}",
			UNKNOWN_ID:   "Unknown ID {{.ID}}",
		},
	}

	ru := &testBundle{
		GREETING: "Привет!",
		ERRORS: &testErrors{
			UNKNOWN_USER: "Неизвестный пользователь {{.Username}}",
		},
	}

	c.Assert(ValidateBundle(en), IsNil)

	en = &testBundle{
		GREETING:      "Hello!",
		EXIT_QUESTION: "Do you really want to exit?",
	}

	c.Assert(ValidateBundle(en), ErrorMatches, `Bundle struct ERRORS is nil`)

	_, err := Fallback(en)
	c.Assert(err, ErrorMatches, `Bundle struct ERRORS is nil`)

	b, err := Fallback(en, ru)
	c.Assert(err, IsNil)
	c.Assert(b.(*testBundle).ERRORS, NotNil)
}

func (s *I18NSuite) TestWith(c *C) {
	data := Data{"Test": 100}

	var txt Text
	c.Assert(txt.With(data), Equals, "")

	txt = "Test"
	c.Assert(txt.With(nil), Equals, "Test")

	txt = "Test {{.Test}"
	c.Assert(txt.With(data), Equals, "Test {{.Test}")

	txt = "Test {.Test}}"
	c.Assert(txt.With(data), Equals, "Test {.Test}}")

	txt = "Test {{.Test1}}"
	c.Assert(txt.With(data), Equals, "Test <no value>")
}

func (s *I18NSuite) TestData(c *C) {
	var data Data

	c.Assert(data.Has("int"), Equals, false)
	c.Assert(data.Plural("EN", "int", "1", "2", "3"), Equals, "")
	c.Assert(data.PrettyNum("int"), Equals, UNKNOWN_VALUE)
	c.Assert(data.PrettySize("int"), Equals, UNKNOWN_VALUE)
	c.Assert(data.PrettyPerc("perc"), Equals, UNKNOWN_VALUE)

	data = Data{
		"int":     123456,
		"int64":   int64(123456),
		"uint":    uint(123456),
		"uint64":  uint64(123456),
		"float64": float64(123456.0),
		"perc":    65.34,
		"string":  "abcd",
		"nil":     nil,
	}

	c.Assert(data.Plural("EN", "int", "1", "2", "3"), Equals, "2")
	c.Assert(data.Plural("EN", "int64", "1", "2", "3"), Equals, "2")
	c.Assert(data.Plural("EN", "uint", "1", "2", "3"), Equals, "2")
	c.Assert(data.Plural("EN", "uint64", "1", "2", "3"), Equals, "2")
	c.Assert(data.Plural("EN", "float64", "1", "2", "3"), Equals, "2")
	c.Assert(data.Plural("EN", "string", "1", "2", "3"), Equals, "1")
	c.Assert(data.Plural("EN", "unknown", "1", "2", "3"), Equals, "1")
	c.Assert(data.Plural("EN", "unknown"), Equals, "")

	c.Assert(data.PrettyNum("int"), Equals, "123,456")
	c.Assert(data.PrettyNum("unknown"), Equals, "???")

	c.Assert(data.PrettySize("int"), Equals, "120.6KB")
	c.Assert(data.PrettySize("int64"), Equals, "120.6KB")
	c.Assert(data.PrettySize("uint"), Equals, "120.6KB")
	c.Assert(data.PrettySize("uint64"), Equals, "120.6KB")
	c.Assert(data.PrettySize("float64"), Equals, "120.6KB")
	c.Assert(data.PrettySize("string"), Equals, "abcd")
	c.Assert(data.PrettySize("nil"), Equals, "???")
	c.Assert(data.PrettySize("unknown"), Equals, "???")

	c.Assert(data.PrettyPerc("perc"), Equals, "65.3%")
	c.Assert(data.PrettyPerc("int"), Equals, "???")
	c.Assert(data.PrettyPerc("unknown"), Equals, "???")
}

func (s *I18NSuite) TestString(c *C) {
	is := Text("Hello")

	c.Assert(is.String(), Equals, "Hello")
	c.Assert(is.S(), Equals, "Hello")
	c.Assert(is.Add("[", "]"), Equals, "[Hello]")

	is = Text("User %s (%d)")

	c.Assert(is.Format("John", 831), Equals, "User John (831)")
	c.Assert(is.F("John", 831), Equals, "User John (831)")
}

func (s *I18NSuite) TestPlurLang(c *C) {
	langs := []string{
		"ACH", "AF", "AK", "AM", "AN", "ANP", "AR", "ARN", "AS", "AST", "AY", "AZ",
		"BE", "BG", "BN", "BO", "BR", "BRX", "BS", "CA", "CGG", "CS", "CSB", "CY",
		"DA", "DE", "DOI", "DZ", "EL", "EN", "EO", "ES", "ESAR", "ET", "EU", "FA",
		"FF", "FI", "FIL", "FO", "FR", "FUR", "FY", "GA", "GD", "GL", "GU", "GUN",
		"HA", "HE", "HI", "HNE", "HR", "HU", "HY", "IA", "ID", "IS", "IT", "JA",
		"JBO", "JV", "KA", "KK", "KL", "KM", "KN", "KO", "KU", "KW", "KY", "LB",
		"LN", "LO", "LT", "LV", "MAI", "MFE", "MG", "MI", "MK", "ML", "MN", "MNI",
		"MNK", "MR", "MS", "MT", "MY", "NAH", "NAP", "NB", "NE", "NL", "NN", "NO",
		"NSO", "OC", "OR", "PA", "PAP", "PL", "PMS", "PS", "PT", "PTBR", "RM", "RO",
		"RU", "RW", "SAH", "SAT", "SCO", "SD", "SE", "SI", "SK", "SL", "SO", "SON",
		"SQ", "SR", "SU", "SV", "SW", "TA", "TE", "TG", "TH", "TI", "TK", "TR", "TT",
		"UG", "UK", "UR", "UZ", "VI", "WA", "WO", "YO", "ZH", "UNKNOWN",
	}

	for _, l := range langs {
		c.Assert(getPluralizerByLang(l), NotNil)
	}
}
