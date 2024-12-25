package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

func ErrorIntercept() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return response.Fail(c, err)
		}
		return nil
	}
}
