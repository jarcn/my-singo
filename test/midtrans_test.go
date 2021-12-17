package test

import (
	"fmt"
	"testing"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var mid_client_key = "SB-Mid-client-p0tMKV9gxVoydUlz"
var mid_server_key = "SB-Mid-server-qgGG6qKkEwylnvYhE44WddNQ"
var c coreapi.Client

func init() {
	c.New(mid_server_key, midtrans.Sandbox)
}

func TestGetToken(t *testing.T) {
	rep, err := c.CardToken("4811 1111 1111 1114", 12, 24, "123", mid_client_key)
	if err != nil {
		fmt.Println(err.GetMessage())
	}
	fmt.Println(rep)
}
