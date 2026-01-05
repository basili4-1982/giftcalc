.PHONY: build clean test lint run docker-build

BINARY_NAME=giftcalc
VERSION=1.0.0

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) -ldflags "-X main.version=$(VERSION)" ./cmd/

clean:
	@echo "Cleaning..."
	rm -rf bin/ coverage.out

test:
	@echo "Running tests..."
	go test ./... -v

test-coverage:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	@echo "Running linter..."
	golangci-lint run

run: build
	@echo "Running $(BINARY_NAME)..."
	./bin/$(BINARY_NAME)

docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):$(VERSION) .

install-deps:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

dev:
	@echo "Starting development mode..."
	go run main.go

bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

release: clean test lint build
	@echo "Release $(VERSION) ready in bin/"

help:
	@echo "Available commands:"
	@echo "  build         - Build the binary"
	@echo "  clean         - Clean build artifacts"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo ""

log-examples:
	@echo "Примеры использования логгера:"
	@echo "  ./giftcalc --log-level debug calculate ..."
	@echo "  ./giftcalc --log-format json calculate ..."
	@echo "  ./giftcalc --verbose calculate ..."

bench-logging:
	@echo "Бенчмарк логирования..."
	go test -bench=BenchmarkLogging -benchmem ./internal/infrastructure/logger/