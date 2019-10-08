# gormDb2Struct

将数据库表结构映射为gorm结构体
目前只支持mysql

## 使用

```shell
go build ./
mv gormDb2Struct $$GOBIN/gormDb2Struct
gormDb2Struct --help
```

## 参数解析

```config
  -c, --config=         配置文件路径(指定此参数时将其他参数将忽略所有配合来源于配置文件)
  -H, --host=127.0.0.1  数据库的host
  --port=3306           数据库端口
  -d, --db=             要链接的数据库
  -t, --table=          要生成的表
  -u, --user=           数据库用户
  -p, --passwd=         数据库密码
  --verbose             显示执行过程
  --package=            设置输出的package名
  --struct=             设置输出的结构体名
  --json                输出json标签
  --gorm                输出文件中添加gorm表名函数
  --guregu              使用guregu空类型
  -o, --output=./       输出的目录
  -h, --help            Show usage message
  --version             Show version
```

## 配置文件解析

```yaml
host: localhost
port: 3306
user: jinzhiwen
passwd : aa942236
DbMap :        # 要进行映射的配置
      db1 :
          - table1
          - table2
      db2 : 
          - table1
          - table2
Verbose : true  # 显示执行过程
JSONAnnotation : true   # 是否加入jsontag
GormAnnotation : true   # 是否加入gormtag
GureguTypes : true      # 是够使用Guregu null type
TargetFile : ./         # 输出映射文件的的根目录

```
