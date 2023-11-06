package main

import (
	"log"
	"os"

	"github.com/Geun-Oh/fiber-react/react-in-go/controllers"
	"github.com/Geun-Oh/fiber-react/react-in-go/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Reading .env file failed")
	}
	initializers.ConnectDB()
}

func main() {
	sqlDB, err := initializers.DBConn.DB()
	defer sqlDB.Close()

	if err != nil {
		panic("Err in SQL connection")
	}

	// Load templates
	engine := html.New("./views", ".html")

	// Create app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	// Configure app
	app.Static("/", "./public")

	// Routing
	app.Get("/api/tasks", controllers.FetchTasks)
	app.Post("/api/tasks", controllers.CreateTask)
	app.Get("/api/tasks/:id", controllers.FetchTask)
	app.Delete("/api/tasks/:id", controllers.DeleteTask)

	app.Get("/", controllers.Home)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
