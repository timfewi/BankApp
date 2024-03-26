postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=konta -e POSTGRES_PASSWORD=1683 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=konta --owner=konta bankApp

dropdb:
	docker exec -it postgres12 dropdb bankApp

migrateup:
	migrate -path db/migration -database "postgresql://konta:1683@localhost:5433/bankApp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://konta:1683@localhost:5433/bankApp?sslmode=disable" -verbose down

server:
	go run main.go

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server


