DB_URL=postgresql://postgres:7TQPNbG2SvQY5zxbJRzq@simple-bank.ckvklkqrqq0g.ap-south-1.rds.amazonaws.com:5432/simple_bank

network:
	docker network create bank-network

postgres:
	docker run --name postgres14 --network go-aws-micro-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=7TQPNbG2SvQY5zxbJRzq -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store



.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema