# DB_URL=postgresql://root:secret@localhost:5432/service?sslmode=disable

network:
	docker network create service-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

creatdb:
	docker exec -it postgres createdb --username=root --owner=root service

migrateup:
	migrate --path db/migration -database ${DB_URL} -verbose up 

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run main.go

.PHONY: network postgres createdb migrateup sqlc test run