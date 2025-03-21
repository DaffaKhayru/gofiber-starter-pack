package main

import (
	"log"

	"github.com/DaffaKhayru/gofiber-starter-pack/routes"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func main() {
	app := fiber.New()

	routes.AuthRoute(app)

	App = app

	log.Fatal(app.Listen(":3000"))
}