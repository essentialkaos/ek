// Package man contains methods for man pages generation
package man

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"
	"time"

	"github.com/essentialkaos/ek/fmtc"
	"github.com/essentialkaos/ek/timeutil"
	"github.com/essentialkaos/ek/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates man page content
func Generate(info *usage.Info, about *usage.About) string {
	var result string

	result += genHeader(about)
	result += genName(about)
	result += genSynopsis(info)
	result += genDescription(info)
	result += genCommands(info)
	result += genOptions(info)
	result += genExamples(info)
	result += genBugTrackerInfo(about)
	result += genLicense(about)
	result += genAuthor(about)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genHeader generates header part
func genHeader(about *usage.About) string {
	return fmt.Sprintf(
		".TH %s 1 \"%s\" \"%s %s\" \"%s Manual\"\n\n",
		strings.ToUpper(about.App),
		timeutil.Format(time.Now(), "%d %b %Y"),
		about.App,
		strings.Replace(about.Version, ".", "\\&.", -1),
		about.App,
	)
}

// genName generates name part
func genName(about *usage.About) string {
	return fmt.Sprintf(
		".SH NAME\n%s \\- %s\n",
		about.App, about.Desc,
	)
}

// genSynopsis generatest synopsis
func genSynopsis(info *usage.Info) string {
	result := ".SH SYNOPSIS\n.sp\n.nf\n"
	result += ".B " + info.Name + " "

	for index, option := range info.Options {
		result += genOptionShort(option)

		if index != 0 && index%4 == 0 && len(info.Options) != index+1 {
			result += "\n" + strings.Repeat(" ", len(info.Name)+1)
		}
	}

	if len(info.Commands) != 0 {
		result += fmt.Sprintf("[\\fB%s\\fR] ", "COMMAND")
	}

	if len(info.Args) != 0 {
		for _, arg := range info.Args {
			result += fmt.Sprintf("\\fI%s\\fR\n", arg)
		}
	} else {
		result += "\n"
	}

	return result + ".fi\n.sp\n"
}

// genOptions generates options part
func genOptions(info *usage.Info) string {
	if len(info.Options) == 0 {
		return ""
	}

	result := ".SH OPTIONS\n"

	for _, option := range info.Options {
		result += genOptionLong(option)
	}

	return result
}

// genCommands generates commands part
func genCommands(info *usage.Info) string {
	if len(info.Commands) == 0 {
		return ""
	}

	curGroup := ""
	result := ".SH COMMANDS\n"

	for _, command := range info.Commands {
		if command.Group != "" && curGroup != command.Group {
			result += fmt.Sprintf(".SS %s\n", command.Group)
			curGroup = command.Group
		}

		result += ".TP\n"
		result += fmt.Sprintf(".B %s", command.Name)

		if len(command.Args) != 0 {
			result += formatCommandArgs(command.Args)
		} else {
			result += "\n"
		}

		result += fmtc.Clean(command.Desc) + "\n"
	}

	return result
}

// genOptionShort generates short info for option
func genOptionShort(option *usage.Option) string {
	if option.Arg != "" {
		return fmt.Sprintf(
			"[\\fB\\-\\-%s\\fR=\\fI%s\\fR] ",
			option.Long, strings.ToUpper(option.Arg),
		)
	} else {
		return fmt.Sprintf("[\\fB\\-\\-%s\\fR] ", option.Long)
	}
}

// genOptionLong generates long info for option
func genOptionLong(option *usage.Option) string {
	result := ".TP\n"

	result += ".BR "

	if option.Short != "" {
		result += fmt.Sprintf("\\-%s \", \" ", option.Short)
	}

	result += fmt.Sprintf("\\-\\-%s", option.Long)

	if option.Arg != "" {
		result += fmt.Sprintf("\\fR=\\fI%s\\fR\n", strings.ToUpper(option.Arg))
	} else {
		result += "\n"
	}

	result += fmtc.Clean(option.Desc) + "\n"

	return result
}

// genDescription generates description part
func genDescription(info *usage.Info) string {
	if info.Spoiler == "" {
		return ""
	}

	return fmt.Sprintf(
		".SH DESCRIPTION\n\n%s\n\n",
		fmtc.Clean(info.Spoiler),
	)
}

// genExamples generates examples part
func genExamples(info *usage.Info) string {
	if len(info.Examples) == 0 {
		return ""
	}

	result := ".SH EXAMPLES\n"

	for index, example := range info.Examples {
		result += ".TP\n"

		if example.Desc != "" {
			result += ".B • " + example.Desc + "\n"
		} else {
			result += fmt.Sprintf(".B • Example %d\n", index+1)
		}

		if !example.Raw {
			result += fmt.Sprintf("%s %s\n", info.Name, example.Cmd)
		} else {
			result += fmt.Sprintf("%s\n", example.Cmd)
		}
	}

	return result
}

// genBugTrackerInfo generates bugs part
func genBugTrackerInfo(about *usage.About) string {
	if about.BugTracker == "" {
		return ""
	}

	return fmt.Sprintf(
		".SH BUGS\n.PD 0\n\nPlease send any comments or bug reports to <\\fB%s\\fP>.\n\n",
		about.BugTracker,
	)
}

// genLicense generates license part
func genLicense(about *usage.About) string {
	if about.License == "" {
		return ""
	}

	license := about.License

	license = strings.Replace(license, "<", `<\fB`, -1)
	license = strings.Replace(license, ">", `\fP>`, -1)

	return fmt.Sprintf(".SH LICENSE\n\n%s.\n\n", license)
}

// genAuthor generates author part
func genAuthor(about *usage.About) string {
	if about.Owner == "" {
		return ""
	}

	if about.Year == 0 {
		return fmt.Sprintf(
			".SH AUTHOR\n\nCopyright (C) %d \\fB%s\\fP\n\n",
			time.Now().Year(), about.Owner,
		)
	}

	return fmt.Sprintf(
		".SH AUTHOR\n\nCopyright (C) %d-%d \\fB%s\\fP\n\n",
		about.Year, time.Now().Year(), about.Owner,
	)
}

// formatCommandArgs formats command arguments
func formatCommandArgs(args []string) string {
	result := ""

	for _, arg := range args {
		if strings.HasPrefix(arg, "?") {
			result += fmt.Sprintf(" \\fR%s\\fP", strings.Replace(arg, "?", "", -1))
		} else {
			result += fmt.Sprintf(" \\fI%s\\fP", arg)
		}
	}

	return result + "\n"
}
