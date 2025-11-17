.PHONY: help run build clean dev

help:
	@echo "Available commands:"
	@echo "  make run      - Run the application"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make dev      - Run with auto-reload (requires Air)"

run:
	go run cmd/web/main.go

build:
	go build -o bin/portfolio cmd/web/main.go

clean:
	rm -rf bin/

dev:
	air
