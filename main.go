// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"3d-print-inventory/database"
	"3d-print-inventory/handler"
	"3d-print-inventory/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})
	// database.Connect()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Access-Control-Allow-Origin, Content-Type, Accept, Accept-Language, Origin, User-Agent, Authorization, rejectUnauthorized",
			AllowMethods: "GET, PUT, OPTIONS, POST, DELETE",
		}))
	database.ConnectToDatabase()
	database.ConnectToRedis()
	router.SetupRoutes(app)
	// Static file server
	app.Static("/download", "./files", fiber.Static{
		ModifyResponse: handler.GetFile,
	})
	// => http://localhost:3000/hello.txt
	// => http://localhost:3000/gopher.gif

	// Start server

	log.Fatal(app.Listen(":8002"))
}
