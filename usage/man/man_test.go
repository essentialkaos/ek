package man

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"testing"
	"time"

	"pkg.re/essentialkaos/ek.v12/timeutil"
	"pkg.re/essentialkaos/ek.v12/usage"

	. "pkg.re/essentialkaos/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ManSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ManSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ManSuite) TestGenerator(c *C) {
	about := &usage.About{
		App:     "TestApp",
		Version: "1.12.34",
		Desc:    "My supper app",
	}

	info := usage.NewInfo()

	c.Assert(Generate(info, about), Not(Equals), "")
}

func (s *ManSuite) TestHeaderGenerator(c *C) {
	about := &usage.About{
		App:     "TestApp",
		Version: "1.12.34",
	}

	now := timeutil.Format(time.Now(), "%d %b %Y")
	header := fmt.Sprintf(".TH TESTAPP 1 \"%s\" \"TestApp 1\\&.12\\&.34\" \"TestApp Manual\"\n\n", now)

	c.Assert(genHeader(about), Equals, header)
}

func (s *ManSuite) TestNameGenerator(c *C) {
	about := &usage.About{
		App:  "TestApp",
		Desc: "My supper app",
	}

	c.Assert(genName(about), Equals, ".SH NAME\nTestApp \\- My supper app\n")
}

func (s *ManSuite) TestSynopsisGenerator(c *C) {
	info := usage.NewInfo("testapp", "args")

	info.AddOption("t:test1", "Test1")
	info.AddOption("T:test2", "Test2", "data")

	info.AddCommand("test", "Test")

	synopsis := ".SH SYNOPSIS\n.sp\n.nf\n"
	synopsis += ".B testapp [\\fB\\-\\-test1\\fR] [\\fB\\-\\-test2\\fR=\\fIDATA\\fR] [\\fBCOMMAND\\fR] \\fIargs\\fR\n"
	synopsis += ".fi\n.sp\n"

	c.Assert(genSynopsis(info), Equals, synopsis)

	info = usage.NewInfo()

	info.AddOption("a:opta", "OptA")
	info.AddOption("b:optb", "OptB")
	info.AddOption("c:optc", "OptC")
	info.AddOption("d:optd", "OptD")
	info.AddOption("e:opte", "OptE")
	info.AddOption("f:optf", "OptF")

	synopsis = ".SH SYNOPSIS\n.sp\n.nf\n"
	synopsis += ".B man.test [\\fB\\-\\-opta\\fR] [\\fB\\-\\-optb\\fR] [\\fB\\-\\-optc\\fR] [\\fB\\-\\-optd\\fR] [\\fB\\-\\-opte\\fR] \n"
	synopsis += "         [\\fB\\-\\-optf\\fR] \n"
	synopsis += ".fi\n.sp\n"

	c.Assert(genSynopsis(info), Equals, synopsis)
}

func (s *ManSuite) TestOptionsGenerator(c *C) {
	info := usage.NewInfo()

	c.Assert(genOptions(info), Equals, "")

	info.AddOption("t:test1", "Test1")
	info.AddOption("T:test2", "Test2", "data")

	options := ".SH OPTIONS\n"
	options += ".TP\n.BR \\-t \", \" \\-\\-test1\nTest1\n"
	options += ".TP\n.BR \\-T \", \" \\-\\-test2\\fR=\\fIDATA\\fR\nTest2\n"

	c.Assert(genOptions(info), Equals, options)
}

func (s *ManSuite) TestCommandsGenerator(c *C) {
	info := usage.NewInfo()

	c.Assert(genCommands(info), Equals, "")

	info.AddCommand("test1", "Test1 command")
	info.AddGroup("Group1")
	info.AddCommand("test2", "Test2 command", "arg1")
	info.AddGroup("Group2")
	info.AddCommand("test3", "Test3 command", "?arg1")

	commands := ".SH COMMANDS\n"
	commands += ".SS Commands\n.TP\n.B test1\nTest1 command\n"
	commands += ".SS Group1\n"
	commands += ".TP\n.B test2 \\fIarg1\\fP\nTest2 command\n"
	commands += ".SS Group2\n"
	commands += ".TP\n.B test3 \\fRarg1\\fP\nTest3 command\n"

	c.Assert(genCommands(info), Equals, commands)
}

func (s *ManSuite) TestDescriptionGenerator(c *C) {
	info := &usage.Info{}

	c.Assert(genDescription(info), Equals, "")

	info.AddSpoiler("Some text.")

	c.Assert(genDescription(info), Equals, ".SH DESCRIPTION\n\nSome text.\n\n")
}

func (s *ManSuite) TestExamplesGenerator(c *C) {
	info := &usage.Info{Name: "app"}

	c.Assert(genExamples(info), Equals, "")

	info.AddExample("test 123", "Test1")
	info.AddExample("test 456")
	info.AddRawExample("app test 789", "Test3")

	examples := ".SH EXAMPLES\n"
	examples += ".TP\n.B • Test1\napp test 123\n"
	examples += ".TP\n.B • Example 2\napp test 456\n"
	examples += ".TP\n.B • Test3\napp test 789\n"

	c.Assert(genExamples(info), Equals, examples)
}

func (s *ManSuite) TestAuthorGenerator(c *C) {
	about := &usage.About{}

	c.Assert(genAuthor(about), Equals, "")

	about = &usage.About{Owner: "John Doe"}

	authorData := fmt.Sprintf(".SH AUTHOR\n\nCopyright (C) %d \\fBJohn Doe\\fP\n\n", time.Now().Year())

	c.Assert(genAuthor(about), Equals, authorData)

	about = &usage.About{Owner: "John Doe", Year: 2000}

	authorData = fmt.Sprintf(
		".SH AUTHOR\n\nCopyright (C) %d-%d \\fBJohn Doe\\fP\n\n",
		about.Year, time.Now().Year(),
	)

	c.Assert(genAuthor(about), Equals, authorData)
}

func (s *ManSuite) TestLicenseGenerator(c *C) {
	about := &usage.About{}

	c.Assert(genLicense(about), Equals, "")

	about = &usage.About{License: "MIT <https://opensource.org/licenses/MIT>"}

	c.Assert(genLicense(about), Equals, ".SH LICENSE\n\nMIT <\\fBhttps://opensource.org/licenses/MIT\\fP>.\n\n")
}

func (s *ManSuite) TestBugTrackerGenerator(c *C) {
	about := &usage.About{}

	c.Assert(genBugTrackerInfo(about), Equals, "")

	about = &usage.About{
		BugTracker: "https://bugs.com",
	}

	info := ".SH BUGS\n.PD 0\n\nPlease send any comments or bug reports to <\\fBhttps://bugs.com\\fP>.\n\n"

	c.Assert(genBugTrackerInfo(about), Equals, info)
}
