package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var SetData = make(map[string]interface{})
var ListData = []string{}
var r = rand.New(rand.NewSource(time.Now().Unix()))

// Redis 在中间件中初始化redis链接
func init() {
	db, _ := strconv.ParseUint("0", 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       "39.105.153.230:16379",
		Password:   "",
		DB:         int(db),
		MaxRetries: 1,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("连接Redis不成功", err)
	}
	RedisClient = client
}

func init() {
	for i := 0; i < 50000; i++ {
		s := strconv.Itoa(i)
		SetData[s] = s
		ListData = append(ListData, s)
	}
}

func TestInitData(t *testing.T) {
	fmt.Println("listdata size=", len(ListData))
	fmt.Println("setdata size=", len(SetData))
}

func TestInsertHSet(t *testing.T) {
	RedisClient.HMSet("id_mset", SetData)
}

func TestInsertList(t *testing.T) {
	RedisClient.LPush("id_list", ListData)
}

func TestLSpentTime(t *testing.T) {
	f := timeSpent(eInList)
	f(strconv.Itoa(60000))
}

func TestMSpentTime(t *testing.T) {
	f := timeSpent(eInSet)
	f(strconv.Itoa(60000))
}

type IntConv func(op string) bool

func timeSpent(inner IntConv) IntConv {
	return func(s string) bool {
		start := time.Now()
		ret := inner(s)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func eInSet(s string) bool {
	return RedisClient.HGet("id_mset", strconv.Itoa(r.Intn(50000))).Val() != ""
}

func eInList(e string) bool {
	list := RedisClient.LRange("id_list", 0, -1).Val()
	for _, v := range list {
		return e == v
	}
	return false
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisClient.LRange("id_list", 0, -1).Val()
	}
}
