# supress output, run `make XXX V=` to be verbose
V := @

OUT_DIR := ./build
CERBERUS_OUT := $(OUT_DIR)/cerberus
CERBERUS_MAIN_PKG := ./cmd/cerberus
PROCESSOR_OUT := $(OUT_DIR)/processor
PROCESSOR_MAIN_PKG := ./cmd/processor
SYNC_OUT := $(OUT_DIR)/sync
SYNC_MAIN_PKG := ./cmd/sync
REWARDS_OUT := $(OUT_DIR)/rewards
REWARDS_MAIN_PKG := ./cmd/rewards

GOBIN := $(shell go env GOPATH)/bin

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

LDFLAGS = -s -w -X github.com/cosmos/cosmos-sdk/version.Name=furya \
	-X github.com/TessorNetwork/cerberus/internal/health.version=$(VERSION) \
	-X github.com/TessorNetwork/cerberus/internal/health.commit=$(COMMIT)


LINTER_NAME := golangci-lint
LINTER_VERSION := v1.29.0

MOCKGEN_NAME := mockgen
MOCKGEN_VERSION := v1.4.3

SWAGGER_NAME := swagger
SWAGGER_VERSION := v0.27.0

default: build

.PHONY: build
build:
	@echo BUILDING $(CERBERUS_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(CERBERUS_OUT) $(CERBERUS_MAIN_PKG)
	@echo BUILDING $(PROCESSOR_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(PROCESSOR_OUT) $(PROCESSOR_MAIN_PKG)
	@echo BUILDING $(SYNC_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(SYNC_OUT) $(SYNC_MAIN_PKG)
	@echo BUILDING $(REWARDS_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(REWARDS_OUT) $(REWARDS_MAIN_PKG)
	@echo DONE

.PHONY: linux
linux: export GOOS := linux
linux: export GOARCH := amd64
linux: LINUX_CERBERUS_OUT := $(CERBERUS_OUT)-$(GOOS)-$(GOARCH)
linux: LINUX_PROCESSOR_OUT := $(PROCESSOR_OUT)-$(GOOS)-$(GOARCH)
linux: LINUX_SYNC_OUT := $(SYNC_OUT)-$(GOOS)-$(GOARCH)
linux: LINUX_REWARDS_OUT := $(REWARDS_OUT)-$(GOOS)-$(GOARCH)
linux:
	@echo BUILDING $(LINUX_CERBERUS_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(LINUX_CERBERUS_OUT) $(CERBERUS_MAIN_PKG)
	@echo BUILDING $(LINUX_PROCESSOR_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(LINUX_PROCESSOR_OUT) $(PROCESSOR_MAIN_PKG)
	@echo BUILDING $(LINUX_SYNC_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(LINUX_SYNC_OUT) $(SYNC_MAIN_PKG)
	@echo BUILDING $(LINUX_REWARDS_OUT)
	$(V) CGO_ENABLED=0 go build -mod=vendor -ldflags "$(LDFLAGS)" -o $(LINUX_REWARDS_OUT) $(REWARDS_MAIN_PKG)
	@echo DONE

.PHONY: image
image:
	docker build -t cerberus-local -f scripts/Dockerfile .

.PHONY: clean
clean:
	$(V)rm -rf $(OUT_DIR)

.PHONY: test
test: GO_TEST_FLAGS := -race
test:
	$(V)go test -mod=vendor -v $(GO_TEST_FLAGS) $(GO_TEST_TAGS) ./...

.PHONY: fulltest
fulltest: GO_TEST_TAGS := --tags=integration
fulltest: test

.PHONY: lint
lint: check-linter-version
	$(V)$(LINTER_NAME) run --config configs/.golangci.yml

.PHONY: generate
generate: check-all
	$(V)go generate -mod=vendor -x ./...


.PHONY: vendor
vendor:
	$(V)go mod tidy
	$(V)go mod vendor

.PHONY: install-linter
install-linter: LINTER_INSTALL_PATH := $(GOBIN)/$(LINTER_NAME)
install-linter:
	@echo INSTALLING $(LINTER_INSTALL_PATH) $(LINTER_VERSION)
	$(V)curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) $(LINTER_VERSION)
	@echo DONE

.PHONY: install-mockgen
install-mockgen: MOCKGEN_INSTALL_PATH := $(GOBIN)/$(MOCKGEN_NAME)
install-mockgen:
	@echo INSTALLING $(MOCKGEN_INSTALL_PATH) $(MOCKGEN_NAME)
	# we need to change dir to use go modules without updating repo deps
	$(V)cd $(TMPDIR) && GO111MODULE=on go get github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)
	@echo DONE

.PHONY: install-swagger
install-swagger: SWAGGER_INSTALL_PATH := $(GOBIN)/$(SWAGGER_NAME)
install-swagger: UNAME_OS := $(shell uname -s)
install-swagger:
	@echo INSTALLING $(SWAGGER_INSTALL_PATH) $(SWAGGER_VERSION)
	# we need to change dir to use go modules without updating repo deps
	$(V)cd $(TMPDIR) && GO111MODULE=on go get github.com/go-swagger/go-swagger/cmd/swagger@$(SWAGGER_VERSION)
	@echo DONE

.PHONY: check-linter-version
check-linter-version: ACTUAL_LINTER_VERSION := $(shell $(LINTER_NAME) --version 2>/dev/null | awk '{print $$4}')
check-linter-version:
	$(V) [ -z $(ACTUAL_LINTER_VERSION) ] && \
	 echo 'Linter is not installed, run `make linter-install`' && \
	 exit 1 || true

	$(V)if [ v$(ACTUAL_LINTER_VERSION) != $(LINTER_VERSION) ] ; then \
		echo $(LINTER_NAME) is version v$(ACTUAL_LINTER_VERSION), want $(LINTER_VERSION) ; \
		echo 'Make sure $$GOBIN has precedence in $$PATH and' \
		'run `make linter-install` to install the correct version' ; \
        exit 1 ; \
	fi

.PHONY: check-mockgen-version
check-mockgen-version: ACTUAL_MOCKGEN_VERSION := $(shell $(MOCKGEN_NAME) --version 2>/dev/null)
check-mockgen-version:
	$(V) [ -z $(ACTUAL_MOCKGEN_VERSION) ] && \
	 echo 'Mockgen is not installed, run `make mockgen-install`' && \
	 exit 1 || true

	$(V)if [ $(ACTUAL_MOCKGEN_VERSION) != $(MOCKGEN_VERSION) ] ; then \
		echo $(MOCKGEN_NAME) is version $(ACTUAL_MOCKGEN_VERSION), want $(MOCKGEN_VERSION) ; \
		echo 'Make sure $$GOBIN has precedence in $$PATH and' \
		'run `make mockgen-install` to install the correct version' ; \
        exit 1 ; \
	fi

.PHONY: check-swagger-version
check-swagger-version: ACTUAL_SWAGGER_VERSION := $(shell $(SWAGGER_NAME) version 2>/dev/null | grep version | cut -c 10-17)
# hack version, see https://github.com/go-swagger/go-swagger/issues/1712#issuecomment-422981313
check-swagger-version: WANT_SWAGGER_VERSION := $(SWAGGER_VERSION)
check-swagger-version:
	$(V) [ -z $(ACTUAL_SWAGGER_VERSION) ] && \
	 echo 'Swagger is not installed, run `make install-swagger`' && \
	 exit 1 || true

	$(V)if [ $(ACTUAL_SWAGGER_VERSION) != $(WANT_SWAGGER_VERSION) ] ; then \
		echo $(SWAGGER_NAME) is version $(ACTUAL_SWAGGER_VERSION), want $(WANT_SWAGGER_VERSION) ; \
		echo 'Make sure $$GOBIN has precedence in $$PATH and' \
		'run `make install-swagger` to install the correct version' ; \
        exit 1 ; \
	fi

.PHONY: check-all
check-all: check-swagger-version check-mockgen-version

.PHONY: install-all
install-all: install-linter install-mockgen install-swagger