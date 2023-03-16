package method

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Offlinefacemch_tokens_post 访问微信的获取授权凭证接口
func httpClient_Offlinefacemch_tokens_post(url string, Authorization string, post_json []byte) string {
	// 创建一个http.Request对象
	body, _ := httpClient_Post(url, http.MethodPost, Authorization, post_json)
	var data struct {
		Token string `json:"token"`
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return ""
	}
	return data.Token
}

// 访问 刷脸用户信息查询接口
func httpClient_Offlinefacemch_organizations_users_GET(url string, Authorization string) string {
	body, _ := httpClient_Post(url, http.MethodGet, Authorization, nil)
	// 以下可以弃用，获取用户信息后，解密用户姓名 直接返回给前端
	return Offlinefacemch_organizations_users_backTofront(body)
}

// 访问微信的刷脸用户信息修改接口
func httpClient_Offlinefacemch_organizations_users_PATCH(url string, b []byte, Authorization string) string {
	// 1.加密上传b数据
	b = []byte(Offlinefacemch_organizations_users_PATCH_backTofront(b))
	body, _ := httpClient_Post(url, http.MethodPatch, Authorization, b)
	if body == nil {
		return ""
	}
	return "修改成功"
}

// 预签约接口
func httpClient_OfflinefaceContractsPresign(url string, Authorization string, b []byte) []byte {
	// 加密用户数据然后返回这里继续
	b = OfflinefaceContractsPresign_backTofront(b)
	body, _ := httpClient_Post(url, http.MethodPost, Authorization, b)
	return body
}

func httpClient_OfflinefaceTransactions(url string, Authorization string, b []byte) []byte {
	body, _ := httpClient_Post(url, http.MethodPost, Authorization, b)
	return body
}
