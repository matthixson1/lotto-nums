# Makefile for Euromillions generator

# Go compiler
GO = go

# Binary name
BINARY = euromillions

# Source files
SOURCES = main.go Euromillions.go

# Build flags
BUILD_FLAGS = -v

# Default target
all: build

# Build the application
build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY) $(SOURCES)

# Run the application
run: build
	./$(BINARY)

# Clean build artifacts
clean:
	rm -f $(BINARY)
	
# Format the code
fmt:
	$(GO) fmt ./...

.PHONY: all build run clean test fmt
