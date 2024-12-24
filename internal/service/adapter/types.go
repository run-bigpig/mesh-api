package adapter

import "io"

type Message struct {
	Role       string  `json:"role,omitempty"`
	Content    any     `json:"content,omitempty"`
	Name       string  `json:"name,omitempty"`
	ToolCalls  []*Tool `json:"tool_calls,omitempty"`
	ToolCallId string  `json:"tool_call_id,omitempty"`
}

type Tool struct {
	Id       string    `json:"id,omitempty"`
	Type     string    `json:"type,omitempty"`
	Function *Function `json:"function"`
}

type Function struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Parameters  any    `json:"parameters,omitempty"`
	Arguments   any    `json:"arguments,omitempty"`
}

type ResponseFormat struct {
	Type       string      `json:"type,omitempty"`
	JsonSchema *JSONSchema `json:"json_schema,omitempty"`
}

type JSONSchema struct {
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Strict      *bool                  `json:"strict,omitempty"`
}

type Audio struct {
	Voice  string `json:"voice,omitempty"`
	Format string `json:"format,omitempty"`
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

type TextRequest struct {
	Messages            []*Message      `json:"messages,omitempty"`
	Model               string          `json:"model,omitempty"`
	Store               bool            `json:"store,omitempty"`
	Metadata            any             `json:"metadata,omitempty"`
	FrequencyPenalty    float64         `json:"frequency_penalty,omitempty"`
	LogitBias           any             `json:"logit_bias,omitempty"`
	Logprobs            bool            `json:"logprobs,omitempty"`
	TopLogprobs         int             `json:"top_logprobs,omitempty"`
	MaxTokens           int             `json:"max_tokens,omitempty"`
	MaxCompletionTokens int             `json:"max_completion_tokens,omitempty"`
	N                   int             `json:"n,omitempty"`
	Modalities          []string        `json:"modalities,omitempty"`
	Prediction          any             `json:"prediction,omitempty"`
	Audio               *Audio          `json:"audio,omitempty"`
	PresencePenalty     float64         `json:"presence_penalty,omitempty"`
	ResponseFormat      *ResponseFormat `json:"response_format,omitempty"`
	Seed                float64         `json:"seed,omitempty"`
	ServiceTier         string          `json:"service_tier,omitempty"`
	Stop                any             `json:"stop,omitempty"`
	Stream              bool            `json:"stream,omitempty"`
	StreamOptions       *StreamOptions  `json:"stream_options,omitempty"`
	Temperature         float64         `json:"temperature,omitempty"`
	TopP                float64         `json:"top_p,omitempty"`
	TopK                int             `json:"top_k,omitempty"`
	Tools               []*Tool         `json:"tools,omitempty"`
	ToolChoice          any             `json:"tool_choice,omitempty"`
	ParallelTooCalls    bool            `json:"parallel_tool_calls,omitempty"`
	User                string          `json:"user,omitempty"`
	FunctionCall        any             `json:"function_call,omitempty"`
	Functions           any             `json:"functions,omitempty"`
	Input               any             `json:"input,omitempty"`
	EncodingFormat      string          `json:"encoding_format,omitempty"`
	Dimensions          int             `json:"dimensions,omitempty"`
	Prompt              any             `json:"prompt,omitempty"`
	Quality             string          `json:"quality,omitempty"`
	Size                string          `json:"size,omitempty"`
	Style               string          `json:"style,omitempty"`
	Instruction         string          `json:"instruction,omitempty"`
	NumCtx              int             `json:"num_ctx,omitempty"`
	Suffix              string          `json:"suffix,omitempty"`
}

type Choice struct {
	Index        int      `json:"index"`
	Delta        *Message `json:"delta"`
	FinishReason string   `json:"finish_reason,omitempty"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionsStreamResponse struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage,omitempty"`
}

type CompletionsStreamResponse struct {
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

type ImageRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt" binding:"required"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	Quality        string `json:"quality,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	Style          string `json:"style,omitempty"`
	User           string `json:"user,omitempty"`
}

type Params struct {
	Api    string
	Sk     string
	Mode   int8
	Stream bool
}

type ConversionParams struct {
	Url    string
	Method string
	Header map[string]string
	Body   io.Reader
}
