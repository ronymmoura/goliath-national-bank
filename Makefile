makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

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
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/gnb?sslmode=disable" -verbose up
	
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/gnb?sslmode=disable" -verbose down


# SQLC
sqlc-init:
	docker run --rm -v $(makeFileDir):/src -w /src sqlc/sqlc init

sqlc-compile:
	docker run --rm -v $(makeFileDir):/src -w /src sqlc/sqlc compile

sqlc-generate:
	docker run --rm -v $(makeFileDir):/src -w /src sqlc/sqlc generate


.PHONY: postgres createdb dropdb migrateup migratedown 
	sqlc-generate sqlc-init sqlc-compile 
	test server