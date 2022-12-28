.PHONY: test
test:
	go test -race -cover -parallel 4 ./...

.PHONY: run-local
run-local:
	@go run main.go

.PHONY: generate
generate:
	go generate ${CURDIR}/...
