package manage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/config"
	"github.com/run-bigpig/mesh-api/internal/gateway/response"
)

func ListAdapters(ctx *fiber.Ctx) error {
	return response.Success(ctx, config.Get().Adapter)
}
