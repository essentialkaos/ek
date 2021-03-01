package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"pkg.re/essentialkaos/ek.v12/usage"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _RESULT_FILES = `# Completion for test
# This completion is automatically generated

_test() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="print clean"
  opts="--option-aaa --option-bbb --option-ccc"
  show_files="true"

  case $prev in
    clean)
      opts="--option-ddd --option-eee"
      COMPREPLY=($(compgen -W "$opts" -- "$cur"))
      return 0
      ;;

  esac

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  if [[ -z "$cmds" && -n "$show_files" ]] ; then
    _filedir && return 0
  fi

  COMPREPLY=($(compgen -W '$(_test_filter "$cmds" "$opts" "$show_files")' -- "$cur"))
}

_test_filter() {
  local cmds="$1"
  local opts="$2"
  local show_files="$3"

  local cmd1 cmd2

  for cmd1 in $1 ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        echo "$2" && return 0
      fi
    done
  done

  if [[ -z "$show_files" ]] ; then
    echo "$opts" && return 0
  fi

  compgen -f -- "${COMP_WORDS[COMP_CWORD]}"
}

complete -F _test test -o filenames
`

const _RESULT_NO_FILES = `# Completion for test
# This completion is automatically generated

_test() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="print clean"
  opts="--option-aaa --option-bbb --option-ccc"
  show_files=""

  case $prev in
    clean)
      opts="--option-ddd --option-eee"
      COMPREPLY=($(compgen -W "$opts" -- "$cur"))
      return 0
      ;;

  esac

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  if [[ -z "$cmds" && -n "$show_files" ]] ; then
    _filedir && return 0
  fi

  COMPREPLY=($(compgen -W '$(_test_filter "$cmds" "$opts" "$show_files")' -- "$cur"))
}

_test_filter() {
  local cmds="$1"
  local opts="$2"
  local show_files="$3"

  local cmd1 cmd2

  for cmd1 in $1 ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        echo "$2" && return 0
      fi
    done
  done

  if [[ -z "$show_files" ]] ; then
    echo "$opts" && return 0
  fi

  compgen -f -- "${COMP_WORDS[COMP_CWORD]}"
}

complete -F _test test 
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BashSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BashSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BashSuite) TestGenerator(c *C) {
	completion := Generate(genTestUsageInfo(true), "test")
	c.Assert(completion, Equals, _RESULT_FILES)

	completion = Generate(genTestUsageInfo(false), "test")
	c.Assert(completion, Equals, _RESULT_NO_FILES)
}

func (s *BashSuite) TestAuxi(c *C) {
	info := usage.NewInfo("")

	c.Assert(isCommandHandlersRequired(info), Equals, false)
	c.Assert(genCommandsHandlers(info), Equals, "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func genTestUsageInfo(withFiles bool) *usage.Info {
	var info *usage.Info

	if withFiles {
		info = usage.NewInfo("", "spec-file")
	} else {
		info = usage.NewInfo("")
	}

	info.AddCommand("print", "Print command")
	info.AddCommand("clean", "Clean command")

	info.AddOption("a:option-aaa", "Test option A")
	info.AddOption("b:option-bbb", "Test option B", "?bbb")
	info.AddOption("option-ccc", "Test option C")

	info.AddOption("d:option-ddd", "Test option D")
	info.AddOption("e:option-eee", "Test option E")

	info.BoundOptions("clean", "d:option-ddd", "e:option-eee", "unknown")

	return info
}
