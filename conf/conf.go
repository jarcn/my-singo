package conf

import (
	"my-singo/cache"
	"my-singo/model"
	"my-singo/process"
	"my-singo/util"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}
	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	// 连接redis
	cache.Redis()
	// 初始化kafka(分区指定为0号有问题)
	brokers := strings.Split(os.Getenv("KAFKA_ADDR"), ",")
	go process.StartConsumer(brokers, os.Getenv("KAFKA_TOPIC"), 0)
}
