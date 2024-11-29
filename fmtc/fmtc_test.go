package fmtc

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/essentialkaos/ek/v13/env"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type FormatSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&FormatSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *FormatSuite) TestColors(c *C) {
	c.Assert(Sprint("{r}W{!}"), Equals, "\x1b[31mW\x1b[0m")
	c.Assert(Sprint("{g}W{!}"), Equals, "\x1b[32mW\x1b[0m")
	c.Assert(Sprint("{y}W{!}"), Equals, "\x1b[33mW\x1b[0m")
	c.Assert(Sprint("{b}W{!}"), Equals, "\x1b[34mW\x1b[0m")
	c.Assert(Sprint("{m}W{!}"), Equals, "\x1b[35mW\x1b[0m")
	c.Assert(Sprint("{c}W{!}"), Equals, "\x1b[36mW\x1b[0m")
	c.Assert(Sprint("{s}W{!}"), Equals, "\x1b[37mW\x1b[0m")
	c.Assert(Sprint("{w}W{!}"), Equals, "\x1b[97mW\x1b[0m")
	c.Assert(Sprint("{r-}W{!}"), Equals, "\x1b[91mW\x1b[0m")
	c.Assert(Sprint("{g-}W{!}"), Equals, "\x1b[92mW\x1b[0m")
	c.Assert(Sprint("{y-}W{!}"), Equals, "\x1b[93mW\x1b[0m")
	c.Assert(Sprint("{b-}W{!}"), Equals, "\x1b[94mW\x1b[0m")
	c.Assert(Sprint("{m-}W{!}"), Equals, "\x1b[95mW\x1b[0m")
	c.Assert(Sprint("{c-}W{!}"), Equals, "\x1b[96mW\x1b[0m")
	c.Assert(Sprint("{s-}W{!}"), Equals, "\x1b[90mW\x1b[0m")
	c.Assert(Sprint("{w-}W{!}"), Equals, "\x1b[97mW\x1b[0m")
}

func (s *FormatSuite) TestBackgrounds(c *C) {
	c.Assert(Sprint("{R}W{!}"), Equals, "\x1b[41mW\x1b[0m")
	c.Assert(Sprint("{G}W{!}"), Equals, "\x1b[42mW\x1b[0m")
	c.Assert(Sprint("{Y}W{!}"), Equals, "\x1b[43mW\x1b[0m")
	c.Assert(Sprint("{B}W{!}"), Equals, "\x1b[44mW\x1b[0m")
	c.Assert(Sprint("{M}W{!}"), Equals, "\x1b[45mW\x1b[0m")
	c.Assert(Sprint("{C}W{!}"), Equals, "\x1b[46mW\x1b[0m")
	c.Assert(Sprint("{S}W{!}"), Equals, "\x1b[47mW\x1b[0m")
	c.Assert(Sprint("{W}W{!}"), Equals, "\x1b[107mW\x1b[0m")
}

func (s *FormatSuite) TestModificators(c *C) {
	c.Assert(Sprint("{!}"), Equals, "\x1b[0m")
	c.Assert(Sprint("{*}W{!}"), Equals, "\x1b[1mW\x1b[0m")
	c.Assert(Sprint("{^}W{!}"), Equals, "\x1b[2mW\x1b[0m")
	c.Assert(Sprint("{_}W{!}"), Equals, "\x1b[4mW\x1b[0m")
	c.Assert(Sprint("{~}W{!}"), Equals, "\x1b[5mW\x1b[0m")
	c.Assert(Sprint("{@}W{!}"), Equals, "\x1b[7mW\x1b[0m")
}

func (s *FormatSuite) TestReset(c *C) {
	c.Assert(Sprint("{*}W{!*}K{!}"), Equals, "\x1b[1mW\x1b[22mK\x1b[0m")
	c.Assert(Sprint("{^}W{!^}K{!}"), Equals, "\x1b[2mW\x1b[22mK\x1b[0m")
	c.Assert(Sprint("{_}W{!_}K{!}"), Equals, "\x1b[4mW\x1b[24mK\x1b[0m")
	c.Assert(Sprint("{~}W{!~}K{!}"), Equals, "\x1b[5mW\x1b[25mK\x1b[0m")
	c.Assert(Sprint("{@}W{!@}K{!}"), Equals, "\x1b[7mW\x1b[27mK\x1b[0m")
}

