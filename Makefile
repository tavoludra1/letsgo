build:
	
	@go build -o bin/letsgo

run: build
	@./bin/letsgo

test:
	@go test -v ./...