build:
	go build -o bin/lestgo

run: build
	./bin/lestgo

test:
	go test -v ./...