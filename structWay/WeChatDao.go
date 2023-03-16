package structWay

type WechatMerchant struct {
	Name string `yaml:"name"`
	// 商户API私钥路径
	PrivateKeyFilepath string `yaml:"private_key_filepath"`
	// 商户证书路径
	MerchantCertificateFilepath string `yaml:"merchant_certificate_filepath"`
	// 商户号
	Mchid string `yaml:"mchid"`
	// 证书序列号
	SerialNo string `yaml:"serial_no"`
}
