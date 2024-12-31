package model

import "time"

type Model struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Class       int8      `db:"class" json:"class"`
	Status      int8      `db:"status" json:"status"`
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}

type FindModelRequest struct {
	Name   string `json:"name"`
	Class  int8   `json:"class"`
	Status int8   `json:"status"` //1:启用 2:禁用
}

type SetModelLineRequest struct {
	ModelIds []int64 `json:"model_ids"`
	LineIds  []int64 `json:"line_ids"`
}
