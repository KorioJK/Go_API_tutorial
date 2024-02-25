package main

import (
	"log"
	"os"

	"example.com/m/controllers"
	"example.com/m/graph"
	"example.com/m/initializers"
	"example.com/m/middleware"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create the GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Create the Gin router
	r := gin.Default()

	// Mount GraphQL endpoints using Gin
	r.POST("/query", middleware.RequireAuth, gin.WrapH(srv))
	r.GET("/", middleware.RequireAuth, gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Define REST API routes using Gin
	r.POST("/post", controllers.PostsCreate)
	r.GET("/post", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)
	r.PUT("/post/:id", controllers.PostsUpdate)
	r.DELETE("/post/:id", controllers.PostsDelete)

	// Users routes
	r.POST("/users", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Start the server
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}
