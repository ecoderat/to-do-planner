package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"to-do-planner/internal/config"
	"to-do-planner/internal/db"
	"to-do-planner/internal/router"
)

func main() {
	db.InitDB()

	app := fiber.New()

	cfg := config.New(db.DB)

	router.RegisterRoutes(app, cfg)

	log.Fatal(app.Listen(":3000"))
}
