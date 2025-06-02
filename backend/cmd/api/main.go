package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Cueclub.live",
	})

	app.Use(cors.New())
	app.Static("/", "../frontend/dist")
	app.Get("/api/v1/mix", getMix)

	log.Fatal(app.Listen(":8080"))
}
