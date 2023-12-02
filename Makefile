include .env.develop

#---build and run---

build-server:
	go build -o ./build

run:
	go run server.go

infra-dev:
	docker compose up --detach

#---code generation---

gen-api:
	cd api; go run github.com/99designs/gqlgen@v0.17.40 generate

gen-db:
	cd db; go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.24.0 generate

#---db migrations---


#Ex: "make create-migration NAME=<name of migration>"
create-migration: 
	go run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2 create -ext sql -dir db/migrations -seq $(NAME)

up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2 -database ${POSTGRESQL_URL} -path db/migrations up

down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2 -database ${POSTGRESQL_URL} -path db/migrations down

#---install dependencies---
#go mod tidy

