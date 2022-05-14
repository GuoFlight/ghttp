package main

import (
	"errors"
	"fmt"
	"github.com/GuoFlight/ghttp"
	"github.com/avast/retry-go"
)

func main() {
	client := ghttp.Client{
		Method: ghttp.MethodGet,
		Url:    "https://www.baidu.com",
		Retry: ghttp.Retry{
			JudgeRetryFunc: func(res *ghttp.Res) error {
				return errors.New("error")		//模拟错误
			},
			Options: []retry.Option{
				retry.Attempts(2),			//最多重试2次
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
