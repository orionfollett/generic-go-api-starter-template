
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

Database migration tool: golang-migrate: https://github.com/golang-migrate/migrate

Infrastructure: Terraform, aws provider most likely, logging, metrics, scaling, kubernetes, networking etc

CICD: github actions since this is popular

Testing: not sure yet

Git


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


Proposed Project Layout:

API, with clean separation to DB layer
Would support swapping out DB layer, or doing a "scientist"
approach where you have multiple DB layers

Infrastructure has clean decoupling from app layer

Where is business logic?

- API is thin layer over DB. API is like the UI, but should be as dumb as possible
it only validates user input so that DB doesn't throw errors. It also handles DB errors
more gracefully and surfaces them in a readable way to the user

- API has all of the business logic, just treating the DB layer as an api that can 
access the data store. 

- mix of both

Pros and Cons
DB being business layer:
Pros:
- DB business rules are more likely to be enforced by all clients to the db
- DB business rules are potentially more performant because any checks for data can 
be done at the db level without a round trip and without data copying

Cons:
- logic can be harder to write in SQL, slower development
- couples our business to the database, making it hard to swap databases - doesnt seem like a big con, 
seems more likely we could switch our API layer, but I guess it's possible we change our database

But if we swap our database in any significant way (like switch from relational to not, then its probably going to be a lot of work no 
matter how you slice it)




----------------------------------------

Good decoupling would say treat your data store as one component
and the API as another, and they should be easily swapped out with something else


Another middle ground approach: put logic where it is most efficient to compute:

Some logic lives best in the DB, other logic is best in the API.

Anything that requiers data already stored or calculated in the DB, should live in the DB, dont make round trips for no reason.

Anything that could be validated before it hits the DB, without querying more data (for example, this field requires a string, but an int was given)
can be handled by the API



Good decoupling from the DB can still be achieved if business logic lives in both, because it is too restrictive to say
that business logic can ONLY exist in one or another. 

In order to support good decoupling...

API functions call DB functions, DB functions never call anything that lives in the API module, it's a one way communication
DB functions in GO are dumb and just map to SQL, this middle layer should largely be handled by sqlc

With that middle layer, it should be easy to replace the API layer easily, and still access all the DB layer stuff
Or vice versa, swapping out the DB layer would mean just supplying all of the DB functions currently defined




