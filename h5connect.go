package umfsdk

import (
	"github.com/cnlisea/crypto"
	"net/http"
	"strconv"
	"strings"
)

func H5Connect(cfg *Config) ([]byte, error) {
	param := map[string]string{
		"service":       "active_scancode_order_new",       // 请求类型
		"mer_id":        cfg.MerId,                         // 商户号
		"ret_url":       cfg.RetUrl,                        // 同步通知地址
		"notify_url":    cfg.NotifyUrl,                     // 异步通知地址
		"goods_id":      cfg.GoodsId,                       // 商品号
		"goods_inf":     cfg.GoodsInfo,                     // 商品描述信息
		"order_id":      cfg.OrderId,                       // 订单号
		"mer_date":      cfg.MerDate.Format("20060102"),    // 订单日期
		"amount":        strconv.FormatInt(cfg.Amount, 10), // 金额
		"mer_priv":      cfg.MerPriv,                       // 私有域
		"user_ip":       cfg.UserIp,                        // 用户IP地址
		"scancode_type": "WECHAT",                          // 类型: 微信
	}

	param = publicParams(param) // 公共参数
	param["sign"] = Sign(param) // 签名

	// http post request
	res, err := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(crypto.BuildQuery(param)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ResponseParse(res.Body)
}
