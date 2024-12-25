package driver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/run-bigpig/mesh-api/internal/config"
	"sync"
)

var (
	sqlDb *sqlx.DB
	once  sync.Once
)

func NewMySQL() {
	once.Do(func() {
		db, err := sqlx.Open("mysql", config.Get().Mysql.DSN)
		if err != nil {
			panic(err)
		}
		db.SetMaxIdleConns(100)
		db.SetMaxOpenConns(200)
		db.SetConnMaxLifetime(10)
		sqlDb = db
	})
}

func GetDb() *sqlx.DB {
	return sqlDb
}
