#!/bin/bash

main() {
  makeLink "$1"
  downloadDeps
  checkInstall
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

  echo -e "Created link $GOPATH/src/${pkg_dir} → $GOPATH/src/github.com/essentialkaos/ek\n"

  ln -sf $GOPATH/src/github.com/essentialkaos/ek $GOPATH/src/${pkg_dir}
}

# Download required dependencies
downloadDeps() {
  go get -v -d golang.org/x/crypto/bcrypt
}

# Check package installation
checkInstall() {
  go install ./...
  exit $?
}

########################################################################################

main "$@"
