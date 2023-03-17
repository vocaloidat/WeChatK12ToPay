# WeChat K12 education version payment merchant backend （微信K12教育版本支付商户后台）

（微信K12教育版本支付商户后台）

介于目前GitHub上面还没有一个好用的go语言开发的微信K12教育商户后台

我将于我的开发开源

后续如果有时间会写好各个接口的使用，如果你是开发人员看一下源代码应该没问题。

想要一起维护和使用可以联系我。

外gmail：vocaloid.at@gmail.com

### 一。使用方法

找到`WeChat_k12.yaml`文件，申请所需要的K12商户资料，填入yaml。

编译golang为二进制文件，放入服务器直接运行。

**yaml为什么没有apiv3apiV3密钥的存放？**

答：

1. apiV3密钥是非常重要的内容，我将apiV3密钥放入了go语言内部，程序启动后访问`/get_certificates`就能自动得到。
2. apiV3密钥需要实时更换，直接用人为修改不现实，我直接放入接口调用，设置定时任务就能完成。

### 二。系统特点

系统实现了直接的请求访问，开发者和使用者都能简单上手修改，建议大家一起维护。

所有的api接口均实现了和微信K12商户后台所需的对应，所有的加密 解密 签名都写好了。

#### 注意系统没有写**验证应答或者回调的签名。**主要事情比较多，有时间再加入验证。

[查看官方文档加入签名认证即可]: https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/qian-ming-yan-zheng



|                          本系统api                           |            介绍            |                       对应微信商户后台                       |
| :----------------------------------------------------------: | :------------------------: | :----------------------------------------------------------: |
|                      /get_certificates                       |      获取平台证书列表      |        https://api.mch.weixin.qq.com/v3/certificates         |
|                /offlinefacemch/organizations                 | 根据机构ID查询机构信息接口 |      /v3/offlinefacemch/organizations?organization_id=       |
|                    /offlinefacemch/tokens                    |      获取授权凭证接口      |                  /v3/offlinefacemch/tokens                   |
| /offlinefacemch/organizations/:organization_id/users/out-user-id/:out_user_id |    刷脸用户信息查询接口    | /v3/offlinefacemch/organizations/{organization_id}/users/out-user-id/{out_user_id} |
| /offlinefacemch/organizations/:organization_id/users/out-user-id/:out_user_id |    刷脸用户信息修改接口    | /v3/offlinefacemch/organizations/{organization_id}/users/out-user-id/{out_user_id} |
|                    /offlineface/authinfo                     |      获取authinfo接口      |                   /v3/offlineface/authinfo                   |
| /offlinefacemch/organizations/:organization_id/users/user-id/:user_id/terminate-contract |  解除刷脸用户签约关系接口  | /v3/offlinefacemch/organizations/{organization_id}/users/user-id/{user_id}/terminate-contract |
|                /offlineface/contracts/presign                |         预签约接口         |              /v3/offlineface/contracts/presign               |
|                  /offlineface/transactions                   |        申请扣款接口        |                 /v3/offlineface/transactions                 |
|             /offlineface/contracts/:contract_id              |        签约查询接口        |      /v3/offlineface/contracts/{contract_id}?appid=XXXX      |
|         /offlineface/face-collections/:collection_id         |      查询重采请求接口      |       /v3/offlineface/face-collections/{collection_id}       |
|                /offlineface/face-collections                 |    查询重采用户列表接口    | /v3/offlineface/face-collections?organization_id={organization_id} |
|     /offlineface/transactions/out-trade-no/:out_trade_no     |          查单接口          | /v3/offlineface/transactions/out-trade-no/{out_trade_no}?sp_mchid={sp_mchid}&sub_mchid={sub_mchid}&business_product_id={business_product_id} |
|                  /offlineface/repayment-url                  |      获取还款链接接口      |                /v3/offlineface/repayment-url                 |

