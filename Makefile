format:
	@go fmt ./...

test:
	@go test -cover -v ./...

cover:
	@go test ./... -coverprofile cover.out
	@go tool cover -html cover.out -o cover.html

run:
	@go run main.go

build:
	@go build -v ./...

install:
	@go install -v ./...

vulncheck:
	@govulncheck ./...
