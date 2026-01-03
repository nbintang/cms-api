package auth


import "github.com/gofiber/fiber/v2"



func RegisterAuthRoutes(app *fiber.App, h AuthHandler){
	api := app.Group("/api")
	auth := api.Group("/auth") 
	auth.Post("/register", h.Register)
}
