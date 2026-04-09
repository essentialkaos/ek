// Package man contains methods for man pages generation
package man

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/timeutil"
	"github.com/essentialkaos/ek/v13/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates man page content
func Generate(info *usage.Info, about *usage.About) string {
	var buf bytes.Buffer

	genHeader(&buf, about)
	genName(&buf, about)
	genSynopsis(&buf, info)
	genDescription(&buf, info)
	genCommands(&buf, info)
	genOptions(&buf, info)
	genExamples(&buf, info)
	genBugTrackerInfo(&buf, about)
	genLicense(&buf, about)
	genAuthor(&buf, about)

	return buf.String()
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genHeader generates header part
func genHeader(buf *bytes.Buffer, about *usage.About) {
	fmt.Fprintf(
		buf, ".TH %s 1 \"%s\" \"%s %s\" \"%s Manual\"\n\n",
		strings.ToUpper(about.App),
		timeutil.Format(time.Now(), "%d %b %Y"),
		about.App,
		strings.ReplaceAll(about.Version, ".", "\\&."),
		about.App,
	)
}

// genName generates name part
func genName(buf *bytes.Buffer, about *usage.About) {
	fmt.Fprintf(
		buf, ".SH NAME\n%s \\- %s\n",
		about.App, about.Desc,
	)
}

// genSynopsis generatest synopsis
func genSynopsis(buf *bytes.Buffer, info *usage.Info) {
	buf.WriteString(".SH SYNOPSIS\n.sp\n.nf\n")
	buf.WriteString(".B " + info.Name + " ")

	for index, option := range info.Options {
		genOptionShort(buf, option)

		if index != 0 && index%4 == 0 && len(info.Options) != index+1 {
			buf.WriteString("\n" + strings.Repeat(" ", len(info.Name)+1))
		}
	}

	if len(info.Commands) != 0 {
		fmt.Fprintf(buf, "[\\fB%s\\fR] ", "COMMAND")
	}

	if len(info.Args) != 0 {
		for _, arg := range info.Args {
			fmt.Fprintf(buf, "\\fI%s\\fR\n", arg)
		}
	} else {
		buf.WriteRune('\n')
	}

	buf.WriteString(".fi\n.sp\n")
}

// genOptions generates options part
func genOptions(buf *bytes.Buffer, info *usage.Info) {
	if len(info.Options) == 0 {
		return
	}

	buf.WriteString(".SH OPTIONS\n")

	for _, option := range info.Options {
		genOptionLong(buf, option)
	}
}

// genCommands generates commands part
func genCommands(buf *bytes.Buffer, info *usage.Info) {
	if len(info.Commands) == 0 {
		return
	}

	curGroup := ""

	buf.WriteString(".SH COMMANDS\n")

	for _, command := range info.Commands {
		if command.Group != "" && curGroup != command.Group {
			fmt.Fprintf(buf, ".SS %s\n", command.Group)
			curGroup = command.Group
		}

		buf.WriteString(".TP\n")
		fmt.Fprintf(buf, ".B %s", command.Name)

		if len(command.Args) != 0 {
			formatCommandArgs(buf, command.Args)
		} else {
			buf.WriteRune('\n')
		}

		buf.WriteString(fmtc.Clean(command.Desc))
		buf.WriteRune('\n')
	}
}

// genOptionShort generates short info for option
func genOptionShort(buf *bytes.Buffer, option *usage.Option) {
	if option.Arg != "" {
		fmt.Fprintf(
			buf, "[\\fB\\-\\-%s\\fR=\\fI%s\\fR] ",
			option.Long, strings.ToUpper(option.Arg),
		)
	} else {
		fmt.Fprintf(buf, "[\\fB\\-\\-%s\\fR] ", option.Long)
	}
}

// genOptionLong generates long info for option
func genOptionLong(buf *bytes.Buffer, option *usage.Option) {
	buf.WriteString(".TP\n.BR ")

	if option.Short != "" {
		fmt.Fprintf(buf, "\\-%s \", \" ", option.Short)
	}

	fmt.Fprintf(buf, "\\-\\-%s", option.Long)

	if option.Arg != "" {
		fmt.Fprintf(buf, "\\fR=\\fI%s\\fR\n", strings.ToUpper(option.Arg))
	} else {
		buf.WriteRune('\n')
	}

	buf.WriteString(fmtc.Clean(option.Desc))
	buf.WriteRune('\n')
}

// genDescription generates description part
func genDescription(buf *bytes.Buffer, info *usage.Info) {
	if info.Spoiler == "" {
		return
	}

	fmt.Fprintf(
		buf, ".SH DESCRIPTION\n\n%s\n\n",
		fmtc.Clean(info.Spoiler),
	)
}

// genExamples generates examples part
func genExamples(buf *bytes.Buffer, info *usage.Info) {
	if len(info.Examples) == 0 {
		return
	}

	buf.WriteString(".SH EXAMPLES\n")

	for index, example := range info.Examples {
		buf.WriteString(".TP\n")

		if example.Desc != "" {
			buf.WriteString(".B • " + example.Desc + "\n")
		} else {
			fmt.Fprintf(buf, ".B • Example %d\n", index+1)
		}

		if !example.IsRaw {
			fmt.Fprintf(buf, "%s %s\n", info.Name, example.Cmd)
		} else {
			fmt.Fprintf(buf, "%s\n", example.Cmd)
		}
	}
}

// genBugTrackerInfo generates bugs part
func genBugTrackerInfo(buf *bytes.Buffer, about *usage.About) {
	if about.BugTracker == "" {
		return
	}

	fmt.Fprintf(
		buf,
		".SH BUGS\n.PD 0\n\nPlease send any comments or bug reports to <\\fB%s\\fP>.\n\n",
		about.BugTracker,
	)
}

// genLicense generates license part
func genLicense(buf *bytes.Buffer, about *usage.About) {
	if about.License == "" {
		return
	}

	license := about.License

	license = strings.ReplaceAll(license, "<", `<\fB`)
	license = strings.ReplaceAll(license, ">", `\fP>`)

	fmt.Fprintf(buf, ".SH LICENSE\n\n%s.\n\n", license)
}

// genAuthor generates author part
func genAuthor(buf *bytes.Buffer, about *usage.About) {
	if about.Owner == "" {
		return
	}

	year := time.Now().Year()

	if about.Year == 0 {
		fmt.Fprintf(
			buf, ".SH AUTHOR\n\nCopyright (C) %d \\fB%s\\fP\n\n",
			year, about.Owner,
		)
	} else {
		fmt.Fprintf(
			buf, ".SH AUTHOR\n\nCopyright (C) %d-%d \\fB%s\\fP\n\n",
			about.Year, year, about.Owner,
		)
	}

}

// formatCommandArgs formats command arguments
func formatCommandArgs(buf *bytes.Buffer, args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "?") {
			fmt.Fprintf(buf, " \\fR%s\\fP", strings.ReplaceAll(arg, "?", ""))
		} else {
			fmt.Fprintf(buf, " \\fI%s\\fP", arg)
		}
	}

	buf.WriteRune('\n')
}
