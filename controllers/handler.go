package controllers

import (
	"assuresfot/harmons/messagebrokerredis/services"

	"github.com/gofiber/fiber/v2"
)

// Handler Init fiber and set service methods and the port
func Handler() {
	app := fiber.New()

	// Methos to controll
	app.Use(services.RemoveQueryAuthService)
	app.Put("*", services.PublishIntoMessageBroker)
	app.Get("*", services.PublishIntoMessageBroker)
	app.Post("*", services.PublishIntoMessageBroker)
	app.Delete("*", services.PublishIntoMessageBroker)

	app.Listen(":4000")
}
