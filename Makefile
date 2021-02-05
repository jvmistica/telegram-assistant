files = $(shell ls *go | grep -v *test.go)

run:
	@go run $(files)
