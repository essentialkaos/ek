#!/bin/bash

################################################################################

NORM=0
BOLD=1
UNLN=4
RED=31
GREEN=32
YELLOW=33
BLUE=34
MAG=35
CYAN=36
GREY=37
DARK=90

CL_NORM="\e[0m"
CL_BOLD="\e[0;${BOLD};49m"
CL_UNLN="\e[0;${UNLN};49m"
CL_RED="\e[0;${RED};49m"
CL_GREEN="\e[0;${GREEN};49m"
CL_YELLOW="\e[0;${YELLOW};49m"
CL_BLUE="\e[0;${BLUE};49m"
CL_MAG="\e[0;${MAG};49m"
CL_CYAN="\e[0;${CYAN};49m"
CL_GREY="\e[0;${GREY};49m"
CL_DARK="\e[0;${DARK};49m"
CL_BL_RED="\e[1;${RED};49m"
CL_BL_GREEN="\e[1;${GREEN};49m"
CL_BL_YELLOW="\e[1;${YELLOW};49m"
CL_BL_BLUE="\e[1;${BLUE};49m"
CL_BL_MAG="\e[1;${MAG};49m"
CL_BL_CYAN="\e[1;${CYAN};49m"
CL_BL_GREY="\e[1;${GREY};49m"

################################################################################

main() {
  if ! type -P go-fuzz-build &> /dev/null ; then
    error "This utility requires go-fuzz-build" $RED
    exit 1
  fi

  if ! grep -q 'go-fuzz' go.mod ; then
    go get github.com/dvyukov/go-fuzz/go-fuzz-dep &> /dev/null
  fi

  local src src_package src_name src_path src_func output func_min

  show "\nBuilding archives for packages…\n"

  for src in $(find . -name "fuzz.go") ; do
    src_package=$(dirname "$src" | sed 's#\.\/##')
    src_path=$(echo "$src" | sed 's#\/fuzz.go##' | sed 's#\.\/##')
    src_name=$(echo "$src_path" | sed 's#\/#-#g')

    if [[ -n "$1" && "$1" != "$src_package" ]] ; then
      continue
    fi

    while read src_func ; do
      src_func=$(echo "$src_func" | cut -f2 -d " " | sed 's/(//')
      
      if [[ "$src_func" == "Fuzz" ]] ; then
        showm " ${CL_GREY}∙ ${CL_BOLD}${src_package}${CL_NORM}… "
        output="${src_name}-fuzz.zip"
        go-fuzz-build -o "$output" "github.com/essentialkaos/ek/v12/${src_path}" &> /dev/null
      else
        showm " ${CL_GREY}∙ ${CL_BOLD}${src_package}${CL_NORM} ${CL_GREY}($src_func)${CL_NORM}… "
        func_min=$(echo "$src_func" | sed 's/Fuzz//' | tr '[A-Z]' '[a-z]')
        output="${src_name}-${func_min}-fuzz.zip"
        go-fuzz-build -func "$src_func" -o "$output" "github.com/essentialkaos/ek/v12/${src_path}" &> /dev/null
      fi

      if [[ $? -eq 0 ]] ; then
        show "${CL_GREEN}✔  ${CL_DARK}($output)${CL_NORM}"
      else
        show "✖ " $RED
      fi

    done < <(grep -Eo '^func Fuzz.*\(' $src)

  done

  show ""

  if grep -q 'go-fuzz' go.mod ; then
    git checkout go.* &> /dev/null
  fi
}

show() {
  if [[ -n "$2" ]] ; then
    echo -e "\e[${2}m${1}\e[0m"
  else
    echo -e "$*"
  fi
}

showm() {
  if [[ -n "$2" ]] ; then
    echo -e -n "\e[${2}m${1}\e[0m"
  else
    echo -e -n "$*"
  fi
}

error() {
  show "$@" 1>&2
}


################################################################################

main "$@"
