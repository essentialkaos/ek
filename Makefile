################################################################################

export EK_TEST_PORT=8080

ifdef VERBOSE ## Print verbose information (Flag)
VERBOSE_FLAG = -v
endif

################################################################################

.DEFAULT_GOAL := help
.PHONY = test fmt vet deps update mod-update mod-download mod-vendor gen-fuzz clean help

################################################################################

deps: mod-download ## Download dependencies

update: mod-update ## Update dependencies to the latest versions

vendor: mod-vendor ## Make vendored copy of dependencies

test: ## Run tests
	go test -covermode=count -tags=unit ./...

gen-fuzz: ## Generate go-fuzz archives for all packages
	bash .scripts/fuzz-gen.sh ${PACKAGE}

mod-update:
ifdef UPDATE_ALL ## Update all dependencies (Flag)
	go get -u $(VERBOSE_FLAG) all
else
	go get -u $(VERBOSE_FLAG) ./...
endif

ifdef COMPAT ## Compatible Go version (String)
	go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT)
else
	go mod tidy $(VERBOSE_FLAG)
endif

	test -d vendor && go mod vendor $(VERBOSE_FLAG) || :

mod-download:
	go mod download

mod-vendor:
	go mod vendor $(VERBOSE_FLAG)

fmt: ## Format source code with gofmt
	find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Runs go vet over sources
	go vet -composites=false -printfuncs=LPrintf,TLPrintf,TPrintf,log.Debug,log.Info,log.Warn,log.Error,log.Critical,log.Print ./...

clean: ## Remove all generated data
	rm -f *.zip

help: ## Show this info
	@echo -e '\n\033[1mTargets:\033[0m\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-14s\033[0m %s\n", $$1, $$2}'
	@echo -e '\n\033[1mVariables:\033[0m\n'
	@grep -E '^ifdef [A-Z_]+ .*?## .*$$' $(abspath $(lastword $(MAKEFILE_LIST))) \
		| sed 's/ifdef //' \
		| awk 'BEGIN {FS = " .*?## "}; {printf "  \033[32m%-14s\033[0m %s\n", $$1, $$2}'
	@echo -e ''
