#! /bin/bash

########################################################################################

# Current major version
VERSION="4"

# Pkg.re package path
PKGRE_PKG="pkg.re/essentialkaos/ek.v${VERSION}"

# List of packages excluded from coverage export
EXCLUDED_PACKAGES=("fsutil system terminal usage netutil")

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

  local pkg has_errors excl_pkg skip_cover

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    skip_cover=""

    if [[ ! -d $dir/$pkg ]] ; then
      continue
    fi

    for excl_pkg in ${EXCLUDED_PACKAGES[@]} ; do
      skip_cover=true
    done

    if [[ $skip_cover ]] ; then
      go test $dir/$pkg

      if [[ $? -ne 0 ]] ; then
        has_errors=true
      fi

      continue
    fi

    go test -covermode=count -coverprofile=coverage.tmp $dir/$pkg

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

  wc -l coverage.txt

  $HOME/gopath/bin/goveralls -coverprofile=coverage.txt -service=travis-ci

  exit 0
}

########################################################################################

main $@
