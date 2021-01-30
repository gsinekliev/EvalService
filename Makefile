.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test ./service

.PHONY: lint
lint:
	golangci-lint run

.PHONY: format
format:
	 gofmt -w .