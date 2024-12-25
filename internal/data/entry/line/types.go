package line

import "time"

type Line struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Host        string    `db:"host" json:"host"`
	Status      int8      `db:"status" json:"status"`
	Auth        string    `db:"auth" json:"auth"`
	Weight      int8      `db:"weight" json:"weight"`
	Adapter     string    `db:"adapter" json:"adapter"`
	ProxyId     int64     `db:"proxy_id" json:"proxy_id"`
	IsProxy     int8      `db:"is_proxy" json:"is_proxy"`
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}

type FindLineRequest struct {
	Name    string `json:"name"`
	Adapter string `json:"adapter"`
	Status  int8   `json:"status"`
}
