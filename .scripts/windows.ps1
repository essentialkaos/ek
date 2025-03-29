$env:EK_TEST_PORT = "8080"

<# Download dependencies #>
function Get-Deps {
  go get -v golang.org/x/crypto/bcrypt
  go get -v github.com/essentialkaos/depsy
}

<# Try to install everything #>
function Install-All {
  go install ./...
}

<# Runt tests #>
function Run-Tests {
  go test ./...
}

Get-Deps
Install-All
Run-Tests