func (s *FormatSuite) TestParsing(c *C) {
	c.Assert(Sprint(""), Equals, "")
	c.Assert(Sprint("{}"), Equals, "{}")
	c.Assert(Sprint("{-}"), Equals, "{-}")
	c.Assert(Sprint("W"), Equals, "W")
	c.Assert(Sprint("{"), Equals, "{")
	c.Assert(Sprint("{r"), Equals, "{r")
	c.Assert(Sprint("{J}W"), Equals, "{J}W")
	c.Assert(Sprint("{r}W"), Equals, "\x1b[31mW\x1b[0m")
	c.Assert(Sprint("{{r}W{!}}"), Equals, "{\x1b[31mW\x1b[0m}")
	c.Assert(Sprint("Test"+string(rune(65533))), Equals, "Test")
}

func (s *FormatSuite) TestIsTag(c *C) {
	c.Assert(IsTag(""), Equals, true)
	c.Assert(IsTag("{r}"), Equals, true)
	c.Assert(IsTag("{r*}"), Equals, true)
	c.Assert(IsTag("{#123}"), Equals, true)
	c.Assert(IsTag("{*}{_}{#123}"), Equals, true)
	c.Assert(IsTag("{%123}"), Equals, true)
	c.Assert(IsTag("{*}"), Equals, true)
	c.Assert(IsTag("{w-}"), Equals, true)
	c.Assert(IsTag("{S*_}"), Equals, true)
	c.Assert(IsTag("{%1F2E3D}"), Equals, true)

	c.Assert(IsTag("{}"), Equals, false)
	c.Assert(IsTag("{-}"), Equals, false)
	c.Assert(IsTag("W"), Equals, false)
	c.Assert(IsTag("{"), Equals, false)
	c.Assert(IsTag("{r"), Equals, false)
	c.Assert(IsTag("{{r}}"), Equals, false)
}

func (s *FormatSuite) Test256Colors(c *C) {
	origTerm := os.Getenv("TERM")
	origColorTerm := os.Getenv("COLORTERM")

	os.Setenv("TERM", "xterm-256color")
	os.Setenv("COLORTERM", "")
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	isColorsSupportChecked = false
	isColors256Supported = false
	isColorsTCSupported = false

	c.Assert(Is256ColorsSupported(), Equals, true)
	c.Assert(IsTrueColorSupported(), Equals, false)

	isColorsSupportChecked = false
	isColors256Supported = false
	isColorsTCSupported = false

	os.Setenv("TERM", "")
	os.Setenv("COLORTERM", "")
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	c.Assert(Is256ColorsSupported(), Equals, false)
	c.Assert(Is256ColorsSupported(), Equals, false)
	c.Assert(IsTrueColorSupported(), Equals, false)

	os.Setenv("TERM", origTerm)
	os.Setenv("COLORTERM", origColorTerm)
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	c.Assert(Sprint("{#214}o{!}"), Equals, "\x1b[38;5;214mo\x1b[0m")
	c.Assert(Sprint("{%214}O{!}"), Equals, "\x1b[48;5;214mO\x1b[0m")

	c.Assert(Sprint("{#}o"), Equals, "{#}o")
	c.Assert(Sprint("{#257}o"), Equals, "{#257}o")
	c.Assert(Sprint("{#-1}o"), Equals, "{#-1}o")
}

