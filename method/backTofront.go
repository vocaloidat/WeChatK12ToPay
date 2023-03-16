package method

import (
	"K12_P/sign"
	"encoding/json"
	"fmt"
)

// 这里主要用于 给前端返回数据使用。

// Offlinefacemch_organizations_users_backTofront 给前端返回用户信息
func Offlinefacemch_organizations_users_backTofront(user_str []byte) string {

	type Student_info struct {
		Class_name string `json:"class_name"`
	}

	type Staff_info struct {
		Occupation string `json:"occupation"`
	}

	type UserInfo struct {
		User_id         string        `json:"user_id"`
		Out_user_id     string        `json:"out_user_id"`
		Organization_id string        `json:"organization_id"`
		User_name       string        `json:"user_name"`
		User_type       string        `json:"user_type"`
		Student_info    *Student_info `json:"student_info"`
		Staff_info      *Staff_info   `json:"staff_info"`
		Status          string        `json:"status"`
		Contract_state  string        `json:"contract_state"`
		Face_image_ok   bool          `json:"face_image_ok"`
		Contract_id     string        `json:"contract_id"`
	}
	c1 := &UserInfo{}
	err := json.Unmarshal(user_str, c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return ""
	}
	// 解密用户姓名
	c1.User_name, _ = sign.PayWeChat.RsaDecryptByPrivateKey(c1.User_name)
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed")
		return ""
	}
	return string(data)
}

func Offlinefacemch_organizations_users_PATCH_backTofront(b []byte) string {
	type Student_info struct {
		Class_name string `json:"class_name"`
	}

	type Staff_info struct {
		Occupation string `json:"occupation"`
	}
	type UserInfo struct {
		user_name    string        `json:"user_name"`
		user_type    string        `json:"user_type"`
		student_info *Student_info `json:"student_info"`
		staff_info   *Staff_info   `json:"staff_info"`
		status       string        `json:"status"`
		phone        string        `json:"phone"`
	}
	u1 := &UserInfo{}
	err := json.Unmarshal(b, u1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return ""
	}
	// 加密姓名和手机号
	u1.user_name, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.user_name)
	u1.phone, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.phone)
	// 将结构体解析为字符串传回去
	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("json marshal failed")
		return ""
	}
	return string(data)
}

func OfflinefaceContractsPresign_backTofront(b []byte) []byte {
	type Identification struct {
		Identification_type   string `json:"identification_type"`
		Identification_number string `json:"identification_number"`
	}
	type Facepay_user struct {
		Out_user_id         string          `json:"out_user_id"`
		Identification_name string          `json:"identification_name"`
		Organization_id     string          `json:"organization_id"`
		Identification      *Identification `json:"identification"`
		Phone               string          `json:"phone"`
	}
	type Limit_bank_card struct {
		Bank_card_number    string          `json:"bank_card_number"`
		Identification_name string          `json:"identification_name"`
		Identification      *Identification `json:"identification"`
		Valid_thru          string          `json:"valid_thru"`
		Bank_type           string          `json:"bank_type"`
		Phone               string          `json:"phone"`
	}
	type UserInfo struct {
		Business_name   string           `json:"business_name"`
		Facepay_user    *Facepay_user    `json:"facepay_user"`
		Limit_bank_card *Limit_bank_card `json:"limit_bank_card"`
		Contract_mode   string           `json:"contract_mode"`
	}
	u1 := &UserInfo{}
	err := json.Unmarshal(b, u1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return nil
	}
	// 加密identification_name,phone
	u1.Facepay_user.Identification_name, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Facepay_user.Identification_name)
	u1.Facepay_user.Phone, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Facepay_user.Phone)
	u1.Facepay_user.Identification.Identification_number, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Facepay_user.Identification.Identification_number)
	u1.Limit_bank_card.Bank_card_number, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Limit_bank_card.Bank_card_number)
	u1.Limit_bank_card.Identification_name, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Limit_bank_card.Identification_name)
	u1.Limit_bank_card.Identification.Identification_number, _ = sign.PayWeChat.RsaEncryptByPublicKey(u1.Limit_bank_card.Identification.Identification_number)

	// 将结构体解析为字符串传回去
	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("json marshal failed")
		return nil
	}
	return data
}
