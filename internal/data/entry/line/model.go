package line

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/run-bigpig/mesh-api/internal/data/driver"
	"github.com/run-bigpig/mesh-api/internal/data/entry"
	"time"
)

var (
	columns            = []string{"name", "host", "status", "auth", "weight", "adapter", "proxy_id", "is_proxy"}
	includeTimeColumns = append(columns, entry.TimeRecords...)
	allColumns         = append([]string{"id"}, includeTimeColumns...)
)

func TableName() string {
	return "line"
}

// InsertOne 插入一条line数据
func InsertOne(ctx context.Context, data *Line) error {
	sb := squirrel.Insert(TableName()).Columns(includeTimeColumns...).Values(data.Name, data.Host, data.Status, data.Auth, data.Weight, data.Adapter, data.ProxyId, data.IsProxy, time.Now(), time.Now())
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}

// FindOne 根据id查找一条line数据
func FindOne(ctx context.Context, id int64) (*Line, error) {
	sb := squirrel.Select(allColumns...).From(TableName()).Where(squirrel.Eq{"id": id})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data Line
	err = driver.GetDb().GetContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindAll 查找所有line数据
func FindAll(ctx context.Context, req *FindLineRequest) ([]*Line, error) {
	sb := squirrel.Select(allColumns...).From(TableName())
	if req.Name != "" {
		sb = sb.Where(squirrel.Like{"name": req.Name})
	}
	if req.Adapter != "" {
		sb = sb.Where(squirrel.Eq{"adapter": req.Adapter})
	}
	if req.Status != 0 {
		sb = sb.Where(squirrel.Eq{"status": req.Status})
	}
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data []*Line
	err = driver.GetDb().SelectContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateOne 更新一条line数据
func UpdateOne(ctx context.Context, data *Line) error {
	sb := squirrel.Update(TableName()).SetMap(map[string]interface{}{
		"name":         data.Name,
		"host":         data.Host,
		"status":       data.Status,
		"auth":         data.Auth,
		"proxy_id":     data.ProxyId,
		"weight":       data.Weight,
		"adapter":      data.Adapter,
		"is_proxy":     data.IsProxy,
		"updated_time": time.Now(),
	}).Where(squirrel.Eq{"id": data.Id})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}

// DeleteOne 删除一条line数据
func DeleteOne(ctx context.Context, id int64) error {
	sb := squirrel.Delete(TableName()).Where(squirrel.Eq{"id": id})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}
