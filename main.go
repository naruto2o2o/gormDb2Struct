package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naruto2o2o/gormDb2Struct/bigcarm"
	"github.com/naruto2o2o/gormDb2Struct/input"
)

var db *sql.DB

func main() {
	input.ListenFlag()
	input.Flags.ValidateFlags()

	db = connect()
	defer db.Close()

	if input.Flags.DbMap != nil && len(*input.Flags.DbMap) > 0 {
		// 按照配置遍历数据库
		for k, v := range *input.Flags.DbMap {
			selectDB(db, k)
			*input.Flags.Db = k
			for _, table := range v {
				input.Flags.Table = &table
				handel(db, input.Flags)
			}
		}

	} else if *input.Flags.Db != "" {
		if *input.Flags.Table != "" {
			selectDB(db, *input.Flags.Db)
			// 进行单张白表映射
			handel(db, input.Flags)
		} else {
			selectDB(db, *input.Flags.Db)
			rows, err := db.Query("select table_name from information_schema.tables where table_schema=?", *input.Flags.Db)

			if err != nil {
				panic(err.Error())
			}

			for rows.Next() {
				rows.Scan(input.Flags.Table)
				handel(db, input.Flags)
			}
		}
	}
}

func handel(db *sql.DB, tags input.Flag) {
	dbName := *tags.Db
	tableName := *tags.Table
	jsonAnnotation := *tags.JSONAnnotation
	gormAnnotation := *tags.GormAnnotation
	gureguTypes := *tags.GureguTypes
	targetFile := *tags.TargetFile + "/" + dbName + "/"

	var structName string
	var packageName string

	if *tags.StructName == "" {
		structName = bigcarm.Marshal(*tags.Table)
	} else {
		structName = *tags.StructName
	}

	if *tags.PackageName == "" {
		packageName = *tags.Db
	} else {
		packageName = *tags.Db
	}

	columnDataTypes, err := GetColumnsFromMysqlTable(db, dbName, tableName)

	if err != nil {
		fmt.Println("Error in selecting column data information from mysql information schema")
		return
	}

	tableCom, err := getTableCom(dbName, tableName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Generate struct string based on columnDataTypes
	struc, err := Generate(*columnDataTypes, tableName, tableCom, structName, packageName, jsonAnnotation, gormAnnotation, gureguTypes)

	if err != nil {
		fmt.Println("Error in creating struct from json: " + err.Error())
		return
	}
	fmt.Println(targetFile)
	err = os.MkdirAll(targetFile, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	targetFile = targetFile + "/" + tableName + ".go"
	if targetFile != "" {
		file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Open File fail: " + err.Error())
			return
		}
		length, err := file.WriteString(string(struc))
		if err != nil {
			fmt.Println("Save File fail: " + err.Error())
			return
		}
		fmt.Printf("wrote %d bytes\n", length)
	} else {
		fmt.Println(string(struc))
	}
}
