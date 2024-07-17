package restyclient

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestyClient() *resty.Client {
	return resty.New().
		SetTimeout(5 * time.Second).
		SetRetryCount(3)
}

func Post[T any](url string, data any, res *Response[T]) error {
	client := NewRestyClient()
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %v", err)
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "9OQQMKFIXMSDSRZYTNC5HKVMRNG0MTXL").
		SetBody(jsonStr).
		SetResult(res).
		Post(url)

	if err != nil {
		return fmt.Errorf("请求发送失败: %v", err)
	}

	// 处理接口返回错误
	if resp.IsError() {
		return fmt.Errorf("请求返回失败，http状态: %v", resp.Status())
	}

	if res.Code != 0 {
		return fmt.Errorf("请求返回失败，错误码: %v, 错误信息: %v", res.Code, res.Msg)
	}
	return nil
}

type Response[T any] struct {
	Code int64  `json:"code"`           // 状态嘛
	Data T      `json:"data,omitempty"` // 响应内容
	Msg  string `json:"msg"`            // 响应说明
}
