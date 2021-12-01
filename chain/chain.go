package chain

import (
	"my-singo/req"
)

// 责任链接口
type Handler interface {
	Process(msg *req.MqNoticeMsgRequest)
	NeedHandler(msg *req.MqNoticeMsgRequest) bool
}

//消息处理链
type MsgHandlerChain struct {
	handlers []Handler
}

//链条中添加处理逻辑
func (msgChain *MsgHandlerChain) AddHandler(handler Handler) {
	msgChain.handlers = append(msgChain.handlers, handler)
}

// 责任链处理逻辑
func (msgChain *MsgHandlerChain) Handler(msg *req.MqNoticeMsgRequest) {
	for _, handler := range msgChain.handlers {
		if handler.NeedHandler(msg) {
			go handler.Process(msg)
		}
	}
}
