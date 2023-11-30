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
#migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2
create-migration:
	migrate create -ext sql -dir db/migrations -seq create_users_table

#---install dependencies---
#install migrate tool
#go mod tidy

