all: test

install:
	go install -v ./...

format:
	for f in testdata/*.xml; do xmllint --format $$f --output $$f; done

test: install format
	go test -short ./...
	go test -coverprofile=cover.out ./...
