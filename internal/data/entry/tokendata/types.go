package tokendata

import "time"

type TokenData struct {
	Id               int64     `db:"id" json:"id"`
	LineId           int64     `db:"line_id" json:"line_id"`
	ModelName        string    `db:"model_name" json:"model_name"`
	PromptTokens     int64     `db:"prompt_tokens" json:"prompt_tokens"`
	CompletionTokens int64     `db:"completion_tokens" json:"completion_tokens"`
	TotalTokens      int64     `db:"total_tokens" json:"total_tokens"`
	CreatedTime      time.Time `db:"created_time" json:"created_time"`
	UpdatedTime      time.Time `db:"updated_time" json:"updated_time"`
}
