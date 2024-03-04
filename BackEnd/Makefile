.PHONY: build run clean

# Name of the output binary
BINARY_NAME := KanbanServer

# Directories
SRC_DIR := .
BIN_DIR := bin

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go
	@echo "Build complete."

run:
	@echo "Running $(BINARY_NAME)..."
	./$(BIN_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	@echo "Cleanup complete."
