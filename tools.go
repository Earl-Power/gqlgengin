//go:build tools
// +build tools

// go: go run github.com/99designs/gqlgen generate
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
