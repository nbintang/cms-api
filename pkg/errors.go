package pkg

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ErrNotFound = errors.New("Not Found")

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	msg := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		statusCode = e.Code
		msg = e.Message
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"error":     msg,
		"status":    statusCode,
		"timestamp": time.Now().Unix(),
	})
}
