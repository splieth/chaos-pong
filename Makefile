BINARY_NAME  := chaos-pong
CLIENT_DIR   := client
SERVER_DIR   := server

.PHONY: all build test run client server fmt vet clean help

all: test build ## Run tests and build

# ── Build ────────────────────────────────────────────────────

build: ## Build the single-player game
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) -v

build-client: ## Build the multiplayer client
	@echo "Building client..."
	@cd $(CLIENT_DIR) && go build -o $(CLIENT_DIR) -v

build-server: ## Build the multiplayer server
	@echo "Building server..."
	@cd $(SERVER_DIR) && go build -o $(SERVER_DIR) -v

build-all: build build-client build-server ## Build all binaries

# ── Run ──────────────────────────────────────────────────────

run: build ## Build and run the single-player game
	@./$(BINARY_NAME)

client: build-client ## Build and run the multiplayer client
	@cd $(CLIENT_DIR) && ./$(CLIENT_DIR)

server: build-server ## Build and run the multiplayer server
	@cd $(SERVER_DIR) && ./$(SERVER_DIR)

# ── Quality ──────────────────────────────────────────────────

test: ## Run all tests
	@go test ./...

test-v: ## Run all tests (verbose)
	@go test -v ./...

test-cover: ## Run tests with coverage report
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@rm -f coverage.out

fmt: ## Format all Go source files
	@gofmt -s -w .

vet: ## Run static analysis
	@go vet ./...

lint: fmt vet ## Format and vet

# ── Misc ─────────────────────────────────────────────────────

clean: ## Remove build artifacts
	@rm -f $(BINARY_NAME)
	@rm -f $(CLIENT_DIR)/$(CLIENT_DIR)
	@rm -f $(SERVER_DIR)/$(SERVER_DIR)
	@echo "Clean."

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
