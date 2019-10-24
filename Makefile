BINARY_NAME=chaos-pong
CLIENT_BINARY=client
SERVER_BINARY=server
.PHONY: server client
all: test build

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./...

run: build
	./$(BINARY_NAME)

client:
	cd ${CLIENT_BINARY}; \
	go build -o ${CLIENT_BINARY} -v; \
	./${CLIENT_BINARY}

server:
	cd ${SERVER_BINARY}; \
	go build -o ${SERVER_BINARY} -v; \
	./${SERVER_BINARY}

fmt:
	gofmt -s -w .
