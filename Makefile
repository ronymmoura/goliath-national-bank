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

.PHONY: postgres createdb dropdb migrateup migratedown