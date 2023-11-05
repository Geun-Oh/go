package main

import (
	"github.com/Geun-Oh/fiber-react/server/database"
	"github.com/Geun-Oh/fiber-react/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

type Resp struct {
	Message string `json:"message"`
}

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(".env loading failed")
	}
	database.ConnectDB()
}

func main() {

	sqlDB, err := database.DBConn.DB()
	defer sqlDB.Close()

	if err != nil {
		panic("Err in SQL connection")
	}

	app := fiber.New()

	app.Use(logger.New())

	router.SetUpRoutes(app)

	app.Listen(":8000")
}
