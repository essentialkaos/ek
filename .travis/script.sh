#! /bin/bash

########################################################################################

main() {
  local dir="$1"
  local pkg

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    [[ ! -d $dir/$pkg ]] && continue

    go test -coverprofile=coverage.tmp -covermode=atomic $dir/$pkg

    if [[ -f coverage.tmp ]] ; then
      cat coverage.tmp >> coverage.txt
      rm -f coverage.tmp
    fi
  done
}

########################################################################################

main $@
