package adapter

import "net/http"

type Adapter interface {
	Relay(params *ConversionParams) (*http.Response, error)
	TextConversion(params *Params, req *TextRequest) (*ConversionParams, error)
	TextResponse(params *Params, resp *http.Response) error
	ImageConversion(params *Params, req *ImageRequest) (*ConversionParams, error)
	ImageResponse(resp *http.Response) error
}
