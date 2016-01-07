test:
	@go test -v ./...

doc:
	@godoc -http=:6060 -goroot=../../../..
