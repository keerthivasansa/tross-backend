package main

import (
	"kv/tross/database"
	"kv/tross/routers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	routers.RegisterUserRoutes(app)
	routers.RegisterGoalsRouter(app)

	log.Fatal(app.Listen(":5000"))
}
