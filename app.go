package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mariumfaheem/PE-GO-PROJECT/database"
	"github.com/mariumfaheem/PE-GO-PROJECT/handlers"
	"log"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {

	flag.Parse()

	database.Connect()

	//Fiber command-line

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Use(recover2.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	app.Static("/", "./static/public")

	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000

}
