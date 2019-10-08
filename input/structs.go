package input

// Flag 参数集合
type Flag struct {
	ConfigFile     *string              `commands:"-c,--config" tag:"配置文件路径(指定此参数时将其他参数将忽略所有配合来源于配置文件)"`
	Host           *string              `commands:"-H,--host" def:"127.0.0.1" tag:"数据库的host"`
	Port           *int                 `commands:"--port" def:"3306" tag:"数据库端口"`
	Db             *string              `commands:"-d,--db"  tag:"要链接的数据库"`
	DbMap          *map[string][]string `commands:"-d,--db"  tag:"要链接的数据库"`
	Table          *string              `commands:"-t,--table" tag:"要生成的表"`
	User           *string              `commands:"-u,--user"  tag:"数据库用户"`
	Passwd         *string              `commands:"-p,--passwd" tag:"数据库密码"`
	Verbose        *bool                `commands:"--verbose" tagY:"显示执行过程"`
	PackageName    *string              `commands:"--package" tag:"设置输出的package名"`
	StructName     *string              `commands:"--struct" tag:"设置输出的结构体名"`
	JSONAnnotation *bool                `commands:"--json" tagY:"输出json标签"`
	GormAnnotation *bool                `commands:"--gorm" tagY:"输出文件中添加gorm表名函数"`
	GureguTypes    *bool                `commands:"--guregu" tagY:"使用guregu空类型"`
	TargetFile     *string              `commands:"-o,--output" def:"./" tag:"输出的目录"`
}
