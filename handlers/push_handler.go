package handlers

import (
	"encoding/json"
	"log"
	"my-singo/req"
)

type PushHandler struct{}

func (p *PushHandler) Process(msg *req.MqNoticeMsgRequest) {
	s, _ := json.Marshal(msg)
	log.Println("PushHandler:", string(s))
}

func (p *PushHandler) NeedHandler(msg *req.MqNoticeMsgRequest) bool {
	return true
}
