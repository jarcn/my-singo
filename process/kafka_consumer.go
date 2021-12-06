package process

import (
	"encoding/json"
	"log"
	"my-singo/chain"
	"my-singo/handlers"
	"my-singo/req"

	"github.com/Shopify/sarama"
)

var processChain *chain.MsgHandlerChain

// 初始化责任链
func init() {
	chain := &chain.MsgHandlerChain{}
	chain.AddHandler(&handlers.InboxHandler{})
	chain.AddHandler(&handlers.PushHandler{})
	chain.AddHandler(&handlers.SnackBarHandler{})
	processChain = chain
}

// 启动kafka消费者
func StartConsumer(brokenAddr []string, topic string, partition int32) {
	consumer, err := sarama.NewConsumer(brokenAddr, nil)
	if err != nil {
		log.Printf("kafka consumer init. Error: %s", err.Error())
		return
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			log.Printf("kafka consumer init. Error: %s", err.Error())
			return
		}
	}()
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Printf("kafka consumer init. Error: %s", err.Error())
		return
	}
	defer func() {
		if err = partitionConsumer.Close(); err != nil {
			log.Printf("kafka consumer init. Error: %s", err.Error())
			return
		}
	}()
	for msg := range partitionConsumer.Messages() {
		req := &req.MqNoticeMsgRequest{}
		convertMsg(string(msg.Value), req)
		processChain.Handler(req)
	}
	log.Println("Starting a new kafka consumer")
}

func convertMsg(msg string, req *req.MqNoticeMsgRequest) {
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}
}
