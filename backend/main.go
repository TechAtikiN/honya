package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/controllers"
	"github.com/techatikin/backend/repositories"
	"github.com/techatikin/backend/routers"
	"github.com/techatikin/backend/services"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Set up Fiber app
	app := fiber.New()

	// Setup cors middleware
	app.Use(cors.New())
	api := app.Group("/api")

	// Set up Repositories
	repo := repositories.BookRepository(db)

	// Set up Services
	svc := services.BookService(repo)

	// Set up Controllers
	ctrl := controllers.BookController(svc)

	// Set up Routers
	routers.SetupBooksRouter(api, ctrl)

	// Start the server
	log.Println("Server starting on port 8080...")
	log.Fatal(app.Listen(":8080"))
}
