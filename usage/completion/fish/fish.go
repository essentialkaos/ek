// Package fish provides methods for generating fish completion
package fish

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strings"

	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/usage"
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

	result = strings.Replace(result, "{{GLOBAL_OPTS}}", genGlobalOptionCompletion(info, name), -1)
	result = strings.Replace(result, "{{COMMANDS}}", genCommandsCompletion(info, name), -1)
	result = strings.Replace(result, "{{COMPNAME}}", name, -1)

	nameSafe := strings.Replace(name, "-", "_", -1)

	result = strings.Replace(result, "{{COMPNAME_SAFE}}", nameSafe, -1)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// genGlobalOptionCompletion generates completion for global options
func genGlobalOptionCompletion(info *usage.Info, name string) string {
	var result string

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

		result += genOptionCompletion(opt, name, "")
	}

	return result
}

// genCommandsCompletion generates completion for all commands
func genCommandsCompletion(info *usage.Info, name string) string {
	if len(info.Commands) == 0 {
		return ""
	}

	var result string

	for _, cmd := range info.Commands {
		result += genCommandCompletion(cmd, name)

		if len(cmd.BoundOptions) == 0 {
			result += "\n"
			continue
		}

		for _, optName := range cmd.BoundOptions {
			opt := info.GetOption(optName)

			if opt == nil {
				result += "\n"
				continue
			}

			result += genOptionCompletion(opt, name, cmd.Name)
		}

		result += "\n"
	}

	return result
}

// genOptionCompletion generates completion for option
func genOptionCompletion(opt *usage.Option, name, cmd string) string {
	var result string

	if cmd == "" {
		result = "complete -f -n '__fish_{{COMPNAME_SAFE}}_no_command' "
	} else {
		result = fmt.Sprintf("complete -f -n '__fish_{{COMPNAME_SAFE}}_using_command %s' ", cmd)
	}

	result += fmt.Sprintf("-c %s ", name)
	result += fmt.Sprintf("-l %s ", opt.Long)

	if opt.Short != "" {
		result += fmt.Sprintf("-s %s ", opt.Short)
	}

	result += fmt.Sprintf("-d '%s'\n", fmtc.Clean(opt.Desc))

	return result
}

// genCommandCompletion generates completion for command
func genCommandCompletion(cmd *usage.Command, name string) string {
	result := "complete -f -n '__fish_{{COMPNAME}}_no_command' "
	result += fmt.Sprintf("-c %s ", name)
	result += fmt.Sprintf("-a '%s' ", cmd.Name)
	result += fmt.Sprintf("-d '%s'\n", fmtc.Clean(cmd.Desc))

	return result
}
