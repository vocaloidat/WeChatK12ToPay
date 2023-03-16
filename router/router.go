package router

import (
	"K12_P/method"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

// 令牌桶
var limiter = rate.NewLimiter(200, 500)

// 限流中间件
func limitMiddleware(c *gin.Context) {
	res := limiter.Reserve()
	delay := res.Delay()
	if delay == 0 {
		c.Next()
	} else {
		time.Sleep(delay) // 等待令牌
		c.Next()
	}
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 允许所有域名跨域访问
	r.Use(cors.Default())
	r.Use(limitMiddleware)
	r.GET("/get_certificates", method.Get_certificates)                                                                                        // 获取平台证书列表
	r.GET("/offlinefacemch/organizations", method.Offlinefacemch_organizations)                                                                // 根据机构ID查询机构信息接口
	r.POST("/offlinefacemch/tokens", method.Offlinefacemch_tokens)                                                                             // 获取授权凭证接口
	r.GET("/offlinefacemch/organizations/:organization_id/users/out-user-id/:out_user_id", method.Offlinefacemch_organizations_users)          // 刷脸用户信息查询接口
	r.PATCH("/offlinefacemch/organizations/:organization_id/users/out-user-id/:out_user_id", method.Offlinefacemch_organizations_users_PATCH)  // 刷脸用户信息修改接口
	r.POST("/offlineface/authinfo", method.Getauthinfo)                                                                                        // 获取authinfo接口
	r.POST("/offlinefacemch/organizations/:organization_id/users/user-id/:user_id/terminate-contract", method.OfflinefacemchTerminateContract) // 解除刷脸用户签约关系接口
	r.POST("/offlineface/contracts/presign", method.OfflinefaceContractsPresign)                                                               // 预签约接口
	r.POST("/offlineface/transactions", method.OfflinefaceTransactions)                                                                        // 申请扣款接口
	r.GET("/offlineface/contracts/:contract_id", method.GetOfflinefaceContracts)                                                               // 签约查询接口
	r.GET("/offlineface/face-collections/:collection_id", method.OfflinefaceFaceCollections)                                                   // 查询重采请求接口
	r.GET("/offlineface/face-collections", method.OfflinefaceFaceCollections_organization_id)                                                  // 查询重采用户列表接口
	r.GET("/offlineface/transactions/out-trade-no/:out_trade_no", method.OfflinefaceTransactionsOutTradeNo)                                    // 查单接口
	r.POST("/offlineface/repayment-url", method.OfflinefaceRepaymentUrl)                                                                       // 获取还款链接接口
	// 启动HTTP服务，默认在0.0.0.0:10888启动服务
	//r.Run(":10888")
	return r
}
