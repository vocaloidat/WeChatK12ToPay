package sign

import (
	"K12_P/constant"
	"K12_P/structWay"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
)

type PayClient struct {
	MchId               string            // 商户号
	ApiV3Key            string            // apiV3密钥
	ApiSerialNo         string            // API证书序列号
	ApiPrivateKey       *rsa.PrivateKey   // API私钥
	ApiCertificate      *x509.Certificate // API证书
	PlatformSerialNo    string            // 平台证书序列号
	PlatformCertificate *x509.Certificate // 平台证书
	HttpClient          *http.Client      // 暂时没用
}

// PayWeChat 直接初始化一个PayWeChat，让他全局可用。
var PayWeChat *PayClient

// InitPayClient 获取平台证书后 将实例化当前访问。
func InitPayClient(wm *structWay.WechatMerchant, str_json string) {
	PayWeChat.MchId = wm.Mchid
	PayWeChat.ApiV3Key = constant.ApiV3Key
	PayWeChat.ApiSerialNo = wm.SerialNo
	PayWeChat.ApiPrivateKey = PrivatePathToPrivateKey(wm.PrivateKeyFilepath)
	PayWeChat.ApiCertificate = PrivatePathTox509Certificate(wm.PrivateKeyFilepath)
	c1 := Init_Certificates_Data(str_json)
	PayWeChat.PlatformSerialNo = c1.Data[0].Serial_no
	// 通过apiv3密钥 解密出来真正的平台证书。
	PlatformCertificate_string := CertificatesDataToApiv3(c1, PayWeChat.ApiV3Key)
	PayWeChat.PlatformCertificate = StringPemToX509(PlatformCertificate_string)
}

// RsaDecryptByPrivateKey 使用商户私钥RSA解密
func (c *PayClient) RsaDecryptByPrivateKey(ciphertext string) (string, error) {
	cipherData, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, c.ApiPrivateKey, cipherData, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// RsaEncryptByPublicKey 使用平台公钥RSA加密
func (c *PayClient) RsaEncryptByPublicKey(plaintext string) (string, error) {
	if c.PlatformSerialNo == "" || c.PlatformCertificate == nil {
		return "", fmt.Errorf("请先初始化平台证书")
	}
	secretMessage := []byte(plaintext)
	rng := rand.Reader

	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, c.PlatformCertificate.PublicKey.(*rsa.PublicKey), secretMessage, nil)
	if err != nil {
		return "", err
	}

	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	return ciphertext, nil
}
