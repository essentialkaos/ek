########################################################################################

export EK_TEST_PORT=8080

########################################################################################

.PHONY = all test fmt deps deps-test

########################################################################################

deps:
	go get -v pkg.re/essentialkaos/go-linenoise.v3
	go get -v golang.org/x/crypto/bcrypt

deps-test:
	go get -v github.com/axw/gocov/gocov
	go get -v pkg.re/check.v1

test:
	go get -v pkg.re/check.v1
	go test -covermode=count ./...

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;
