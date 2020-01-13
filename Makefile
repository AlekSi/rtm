all: test

test:
	go test -v -coverprofile=cover.out
