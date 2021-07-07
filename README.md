# xorm tools

基于https://github.com/go-xorm/cmd

添加配置
int64Cols：该字段用于逆向生成go文件时，将指定字段指定为int64

虽然原工具可自动将bigint转成int64，但习惯将时间戳字段为int类型。