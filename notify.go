package umfsdk

import (
	"bytes"
	"sort"
)

type NotifyRequest struct {
	Service    string `json:"service"`     // 接口名称
	Charset    string `json:"charset"`     // 参数字符编码集
	SignType   string `json:"sign_type"`   // 签名方式
	Sign       string `json:"sign"`        // 签名
	Version    string `json:"version"`     // 版本号
	MerId      string `json:"mer_id"`      // 商户编号
	TradeNo    string `json:"trade_no"`    // 联动交易号
	GoodsId    string `json:"goods_id"`    // 商品号
	OrderId    string `json:"order_id"`    // 订单号
	MerDate    string `json:"mer_date"`    // 原商户订单日期 YYYYMMDD
	PayDate    string `json:"pay_date"`    // 支付日期 YYYYMMDD
	Amount     string `json:"amount"`      // 付款金额 单位分
	AmtType    string `json:"amt_type"`    // 付款币种
	PayType    string `json:"pay_type"`    // 支付方式
	SettleDate string `json:"settle_date"` // 对账日期 YYYYMMDD
	MerPriv    string `json:"mer_priv"`    // 商户私有域
	TradeState string `json:"trade_state"` // 订单状态
	ErrorCode  string `json:"error_code"`  // 交易错误码
}

const (
	NotifyRequestTradeStateSuccess = "TRADE_SUCCESS"  // 交易成功
	NotifyRequestTradeStateWaitPay = "WAIT_BUYER_PAY" // 交易创建,等待支付
	NotifyRequestTradeStateClose   = "TRADE_CLOSE"    // 交易关闭(商户支付或者查询已经过期的订单后,订单状态才会改变为交易关闭)
	NotifyRequestTradeStateCancel  = "TRADE_CANCEL"   // 交易撤销
	NotifyRequestTradeStateFail    = "TRADE_FAIL"     // 交易失败
)

func NotifyResponse(cfg *Config) string {
	m := map[string]string{
		"mer_date":  cfg.MerDate.Format("20060102"), // 订单日期
		"mer_id":    cfg.MerId,                      // 商户编号
		"order_id":  cfg.OrderId,                    // 订单号
		"ret_code":  "0000",                         // 返回状态码
		"ret_msg":   "success",                      // 返回状态信息
		"version":   "4.0",                          // 版本号
		"sign_type": "RSA",                          // 签名方式
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(m[k])
		b.WriteString("&")
	}
	b.WriteString("sign=")
	b.WriteString(Sign(m))

	var ret bytes.Buffer
	ret.WriteString(`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">`)
	ret.WriteString("<HTML>")
	ret.WriteString(`  <HEAD><META NAME="MobilePayPlatform" CONTENT="`)
	ret.Write(b.Bytes())
	ret.WriteString(`"/></HEAD>`)
	ret.WriteString("  <BODY>")
	ret.WriteString("</BODY>")
	ret.WriteString("</HTML>")

	return ret.String()
}
