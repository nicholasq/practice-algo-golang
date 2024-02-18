# Makefile
.PHONY: build clean run test

GO_FILES=$(wildcard *.go)
OUTPUT_NAME=myapp

# Build rule
build: $(GO_FILES)
	go build -o $(OUTPUT_NAME) $(GO_FILES)

# Clean rule
clean:
	go clean

# Run rule
run:
	go run $(GO_FILES)

# Test rule
test:
	go test ./...