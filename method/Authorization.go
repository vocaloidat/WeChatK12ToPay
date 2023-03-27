package method

import (
	"K12_P/constant"
	"K12_P/sign"
	"K12_P/structWay"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
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

// verifySignature 这个函数接受三个参数：签名字符串，应答签名字符串和平台证书。它返回一个布尔值和一个错误。如果签名验证成功，则返回true，否则返回false，并返回一个错误。
func verifySignature(signatureString string, responseSignature string, platformCertificate *x509.Certificate) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureString)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature string: %v", err)
	}

	responseSignatureBytes, err := base64.StdEncoding.DecodeString(responseSignature)
	if err != nil {
		return false, fmt.Errorf("failed to decode response signature string: %v", err)
	}

	hashed := sha256.Sum256(signatureBytes)

	err = rsa.VerifyPKCS1v15(platformCertificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], responseSignatureBytes)
	if err != nil {
		return false, fmt.Errorf("failed to verify signature: %v", err)
	}

	return true, nil
}
