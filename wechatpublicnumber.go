package umfsdk

import (
	"bytes"
	"strconv"

	"github.com/cnlisea/crypto"
)

func WeChatPublicNumber(cfg *Config) string {
	param := map[string]string{
		"service":          "publicnumber_and_verticalcode",   // 请求类型
		"mer_id":           cfg.MerId,                         // 商户号
		"ret_url":          cfg.RetUrl,                        // 同步通知地址
		"notify_url":       cfg.NotifyUrl,                     // 异步通知地址
		"amt_type":         "RMB",                             // 付款币种
		"goods_id":         cfg.GoodsId,                       // 商品号
		"goods_inf":        cfg.GoodsInfo,                     // 商品描述信息
		"order_id":         cfg.OrderId,                       // 订单号
		"mer_date":         cfg.MerDate.Format("20060102"),    // 订单日期
		"amount":           strconv.FormatInt(cfg.Amount, 10), // 金额
		"mer_priv":         cfg.MerPriv,                       // 私有域
		"user_ip":          cfg.UserIp,                        // 用户IP地址
		"is_public_number": "Y",                               // 是否是公众号支付
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
