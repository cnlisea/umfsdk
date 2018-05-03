package umfsdk

import (
	"bytes"
	"sort"
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
