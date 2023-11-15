
## Tools

SQLC: https://docs.sqlc.dev/en/latest/index.html
This tool is controlled by the sqlc.yaml file
It autogenerates DB Layer code to execute queries that you write in SQL, and map them to 
typesafe Go structs. Generally each table gets mapped to a struct, and each query to a function.
SQLC works alongside a DB driver to achieve this. 

PGX: https://github.com/jackc/pgx
This is a db driver for postgres so that the go standard library can use postgres.
This is a low level library to establish a db connection and execute queries.

Other Tools:
API framework? Or just use standard library

gqlgen?

Database migration tool: leaning towards golang-migrate: https://github.com/golang-migrate/migrate

Logging, std library

Infrastructure: Terraform, aws provider most likely

CICD: github actions since this is popular

Testing: not sure yet




## Build Application
go build -o output/build.exe

## Run Main Application
go run main.go

## Docker Compose
'''docker compose up''' To create the db in a locally running container: connection string is: "postgresql://postgres:example@localhost:5432"



## Goal
Goal for this repo is to be a gold standard starter repo to create a go backend application.

While there are many powerful tools for doing backend development, the goal here is to abstract that away
similar to how a Linux distro packages up a bunch of composable tools into a cohesive software package, 
this repo should package a bunch of tools for backend development into a cohesive platform that
go backends can be developed on.

This will be an opinionated repo that provides standard tools for database migrations,
infrastructure, containerization, etc.

Goal would be:
- all infrastructure needed to deploy, codified and declarative?
- relation database management, migrations, schema, queries
- unit tests, integration tests, end to end tests
- logging and error handling
- DB access layer
- API layer
- standard make commands defined for local testing, env setup, etc

Codebase should be decoupled, so that individual parts may be swapped out
Dependencies should be minimal so that code doesn't rot as quickly

- documentation on tradeoffs, project setup, etc
- configuration should be managed by config files or env files, and should not involve modifying code

- generally should follow 12 factor app practices



