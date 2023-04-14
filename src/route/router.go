package route

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/meetup/graph"
	"github.com/meetup/graph/service"
	"github.com/meetup/internal"
	"log"
	"net/http"
	"os"
)

func GetRouter() {
	port := os.Getenv("PORT")
	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv: service.NewUserService(),
	}}))

	//query {
	//	user(name: "nagamatsu") {
	//		id
	//		username
	//		email
	//	}
	//}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
