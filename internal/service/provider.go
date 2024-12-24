package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/deepseek"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/siliconflow"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/tongyi"
)

func AdapterProvider(ctx *fiber.Ctx, name string) adapter.Adapter {
	switch name {
	case adapter.Deepseek:
		return deepseek.New(ctx)
	case adapter.Tongyi:
		return tongyi.New(ctx)
	case adapter.SiliconFlow:
		return siliconflow.New(ctx)
	case adapter.MoonShot:
		return siliconflow.New(ctx)
	}
	return nil
}
