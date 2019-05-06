package bash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"pkg.re/essentialkaos/ek.v10/usage"

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
  opts="--option-aaa --option-bbb --option-ccc --option-ddd --option-eee"

  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  if [[ -z "$prev" ]] ; then
    COMPREPLY=($(compgen -W "$cmds" -- "$cur"))
  fi
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

	info.BoundOptions("clean", "d:option-ddd", "e:option-eee")

	return info
}
