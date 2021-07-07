# xorm tools

基于https://github.com/go-xorm/cmd

虽然原工具可自动将bigint转成int64，但习惯将时间戳字段为int类型。

添加配置
int64Cols：该字段用于逆向生成go文件时，将指定字段指定为int64
eg:
```
lang=go
genJson=1
int64Cols=updated_at,created_at
```
该配置在逆向时，可将"updated_at"，"created_at"指定为int64。

reverse
```
xorm-cmd reverse mysql USER:PWD@tcp(XXX:3306)/DB?charset=utf8 tpl res TABLE
```
