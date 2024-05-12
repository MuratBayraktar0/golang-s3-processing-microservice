package main

import (
	"fmt"
	"os"
	"product_service/internal/app"

	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("ENV")
	switch env {
	case "dev":
		fmt.Println("Running in development mode")
	case "prod":
		fmt.Println("Running in production mode")
	default:
		env = "dev"
		fmt.Println("Running in development mode")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		fmt.Printf("Failed to load .env.%s file\n", env)
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	app.StartProductService(dbConnection, dbName, serverPort)
}
