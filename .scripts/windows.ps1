<# Download dependencies #>
Function Download-Deps {
  go get -v golang.org/x/crypto/bcrypt
  go get -v github.com/essentialkaos/depsy
}

<# Try to install everything #>
Function Check-Install {
  go install ./...
}

Download-Deps
Check-Install
