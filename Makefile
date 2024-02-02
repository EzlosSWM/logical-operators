# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	make --jobs=2 css air

css: 
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

air:
	~/go/bin/air

.PHONY: all build run test clean
