MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
DB_URL := postgres://root:secret@localhost:5432/gnb?sslmode=disable

setup: postgres createdb migrateup

test:
	go test -cover ./...

server:
	go run main.go

# DB
postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root gnb

dropdb:
	docker exec -it postgres dropdb gnb

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1
	
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
	
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1


# SQLC
sqlc-init:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc init

sqlc-compile:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc compile

sqlc-generate:
	docker run --rm -v $(MAKEFILE_DIR):/src -w /src sqlc/sqlc generate


# MOCK
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ronymmoura/goliath-national-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 
	sqlc-generate sqlc-init sqlc-compile 
	test server mock