format:
	@go fmt ./...

test:
	@go test -cover -v ./...

cover:
	@go test ./... -coverprofile cover.out

run:
	@go run main.go

build:
	@go build -v ./...

install:
	@go install -v ./...
