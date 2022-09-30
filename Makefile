# Makefile

init:
	go get github.com/golang/protobuf/protoc-gen-go@v1.5.2
	go mod download

generate:
	go generate ./...

# 사용할 때에는 터미널에서 'make run'을 입력하면 된다.
run: generate
	@go run cmd/main.go