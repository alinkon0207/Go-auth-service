package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/quible-backend/auth-service/config"
	"gitlab.com/quible-backend/auth-service/pkg/repository/user"

    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

    "gitlab.com/quible-backend/auth-service/docs"
)

func main() {
	// separate the code from the 'main' function.
	// all code that available in main function were not testable
	Server()
}

func Server() {
	// prepare gin
	gin.SetMode(gin.ReleaseMode)

	// gin setup
	r := gin.Default()
	r.Use(cors.Default())

	// prepare postgresql database
	dbPool, _, err := config.NewDBPool(config.DatabaseConfig{
		Username: "scraper",
		Password: "!QAZxsw2",
		Hostname: "localhost",
		Port:     "5432",
		DBName:   "scraper",
	})

	// log for error if error occur while connecting to the database
	if err != nil {
		log.Fatalf("unexpected error while tried to connect to database: %v\n", err)
	}

	defer dbPool.Close()

	// setup api
	database := user.NewRepository(dbPool)
	service := user.NewService(database)
	controller := user.NewController(service)

	user.Routes(r, controller)

    // swagger setup and routing
	docs.SwaggerInfo.Title = "Quible auth-service"
	docs.SwaggerInfo.Description = "Authentication and authorization service of Quible.io"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = "www.quible.io"
	docs.SwaggerInfo.BasePath = "api/auth"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
    r.GET("/api/auth/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run the server
	log.Fatalf("%v", r.Run(":8001"))
}