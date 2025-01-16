package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/GuoFlight/ghttp"
	"github.com/avast/retry-go"
	"net/http"
)

func main() {
	client := ghttp.Client{
		Method: ghttp.MethodGet,
		Url:    "https://www.baidu.com",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略证书验证
			},
		},
		Retry: ghttp.Retry{
			JudgeRetryFunc: func(res *ghttp.Res) error {
				return errors.New("error") // 模拟错误
			},
			Options: []retry.Option{
				retry.Attempts(2), // 最多重试2次
			},
		},
	}
	res, err := client.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Detail.StatusCode)
}
