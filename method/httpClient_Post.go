package method

import (
	"K12_P/constant"
	"K12_P/sign"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// isCheck 代表是否要进行验证
func httpClient_Post(url string, httpMethod string, Authorization string, post_json []byte, isCheck bool) ([]byte, int) {
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
	// 判断是否需要验证微信会调过来的签名
	if isCheck {
		// 需要验证
		Wechatpay_Serial := resp.Header.Get("Wechatpay-Serial")
		Wechatpay_Signature := resp.Header.Get("Wechatpay-Signature")
		Wechatpay_Timestamp := resp.Header.Get("Wechatpay-Timestamp")
		Wechatpay_Nonce := resp.Header.Get("Wechatpay-Nonce")
		body_str := string(body)
		// 生成签名串
		signature_string := Wechatpay_Timestamp + `\n` + Wechatpay_Nonce + `\n` + body_str + `\n`
		// 对Wechatpay_Signature字段值进行Base64解码
		decodedSignature, err := base64.StdEncoding.DecodeString(Wechatpay_Signature)
		if err != nil {
			fmt.Println("解码失败", err)
			return nil, 404
		}
		// 得到应答签名
		responseSignature := string(decodedSignature)
		// 得到签名证书
		c1 := sign.Init_Certificates_Data(body_str)
		PlatformSerialNo1 := c1.Data[0].Serial_no
		PlatformSerialNo2 := c1.Data[1].Serial_no
		if Wechatpay_Serial != PlatformSerialNo1 && Wechatpay_Serial != PlatformSerialNo2 {
			fmt.Println("解码失败", err)
			return nil, 404
		}
		// 通过apiv3密钥 解密出来真正的平台证书。
		PlatformCertificate_string := sign.CertificatesDataToApiv3(c1, constant.ApiV3Key)
		PlatformCertificate := sign.StringPemToX509(PlatformCertificate_string)
		isok, _ := verifySignature(signature_string, responseSignature, PlatformCertificate)
		if !isok {
			fmt.Println("解码失败", err)
			return nil, 404
		}
	} else {
		// 不需要验证
	}
	return body, resp.StatusCode
}
