// Package fish provides methods for generating fish completion
package fish

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

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _FISH_TEMPLATE = `# Completion for {{COMPNAME}}
# This completion is automatically generated

function __fish_{{COMPNAME_SAFE}}_no_command
  set cmd (commandline -opc)
  if [ (count $cmd) -eq 1 ]
    return 0
  end
  return 1
end

function __fish_{{COMPNAME_SAFE}}_using_command
  set cmd (commandline -opc)
  if [ (count $cmd) -gt 1 ]
    if [ $argv[1] = $cmd[2] ]
      return 0
    end
  end
  return 1
end

{{GLOBAL_OPTS}}
{{COMMANDS}}`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates fish completion code
func Generate(info *usage.Info, name string) string {
	result := _FISH_TEMPLATE

	result = strings.ReplaceAll(result, "{{GLOBAL_OPTS}}", genGlobalOptionCompletion(info, name))
	result = strings.ReplaceAll(result, "{{COMMANDS}}", genCommandsCompletion(info, name))
	result = strings.ReplaceAll(result, "{{COMPNAME}}", name)

	nameSafe := strings.ReplaceAll(name, "-", "_")

	result = strings.ReplaceAll(result, "{{COMPNAME_SAFE}}", nameSafe)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genGlobalOptionCompletion generates completion for global options
func genGlobalOptionCompletion(info *usage.Info, name string) string {
	var buf bytes.Buffer

	nonGlobalOptions := make(map[string]bool)

	for _, cmd := range info.Commands {
		for _, opt := range cmd.BoundOptions {
			nonGlobalOptions[opt] = true
		}
	}

	for _, opt := range info.Options {
		if nonGlobalOptions[opt.Long] {
			continue
		}

		buf.WriteString(genOptionCompletion(opt, name, ""))
	}

	return buf.String()
}

// genCommandsCompletion generates completion for all commands
func genCommandsCompletion(info *usage.Info, name string) string {
	if len(info.Commands) == 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, cmd := range info.Commands {
		buf.WriteString(genCommandCompletion(cmd, name))

		if len(cmd.BoundOptions) == 0 {
			buf.WriteRune('\n')
			continue
		}

		for _, optName := range cmd.BoundOptions {
			opt := info.GetOption(optName)

			if opt == nil {
				buf.WriteRune('\n')
				continue
			}

			buf.WriteString(genOptionCompletion(opt, name, cmd.Name))
		}

		buf.WriteRune('\n')
	}

	return buf.String()
}

// genOptionCompletion generates completion for option
func genOptionCompletion(opt *usage.Option, name, cmd string) string {
	var buf bytes.Buffer

	if cmd == "" {
		buf.WriteString("complete -f -n '__fish_{{COMPNAME_SAFE}}_no_command' ")
	} else {
		fmt.Fprintf(&buf, "complete -f -n '__fish_{{COMPNAME_SAFE}}_using_command %s' ", cmd)
	}

	fmt.Fprintf(&buf, "-c %s ", name)
	fmt.Fprintf(&buf, "-l %s ", opt.Long)

	if opt.Short != "" {
		fmt.Fprintf(&buf, "-s %s ", opt.Short)
	}

	fmt.Fprintf(&buf, "-d '%s'\n", fmtc.Clean(opt.Desc))

	return buf.String()
}

// genCommandCompletion generates completion for command
func genCommandCompletion(cmd *usage.Command, name string) string {
	var buf bytes.Buffer

	buf.WriteString("complete -f -n '__fish_{{COMPNAME_SAFE}}_no_command' ")

	fmt.Fprintf(&buf, "-c %s ", name)
	fmt.Fprintf(&buf, "-a '%s' ", cmd.Name)
	fmt.Fprintf(&buf, "-d '%s'\n", fmtc.Clean(cmd.Desc))

	return buf.String()
}
