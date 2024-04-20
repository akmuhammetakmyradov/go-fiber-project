package handlers

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register(router fiber.Router)
}
