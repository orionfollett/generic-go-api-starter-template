#---build and run---

build-server:
	go build -o ./build

run:
	go run server.go

infra-dev:
	docker compose up --detach

#---code generation---

gen-api:
	cd api; go run github.com/99designs/gqlgen generate

gen-db:
	cd db; go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

#---db migrations---

#This command takes the form: "make create-migration NAME=<enter human readable migration name here>"
#It will create a blank up and down migration file for you in the correct directory
create-migration:
	go run github.com/golang-migrate/migrate/v4/cmd/migrate create -ext sql -dir db/migrations -seq $(NAME)

#---install dependencies---
#go mod tidy