func (s *FormatSuite) Test24BitColors(c *C) {
	origTerm := os.Getenv("TERM")
	origColorTerm := os.Getenv("COLORTERM")

	os.Setenv("TERM", "xterm-256color")
	os.Setenv("COLORTERM", "truecolor")
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	isColorsSupportChecked = false
	isColors256Supported = false
	isColorsTCSupported = false

	c.Assert(IsTrueColorSupported(), Equals, true)
	isColorsSupportChecked = false
	c.Assert(Is256ColorsSupported(), Equals, true)
	isColorsSupportChecked = false
	c.Assert(IsColorsSupported(), Equals, true)
	c.Assert(IsColorsSupported(), Equals, true)

	os.Setenv("TERM", "")
	os.Setenv("COLORTERM", "")
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	isColorsSupportChecked = false
	isColors256Supported = false
	isColorsTCSupported = false

	c.Assert(IsTrueColorSupported(), Equals, false)
	c.Assert(IsTrueColorSupported(), Equals, false)

	os.Setenv("TERM", origTerm)
	os.Setenv("COLORTERM", origColorTerm)
	termEnvVar = env.Var("TERM")
	colorTermEnvVar = env.Var("COLORTERM")

	c.Assert(Sprint("{#f1c1b2}o{!}"), Equals, "\x1b[38;2;241;193;178mo\x1b[0m")
	c.Assert(Sprint("{%1F2E3D}O{!}"), Equals, "\x1b[48;2;31;46;61mO\x1b[0m")

	c.Assert(Sprint("{#}o"), Equals, "{#}o")
	c.Assert(Sprint("{#gggggg}o"), Equals, "{#gggggg}o")
	c.Assert(Sprint("{#-1}o"), Equals, "{#-1}o")
}

func (s *FormatSuite) TestModDisable(c *C) {
	os.Setenv("FMTC_FLAG", "1")

	boldDisableEnvVar = env.Var("FMTC_FLAG")
	italicDisableEnvVar = env.Var("FMTC_FLAG")
	blinkDisableEnvVar = env.Var("FMTC_FLAG")

	c.Assert(Sprint("{*}test{!}"), Equals, "test\x1b[0m")
	c.Assert(Sprint("{&}test{!}"), Equals, "test\x1b[0m")
	c.Assert(Sprint("{~}test{!}"), Equals, "test\x1b[0m")

	boldDisableEnvVar = env.Var("FMTC_NO_BOLD")
	italicDisableEnvVar = env.Var("FMTC_NO_ITALIC")
	blinkDisableEnvVar = env.Var("FMTC_NO_BLINK")
}

func (s *FormatSuite) TestNamedColors(c *C) {
	RemoveColor("myTest_1")
	parseNamedColor("?myTest_1", false)

	c.Assert(NameColor("", "{r}"), ErrorMatches, `Can't add named color: name can't be empty`)
	c.Assert(NameColor("test", ""), ErrorMatches, `Can't add named color: tag can't be empty`)
	c.Assert(NameColor("test", "{H}"), ErrorMatches, `Can't add named color: "{H}" is not valid color tag`)
	c.Assert(NameColor("test%", "{r}"), ErrorMatches, `Can't add named color: "test%" is not valid name`)

	NameColor("myTest_1", "{r}")
	c.Assert(Sprint("{?myTest_1}o{!}"), Equals, "\x1b[31mo\x1b[0m")

	NameColor("myTest_1", "{#214}")
	c.Assert(Sprint("{?myTest_1}o{!}"), Equals, "\x1b[38;5;214mo\x1b[0m")

	NameColor("myTest_1", "{#f1c1b2}")
	c.Assert(Sprint("{?myTest_1}o{!}"), Equals, "\x1b[38;2;241;193;178mo\x1b[0m")

	NameColor("myTest_1", "{#f1c1b2}{_}{&}")
	c.Assert(Sprint("{?myTest_1}o{!}"), Equals, "\x1b[38;2;241;193;178m\x1b[4m\x1b[3mo\x1b[0m")

	RemoveColor("myTest_1")
	c.Assert(Sprint("{?myTest_1}o{!}"), Equals, "o\x1b[0m")

	c.Assert(Sprint("{?}o"), Equals, "{?}o")
	c.Assert(Sprint("{?<}o"), Equals, "{?<}o")
	c.Assert(Sprint("{?mytest+}o"), Equals, "{?mytest+}o")
}

