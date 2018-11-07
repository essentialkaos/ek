########################################################################################

export EK_TEST_PORT=8080

########################################################################################

.DEFAULT_GOAL := help
.PHONY = test fmt deps deps-test clean help

########################################################################################

deps: ## Download dependencies
	git config --global http.https://pkg.re.followRedirects true
	go get -v pkg.re/essentialkaos/go-linenoise.v3
	go get -v golang.org/x/crypto/bcrypt

deps-test: ## Download dependencies for tests
	git config --global http.https://pkg.re.followRedirects true
	go get -v github.com/axw/gocov/gocov
	go get -v pkg.re/check.v1

test: ## Run tests
	git config --global http.https://pkg.re.followRedirects true
	go get -v pkg.re/check.v1
	go test -covermode=count ./...

gen-fuzz: ## Generate go-fuzz archives for all packages
	bash .scripts/fuzz-gen.sh ${PACKAGE}

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Run 'go vet' over sources
	go vet -printfuncs=Print,Printf,Println,Fprint,Fprintf,Fprintln,Sprint,Sprintf,Sprintln,Error,Errorf,Fatal,Fatalf,Panic,Panicf,Panicln ./...

clean: ## Remove all generated data
	rm -f *.zip

help: ## Show this info
	@echo -e '\nSupported targets:\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
