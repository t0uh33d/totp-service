include .env
export DB_URL
export EC2_DB_URL

DB_MIGRATION= ./db/migrations
# Go command
GOCMD=go
# Build, clean, test, and get commands
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
# Binary names, adjusted to place the output in ./bin/
BINARY_NAME=totp_service
BINARY_PATH=./bin/$(BINARY_NAME)
BINARY_UNIX=$(BINARY_PATH)_unix
BACKEND_SERVICE_ENTRYPOINT=./cmd/start/main.go

# Default to run tests and then build
all: test build

# Compile the binary to bin/ directory
build:
	$(GOBUILD) -o $(BINARY_PATH) -v $(BACKEND_SERVICE_ENTRYPOINT)

# Run tests across the project
test:
	$(GOTEST) -v ./...

# Clean up the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_PATH)
	rm -f $(BINARY_UNIX)

# Build and run the project
run: build
	$(BINARY_PATH)

# Handle dependencies using Go modules
deps:
	$(GOMOD) tidy
	$(GOMOD) verify

# Cross compilation for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BINARY_PATH) $(BACKEND_SERVICE_ENTRYPOINT)

up:
	@goose -dir=$(DB_MIGRATION) postgres $$DB_URL up

down:
	@goose -dir=$(DB_MIGRATION) postgres $$DB_URL down

ec2-db-up:
	goose -dir=$(DB_MIGRATION) postgres $$EC2_DB_URL up

ec2-db-down:
	@goose -dir=$(DB_MIGRATION) postgres $$EC2_DB_URL down

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build-htmx:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
