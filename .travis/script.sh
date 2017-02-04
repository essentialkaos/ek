#! /bin/bash

########################################################################################

# List of packages excluded from coverage export
EXCLUDED_PACKAGES=("fsutil system terminal usage netutil")

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
  go get -v github.com/mattn/goveralls
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

  echo "Created link $GOPATH/src/${pkg_dir} → $GOPATH/src/github.com/essentialkaos/ek"

  ln -sf $GOPATH/src/github.com/essentialkaos/ek $GOPATH/src/${pkg_dir}
}

# Test packaages and save coverage info to file
#
# 1: Dir with sources (String)
#
testWithCover() {
  local dir="$1"

  local pkg has_errors excl_pkg skip_cover

  rm -f coverage.tmp coverage.txt &> /dev/null

  if [[ -z "$EK_TEST_PORT" ]] ; then
    export EK_TEST_PORT=8080
  fi

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

  goveralls -service travis-ci -repotoken $COVERALLS_TOKEN -coverprofile coverage.txt
}

########################################################################################

main "$@"
