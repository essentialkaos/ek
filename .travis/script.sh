#! /bin/bash

########################################################################################

main() {
  local dir="$1"

  if [[ ! -d $dir ]] ; then
    exit 1
  fi

  local pkg has_errors

  local count=0

  rm -f coverage.tmp coverage.txt &> /dev/null

  for pkg in $(ls -1 $dir) ; do
    if [[ ! -d $dir/$pkg ]] ; then
      continue
    fi

    while : ; do
      testPackage "$dir" "$pkg"

      if [[ $? -ne 0 ]] ; then
        # Workaround for perediocal errors on TravisCI containers
        if [[ "$pkg" == "req" && $count -ne 3 ]] ; then
          ((count++))
          continue
        else
          has_errors=true
          break
        fi
      fi

      break
    done

    if [[ -f coverage.tmp ]] ; then
      cat coverage.tmp >> coverage.txt
    fi
  done

  if [[ $has_errors ]] ; then
    exit 1
  fi
}

testPackage() {
  local dir="$1"
  local pkg="$2"

  rm -f coverage.tmp

  go test -coverprofile=coverage.tmp -covermode=atomic $dir/$pkg

  return $?
}

########################################################################################

main $@
