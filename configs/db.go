package configs

import "time"

const (
	db         = "mysql"
	connString = "root:asdasd123@tcp(localhost:3306)/uangq_db"
)

type DbConfiguration struct {
	Db              string
	ConnString      string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
}

var DbConfig = DbConfiguration{
	Db:              db,
	ConnString:      connString,
	MaxOpenConn:     100,
	MaxIdleConn:     10,
	ConnMaxLifetime: 0,
}
