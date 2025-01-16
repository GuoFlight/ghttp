package ghttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_Do(t *testing.T) {
	client := Client{
		Method: http.MethodGet,
		Url:    "https://www.baidu.com",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略证书验证
			},
		},
	}
	resHttp, err := client.Do()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(string(resHttp.Body))
}
