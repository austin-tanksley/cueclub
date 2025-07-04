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

	app.Static("/", "./frontend/dist")
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "https://cueclub.onrender.com",
		// AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/api/v1/mix", getMix)

	log.Fatal(app.Listen(":8080"))
}
