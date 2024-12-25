package relay

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/service"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
)

func TextRelayHandler(ctx *fiber.Ctx) error {
	var req adapter.TextRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	if req.Model == "" {
		ctx.Status(fiber.StatusBadRequest).Send([]byte("model is required"))
		return nil
	}
	relay := service.AdapterProvider(ctx, "siliconflow")
	adapterParams := &adapter.Params{
		Api:    "https://api.moonshot.cn",
		Sk:     "sk-S9fP5nRJVrVuO8jWo831JORCpYUTYzZkC8jLDp8YlAX2FUyG",
		Mode:   getMode(ctx),
		Stream: req.Stream,
	}
	//转换请求
	params, err := relay.TextConversion(adapterParams, &req)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	//转发
	resp, err := relay.Relay(params)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	//处理响应
	return relay.TextResponse(adapterParams, resp)
}

// 判断请求的mode
func getMode(ctx *fiber.Ctx) int8 {
	switch ctx.Path() {
	case "/v1/chat/completions":
		return adapter.ChatCompletions
	case "/v1/completions":
		return adapter.Completions
	default:
		return 0
	}
}
