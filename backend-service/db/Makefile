createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:15-alpine

migrateup:
	migrate -path migrations -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: createdb dropdb postgres migrateup migratedown