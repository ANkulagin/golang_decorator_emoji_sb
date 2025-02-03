BINARY_NAME = decorator
PKG = ./...
COVERAGE_FILE = coverage.out
COVERAGE_HTML = coverage.html

.PHONY: all build run test coverage help

build: ## Build the project binary
	@echo "Building the project..."
	go build  -o $(BINARY_NAME) cmd/decorator/main.go


run: build ## Run the built binary (decorator)
	@echo "Running the application..."
	./$(BINARY_NAME)


test: ## Run unit tests with verbose output
	@echo "Running tests..."
	go test $(PKG) -v


coverage: ## Generate code coverage report (HTML)
	@echo "Running tests with coverage..."
	go test $(PKG) -coverprofile=$(COVERAGE_FILE)
	@echo "Generating HTML coverage report..."
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

help: ## Show a certificate of affordable commands
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<цель>\033[0m\n\nДоступные цели:\n"} /^[a-zA-Z_-]+:.*##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
