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

  pushd $dir &> /dev/null

    EK_TEST_PORT=8080 gocov test ./... | gocov report

    if [[ $? -ne 0 ]] ; then
      has_errors=true
    fi

  popd &> /dev/null

  if [[ $has_errors ]] ; then
    exit 1
  fi

  exit 0
}

########################################################################################

main $@
