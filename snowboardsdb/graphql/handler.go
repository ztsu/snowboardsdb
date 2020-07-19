package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"net/http"
)

func Handler(services *Stores) http.Handler {
	return ContextWithDataLoaders(
		services,
		handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: NewRootResolver(services)})),
	)
}
