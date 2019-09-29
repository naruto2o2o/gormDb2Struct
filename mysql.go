package main

import (
	"database/sql"
	"strconv"

	"github.com/naruto2o2o/gormDb2Struct/input"
)

// MakeDataSourceName 拼接SourceName
func MakeDataSourceName(user string, passwd string, host string, port int, database string) string {
	return user + ":" + passwd + "@tcp(" + host + ":" + strconv.Itoa(port) + ")?&parseTime=True"
}

func connect() {
	db, err := sql.Open("mysql", makeDataSourceName(*input.Flags.User, *input.Flags.Passwd, *input.Flags.Host, *input.Flags.Port, *input.Flags.Db))

	if err != nil {
		panic("链接数据库错误" + err)
	}

	return db
}
