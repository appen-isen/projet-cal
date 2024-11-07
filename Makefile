# Variables
BACKEND_DIR=backend
BACKEND_BINARY=backend

# Règles
all: build-backend start-backend

build-backend:
	@echo "Building backend..."
	cd $(BACKEND_DIR) && go build -o $(BACKEND_BINARY)

start-backend:
	@echo "Starting backend..."
	cd $(BACKEND_DIR) && ./$(BACKEND_BINARY)

.PHONY: all build-backend start-backend