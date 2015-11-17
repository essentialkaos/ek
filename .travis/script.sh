#! /bin/bash

########################################################################################

main() {
  local dir="$1"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  local pkg has_errors

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    if [[ ! -d $dir/$pkg ]] ; then
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
