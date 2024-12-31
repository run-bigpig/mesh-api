package relay

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/common"
	"github.com/run-bigpig/mesh-api/internal/service"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
)

func ImageHandler(ctx *fiber.Ctx) error {
	var req adapter.ImageRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}
	if req.Model == "" {
		return errors.New("model is required")
	}
	line, err := common.RandomLineByWeight(ctx.Context(), req.Model)
	if err != nil {
		return err
	}
	relay := service.AdapterProvider(ctx, line.Adapter)
	adapterParams := &adapter.Params{
		Api:  line.Host,
		Sk:   line.Auth,
		Mode: adapter.ImageGeneration,
	}
	//转换图片请求
	params, err := relay.ImageConversion(adapterParams, &req)
	if err != nil {
		return err
	}
	//转发
	resp, err := relay.Relay(params)
	if err != nil {
		return err
	}
	//处理图片响应
	return relay.ImageResponse(resp)
}
