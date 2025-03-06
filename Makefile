################################################################################

export EK_TEST_PORT=8080
export CGO_ENABLED=1

################################################################################

ifdef VERBOSE ## Print verbose information (Flag)
VERBOSE_FLAG = -v
endif

ifdef PROXY ## Force proxy usage for downloading dependencies (Flag)
export GOPROXY=https://proxy.golang.org/cached-only,direct
endif

export CGO_ENABLED=1

MAKEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
GITREV ?= $(shell test -s $(MAKEDIR)/.git && git rev-parse --short HEAD)

################################################################################

.DEFAULT_GOAL := help
.PHONY = fmt vet deps update test init vendor gen-fuzz tidy mod-init mod-update mod-download mod-vendor help

################################################################################

init: mod-init ## Initialize new module

deps: mod-download ## Download dependencies

update: mod-update ## Update dependencies to the latest versions

vendor: mod-vendor ## Make vendored copy of dependencies

test: ## Run tests
	@echo "[36;1mStarting tests…[0m"
ifdef COVERAGE_FILE ## Save coverage data into file (String)
	@go test $(VERBOSE_FLAG) -tags=unit -covermode=count -coverprofile=$(COVERAGE_FILE) ./...
else
	@go test $(VERBOSE_FLAG) -tags=unit -covermode=count ./...
endif

gen-fuzz: ## Generate archives for fuzz testing
	@which go-fuzz-build &>/dev/null || go install github.com/dvyukov/go-fuzz/go-fuzz-build@latest
	@echo "[36;1mGenerating fuzzing data…[0m"
	@go-fuzz-build -o cron-fuzz.zip github.com/essentialkaos/ek/cron
	@go-fuzz-build -o fmtc-fuzz.zip github.com/essentialkaos/ek/fmtc
	@go-fuzz-build -o knf-fuzz.zip github.com/essentialkaos/ek/knf
	@go-fuzz-build -o strutil-fuzz.zip github.com/essentialkaos/ek/strutil
	@go-fuzz-build -o system-fuzz.zip github.com/essentialkaos/ek/system
	@go-fuzz-build -o timeutil-fuzz.zip github.com/essentialkaos/ek/timeutil
	@go-fuzz-build -o version-fuzz.zip github.com/essentialkaos/ek/version

tidy: ## Cleanup dependencies
	@echo "[32m•[0m[90m•[0m [36;1mTidying up dependencies…[0m"
ifdef COMPAT ## Compatible Go version (String)
	@go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT) -go=$(COMPAT)
else
	@go mod tidy $(VERBOSE_FLAG)
endif
	@echo "[32m••[0m [36;1mUpdating vendored dependencies…[0m"
	@test -d vendor && rm -rf vendor && go mod vendor $(VERBOSE_FLAG) || :

mod-init:
	@echo "[32m•[0m[90m••[0m [36;1mModules initialization…[0m"
	@rm -f go.mod go.sum
ifdef MODULE_PATH ## Module path for initialization (String)
	@go mod init $(MODULE_PATH)
else
	@go mod init
endif

	@echo "[32m••[0m[90m•[0m [36;1mDependencies cleanup…[0m"
ifdef COMPAT ## Compatible Go version (String)
	@go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT) -go=$(COMPAT)
else
	@go mod tidy $(VERBOSE_FLAG)
endif
	@echo "[32m•••[0m [36;1mStripping toolchain info…[0m"
	@grep -q 'toolchain ' go.mod && go mod edit -toolchain=none || :

mod-update:
	@echo "[32m•[0m[90m•••[0m [36;1mUpdating dependencies…[0m"
ifdef UPDATE_ALL ## Update all dependencies (Flag)
	@go get -u $(VERBOSE_FLAG) all
else
	@go get -u $(VERBOSE_FLAG) ./...
endif

	@echo "[32m••[0m[90m••[0m [36;1mStripping toolchain info…[0m"
	@grep -q 'toolchain ' go.mod && go mod edit -toolchain=none || :

	@echo "[32m•••[0m[90m•[0m [36;1mDependencies cleanup…[0m"
ifdef COMPAT
	@go mod tidy $(VERBOSE_FLAG) -compat=$(COMPAT)
else
	@go mod tidy $(VERBOSE_FLAG)
endif

	@echo "[32m••••[0m [36;1mUpdating vendored dependencies…[0m"
	@test -d vendor && rm -rf vendor && go mod vendor $(VERBOSE_FLAG) || :

mod-download:
	@echo "[36;1mDownloading dependencies…[0m"
	@go mod download

mod-vendor:
	@echo "[36;1mVendoring dependencies…[0m"
	@rm -rf vendor && go mod vendor $(VERBOSE_FLAG) || :

fmt: ## Format source code with gofmt
	@echo "[36;1mFormatting sources…[0m"
	@find . -name "*.go" -exec gofmt -s -w {} \;

vet: ## Runs 'go vet' over sources
	@echo "[36;1mRunning 'go vet' over sources…[0m"
	@go vet -composites=false -printfuncs=LPrintf,TLPrintf,TPrintf,log.Debug,log.Info,log.Warn,log.Error,log.Critical,log.Print ./...

help: ## Show this info
	@echo -e '\n\033[1mTargets:\033[0m\n'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[33m%-8s\033[0m  %s\n", $$1, $$2}'
	@echo -e '\n\033[1mVariables:\033[0m\n'
	@grep -E '^ifdef [A-Z_]+ .*?## .*$$' $(abspath $(lastword $(MAKEFILE_LIST))) \
		| sed 's/ifdef //' \
		| sort -h \
		| awk 'BEGIN {FS = " .*?## "}; {printf "  \033[32m%-13s\033[0m  %s\n", $$1, $$2}'
	@echo -e ''

################################################################################
