package method

import (
	"K12_P/sign"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func httpClient_Post(url string, httpMethod string, Authorization string, post_json []byte) ([]byte, int) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(post_json))
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	// 设置请求头部
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Wechatpay-Serial", sign.PayWeChat.PlatformSerialNo)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	// 关闭响应体
	defer resp.Body.Close()
	//resp.StatusCode 相应状态 200 404 401
	// 读取响应体数据
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	return body, resp.StatusCode
}
