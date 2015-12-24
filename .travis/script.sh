#! /bin/bash

########################################################################################

# Main func
#
# *: All arguments passed to script
#
main() {
  local dir="$1"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  makeLink
  testWithCover "$dir"
}

# Create links for pkg.re import paths
#
makeLink() {
  # TravicCI download last stable version of ek, but it not ok
  # remove downloaded version for linking with current version for test
  if [[ -e $GOPATH/src/pkg.re/essentialkaos/ek.v1 ]] ; then
    echo "Directory pkg.re/essentialkaos/ek.v1 removed"
    rm -rf $GOPATH/src/pkg.re/essentialkaos/ek.v1
  fi

  mkdir -p $GOPATH/src/pkg.re/essentialkaos
  
  echo "Created link $GOPATH/src/pkg.re/essentialkaos/ek.v1 -> $GOPATH/src/github.com/essentialkaos/ek"
  ln -sf $GOPATH/src/github.com/essentialkaos/ek $GOPATH/src/pkg.re/essentialkaos/ek.v1
}

# Test packaages and save coverage info to file
#
# 1: Dir with sources (String)
#
testWithCover() {
  local dir="$1"

  local pkg has_errors

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    if [[ ! -d $dir/$pkg ]] ; then
      continue
    fi

    # fsutil currently is hard to test
    # we test this package by hands
    if [[ "$pkg" == "fsutil" ]] ; then
      continue
    fi

    go test -coverprofile=coverage.tmp -covermode=atomic $dir/$pkg

    if [[ $? -ne 0 ]] ; then
      has_errors=true
    fi

    if [[ -f coverage.tmp ]] ; then
      cat coverage.tmp >> coverage.txt
      rm -f coverage.tmp
    fi
  done

  if [[ $has_errors ]] ; then
    exit 1
  fi

  exit 0
}

########################################################################################

main $@
