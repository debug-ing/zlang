APP_NAME := zlang
SRC := cmd/zlang.go
OUTPUT := $(APP_NAME)

build: $(SRC)
	@echo "Building the project..."
	go build -o $(OUTPUT) $(SRC)

clean:
	@echo "Cleaning up..."
	rm $(OUTPUT)

help:
	@echo "Makefile targets:"
	@echo "  build   - Build the Go project"
	@echo "  run     - Build and run the project"
	@echo "  clean   - Remove build file"
	@echo "  help    - Show this help message"