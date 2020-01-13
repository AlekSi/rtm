all: test

install:
	go install -v ./...

test: install
	go test -coverprofile=cover.out ./...
