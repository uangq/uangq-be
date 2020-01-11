package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"uangq-be/common/logger"
	"uangq-be/configs"
)

type DbService struct {
	Db *sql.DB
}

var MyDbService DbService

func InitDB() {
	db, err := sql.Open(configs.DbConfig.Db, configs.DbConfig.ConnString)

	if err != nil {
		logger.Log("Can't connect to db." + err.Error())
	}

	db.SetConnMaxLifetime(configs.DbConfig.ConnMaxLifetime)
	db.SetMaxIdleConns(configs.DbConfig.MaxIdleConn)
	db.SetMaxOpenConns(configs.DbConfig.MaxOpenConn)

	MyDbService = DbService{
		Db: db,
	}
}
