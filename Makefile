```makefile
.PHONY: all build clean test lint run docker-build install-deps dev bench release help

# ==============================================================================
# КОНФИГУРАЦИЯ
# ==============================================================================

BINARY_NAME := giftcalc
VERSION := 1.0.0
BUILD_DIR := bin
COVERAGE_FILE := coverage.out
DOCKER_TAG := $(BINARY_NAME):$(VERSION)

# ==============================================================================
# ОСНОВНЫЕ ЦЕЛИ
# ==============================================================================

all: build

# ------------------------------------------------------------------------------
# Сборка
# ------------------------------------------------------------------------------

build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME):
	@echo "Building $(BINARY_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $@ -ldflags "-X main.version=$(VERSION)" ./cmd/

# ------------------------------------------------------------------------------
# Очистка
# ------------------------------------------------------------------------------

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR) $(COVERAGE_FILE)

# ------------------------------------------------------------------------------
# Тестирование
# ------------------------------------------------------------------------------

test:
	@echo "Running tests..."
	go test ./... -v

test-coverage: $(COVERAGE_FILE)
	@echo "Opening coverage report..."
	go tool cover -html=$(COVERAGE_FILE)

$(COVERAGE_FILE):
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=$@

# ------------------------------------------------------------------------------
# Линтинг
# ------------------------------------------------------------------------------

lint:
	@echo "Running linter..."
	golangci-lint run

# ------------------------------------------------------------------------------
# Запуск
# ------------------------------------------------------------------------------

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# ------------------------------------------------------------------------------
# Docker
# ------------------------------------------------------------------------------

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_TAG) .

# ------------------------------------------------------------------------------
# Зависимости
# ------------------------------------------------------------------------------

install-deps:
	@echo "Installing dependencies..."
	go mod download
	@echo "Installing golangci-lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# ------------------------------------------------------------------------------
# Разработка
# ------------------------------------------------------------------------------

dev:
	@echo "Starting development mode..."
	go run ./cmd/

# ------------------------------------------------------------------------------
# Бенчмарки
# ------------------------------------------------------------------------------

bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

bench-logging:
	@echo "Running logging benchmarks..."
	go test -bench=BenchmarkLogging -benchmem ./internal/infrastructure/logger/

# ------------------------------------------------------------------------------
# Релиз
# ------------------------------------------------------------------------------

release: clean test lint build
	@echo "Release $(VERSION) ready in $(BUILD_DIR)/"

# ------------------------------------------------------------------------------
# Документация
# ------------------------------------------------------------------------------

help:
	@echo "Available commands:"
	@echo "  build           - Build the binary"
	@echo "  clean           - Clean build artifacts"
	@echo "  test            - Run tests"
	@echo "  test-coverage   - Run tests with coverage report"
	@echo "  lint            - Run linter"
	@echo "  run             - Build and run the binary"
	@echo "  docker-build    - Build Docker image"
	@echo "  install-deps    - Install dependencies"
	@echo "  dev             - Run in development mode"
	@echo "  bench           - Run benchmarks"
	@echo "  bench-logging   - Run logging benchmarks"
	@echo "  release         - Clean, test, lint and build"
	@echo "  help            - Show this help message"
	@echo "  log-examples    - Show logger usage examples"

log-examples:
	@echo "Logger usage examples:"
	@echo "  ./$(BINARY_NAME) --log-level debug calculate ..."
	@echo "  ./$(BINARY_NAME) --log-format json calculate ..."
	@echo "  ./$(BINARY_NAME) --verbose calculate ..."
```