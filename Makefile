files = $(shell ls *.go | grep -v "_test.go")

run:
	@go run $(files)
