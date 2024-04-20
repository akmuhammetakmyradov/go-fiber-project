# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=testmain
BIN_DIR=bin
MAIN_FILE=cmd/main.go

# db
POSTGRESQL_URL := "postgresql://postgres:12345@localhost:5432/test?sslmode=disable"
MIGRATE_VERSION := ${v}

# Flags
LDFLAGS=-ldflags "-s -w"

build:
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BIN_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) ./...

migrate:
	migrate create -ext sql -dir database/migrations -seq init_test

migrateup:
	migrate -database ${POSTGRESQL_URL} -path database/migrations -verbose up ${MIGRATE_VERSION}

migratedown:
	migrate -database ${POSTGRESQL_URL} -path database/migrations -verbose down ${MIGRATE_VERSION}

migratefix:
	migrate -database ${POSTGRESQL_URL} -path database/migrations -verbose force ${MIGRATE_VERSION}

composebuild:
	sudo docker-compose build

composeup:
	sudo docker-compose up -d

composedown:
	sudo docker-compose down

.PHONY: migrate migrateup migratedown migratefix composebuild composeup composedown build clean deps