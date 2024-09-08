postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres simple-bank

createmigrate:
	migrate create -ext sql -dir db/migration -seq <<MIGRATION NAME>>

migrateup1:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable" -verbose down 1

migrateup:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable" -verbose down -all

testdbmigrateup:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank-test-db?sslmode=disable" -verbose up

testdbmigratedown:
	migrate -path db/migration -database "postgres://postgres:1@localhost:5432/simple-bank-test-db?sslmode=disable" -verbose down -all

dropdb:
	docker exec -it postgres12 dropdb simple-bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

# dev
devserver:
	air

# prod
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock
