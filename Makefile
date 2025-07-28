build:
	go build -o gopack .

test:
	go test ./...

test-verbose:
	go test -v ./...

clean:
	rm -f gopack
	rm -rf dist/

example:
	./gopack examples/app.js

example-prod:
	./gopack --entry examples/app.js --mode production

install:
	go install .

.PHONY: build test test-verbose clean example example-prod install
