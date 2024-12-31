package relay

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/common"
	"github.com/run-bigpig/mesh-api/internal/service"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
)

func TextRelayHandler(ctx *fiber.Ctx) error {
	var req adapter.TextRequest
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
		Api:         line.Host,
		Sk:          line.Auth,
		Mode:        getMode(ctx),
		Stream:      req.Stream,
		TokenRecord: make(chan *adapter.Usage),
	}
	//记录token
	common.RecordToken(req.Model, line.Id, adapterParams.TokenRecord)
	//转换文本请求
	params, err := relay.TextConversion(adapterParams, &req)
	if err != nil {
		return err
	}
	//转发
	resp, err := relay.Relay(params)
	if err != nil {
		return err
	}
	//处理文本响应
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
