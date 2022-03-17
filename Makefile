################################################################################

export EK_TEST_PORT=8080

################################################################################

.DEFAULT_GOAL := help
.PHONY = test fmt deps deps-test clean help

################################################################################

deps: ## Download dependencies
	go get -v github.com/essentialkaos/go-linenoise
	go get -v golang.org/x/crypto/bcrypt

deps-test: ## Download dependencies for tests
	go get -v github.com/axw/gocov/gocov
	go get -v github.com/essentialkaos/check
	go get -v golang.org/x/sys/unix

test: ## Run tests
	go test -covermode=count -tags=unit ./...

gen-fuzz: ## Generate go-fuzz archives for all packages
	bash .scripts/fuzz-gen.sh ${PACKAGE}

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Runs go vet over sources
	go vet -composites=false -printfuncs=LPrintf,TLPrintf,TPrintf,log.Debug,log.Info,log.Warn,log.Error,log.Critical,log.Print ./...

clean: ## Remove all generated data
	rm -f *.zip

help: ## Show this info
	@echo -e '\nSupported targets:\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-12s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
