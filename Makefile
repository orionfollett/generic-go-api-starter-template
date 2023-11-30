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
	cd db; sqlc generate


