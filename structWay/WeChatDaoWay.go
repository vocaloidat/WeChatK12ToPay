package structWay

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func GetWeChatDao() WechatMerchant {
	// 读取yaml文件内容到字节数组中
	data, err := os.ReadFile("./WeChat_k12.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个WechatMerchant类型的变量
	var wm WechatMerchant

	// 使用yaml.Unmarshal函数解析字节数组到变量中
	err = yaml.Unmarshal(data, &wm)
	if err != nil {
		log.Fatal(err)
	}

	// 打印变量中的参数值
	fmt.Println("商户API私钥路径:", wm.PrivateKeyFilepath)
	fmt.Println("商户证书路径:", wm.MerchantCertificateFilepath)

	return wm
}
