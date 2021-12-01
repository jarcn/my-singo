package handlers

import (
	"encoding/json"
	"log"
	"my-singo/req"
)

type InboxHandler struct{}

func (i *InboxHandler) Process(msg *req.MqNoticeMsgRequest) {
	s, _ := json.Marshal(msg)
	log.Println("inbox handler reciver msg:", string(s))
}

func (i *InboxHandler) NeedHandler(msg *req.MqNoticeMsgRequest) bool {
	return true
}
