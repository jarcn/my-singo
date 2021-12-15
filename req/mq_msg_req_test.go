package req

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/Shopify/sarama"
)

func TestKafkaP(t *testing.T) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"106.55.104.223:19092", "106.55.104.223:19093", "106.55.104.223:19094"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic:     "qiyee-job-msg-push",
		Partition: 0,
		Key:       sarama.StringEncoder("notice"),
	}

	msgBean := MqNoticeMsgRequest{
		Title:        "notice_msg",
		Content:      "test notice msg",
		Summary:      "msg test 1234",
		ToUserId:     123456789,
		ToUserName:   "chenjia",
		MsgType:      28001,
		JobId:        1111,
		JobOfId:      1234567890,
		JobName:      "开发测试",
		BusinessName: "数智方",
	}
	msgJson, err := json.Marshal(msgBean)
	if err != nil {
		log.Println("[处理失败] JSON解析失败：", msgJson)
	}
	//将字符串转换为字节数组
	msg.Value = sarama.ByteEncoder(msgJson)
	//fmt.Println(value)
	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message Fail")
	}
	fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
}

func TestHex(t *testing.T) {

	s := "a\r\nb"
	t.Log("----->>>>", s)
	shex := fmt.Sprintf("%x", s)
	t.Log("--->>>", shex)
	nStr := strings.ReplaceAll(shex, "0d0a", "")
	t.Log(nStr)
	v, _ := strconv.Atoi(nStr)
	t.Log("====>>>>", v)

}
