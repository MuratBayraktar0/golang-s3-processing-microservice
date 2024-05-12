package app

import (
	"log"
	"net/http"
	"product_service/internal/domain/service"
	"product_service/internal/handler"
	"product_service/internal/infrastructure/database"
	"product_service/internal/infrastructure/server"
	"product_service/internal/repository"
)

func StartProductService(dbConnection, dbName, serverPort string) {

	// Connect to MongoDB
	mongo, err := database.NewMongoDB(dbConnection, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Defer disconnect from MongoDB
	defer mongo.Disconnect()

	// Initialize Product Repository
	repo := repository.NewProductRepository(mongo, "products")

	// Initialize Product Service
	service := service.NewProductService(repo)

	// Initialize Product Handler
	productHandler := handler.NewProductHandler(service)

	// Initialize HTTP Server
	server := server.NewHTTPServer(serverPort)
	server.RegisterHandler(http.MethodGet, "/products", productHandler.GetProducts)
	server.RegisterHandler(http.MethodGet, "/products/:id", productHandler.GetProduct)
	server.Start()
}
