package deepseek

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/openai"
	"net/http"
)

type DeepSeek struct {
	ctx *fiber.Ctx
	oai *openai.OpenAI
}

func New(ctx *fiber.Ctx) *DeepSeek {
	return &DeepSeek{ctx: ctx, oai: openai.New(ctx)}
}

func (d *DeepSeek) Relay(params *adapter.ConversionParams) (*http.Response, error) {
	return d.oai.Relay(params)
}

func (d *DeepSeek) TextConversion(params *adapter.Params, req *adapter.TextRequest) (*adapter.ConversionParams, error) {
	if params.Mode == adapter.Completions {
		d.ctx.Path("/beta/completions")
	}
	return d.oai.TextConversion(params, req)
}

func (d *DeepSeek) TextResponse(params *adapter.Params, resp *http.Response) error {
	return d.oai.TextResponse(params, resp)
}

func (d *DeepSeek) ImageConversion(params *adapter.Params, req *adapter.ImageRequest) (*adapter.ConversionParams, error) {
	return nil, errorcode.ErrorCodeNotImplemented
}

func (d *DeepSeek) ImageResponse(resp *http.Response) error {
	return errorcode.ErrorCodeNotImplemented
}
