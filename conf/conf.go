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

// 从本地读取环境变量
func init() {
	godotenv.Load()
}

// Init 初始化配置项
func Init() {
	initLogger()
	initDb()
	initKafkaConsumer()
	initi18n()
}

// 设置日志级别
func initLogger() {
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
}

// 读取翻译文件
func initi18n() {
	if err := InitLocales("zh", "conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("汉语翻译文件加载失败", err)
	}
	if err := InitLocales("en", "conf/locales/us-en.yaml"); err != nil {
		util.Log().Panic("英语翻译文件加载失败", err)
	}
}

// 初始化kafka(分区指定为0号有问题)
func initKafkaConsumer() {
	brokers := strings.Split(os.Getenv("KAFKA_ADDR"), ",")
	go process.StartConsumer(brokers, os.Getenv("KAFKA_TOPIC"), 0)
}

// 连接数据库
// 连接redis
func initDb() {
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
