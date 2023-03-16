package sign

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
)

// 进行SHA256 with RSA签名
/**
privatePath 私钥文件所在目录
str 所需要SHA256 with RSA签名的字符串
*/
func PrivateKeyToSHA256ToString(privatePath string, str string) []byte {
	// 读取私钥文件
	pemData, err := os.ReadFile(privatePath)
	if err != nil {
		panic(err)
	}

	// 解析PEM格式的私钥
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		panic("invalid pem file")
	}

	// 转换为RSA私钥对象
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	rsaKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		panic("not a rsa key")
	}

	// 对字符串进行SHA256 with RSA签名
	data := []byte(str)
	signature, err := rsa.SignPKCS1v15(nil, rsaKey, crypto.SHA256, data)
	if err != nil {
		panic(err)
	}
	return signature
}

// StringToBase64 对字节进行Base64编码得到签名值
func StringToBase64(str []byte) string {
	// 对字符串进行Base64编码得到签名值
	signatureValue := base64.StdEncoding.EncodeToString(str)
	return signatureValue
}

// PrivatePathToPrivateKey 传入pem 直接返回私钥对象
func PrivatePathToPrivateKey(privatePath string) *rsa.PrivateKey {
	// 读取私钥文件
	pemData, err := os.ReadFile(privatePath)
	if err != nil {
		panic(err)
	}

	// 解析PEM格式的私钥
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		panic("invalid pem file")
	}

	// 转换为RSA私钥对象
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	rsaKey, ok := privKey.(*rsa.PrivateKey)
	if !ok {
		panic("not a rsa key")
	}

	return rsaKey
}

// 传入pem路径 直接转为x509Certificate 对象
func PrivatePathTox509Certificate(privatePath string) *x509.Certificate {
	// 读取证书文件
	certFile, err := os.Open(privatePath)
	if err != nil {
		log.Fatal(err)
	}
	defer certFile.Close()

	// 解码PEM格式的证书
	pemBytes, err := io.ReadAll(certFile)
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		log.Fatal("invalid PEM file")
	}

	// 转化为x509.Certificate类型
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	// 打印证书信息
	fmt.Println(cert.Subject.CommonName)
	return cert
}

// StringPemToX509 将可用的字符串转为x509.Certificate类型
func StringPemToX509(str_pem string) *x509.Certificate {
	// 假设str_pem是一个有效的PEM编码的证书字符串
	certBytes := []byte(str_pem)
	block, _ := pem.Decode(certBytes) // 解码PEM块
	if block == nil {
		// 处理错误
	}
	cert, err := x509.ParseCertificate(block.Bytes) // 解析证书
	if err != nil {
		// 处理错误
	}
	return cert
}

// 加密报文 使用apiv3密钥解密出来平台证书
func CertificatesDataToApiv3(c1 *Certificates_Data, apiv3key string) string {
	PlatformCertificate_string := string(decrypt(c1.Data[0].Encrypt_certificate.Nonce, c1.Data[0].Encrypt_certificate.Ciphertext, c1.Data[0].Encrypt_certificate.Associated_data, apiv3key))
	return PlatformCertificate_string
}

func decrypt(nonce, ciphertext, associated_data string, apiv3key string) []byte {
	key := apiv3key

	key_bytes := []byte(key)
	nonce_bytes := []byte(nonce)
	ad_bytes := []byte(associated_data)
	data, _ := base64.StdEncoding.DecodeString(ciphertext)

	block, _ := aes.NewCipher(key_bytes)
	gcm, _ := cipher.NewGCM(block)
	plaintext, _ := gcm.Open(nil, nonce_bytes, data, ad_bytes)
	return plaintext
}
