SOURCE_FILES := $(shell find . -type f -name '*.go')

VERSION ?= $(shell git describe | cut -c2-)


policy.wasm: $(SOURCE_FILES) go.mod go.sum
	GOOS=wasip1 GOARCH=wasm go build -gcflags=all="-B" -ldflags="-w -s" -o policy.wasm
	wasm-opt --enable-bulk-memory -Oz -o policy.wasm policy.wasm 


annotated-policy.wasm: policy.wasm metadata.yml
	kwctl annotate -m metadata.yml -u README.md -o annotated-policy.wasm policy.wasm

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

.PHONY: clean
clean:
	go clean
	rm -f policy.wasm annotated-policy.wasm

test:
	echo "No tests implemented - refer to e2e-tests"

.PHONY: e2e-tests
e2e-tests: annotated-policy.wasm
	bats e2e.bats
