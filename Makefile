################################################################################

export EK_TEST_PORT=8080

################################################################################

.DEFAULT_GOAL := help
.PHONY = test fmt deps deps-test mod-init mod-update mod-download mod-vendor clean help

################################################################################

deps: mod-download ## Download dependencies

deps-test: deps ## Download dependencies for tests

mod-init: ## Initialize new module
	go mod init
	go mod tidy

mod-update: ## Update modules to their latest versions
	go get -u
	go mod tidy

mod-download: ## Download modules to local cache
	go mod download

mod-vendor: ## Make vendored copy of dependencies
	go mod vendor

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
