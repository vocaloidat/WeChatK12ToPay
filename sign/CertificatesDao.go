package sign

import (
	"encoding/json"
	"fmt"
)

type Certificates_Data struct {
	Data []*Certificates `json:"data"`
}

// Certificates 用来接收微信支付的平台证书
type Certificates struct {
	Serial_no           string               `json:"serial_no"`
	effective_time      string               `json:"effective_time"`
	expire_time         string               `json:"expire_time"`
	Encrypt_certificate *Encrypt_certificate `json:"encrypt_certificate"`
}

type Encrypt_certificate struct {
	Algorithm       string `json:"algorithm"`
	Nonce           string `json:"nonce"`
	Associated_data string `json:"associated_data"`
	Ciphertext      string `json:"ciphertext"`
}

// Init_Certificates_Data 获取的平台证书列表转为 结构体使用
func Init_Certificates_Data(str string) *Certificates_Data {
	c1 := &Certificates_Data{}
	err := json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return nil
	}
	return c1
}
