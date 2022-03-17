#!/bin/bash

main() {
  downloadDeps
  checkInstall
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
