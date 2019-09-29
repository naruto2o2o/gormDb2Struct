package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naruto2o2o/gormDb2Struct/input"
)

// func main() {

// 	if *verbose {
// 		fmt.Println("正在连接mysqlserver " + *host + ":" + strconv.Itoa(*prot))
// 	}

// 	if db == nil || *db == "" {
// 		fmt.Println("请选择数据库")
// 		return
// 	}

// 	if table == nil || *table == "" {
// 		fmt.Println("请选择要同步的表")
// 		return
// 	}

// 	columnDataTypes, err := GetColumnsFromMysqlTable(*user, *passwd, *host, *prot, *db, *table)

// 	if err != nil {
// 		fmt.Println("Error in selecting column data information from mysql information schema")
// 		return
// 	}

// 	// If structName is not set we need to default it
// 	if structName == nil || *structName == "" {
// 		*structName = "newstruct"
// 	}
// 	// If packageName is not set we need to default it
// 	if packageName == nil || *packageName == "" {
// 		*packageName = "newpackage"
// 	}
// 	// Generate struct string based on columnDataTypes
// 	struc, err := Generate(*columnDataTypes, *table, *structName, *packageName, *jsonAnnotation, *gormAnnotation, *gureguTypes)

// 	if err != nil {
// 		fmt.Println("Error in creating struct from json: " + err.Error())
// 		return
// 	}
// 	if targetFile != nil && *targetFile != "" {
// 		file, err := os.OpenFile(*targetFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 		if err != nil {
// 			fmt.Println("Open File fail: " + err.Error())
// 			return
// 		}
// 		length, err := file.WriteString(string(struc))
// 		if err != nil {
// 			fmt.Println("Save File fail: " + err.Error())
// 			return
// 		}
// 		fmt.Printf("wrote %d bytes\n", length)
// 	} else {
// 		fmt.Println(string(struc))
// 	}

// }

var db *sql.DB

func main() {
	input.ListenFlag()
	input.Flags.ValidateFlags()

	db, err := sql.Open("mysql")

	if len(*input.Flags.DbMap) > 0 {
		// 按照配置遍历数据库
	} else if *input.Flags.Db != "" {
		if *input.Flags.Table != "" {
			// 进行单张白表映射
		} else {
			// 遍历数据库下所有的表
		}
	}

}
