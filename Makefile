.PHONY: test
test:
	go test -race -cover -parallel 4 ./...

.PHONY: run-local
run-local:
	go run cmd/gateway/main.go v1
