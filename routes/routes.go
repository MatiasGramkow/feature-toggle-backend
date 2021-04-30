package routes

import (
	"github.com/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setup for endpoints
func Setup(app *fiber.App) {
	// Auth Endpoints
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)

	// Forgot Endpoints
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)

	// Feature Endpoints
	app.Post("/api/createFeature", controllers.CreateFeature)
	app.Get("/api/features", controllers.Features)
	app.Post("/api/updateFeature/:id/:is_active/:deactivation_reason", controllers.UpdateFeature)
}
