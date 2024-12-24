package openai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"io"
	"net/http"
	"strings"
)

const (
	dataPrefix       = "data: "
	done             = "[DONE]"
	dataPrefixLength = len(dataPrefix)
)

type OpenAI struct {
	ctx *fiber.Ctx
}

func New(ctx *fiber.Ctx) *OpenAI {
	return &OpenAI{
		ctx: ctx,
	}
}

func (o *OpenAI) TextConversion(params *adapter.Params, req *adapter.TextRequest) (*adapter.ConversionParams, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", params.Sk),
	}
	if req.Stream {
		header["Accept"] = "text/event-stream"
	}
	return &adapter.ConversionParams{
		Url:    fmt.Sprintf("%s%s", params.Api, o.ctx.Path()),
		Method: o.ctx.Method(),
		Body:   bytes.NewReader(body),
		Header: header,
	}, nil
}

func (o *OpenAI) TextResponse(params *adapter.Params, resp *http.Response) error {
	if params.Stream {
		return o.StreamResponse(resp)
	}
	return o.NormalResponse(resp)
}

func (o *OpenAI) ImageConversion(params *adapter.Params, req *adapter.ImageRequest) (*adapter.ConversionParams, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", params.Sk),
	}
	return &adapter.ConversionParams{
		Url:    fmt.Sprintf("%s%s", params.Api, o.ctx.Path()),
		Method: o.ctx.Method(),
		Body:   bytes.NewReader(body),
		Header: header,
	}, nil
}

func (o *OpenAI) ImageResponse(resp *http.Response) error {
	return o.NormalResponse(resp)
}

func (o *OpenAI) Relay(params *adapter.ConversionParams) (*http.Response, error) {
	req, err := http.NewRequest(params.Method, params.Url, params.Body)
	if err != nil {
		return nil, err
	}
	for key, value := range params.Header {
		req.Header.Set(key, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NormalResponse 普通响应
func (o *OpenAI) NormalResponse(resp *http.Response) error {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = o.ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	return o.ctx.Status(resp.StatusCode).Send(responseBody)
}

// StreamResponse 流式响应
func (o *OpenAI) StreamResponse(resp *http.Response) error {
	//设置header
	o.ctx.Set("Content-Type", "text/event-stream")
	o.ctx.Set("Cache-Control", "no-cache")
	o.ctx.Set("Connection", "keep-alive")
	o.ctx.Set("Transfer-Encoding", "chunked")
	//设置body
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)
	var closed bool
	doneRendered := false
	//设置body stream
	o.ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		defer resp.Body.Close()
		for scanner.Scan() {
			if closed {
				return
			}
			data := scanner.Text()
			if len(data) < dataPrefixLength {
				continue
			}
			if data[:dataPrefixLength] != dataPrefix && data[:dataPrefixLength] != done {
				continue
			}
			if strings.HasPrefix(data[dataPrefixLength:], done) {
				adapter.WriteStringData(w, data, &closed)
				doneRendered = true
				continue
			}
			adapter.WriteStringData(w, data, &closed)
		}
		if err := scanner.Err(); err != nil {
			log.Error("error reading stream: " + err.Error())
		}

		if !doneRendered {
			adapter.WriteStringData(w, done, &closed)
		}
	})
	return nil
}
