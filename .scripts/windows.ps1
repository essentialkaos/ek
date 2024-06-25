<# Download dependencies #>
function Get-Deps {
  go get -v golang.org/x/crypto/bcrypt
  go get -v github.com/essentialkaos/depsy
}

<# Try to install everything #>
function Install-All {
  go install ./...
}

Get-Deps
Install-All
