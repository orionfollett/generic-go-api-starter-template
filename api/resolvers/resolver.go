package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	sqlc "orion/generic-api-starter/db/generated"
)

type Resolver struct {
	QUERIES *sqlc.Queries
}