func (s *FormatSuite) TestDisabled(c *C) {
	DisableColors = true

	c.Assert(Sprint("{r}W{!}"), Equals, "W")
	c.Assert(Sprint("{g}W{!}"), Equals, "W")
	c.Assert(Sprint("{y}W{!}"), Equals, "W")
	c.Assert(Sprint("{b}W{!}"), Equals, "W")
	c.Assert(Sprint("{m}W{!}"), Equals, "W")
	c.Assert(Sprint("{c}W{!}"), Equals, "W")
	c.Assert(Sprint("{s}W{!}"), Equals, "W")
	c.Assert(Sprint("{R}W{!}"), Equals, "W")
	c.Assert(Sprint("{G}W{!}"), Equals, "W")
	c.Assert(Sprint("{Y}W{!}"), Equals, "W")
	c.Assert(Sprint("{B}W{!}"), Equals, "W")
	c.Assert(Sprint("{M}W{!}"), Equals, "W")
	c.Assert(Sprint("{C}W{!}"), Equals, "W")
	c.Assert(Sprint("{S}W{!}"), Equals, "W")
	c.Assert(Sprint("{S*_}W{!}"), Equals, "W")

	c.Assert(Sprint("Test {config} value"), Equals, "Test {config} value")

	DisableColors = false
}

func (s *FormatSuite) TestClean(c *C) {
	c.Assert(Clean("{r}W{!}"), Equals, "W")
	c.Assert(Clean("{g}W{!}"), Equals, "W")
	c.Assert(Clean("{y}W{!}"), Equals, "W")
	c.Assert(Clean("{b}W{!}"), Equals, "W")
	c.Assert(Clean("{m}W{!}"), Equals, "W")
	c.Assert(Clean("{c}W{!}"), Equals, "W")
	c.Assert(Clean("{s}W{!}"), Equals, "W")
	c.Assert(Clean("{R}W{!}"), Equals, "W")
	c.Assert(Clean("{G}W{!}"), Equals, "W")
	c.Assert(Clean("{Y}W{!}"), Equals, "W")
	c.Assert(Clean("{B}W{!}"), Equals, "W")
	c.Assert(Clean("{M}W{!}"), Equals, "W")
	c.Assert(Clean("{C}W{!}"), Equals, "W")
	c.Assert(Clean("{S}W{!}"), Equals, "W")
	c.Assert(Clean("{S*_}W{!}"), Equals, "W")
}

func (s *FormatSuite) TestRender(c *C) {
	c.Assert(Render("{S*_}W{!}"), Equals, "\x1b[47;1;4mW\x1b[0m")
}

func (s *FormatSuite) TestMethods(c *C) {
	c.Assert(Errorf("Test %s", "OK"), DeepEquals, errors.New("Test OK"))
	c.Assert(Sprintf("Test %s", "OK"), Equals, "Test OK")
	c.Assert(Sprintln("Test OK"), Equals, "Test OK\n")

	w := bytes.NewBufferString("")

	Fprint(w, "TEST")

	c.Assert(w.String(), Equals, "TEST")

	w = bytes.NewBufferString("")

	Fprintln(w, "TEST")

	c.Assert(w.String(), Equals, "TEST\n")

	w = bytes.NewBufferString("")

	Fprintf(w, "TEST %s", "OK")

	c.Assert(w.String(), Equals, "TEST OK")

	Printf("Printf: %s\n", "OK")
	Print("Print: OK\n")

	LPrint(11, "LPrintf: OK NOTOK")
	NewLine()
	LPrintf(11, "LPrintf: %s NOTOK", "OK")
	NewLine(2)
	LPrintln(12, "LPrintln: OK NOTOK")
	NewLine(-100)
}

func (s *FormatSuite) TestAux(c *C) {
	TPrint("TPrint: OK\n")
	TPrintf("TPrint: %s", "OK")
	TPrintf("")
	TPrintf("TPrint: %s", "OK")

	TPrintln("TPrint: OK")

	TLPrint(11, "TLPrint: OK NOTOK")
	TLPrintf(11, "TLPrint: %s NOTOK", "OK")
	TLPrintf(11, "")
	TLPrintf(11, "TLPrint: %s NOTOK", "OK")

	TLPrintln(11, "TLPrint: OK NOTOK")

	Bell()
}

