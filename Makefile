#---build and run---
build-server:
	go build -o ./build

run:
	go run server.go

infra-dev:
	docker compose up --detach

#---code generation---
api-gen:
	cd api; go run github.com/99designs/gqlgen generate

db-gen:
	cd db; sqlc generate

all-gen:
	cd db; sqlc generate; 
	cd api; go run github.com/99designs/gqlgen generate



