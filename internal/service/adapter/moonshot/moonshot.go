package moonshot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/errorcode"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/openai"
	"net/http"
)

type Moonshot struct {
	ctx *fiber.Ctx
	oai *openai.OpenAI
}

func New(ctx *fiber.Ctx) *Moonshot {
	return &Moonshot{ctx: ctx, oai: openai.New(ctx)}
}

func (m *Moonshot) TextConversion(params *adapter.Params, req *adapter.TextRequest) (*adapter.ConversionParams, error) {
	return m.oai.TextConversion(params, req)
}

func (m *Moonshot) Relay(params *adapter.ConversionParams) (*http.Response, error) {
	return m.oai.Relay(params)
}

func (m *Moonshot) TextResponse(params *adapter.Params, resp *http.Response) error {
	return m.oai.TextResponse(params, resp)
}

func (m *Moonshot) ImageConversion(params *adapter.Params, req *adapter.ImageRequest) (*adapter.ConversionParams, error) {
	return nil, errorcode.ErrorCodeNotImplemented
}

func (m *Moonshot) ImageResponse(resp *http.Response) error {
	return errorcode.ErrorCodeNotImplemented
}
