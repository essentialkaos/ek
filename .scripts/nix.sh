#! /bin/bash

########################################################################################

# Main func
#
# *: All arguments passed to script
#
main() {
  local dir="${1:-.}"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  testPackages "$dir"

  exit 0
}

# Test all packages from lis and save coverage info into the file
#
# 1: Dir with sources (String)
#
# Code: No
# Echo: No
testPackages() {
  local dir="$1"

  # Remove coverage output if exist
  rm -f coverage.tmp coverage.txt &> /dev/null

  # Fix coverage header
  echo "mode: count" > coverage.txt

  if [[ -z "$EK_TEST_PORT" ]] ; then
    export EK_TEST_PORT=8080
  fi

  local has_errors os_flag cover_enabled package_dir
  local package_list=".scripts/packages.list"

  while read -r package ; do

    read -r os_flag cover_enabled package_dir <<< "$package"

    if ! isOSFit "$os_flag" ; then
      continue
    fi

    if [[ "$cover_enabled" == "!" ]] ; then
      continue
    fi

    if [[ "$cover_enabled" == "-" ]] ; then
      if ! go test "$dir/$package_dir" -covermode=count -tags=unit ; then
        has_errors=true
      fi
    else
      if ! go test -covermode=count -tags=unit -coverprofile=coverage.tmp "$dir/$package_dir" ; then
        has_errors=true
      fi

      if [[ -f coverage.tmp ]] ; then
        grep -Ev '^mode:' coverage.tmp >> coverage.txt
        rm -f coverage.tmp
      fi
    fi

  done < <(awk 1 "$package_list")

  if [[ -n "$has_errors" ]] ; then
    exit 1
  fi
}

# Check if tests works on current runner OS
#
# 1: OS flag (String)
#
# Code: Yes
# Echo: No
isOSFit() {
  local os_flag="$1"

  if [[ "$os_flag" == "*" ]] ; then
    return 0
  fi

  if [[ "$os_flag" == "L" && "$RUNNER_OS" == "Linux" ]] ; then
    return 0
  fi

  if [[ "$os_flag" == "M" && "$RUNNER_OS" == "macOS" ]] ; then
    return 0
  fi

  return 1
}

########################################################################################

main "$@"
