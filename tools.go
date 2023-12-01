//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)

/*
This file is a container for dependencies that our CLI tools use.
This makes it easy to just run 'go mod tidy' and then use make commands
to run all of the tools needed to develop in this repo.

I would like to decouple all of the dependencies this creates from the actual project go.mod file though.
I will need to explore what it takes to have multple modules, so that I can have a 'tools' module
and a module that actually has the important program. Most dependencies in go.mod are actually just dependencies
for these CLI tools, not the actual program. And all of these CLI tools just autogenerate code that I could
technically write myself fairly easily.
*/
