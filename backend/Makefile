# Variables
NAME ?= flowforge_be
BIN_DIR ?= ./bin
SRC_DIR ?= ./

# Clean up the binary directory
clean:
	@rm -rf $(BIN_DIR)/*
	@echo "Cleaned up binaries in $(BIN_DIR)"

# Build the project
build: clean
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(NAME) $(SRC_DIR)
	@echo "Build completed: $(BIN_DIR)/$(NAME)"

# Run the built binary
run: build
	@chmod +x $(BIN_DIR)/$(NAME)
	@$(BIN_DIR)/$(NAME)

# Default target
all: build

.PHONY: clean build run all
