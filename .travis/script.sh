#! /bin/bash

########################################################################################

# Current major version
VERSION="4"

# Pkg.re package path
PKGRE_PKG="pkg.re/essentialkaos/ek.v${VERSION}"

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
makeLink() {
  # TravicCI download last stable version of ek, but it not ok
  # remove downloaded version for linking with current version for test
  if [[ -e $GOPATH/src/${PKGRE_PKG} ]] ; then
    echo "Directory ${PKGRE_PKG} removed"
    rm -rf $GOPATH/src/${PKGRE_PKG}
  fi

  mkdir -p $GOPATH/src/pkg.re/essentialkaos
  
  echo "Created link $GOPATH/src/${PKGRE_PKG} -> $GOPATH/src/github.com/essentialkaos/ek"
  ln -sf $GOPATH/src/github.com/essentialkaos/ek $GOPATH/src/${PKGRE_PKG}
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

    # fsutil and system packages currently is hard to implement unit testing
    # we test this package by hands
    if [[ "$pkg" == "fsutil" || "$pkg" == "system" ]] ; then
      go test $dir/$pkg

      if [[ $? -ne 0 ]] ; then
        has_errors=true
      fi

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
