package method

import (
	"K12_P/constant"
	"K12_P/sign"
	"K12_P/structWay"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Get_certificates 获取平台证书列表
func Get_certificates(c *gin.Context) {
	// 设置请求URL
	url := constant.ApiDomain + constant.APIGetplatformcertificates
	wm := structWay.GetWeChatDao()
	Authorization := GetAuthorization(http.MethodGet, constant.APIGetplatformcertificates)
	body, code := httpClient_Post(url, http.MethodGet, Authorization, nil, true)
	if code == 404 {
		c.JSON(200, gin.H{
			"message": "获取平台证书列表失败",
		})
		return
	}
	fmt.Println("response body:", string(body))
	// 获取证书列表（微信官方平台证书）后 ，需要使用APIv3密钥来解密出来 实例化PayWeChat *PayClient为了后续接口连接
	sign.InitPayClient(&wm, string(body))
	// 在这里进行验证
	c.JSON(200, gin.H{
		"message": "获取平台证书列表完成",
	})
}

func Offlinefacemch_organizations(c *gin.Context) {
	//username := c.Query("username")
	organization_id := c.Query("organization_id") // 机构ID
	url := constant.ApiDomain + constant.APIQueryOrganizationInfoById + organization_id
	// 组装Authorization内容。
	Authorization := GetAuthorization(http.MethodGet, constant.APIQueryOrganizationInfoById)
	body, _ := httpClient_Post(url, http.MethodGet, Authorization, nil, true)
	fmt.Println("response body:", string(body))
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func Offlinefacemch_tokens(c *gin.Context) {

	url := constant.ApiDomain + constant.APIOfflinefacemchTokens
	// json传递post访问数据
	// 注意：下面为了举例子方便，暂时忽略了错误处理
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 反序列化

	Authorization := GetAuthorization(http.MethodPost, constant.APIOfflinefacemchTokens)

	token := httpClient_Offlinefacemch_tokens_post(url, Authorization, b)

	c.JSON(http.StatusOK, gin.H{
		"isok":  "ok",
		"token": token,
	})
}

func Offlinefacemch_organizations_users(c *gin.Context) {
	organization_id := c.Param("organization_id")
	out_user_id := c.Param("out_user_id")
	url := constant.ApiDomain + constant.APIOfflinefacemchOrganizationsUser
	newURL := strings.ReplaceAll(url, "{organization_id}", organization_id)
	newURL = strings.ReplaceAll(newURL, "{out_user_id}", out_user_id)
	fmt.Println(newURL)
	Authorization := GetAuthorization(http.MethodGet, constant.APIOfflinefacemchOrganizationsUser)
	body := httpClient_Offlinefacemch_organizations_users_GET(newURL, Authorization)
	var data interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func Offlinefacemch_organizations_users_PATCH(c *gin.Context) {
	organization_id := c.Param("organization_id")
	out_user_id := c.Param("out_user_id")
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	url := constant.ApiDomain + constant.APIOfflinefacemchOrganizationsUser
	newURL := strings.ReplaceAll(url, "{organization_id}", organization_id)
	newURL = strings.ReplaceAll(newURL, "{out_user_id}", out_user_id)
	Authorization := GetAuthorization(http.MethodPatch, constant.APIOfflinefacemchOrganizationsUser)
	// 要记得加密上传
	body := httpClient_Offlinefacemch_organizations_users_PATCH(newURL, b, Authorization)

	var data interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func Getauthinfo(c *gin.Context) {
	b, _ := c.GetRawData()
	url := constant.ApiDomain + constant.APIGetauthinfo
	Authorization := GetAuthorization(http.MethodPost, constant.APIGetauthinfo)
	body, _ := httpClient_Post(url, http.MethodPost, Authorization, b, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefacemchTerminateContract(c *gin.Context) {
	organization_id := c.Param("organization_id")
	user_id := c.Param("user_id")

	url := constant.ApiDomain + constant.APIOfflinefacemchTerminateContract
	newURL := strings.ReplaceAll(url, "{organization_id}", organization_id)
	newURL = strings.ReplaceAll(newURL, "{user_id}", user_id)

	Authorization := GetAuthorization(http.MethodPost, constant.APIOfflinefacemchTerminateContract)
	_, code := httpClient_Post(newURL, http.MethodPost, Authorization, nil, true)
	if code != 204 {
		c.JSON(http.StatusOK, gin.H{
			"isok": "no",
			"data": "解除刷脸用户签约关系接口失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": "解除刷脸用户签约关系接口成功",
	})
}

func OfflinefaceContractsPresign(c *gin.Context) {
	b, _ := c.GetRawData()
	url := constant.ApiDomain + constant.APIOfflinefaceContractsPresign
	Authorization := GetAuthorization(http.MethodPost, constant.APIOfflinefaceContractsPresign)
	body := httpClient_OfflinefaceContractsPresign(url, Authorization, b)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefaceTransactions(c *gin.Context) {
	b, _ := c.GetRawData()
	url := constant.ApiDomain + constant.APIOfflinefaceTransactions
	Authorization := GetAuthorization(http.MethodPost, constant.APIOfflinefaceTransactions)
	body := httpClient_OfflinefaceTransactions(url, Authorization, b)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func GetOfflinefaceContracts(c *gin.Context) {
	contract_id := c.Param("contract_id")
	appid := c.Query("appid")
	url := constant.ApiDomain + constant.APIGetOfflinefaceContracts
	newURL := strings.ReplaceAll(url, "{contract_id}", contract_id)
	newURL = strings.ReplaceAll(newURL, "XXXX", appid)
	Authorization := GetAuthorization(http.MethodGet, constant.APIGetOfflinefaceContracts)
	body, _ := httpClient_Post(newURL, http.MethodGet, Authorization, nil, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefaceFaceCollections(c *gin.Context) {
	collection_id := c.Param("collection_id")
	url := constant.ApiDomain + constant.APIOfflinefaceFaceCollections
	newURL := strings.ReplaceAll(url, "{collection_id}", collection_id)
	Authorization := GetAuthorization(http.MethodGet, constant.APIOfflinefaceFaceCollections)
	body, _ := httpClient_Post(newURL, http.MethodGet, Authorization, nil, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefaceFaceCollections_organization_id(c *gin.Context) {
	organization_id := c.Query("organization_id")
	offset := c.Query("offset")
	limit := c.Query("limit")
	url := constant.ApiDomain + constant.APIOfflinefaceFaceCollections_organization_id
	newURL := strings.ReplaceAll(url, "{organization_id}", organization_id)
	newURL = newURL + "&offset=" + offset + "&limit=" + limit
	Authorization := GetAuthorization(http.MethodGet, constant.APIOfflinefaceFaceCollections)
	body, _ := httpClient_Post(newURL, http.MethodGet, Authorization, nil, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefaceTransactionsOutTradeNo(c *gin.Context) {
	out_trade_no := c.Param("out_trade_no")
	sp_mchid := c.Query("sp_mchid")
	sub_mchid := c.Query("sub_mchid")
	business_product_id := c.Query("business_product_id")
	url := constant.ApiDomain + constant.APIOfflinefaceTransactionsOutTradeNo
	Authorization := GetAuthorization(http.MethodGet, constant.APIOfflinefaceTransactionsOutTradeNo)
	newURL := strings.ReplaceAll(url, "{out_trade_no}", out_trade_no)
	newURL = strings.ReplaceAll(newURL, "{sp_mchid}", sp_mchid)
	newURL = strings.ReplaceAll(newURL, "{sub_mchid}", sub_mchid)
	newURL = strings.ReplaceAll(newURL, "{business_product_id}", business_product_id)
	body, _ := httpClient_Post(newURL, http.MethodGet, Authorization, nil, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}

func OfflinefaceRepaymentUrl(c *gin.Context) {
	b, _ := c.GetRawData()
	url := constant.ApiDomain + constant.APIOfflinefaceRepaymentUrl
	Authorization := GetAuthorization(http.MethodPost, constant.APIOfflinefaceRepaymentUrl)
	body, _ := httpClient_Post(url, http.MethodPost, Authorization, b, true)
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"isok": "ok",
		"data": data,
	})
}
