build:
	go build cmd/tcplat.go

test-probe:
	go generate ./...
	go test -v -count=1 ./internal/probe/

