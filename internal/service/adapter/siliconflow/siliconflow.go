package siliconflow

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"github.com/run-bigpig/mesh-api/internal/service/adapter/openai"
	"net/http"
)

type SiliconFlow struct {
	ctx *fiber.Ctx
	oai *openai.OpenAI
}

func New(ctx *fiber.Ctx) *SiliconFlow {
	return &SiliconFlow{
		ctx: ctx,
		oai: openai.New(ctx),
	}
}

func (s *SiliconFlow) TextConversion(params *adapter.Params, req *adapter.TextRequest) (*adapter.ConversionParams, error) {
	return s.oai.TextConversion(params, req)
}

func (s *SiliconFlow) Relay(params *adapter.ConversionParams) (*http.Response, error) {
	return s.oai.Relay(params)
}
func (s *SiliconFlow) TextResponse(params *adapter.Params, resp *http.Response) error {
	return s.oai.TextResponse(params, resp)
}

func (s *SiliconFlow) ImageConversion(params *adapter.Params, req *adapter.ImageRequest) (*adapter.ConversionParams, error) {
	return s.oai.ImageConversion(params, req)
}

func (s *SiliconFlow) ImageResponse(resp *http.Response) error {
	return s.oai.ImageResponse(resp)
}
