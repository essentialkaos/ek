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

# shellcheck disable=SC2034
CL_NORM="\e[${NORM}m"
# shellcheck disable=SC2034
CL_BOLD="\e[${BOLD}m"
# shellcheck disable=SC2034
CL_ITLC="\e[${ITLC}m"
# shellcheck disable=SC2034
CL_UNLN="\e[${UNLN}m"
# shellcheck disable=SC2034
CL_RED="\e[${RED}m"
# shellcheck disable=SC2034
CL_GREEN="\e[${GREEN}m"
# shellcheck disable=SC2034
CL_YELLOW="\e[${YELLOW}m"
# shellcheck disable=SC2034
CL_BLUE="\e[${BLUE}m"
# shellcheck disable=SC2034
CL_MAG="\e[${MAG}m"
# shellcheck disable=SC2034
CL_CYAN="\e[${CYAN}m"
# shellcheck disable=SC2034
CL_GREY="\e[${GREY}m"
# shellcheck disable=SC2034
CL_DARK="\e[${DARK}m"

################################################################################

main() {
  if ! type -P go-fuzz-build &> /dev/null ; then
    show "\nInstalling go-fuzz tooling…"
    go install github.com/dvyukov/go-fuzz/go-fuzz@latest &> /dev/null
    go install github.com/dvyukov/go-fuzz/go-fuzz-build@latest &> /dev/null
    show "go-fuzz tooling successfully installed\n" $GREEN
  fi

  if ! grep -q 'go-fuzz' go.mod ; then
    go get github.com/dvyukov/go-fuzz/go-fuzz-dep &> /dev/null
  fi

  local src src_package src_name src_path src_func output func_min

  show "\nBuilding archives for packages…\n"

  while read -r src ; do
    src_package=$(dirname "$src" | sed 's#\.\/##')
    src_path=$(echo "$src" | sed 's#\/fuzz.go##' | sed 's#\.\/##')
    src_name="${src_path//\//-}"

    if [[ -n "$1" && "$1" != "$src_package" ]] ; then
      continue
    fi

    while read -r src_func ; do
      src_func=$(echo "$src_func" | cut -f2 -d " " | sed 's/(//')
      
      if [[ "$src_func" == "Fuzz" ]] ; then
        showm " ${CL_GREY}∙ ${CL_BOLD}${src_package}${CL_NORM}… "
        output="${src_name}-fuzz.zip"
        go-fuzz-build -o "$output" "github.com/essentialkaos/ek/v13/${src_path}" &> /dev/null
      else
        showm " ${CL_GREY}∙ ${CL_BOLD}${src_package}${CL_NORM} ${CL_GREY}($src_func)${CL_NORM}… "
        func_min=$(echo "$src_func" | sed 's/Fuzz//' | tr '[:upper:]' '[:lower:]')
        output="${src_name}-${func_min}-fuzz.zip"
        go-fuzz-build -func "$src_func" -o "$output" "github.com/essentialkaos/ek/v13/${src_path}" &> /dev/null
      fi

      # shellcheck disable=SC2181
      if [[ $? -eq 0 ]] ; then
        show "${CL_GREEN}✔  ${CL_DARK}($output)${CL_NORM}"
      else
        show "✖ " $RED
      fi

    done < <(grep -Eo '^func Fuzz.*\(' "$src")
  done < <(find . -name "fuzz.go")

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