func (s *FormatSuite) TestIfHelper(c *C) {
	w := bytes.NewBufferString("")

	If(false).Print("Print: NOT OK\n")
	If(false).Println("Println: NOT OK")
	If(false).Printf("Printf: %s\n", "NOT OK")
	If(false).Printfn("Printfn: %s\n", "NOT OK")
	If(false).Fprint(w, "Fprint: NOT OK\n")
	If(false).Fprintln(w, "Fprintln: NOT OK")
	If(false).Fprintf(w, "Fprintf: %s\n", "NOT OK")
	If(false).Fprintfn(w, "Fprintfn: %s", "NOT OK")
	If(false).Sprint("Sprint: NOT OK\n")
	If(false).Sprintln("Sprintln: NOT OK")
	If(false).Sprintf("Sprintf: %s\n", "NOT OK")
	If(false).Sprintfn("Sprintfn: %s", "NOT OK")
	If(false).TPrint("TPrint: NOT OK\n")
	If(false).TPrintln("TPrintln: NOT OK")
	If(false).TPrintf("TPrintf: %s\n", "NOT OK")
	If(false).LPrint(100, "LPrint: NOT OK\n")
	If(false).LPrintln(100, "LPrintln: NOT OK")
	If(false).LPrintf(100, "LPrintf: %s\n", "NOT OK")
	If(false).LPrintfn(100, "LPrintfn: %s", "NOT OK")
	If(false).TLPrint(100, "TLPrint: NOT OK\n")
	If(false).TLPrintln(100, "TLPrintln: NOT OK")
	If(false).TLPrintf(100, "TLPrintf: %s\n", "NOT OK")
	If(false).NewLine()
	If(false).Bell()

	If(true).Print("Print: OK\n")
	If(true).Println("Println: OK")
	If(true).Printf("Printf: %s\n", "OK")
	If(true).Printfn("Printfn: %s\n", "OK")
	If(true).Fprint(w, "Fprint: OK\n")
	If(true).Fprintln(w, "Fprintln: OK")
	If(true).Fprintf(w, "Fprintf: %s\n", "OK")
	If(true).Fprintfn(w, "Fprintfn: %s", "OK")
	If(true).Sprint("Sprint: OK\n")
	If(true).Sprintln("Sprintln: OK")
	If(true).Sprintf("Sprintf: %s\n", "OK")
	If(true).Sprintfn("Sprintf: %s", "OK")
	If(true).TPrint("TPrint: OK\n")
	If(true).TPrintln("TPrintln: OK")
	If(true).TPrintf("TPrintf: %s\n", "OK")
	If(true).LPrint(100, "LPrint: OK\n")
	If(true).LPrintln(100, "LPrintln: OK")
	If(true).LPrintf(100, "LPrintf: %s\n", "OK")
	If(true).LPrintfn(100, "LPrintfn: %s", "OK")
	If(true).TLPrint(100, "TLPrint: OK\n")
	If(true).TLPrintln(100, "TLPrintln: OK")
	If(true).TLPrintf(100, "TLPrintf: %s\n", "OK")
	If(true).NewLine()
	If(true).Bell()
}

func (s *FormatSuite) TestFuzzFixes(c *C) {
	c.Assert(isValidTag("!!!"), Equals, false)
	c.Assert(isValidTag("---"), Equals, false)
	c.Assert(isValidTag("-!"), Equals, false)
	c.Assert(isValidTag("!--"), Equals, false)
	c.Assert(tag2ANSI("-!", false), Equals, "")
}

func (s *FormatSuite) BenchmarkSimple(c *C) {
	for i := 0; i < c.N; i++ {
		Sprint("Test {r}1{!}!")
	}
}

func (s *FormatSuite) Benchmark256(c *C) {
	for i := 0; i < c.N; i++ {
		Sprint("Test {#123}1{!}!")
	}
}

func (s *FormatSuite) Benchmark24bit(c *C) {
	for i := 0; i < c.N; i++ {
		Sprint("Test {#fac1bd}1{!}!")
	}
}

func (s *FormatSuite) BenchmarkNamed(c *C) {
	NameColor("myTest_1", "{r}")

	for i := 0; i < c.N; i++ {
		Sprint("Test {?myTest_1}1{!}!")
	}

	RemoveColor("myTest_1")
}

func (s *FormatSuite) BenchmarkAll(c *C) {
	NameColor("myTest_1", "{r}")

	for i := 0; i < c.N; i++ {
		Sprint("Test {r}1{!} {#123}2{!} {#fac1bd}3{!} {?myTest_1}4{!}!")
	}

	RemoveColor("myTest_1")
}
