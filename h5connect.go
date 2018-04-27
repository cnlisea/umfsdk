package umfsdk

import (
	"bytes"
	"github.com/cnlisea/crypto"
	"strconv"
)

func H5Connect(cfg *Config, auto bool) string {
	autoParam := "FALSE"
	if auto {
		autoParam = "TRUE"
	}
	param := map[string]string{
		"service":       "active_scancode_order_new",       // 请求类型
		"mer_id":        cfg.MerId,                         // 商户号
		"ret_url":       cfg.RetUrl,                        // 同步通知地址
		"notify_url":    cfg.NotifyUrl,                     // 异步通知地址
		"goods_id":      cfg.GoodsId,                       // 商品号
		"goods_inf":     cfg.GoodsInfo,                     // 商品描述信息
		"order_id":      cfg.OrderId,                       // 订单号
		"mer_date":      cfg.MerData.Format("20060102"),    // 订单日期
		"amount":        strconv.FormatInt(cfg.Amount, 10), // 金额
		"mer_priv":      cfg.MerPriv,                       // 私有域
		"user_ip":       cfg.UserIp,                        // 用户IP地址
		"scancode_type": "WECHAT",                          // 类型: 微信
		"auto_pay":      autoParam,                         // 是否需要自动跳转支付
	}
	param = publicParams(param) // 公共参数
	param["sign"] = Sign(param) // 签名

	// url?k1=v1...
	var b bytes.Buffer
	b.WriteString(requestUrl)
	b.WriteString("?")
	b.WriteString(crypto.BuildQuery(param))

	return b.String()
}
