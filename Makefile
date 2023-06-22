run:
	go run ./cmd/web
unit-tests:
	go test -v ./cmd/web
clean-cache:
	go clean -testcache