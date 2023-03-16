package method

import (
	"K12_P/constant"
	"K12_P/sign"
	"K12_P/structWay"
	"math/rand"
	"time"
)

// GetAuthorization 一生成访问签名 返回字符串
func GetAuthorization(method string, constantApi string) string {
	// 获取发起请求时的系统当前时间戳
	timestamp := string(time.Now().Unix()) // 获取当前时间戳

	// 生成一个请求随机串
	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 使用r来生成随机数
	randNum := string(r.Intn(10001)) // 生成0到99之间的一个整数

	// 获取请求中的请求报文主体（request body）
	/**
	请求方法为GET时，报文主体为空。
	当请求方法为POST或PUT时，请使用真实发送的JSON报文。
	图片上传API，请使用meta对应的JSON报文。
	*/

	//按照前述规则，构造的请求签名串为：
	signString := method + `\n` + constantApi + `\n` + timestamp + `\n` + randNum + `\n`
	// 获取yaml文件
	wm := structWay.GetWeChatDao()
	// 计算sha256 获取签名结果
	signature := sign.PrivateKeyToSHA256ToString(wm.PrivateKeyFilepath, signString)

	// 使用base64 得到Base64编码签名值
	signatureValue := sign.StringToBase64(signature)

	// 组装Authorization内容。
	Authorization := constant.AlgorithmAEADAES256GCM + " " +
		"mchid=" + sign.PayWeChat.MchId + " " + "serial_no=" + sign.PayWeChat.ApiSerialNo + " " + "nonce_str=" + randNum + " " + "timestamp=" + timestamp + " " +
		"signature=" + signatureValue

	return Authorization
}
