RED = \033[0;31m
GREEN = \033[0;32m
NC = \033[0m


all: deps build

build: build_client build_server

deps:
	@echo "Start cloning dependencies...$(NC)"
	@go get ./...
	@echo "$(RED)Exiting dependency cloning process.$(NC)"

build_client:
	@echo "Start building Client...$(NC)"
	@go build -o client src/app/client/main.go
	@echo "$(RED)Exiting client's building process.$(NC)"

build_server:
	@echo "Start building Server...$(NC)"
	@go build -o server src/app/server/main.go
	@echo "$(RED)Exiting server's building process.$(NC)"

clean:
	@rm client server
	@echo "$(RED)Client and server binary files removed.$(NC)"

.PHONY: all build clean build_client build_server deps
