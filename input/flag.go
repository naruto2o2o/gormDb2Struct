package input

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/droundy/goopt"
	"github.com/howeyc/gopass"
)

// Flags 参数结构体
var Flags flag

// ListenFlag 监听flag
func ListenFlag() {
	err := Flags.dealTag()

	if err != nil {
		fmt.Println(err)
		return
	}
	goopt.Version = "1.0.0"
	goopt.Summary = " [-H] [--port] -d -t -u -p  --package pkgName --struct structName --database databaseName --table tableName"
	goopt.Parse(nil)
}

func (f *flag) dealTag() error {
	fValue := reflect.ValueOf(f)
	fType := reflect.TypeOf(f)

	for i := 0; i < fValue.Elem().NumField(); i++ {
		fTypeI := fType.Elem().Field(i)
		fValueI := fValue.Elem().Field(i)

		commands := fTypeI.Tag.Get("commands")

		if commands == "" {
			continue
		}

		commandsArr := strings.Split(commands, ",")

		def := fTypeI.Tag.Get("def")
		tag := fTypeI.Tag.Get("tag")
		tagY := fTypeI.Tag.Get("tagY")

		switch fTypeI.Type.String() {
		case "*string":
			ff := goopt.String(commandsArr, def, tag)
			fValueI.Set(reflect.ValueOf(ff))
		case "*int":
			defI, err := strconv.Atoi(def)

			if err != nil {
				return err
			}

			ff := goopt.Int(commandsArr, defI, tag)
			fValueI.Set(reflect.ValueOf(ff))
		case "*bool":
			ff := goopt.Flag(commandsArr, []string{}, tagY, "")
			fValueI.Set(reflect.ValueOf(ff))
		}
	}

	return nil
}

// ValidateFlags 验证参数
func (f *flag) ValidateFlags() {

	if *f.ConfigFile != "" {
		loadConfig(*f.ConfigFile)
	}

	if *f.Host == "" {
		log.Fatal("Host是必传参数")
	}

	if *f.Port == 0 {
		log.Fatal("Port是必传参数")
	}

	if *f.User == "" {
		log.Fatal("User是必传参数")
	}

	if *f.Db == "" && len(*f.DbMap) == 0 {
		log.Fatal("Db是必传参数")
	}

	if len(*f.DbMap) > 0 && (*f.Db != "" && *f.Table != "") {
		log.Fatal("不可以在命令中与配置文件中同时指定db")
	}

	if *f.Passwd == "" {
		fmt.Print("Password: ")
		pass, err := gopass.GetPasswd()
		stringPass := string(pass)
		f.Passwd = &stringPass
		if err != nil {
			fmt.Println("密码错误: " + err.Error())
			return
		}
	}

	if *f.Verbose {
		fmt.Println("正在连接mysqlserver " + *f.Host + ":" + strconv.Itoa(*f.Port))
	}
}
