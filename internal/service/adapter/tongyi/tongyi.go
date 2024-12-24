package tongyi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/openai"
	"net/http"
)

type Tongyi struct {
	ctx *fiber.Ctx
	oai *openai.OpenAI
}

func New(ctx *fiber.Ctx) *Tongyi {
	return &Tongyi{
		ctx: ctx,
		oai: openai.New(ctx),
	}
}

func (t *Tongyi) Relay(params *adapter.ConversionParams) (*http.Response, error) {
	return t.oai.Relay(params)
}

func (t *Tongyi) TextConversion(params *adapter.Params, req *adapter.TextRequest) (*adapter.ConversionParams, error) {
	path := t.ctx.Path()
	if params.Mode == adapter.Completions {
		path = "/v1/chat/completions"
		req.Messages = []*adapter.Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: req.Prompt,
			},
		}
	}
	t.ctx.Path(fmt.Sprintf("%s%s", "/compatible-mode", path))
	return t.oai.TextConversion(params, req)
}

func (t *Tongyi) TextResponse(params *adapter.Params, resp *http.Response) error {
	return t.oai.TextResponse(params, resp)
}

func (t *Tongyi) ImageConversion(params *adapter.Params, req *adapter.ImageRequest) (*adapter.ConversionParams, error) {
	return t.oai.ImageConversion(params, req)
}

func (t *Tongyi) ImageResponse(resp *http.Response) error {
	return t.oai.ImageResponse(resp)
}
