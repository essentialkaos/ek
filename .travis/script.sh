#! /bin/bash

########################################################################################

main() {
  local dir="$1"

  [[ ! -d $dir ]] && exit 1

  local pkg has_errors

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    [[ ! -d $dir/$pkg ]] && continue

    go test -coverprofile=coverage.tmp -covermode=atomic $dir/$pkg

    [[ $? -ne 0 ]] && has_errors=true

    if [[ -f coverage.tmp ]] ; then
      cat coverage.tmp >> coverage.txt
      rm -f coverage.tmp
    fi
  done

  [[ $has_errors ]] && exit 1
}

########################################################################################

main $@
