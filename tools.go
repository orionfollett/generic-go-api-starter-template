//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
