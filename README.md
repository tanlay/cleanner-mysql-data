# cleanner-mysql-data

## 功能
- 查询需要删除的数量
- 查询删除的ids
- 根据查询到的ids删除数据

## 安装依赖
```shell
go mod tidy
go build
```

## 删除命令
```go
cleanner-mysql-data hand-clean -c config.toml -d "2006-01-02 15:04:05"
```

###conf.toml配置文件
```toml
# 如：总数据20亿次，每次串行删除2000万（per_total_count）条；
# 2000万条数据，一次删除5万（batch_count）条；
# 开启20个goroutine并发删除5万数据

# 每次串行删除的数量,
per_total_count = 100000
# 每批删除命令一次删除条数
batch_count= 5000
# 并发协程数
go_count=20
# 每删除一批数据，间隔多少秒删除下一批
interval_time=2
[logger]
env="prod"
level="debug"
output = "log.txt"
[database]
judgery_dsn="mysql://root:123456@tcp(localhost:3306)/judgery?parseTime=True"
data_report_enable=true
```
