#! /bin/bash

########################################################################################

# Main func
#
# *: All arguments passed to script
#
main() {
  local version="$1"
  local dir="${1:-.}"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  makeLink "$version"
  testWithCover "$dir"
}

# Create links for pkg.re import paths
makeLink() {
  local version="$1"
  local pkg_dir="pkg.re/essentialkaos/ek.v${version}"

  # TravicCI download last stable version of ek, but it not ok
  # remove downloaded version for linking with current version for test
  if [[ -e $GOPATH/src/${pkg_dir} ]] ; then
    echo "Directory ${pkg_dir} removed"
    rm -rf $GOPATH/src/${pkg_dir}
  fi

  mkdir -p $GOPATH/src/pkg.re/essentialkaos

  echo "Created link $GOPATH/src/${pkg_dir} -> $GOPATH/src/github.com/essentialkaos/ek"

  ln -sf $GOPATH/src/github.com/essentialkaos/ek $GOPATH/src/${pkg_dir}
}

# Test packaages and save coverage info to file
#
# 1: Dir with sources (String)
#
testWithCover() {
  local dir="$1"

  local pkg has_errors

  pushd "$dir" &> /dev/null

    EK_TEST_PORT=8080 gocov test ./... | gocov report

    if [[ $? -ne 0 ]] ; then
      has_errors=true
    fi

  popd &> /dev/null

  if [[ -n "$has_errors" ]] ; then
    exit 1
  fi

  exit 0
}

########################################################################################

main "$@"
