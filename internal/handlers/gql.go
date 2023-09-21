package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/Earl-Power/gqlgengin/internal/gql"
	"github.com/Earl-Power/gqlgengin/internal/gql/resolvers"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}
	h := handler.GraphQL(gql.NewExecutableSchema(c))
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
