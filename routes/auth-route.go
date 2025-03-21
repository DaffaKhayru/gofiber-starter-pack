package routes

import (
	"github.com/DaffaKhayru/gofiber-starter-pack/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/signup", controllers.Signup)
	api.Post("/login", controllers.Login)
	api.Post("/signout", controllers.Signout)
}