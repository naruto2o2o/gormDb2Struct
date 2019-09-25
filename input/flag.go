package input

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type flag struct {
	Host           *string `commands:"-h,--host" tag:"数据库的host"`
	Port           *int
	Db             *string
	Table          *string
	User           *string
	Passwd         *string
	Verbose        *bool   // Verbose        输出执行过程
	PackageName    *string // PackageName    导出文件的包名
	StructName     *string // StructName     导出文件中结构体的名字
	JSONAnnotation *bool   // JSONAnnotation 是否输出json标签
	GormAnnotation *bool   // GormAnnotation 加入gorm 获取表名函数
	GureguTypes    *bool   // GureguTypes    是否使用gure null类型
	TargetFile     *string // TargetFile     输出目录
}

// ListenFlag 监听flag
func ListenFlag() {
	var f flag
	err := f.dealTag()

	if err != nil {
		fmt.Println(err)
	}
	// 数据库相关配置
	// f.Host = goopt.String([]string{"-h", "--host"}, "", "数据库的host")
	// f.Port = goopt.Int([]string{"-P", "--port"}, 3306, "数据库端口")
	// f.Table = goopt.String([]string{"-t", "--table"}, "", "表名")
	// f.Db = goopt.String([]string{"-d", "--database"}, "nil", "设置数据库")
	// f.User = goopt.String([]string{"-u", "--user"}, "user", "用户")
	// f.Passwd = goopt.String([]string{"-p", "--password"}, "", "mysql密码")

	// // 输出文件格式相关配置
	// f.Verbose = goopt.Flag([]string{"-v", "--verbose"}, []string{}, "允许没逼用的输出(执行过程输出)", "")
	// f.PackageName = goopt.String([]string{"--package"}, "", "输出文件的包名")
	// f.StructName = goopt.String([]string{"--struct"}, "", "结构体名称")
	// f.JSONAnnotation = goopt.Flag([]string{"--json"}, []string{"--no-json"}, "添加json标签(默认)", "禁用json标签")
	// f.GormAnnotation = goopt.Flag([]string{"--gorm"}, []string{}, "添加gorm表名方法", "")
	// f.GureguTypes = goopt.Flag([]string{"--guregu"}, []string{}, "使用gure null类型", "")

	// // 输出的目标目录
	// f.TargetFile = goopt.String([]string{"-o", "--output"}, "", "输出目录设置")

}

func (f *flag) dealTag() error {
	ft := reflect.TypeOf(*f)

	for i := 0; i < ft.NumField(); i++ {
		fti := ft.Field(i)
		commands := fti.Tag.Get("commands")

		if commands == "" {
			return errors.New("tag的标签是必须的")
		}

		commandsArr := strings.Split(commands, ",")

		def := fti.Tag.Get("def")
		tag := fti.Tag.Get("tag")

		switch fti.Type.Name() {
		case "*string":
			// fti.Type. goopt.String(commandsArr, def, tag)
		}

	}

	return nil
}
