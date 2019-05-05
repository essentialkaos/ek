// Package bash provides methods for generating bash completion
package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"strings"

	"pkg.re/essentialkaos/ek.v10/usage"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// _BASH_TEMPLATE is template used for completion generation
const _BASH_TEMPLATE = `# Completion for {{COMPNAME}}
# This completion is automatically generated

_{{COMPNAME_SAFE}}() {
  local cur prev cmds opts

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="{{COMMANDS}}"
  opts="{{OPTIONS}}"

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  if [[ -z "$prev" ]] ; then
    COMPREPLY=($(compgen -W "$cmds" -- "$cur"))
  fi
}

complete -F _{{COMPNAME_SAFE}} {{COMPNAME}} -o nosort
`

// ////////////////////////////////////////////////////////////////////////////////// //

// Generate generates Bash completion code
func Generate(info *usage.Info, name string) string {
	result := _BASH_TEMPLATE

	commands := strings.Join(getCommandsSlice(info), " ")
	options := strings.Join(getOptionsSlice(info), " ")

	result = strings.Replace(result, "{{COMMANDS}}", commands, -1)
	result = strings.Replace(result, "{{OPTIONS}}", options, -1)
	result = strings.Replace(result, "{{COMPNAME}}", name, -1)

	nameSafe := strings.Replace(name, "-", "_", -1)

	result = strings.Replace(result, "{{COMPNAME_SAFE}}", nameSafe, -1)

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getCommandsSlice returns slice with commands
func getCommandsSlice(info *usage.Info) []string {
	var result []string

	for _, command := range info.Commands {
		result = append(result, command.Name)
	}

	return result
}

// getOptionsSlice returns slice with long options
func getOptionsSlice(info *usage.Info) []string {
	var result []string

	for _, option := range info.Options {
		result = append(result, "--"+option.Long)
	}

	return result
}
