#! /bin/bash

########################################################################################

# Main func
#
# *: All arguments passed to script
#
main() {
  local version="$1"
  local dir="${2:-.}"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  installGoveralls
  makeLink "$version"
  testWithCover "$dir"
}

# Install goveralls
installGoveralls() {
  echo "Installing latest version of Goveralls..."
  go get -v github.com/mattn/goveralls
  echo ""
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

# Test packaages and save coverage info to file
#
# 1: Dir with sources (String)
testWithCover() {
  local dir="$1"

  # Remove coverage output if exist
  rm -f coverage.tmp coverage.txt &> /dev/null

  # Fix coverage header
  echo "mode: count" > coverage.txt

  if [[ -z "$EK_TEST_PORT" ]] ; then
    export EK_TEST_PORT=8080
  fi

  local has_errors cover_enabled package_dir
  local package_list=".travis/packages.list"

  while read package ; do

    read cover_enabled package_dir <<< "$package"

    if [[ "$cover_enabled" == "-" ]] ; then
      go test $dir/$package_dir -covermode=count

      if [[ $? -ne 0 ]] ; then
        has_errors=true
      fi

      continue
    fi

    go test -covermode=count -coverprofile=coverage.tmp $dir/$package_dir

    if [[ $? -ne 0 ]] ; then
      has_errors=true
    fi

    if [[ -f coverage.tmp ]] ; then
      egrep -v '^mode:' coverage.tmp >> coverage.txt
      rm -f coverage.tmp
    fi

  done < <(awk 1 $package_list)

  echo -e "\nSending data to Coveralls..."

  goveralls -service travis-ci -repotoken $COVERALLS_TOKEN -coverprofile coverage.txt

  exit 0
}

########################################################################################

main "$@"
