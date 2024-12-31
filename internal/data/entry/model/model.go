package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/run-bigpig/mesh-api/internal/data/driver"
	"github.com/run-bigpig/mesh-api/internal/data/entry"
	"time"
)

var (
	columns            = []string{"name", "class", "status"}
	includeTimeColumns = append(columns, entry.TimeRecords...)
	allColumns         = append([]string{"id"}, includeTimeColumns...)
)

func TableName() string {
	return "model"
}

// InsertOne 插入一条model数据
func InsertOne(ctx context.Context, data *Model) error {
	sb := squirrel.Insert(TableName()).Columns(includeTimeColumns...).Values(data.Name, data.Class, data.Status, time.Now(), time.Now())
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}

// FindOne 根据id查找一条model数据
func FindOne(ctx context.Context, id int64) (*Model, error) {
	sb := squirrel.Select(allColumns...).From(TableName()).Where(squirrel.Eq{"id": id})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data Model
	err = driver.GetDb().GetContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindOneByName 根据name查找一条model数据
func FindOneByName(ctx context.Context, name string) (*Model, error) {
	sb := squirrel.Select(allColumns...).From(TableName()).Where(squirrel.Eq{"name": name})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data Model
	err = driver.GetDb().GetContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindAll 查找所有model数据
func FindAll(ctx context.Context, req *FindModelRequest) ([]*Model, error) {
	sb := squirrel.Select(allColumns...).From(TableName())
	if req.Name != "" {
		sb = sb.Where(squirrel.Like{"name": req.Name})
	}
	if req.Class != 0 {
		sb = sb.Where(squirrel.Eq{"class": req.Class})
	}
	if req.Status != 0 {
		sb = sb.Where(squirrel.Eq{"status": req.Status})
	}
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data []*Model
	err = driver.GetDb().SelectContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateOne 更新一条model数据
func UpdateOne(ctx context.Context, data *Model) error {
	sb := squirrel.Update(TableName()).Set("name", data.Name).Set("class", data.Class).Set("status", data.Status).Set("updated_time", time.Now()).Where(squirrel.Eq{"id": data.Id})
	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	_, err = driver.GetDb().ExecContext(ctx, query, args...)
	return err
}

// DeleteOne 删除一条model数据
func DeleteOne(ctx context.Context, id int64) error {
	tx, err := driver.GetDb().BeginTxx(ctx, nil)
	sb := squirrel.Delete(TableName()).Where(squirrel.Eq{"id": id})
	query, args, err := sb.ToSql()
	if err != nil {
		return tx.Rollback()
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		return tx.Rollback()
	}
	sb = squirrel.Delete("model_line").Where(squirrel.Eq{"model_id": id})
	query, args, err = sb.ToSql()
	if err != nil {
		return tx.Rollback()
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

// SetModelLine 批量设置model的line
func SetModelLine(ctx context.Context, req *SetModelLineRequest) error {
	tx, err := driver.GetDb().BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	deleteSb := squirrel.Delete("model_line").Where(squirrel.Eq{"model_id": req.ModelIds})
	query, args, err := deleteSb.ToSql()
	if err != nil {
		return tx.Rollback()
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		return tx.Rollback()
	}
	for _, modelId := range req.ModelIds {
		insertSb := squirrel.Insert("model_line").Columns("model_id", "line_id")
		for _, lineId := range req.LineIds {
			insertSb = insertSb.Values(modelId, lineId)
		}
		query, args, err = insertSb.ToSql()
		if err != nil {
			return tx.Rollback()
		}
		_, err = tx.Exec(query, args...)
		if err != nil {
			return tx.Rollback()
		}
	}
	return tx.Commit()
}

// FindModelLine 查找model的line
func FindModelLine(ctx context.Context, modelId int64) ([]int64, error) {
	sb := squirrel.Select("line_id").From("model_line").Where(squirrel.Eq{"model_id": modelId})
	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}
	var data []int64
	err = driver.GetDb().SelectContext(ctx, &data, query, args...)
	if err != nil {
		return nil, err
	}
	return data, nil
}
