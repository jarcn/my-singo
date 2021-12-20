package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestDisbursement(t *testing.T) {

	url := "https://bigflip.id/big_sandbox_api/v2/disbursement"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("account_number", "5465327020")
	_ = writer.WriteField("bank_code", "bni")
	_ = writer.WriteField("amount", "100000")
	_ = writer.WriteField("remark", "chenjia test")
	_ = writer.WriteField("recipient_city", "391")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("idempotency-key", "$2y$13$sphRq9pUy2eaBKT9WV98NeJYSiQViNeEPE2vBQ9RR.b3dAdpHvPge")
	req.Header.Add("Authorization", "Basic SkRKNUpERXpKRGN3UTFKMk9WVm1PRnBuYjNOcmFXMWFiMlF2VVdVeGFFNVNXSEpoVEUxYU5tRlNZbFI2TlZWblpFRTRUR1paV1drNU4wazI6Og==")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
