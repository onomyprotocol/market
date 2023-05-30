PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell git describe --abbrev=6 --dirty --always --tags)
COMMIT := $(shell git log -1 --format='%H')
IMPORT_PREFIX=github.com/pendulum-labs
SCAN_FILES := $(shell find . -type f -name '*.go' -not -name '*mock.go' -not -name '*_gen.go' -not -path "*/vendor/*")

.PHONY: build
build: go.sum
		go build ./cmd/marketd

.PHONY: build_standalone
build_standalone: go.sum
		go build ./cmd/market_standaloned

.PHONY: test
test:
	@go test -mod=readonly $(PACKAGES)

.PHONY: lint
lint:
	golangci-lint -c .golangci.yml run
	gofmt -d -s $(SCAN_FILES)

.PHONY: format
format:
	gofumpt -lang=1.6 -extra -s -w $(SCAN_FILES)
	gogroup -order std,other,prefix=$(IMPORT_PREFIX) -rewrite $(SCAN_FILES)

###############################################################################
###                                Protobuf                                 ###
###############################################################################

.PHONY: proto-gen-all
proto-gen-all: proto-gen-go proto-gen-openapi

.PHONY: proto-gen-openapi
proto-gen-openapi:
	bash ./dev/scripts/protoc-swagger-gen.sh

.PHONY: proto-gen-go
proto-gen-go:
	bash ./dev/scripts/protocgen.sh
	go mod tidy
	make format

.PHONY: proto-lint
proto-lint:
	buf lint proto --config buf.yaml
