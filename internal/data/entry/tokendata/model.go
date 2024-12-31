package tokendata

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/run-bigpig/mesh-api/internal/data/driver"
	"github.com/run-bigpig/mesh-api/internal/data/entry"
	"time"
)

var (
	columns            = []string{"model_name", "line_id", "prompt_tokens", "completion_tokens", "total_tokens"}
	includeTimeColumns = append(columns, entry.TimeRecords...)
	allColumns         = append([]string{"id"}, includeTimeColumns...)
)

func TableName() string {
	return "token_data"
}

func InsertOne(ctx context.Context, data *TokenData) error {
	sb := squirrel.Insert(TableName()).Columns(includeTimeColumns...).Values(data.ModelName, data.LineId, data.PromptTokens, data.CompletionTokens, data.TotalTokens, time.Now(), time.Now())
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}
