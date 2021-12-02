package chain

import (
	"my-singo/handlers"
	"my-singo/req"
	"testing"
	"time"
)

func TestChain(t *testing.T) {
	chain := &MsgHandlerChain{}
	chain.AddHandler(&handlers.InboxHandler{})
	chain.AddHandler(&handlers.PushHandler{})
	chain.AddHandler(&handlers.SnackBarHandler{})
	chain.Handler(&req.MqNoticeMsgRequest{})
	time.Sleep(time.Second * 1)
}
