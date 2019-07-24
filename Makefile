MAIN_SRCS := $(shell find . -type f -name '*.go' -not -name '*_test.go')

.PHONY: run
run:
	go run $(MAIN_SRCS)

.PHONY: test
test:
	go test -count=1 $${TEST_ARGS} ./...

.PHONY: testrace
testrace:
	go test -count=1 $${TEST_ARGS} -race ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: vet
vet:
	go vet ./...
