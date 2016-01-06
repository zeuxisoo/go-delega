all:
	@echo "make test"

test:
	@go test ./...

doc:
	@godoc -http=:6060 -goroot=../../../..
