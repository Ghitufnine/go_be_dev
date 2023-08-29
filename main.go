package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"go_be_dev/auth"
	"go_be_dev/database"
	"go_be_dev/transaction"
	"log"
)

func setUpRoutes(app *fiber.App) {
	api := app.Group("/api")

	Auth := api.Group("/auth")
	Auth.Post("/login", auth.PostDoLogin)

	// Apply JWT middleware to protected routes
	name := "projects/77840544412/secrets/fleetify_api_jwt_key/versions/latest"
	secretKey := auth.AccessSecretVersion(name)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secretKey),
	}))

	Transaction := api.Group("/transaction")
	Transaction.Post("/clock_in", transaction.PostClockIn)
	Transaction.Post("/clock_out", transaction.PostClockOut)
	// Add your protected endpoints here
}

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{})

	// Add recover middleware
	app.Use(recover.New())

	// Configuring CORS middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
