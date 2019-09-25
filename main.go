package main

import (
	"github.com/Shelnutt2/db2struct/input"
	"github.com/droundy/goopt"
	_ "github.com/go-sql-driver/mysql"
)

const defaultConfigPath = "../configs/db.yaml"

// 指定数据库配置文件路径(只包含连接相关配置)
var configPath = goopt.String([]string{"-c", "--config"}, "", "配置文件的路径")

// func init() {

// 	goopt.Version = "0.0.2"
// 	goopt.Summary = "db2struct [-H] [-p] [-v] --package pkgName --struct structName --database databaseName --table tableName"

// 	//Parse options
// 	goopt.Parse(nil)

// }

// func main() {

// 	// Username is required
// 	if user == nil || *user == "user" {
// 		fmt.Println("请输入mysql用户名 --user=name")
// 		return
// 	}

// 	if passwd != nil && *passwd == "" {
// 		fmt.Print("Password: ")
// 		pass, err := gopass.GetPasswd()
// 		stringPass := string(pass)
// 		passwd = &stringPass
// 		if err != nil {
// 			fmt.Println("密码错误: " + err.Error())
// 			return
// 		}
// 	} else if passwd == nil {
// 		p := ""
// 		passwd = &p
// 	}

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

func main() {
	input.ListenFlag()
}
