package umfsdk

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func QueryHistoryOrder(cfg *Config) (map[string]string, error) {
	param := map[string]string{
		"service":  "mer_order_info_query",         // 请求类型
		"mer_id":   cfg.MerId,                      // 商户号
		"mer_date": cfg.MerDate.Format("20060102"), // 订单日期
		"order_id": cfg.OrderId,                    // 订单号
		"trade_no": cfg.TradeNo,                    // 交易号
	}
	param = publicParams(param)
	param["sign"] = Sign(param) // 签名

	// generate form
	data := make(url.Values, len(param))
	for k, v := range param {
		if v == "" {
			continue
		}
		data.Add(k, v)
	}

	// http post form request
	res, err := http.PostForm(requestUrl, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// param response
	parseData, err := ResponseParse(res.Body)
	if err != nil {
		return nil, err
	}

	// sign verify
	if !VerifySignQuery(*(*string)(unsafe.Pointer(&parseData))) {
		return nil, errors.New("Invalid sign data ")
	}

	vals := strings.Split(string(parseData), "&")
	ret := make(map[string]string, len(vals))
	for i := range vals {
		s := strings.Split(vals[i], "=")
		if len(s) < 2 {
			continue
		}
		ret[s[0]] = s[1]
	}

	return ret, nil
}
