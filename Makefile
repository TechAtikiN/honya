.DEFAULT_GOAL := help
.PHONY: help install install-fe install-be lint lint-fe lint-be test test-fe test-be docker-up docker-down docker-clean seed

# ============= Variables =============

DOCKER_COMPOSE_FILE := docker-compose.yml
FRONTEND_DIR := frontend
BACKEND_DIR := backend
SEED_SCRIPT := scripts/seed/main.go

# =========== Install =============

install-fe: ## Install frontend dependencies
	cd ${FRONTEND_DIR} && pnpm install 

install-be: ## Install backend dependencies
	cd ${BACKEND_DIR} && go mod download

install: ## Install all dependencies
	@echo "Installing frontend and backend dependencies..."
	@$(MAKE) install-fe
	@$(MAKE) install-be
	@echo "Dependencies installed successfully."

# ============= Lint =============

lint-fe: ## Lint frontend code
	cd ${FRONTEND_DIR} && pnpm lint

lint-be: ## Lint backend code
	cd ${BACKEND_DIR} && golangci-lint run

lint: ## Lint all code
	@echo "Linting frontend and backend code..."
	@$(MAKE) lint-fe
	@$(MAKE) lint-be
	@echo "Linting completed successfully."

# ============= Test =============

test-fe: ## Test frontend code
	cd ${FRONTEND_DIR} && pnpm test

test-be: ## Test backend code
	cd ${BACKEND_DIR} && go test ./...

test: ## Test all code 
	@echo "Running frontend and backend tests..."
	@$(MAKE) test-fe
	@$(MAKE) test-be
	@echo "All tests completed successfully."

# ============= Docker =============

docker-up: ## Start Docker containers
	docker compose -f ${DOCKER_COMPOSE_FILE} up -d 

docker-down: ## Stop Docker containers
	docker compose -f ${DOCKER_COMPOSE_FILE} stop

docker-clean : ## Stop and remove Docker containers, networks, images, and volumes
	docker compose -f ${DOCKER_COMPOSE_FILE} down --rmi all -v --remove-orphans

# ============= Seed =============

seed: ## Seed the database with initial data
	cd ${BACKEND_DIR} && go run ${SEED_SCRIPT}

# ============= Help  =============

help: ## Show this help message.
	@echo "\033[36mAvailable commands:\033[0m"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' | sort
	