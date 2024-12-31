package common

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/mesh-api/internal/data/entry/line"
	"github.com/run-bigpig/mesh-api/internal/data/entry/model"
	"github.com/run-bigpig/mesh-api/internal/data/entry/tokendata"
	"github.com/run-bigpig/mesh-api/internal/service/adapter"
	"math/rand"
	"time"
)

// RandomLineByWeight 根据线路权重随机选择线路
func RandomLineByWeight(ctx context.Context, modelName string) (*line.Line, error) {
	// 根据model获取线路列表
	m, err := model.FindOneByName(ctx, modelName)
	if err != nil {
		return nil, err
	}
	//通过模型id获取线路Ids
	lineIds, err := model.FindModelLine(ctx, m.Id)
	if err != nil {
		return nil, err
	}
	if len(lineIds) == 0 {
		return nil, errors.New("no line")
	}
	//通过线路Ids获取线路列表
	lines, err := line.FindLineByIds(ctx, lineIds)
	if err != nil {
		return nil, err
	}
	totalWeight := 0
	for _, l := range lines {
		totalWeight += int(l.Weight)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomWeight := r.Intn(totalWeight)

	for _, l := range lines {
		randomWeight -= int(l.Weight)
		if randomWeight < 0 {
			return l, nil
		}
	}
	return lines[0], nil
}

// RecordToken 异步记录token使用到数据库 监听通道获取使用量
func RecordToken(modelName string, lineId int64, tokenRecord chan *adapter.Usage) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	go func() {
		defer cancel()
		select {
		case usage, ok := <-tokenRecord:
			if !ok {
				return
			}
			err := tokendata.InsertOne(ctx, &tokendata.TokenData{
				ModelName:        modelName,
				LineId:           lineId,
				PromptTokens:     int64(usage.PromptTokens),
				CompletionTokens: int64(usage.CompletionTokens),
				TotalTokens:      int64(usage.TotalTokens),
			})
			if err != nil {
				log.Errorf("RecordToken error: %v", err)
			}
		case <-ctx.Done():
			log.Infof("RecordToken timeout")
		}
	}()
}
