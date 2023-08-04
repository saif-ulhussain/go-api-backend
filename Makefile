SHELL := /bin/bash

export $(shell sed 's/=.*//' .env)

create-database:
	docker run --name postgresql -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -e POSTGRES_DB=go-api-backend-db -d postgres

create-database-test:
	docker run --name postgresql-test -e POSTGRES_PASSWORD=mysecretpassword -p 5433:5432 -e POSTGRES_DB=go-api-backend-db-test -d postgres

drop-database:
	docker exec -it postgresql dropdb -U postgres go-api-backend-db

migrate-up:
	migrate -path migrations -database "postgresql://postgres:mysecretpassword@localhost:5432/go-api-backend-db?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:mysecretpassword@localhost:5432/go-api-backend-db?sslmode=disable" -verbose down
#	@go run -mod=readonly cmd/api/main.go migrate -path internal/db/migrations -database "postgresql://$$DB_USER:DbS3rVe@1@localhost:5432/postgres?sslmode=disable" -verbose down

