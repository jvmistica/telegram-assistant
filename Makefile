.PHONY: format test cover run build install vulncheck trivyscan

format:
	@go fmt ./...

test:
	@go test ./... -coverprofile cover.out

cover:
	@go tool cover -html cover.out -o cover.html

run:
	@go run main.go

build:
	@go build -v ./...

install:
	@go install -v ./...

vulncheck:
	@govulncheck ./...

trivyscan:
	@trivy fs . --scanners=vuln