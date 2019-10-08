package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/naruto2o2o/gormDb2Struct/input"
)

// MakeDataSourceName 拼接SourceName
func MakeDataSourceName(user, passwd, host string, port int) string {
	return user + ":" + passwd + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/?&parseTime=True"
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", MakeDataSourceName(*input.Flags.User, *input.Flags.Passwd, *input.Flags.Host, *input.Flags.Port))

	if err != nil {
		panic(fmt.Errorf("链接数据库错误%s", err))
	}

	return db
}

func selectDB(db *sql.DB, DBName string) {
	fmt.Println(DBName)
	_, err := db.Exec("use " + DBName)

	if err != nil {
		panic(err.Error())
	}
}
