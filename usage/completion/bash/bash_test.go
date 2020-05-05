package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"pkg.re/essentialkaos/ek.v12/usage"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const _RESULT = `# Completion for test
# This completion is automatically generated

_test() {
  local cur prev cmds opts

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="print clean"
  opts="--option-aaa --option-bbb --option-ccc"

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

  COMPREPLY=($(compgen -W '$(_test_filter "$cmds" "$opts")' -- "$cur"))
}

_test_filter() {
  if [[ -z "$1" ]] ; then
    echo "$2" && return 0
  fi

  local cmd1 cmd2

  for cmd1 in $1 ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        echo "$2" && return 0
      fi
    done
  done

  echo "$1" && return 0
}

complete -F _test test -o nosort
`

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type BashSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&BashSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *BashSuite) TestGenerator(c *C) {
	data := Generate(genTestUsageInfo(), "test")
	c.Assert(data, Equals, _RESULT)
}

func (s *BashSuite) TestAuxi(c *C) {
	info := usage.NewInfo("")

	c.Assert(isCommandHandlersRequired(info), Equals, false)
	c.Assert(genCommandsHandlers(info), Equals, "")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func genTestUsageInfo() *usage.Info {
	info := usage.NewInfo("")

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
