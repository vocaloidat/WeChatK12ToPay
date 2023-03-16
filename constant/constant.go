package constant

/**
用来存储微信商户支付的api
*/

// ApiDomain 微信商户的域名
const ApiDomain = "https://api.mch.weixin.qq.com"

// ApiV3Key apiV3密钥
const ApiV3Key = "apiV3密钥"

// AlgorithmAEADAES256GCM api签名认证类型
const AlgorithmAEADAES256GCM = "WECHATPAY2-SHA256-RSA2048"

// K12离线团餐类接口
const (
	APIQueryOrganizationInfoById                  = "/v3/offlinefacemch/organizations?organization_id="                                             //根据机构ID查询机构信息
	APIGetplatformcertificates                    = "/v3/certificates"                                                                              //获取平台证书
	APIOfflinefacemchTokens                       = "/v3/offlinefacemch/tokens"                                                                     //获取授权凭证接口
	APIOfflinefacemchOrganizationsUser            = "/v3/offlinefacemch/organizations/{organization_id}/users/out-user-id/{out_user_id}"            //刷脸用户信息查询接口 (查询和修改)
	APIGetauthinfo                                = "/v3/offlineface/authinfo"                                                                      //获取authinfo接口 建议authinfo每1小时内获取一次。
	APIOfflinefacemchTerminateContract            = "/v3/offlinefacemch/organizations/{organization_id}/users/user-id/{user_id}/terminate-contract" //解除刷脸用户签约关系接口
	APIOfflinefaceContractsPresign                = "/v3/offlineface/contracts/presign"                                                             //预签约接口
	APIOfflinefaceTransactions                    = "/v3/offlineface/transactions"                                                                  //申请扣款接口
	APIGetOfflinefaceContracts                    = "/v3/offlineface/contracts/{contract_id}?appid=XXXX"                                            //签约查询接口
	APIOfflinefaceFaceCollections                 = "/v3/offlineface/face-collections/{collection_id}"                                              //查询重采请求接口
	APIOfflinefaceFaceCollections_organization_id = "/v3/offlineface/face-collections?organization_id={organization_id}"                            //查询重采用户列表接口
	//查单接口
	APIOfflinefaceTransactionsOutTradeNo = "/v3/offlineface/transactions/out-trade-no/{out_trade_no}?sp_mchid={sp_mchid}&sub_mchid={sub_mchid}&business_product_id={business_product_id}"
	APIOfflinefaceRepaymentUrl           = "/v3/offlineface/repayment-url" //获取还款链接接口
)
