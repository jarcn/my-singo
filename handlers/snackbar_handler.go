package handlers

import (
	"encoding/json"
	"log"
	"my-singo/req"
)

type SnackBarHandler struct{}

//处理snackbar消息
func (sbh *SnackBarHandler) Process(msg *req.MqNoticeMsgRequest) {
	s, _ := json.Marshal(msg)
	log.Println("SnackBarHandler:", string(s))
}

//判断是否需要处理 snackbar 消息
func (sbh *SnackBarHandler) NeedHandler(msg *req.MqNoticeMsgRequest) bool {
	return true
}
