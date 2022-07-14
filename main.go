package main

import (
	"graphql-golang/config"
	"graphql-golang/graph"
	"graphql-golang/graph/model"
	"graphql-golang/service"
	"graphql-golang/utils"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file

	server := config.NewServer()

	c := graph.Config{Resolvers: &graph.Resolver{
		StudentService: service.NewStudentService(server),
		AuthService:    service.NewAuthService(server),
		FileService:    service.NewFileService(server),
		ChatMessages:   []*model.Message{},
		ChatObservers:  map[string]chan []*model.Message{},
	}}
	c.Directives.JwtAuth = server.JwtAuth
	c.Directives.HasRole = server.HasRole

	h := handler.New(graph.NewExecutableSchema(c))
	h.AddTransport(transport.POST{})

	// disabling introspection on production
	if os.Getenv("ENV") != "production" {
		h.Use(extension.Introspection{})
	}

	// config websocket
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	h.AddTransport(transport.MultipartForm{
		MaxMemory:     50000,
		MaxUploadSize: 50000,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// mwServerHeader display server copyright
func mwServerHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Server", "GM API https://guillaume-morin.fr")
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(mwServerHeader())
	r.Use(config.AuthMiddleware())
	r.Use(gzip.Gzip(gzip.BestCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://" + os.Getenv("API_DOMAIN") + os.Getenv("API_PORT")},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "jwtToken"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.StaticFile("/favicon.ico", "favicon.ico")
	r.StaticFS("/public", http.Dir(utils.Dir()+"/public"))

	r.Any("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
