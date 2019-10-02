BINARY_NAME=chaos-pong

all: test build

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./...

run: build
	./$(BINARY_NAME)

fmt:
	gofmt -s -w .
