package umfsdk

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

func QueryHistoryOrder(cfg *Config) (map[string]string, error) {
	param := map[string]string{
		"service":  "mer_order_info_query",         // 请求类型
		"mer_id":   cfg.MerId,                      // 商户号
		"mer_date": cfg.MerData.Format("20060102"), // 订单日期
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
	parseData, err := QueryHistoryOrderResponseParse(res.Body)
	if err != nil {
		return nil, err
	}

	// sign verify
	if !VerifySign(*(*string)(unsafe.Pointer(&parseData))) {
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

func QueryHistoryOrderResponseParse(data io.Reader) ([]byte, error) {
	doc, err := html.Parse(data)
	if err != nil {
		return nil, err
	}

	var search func(*html.Node) *html.Node
	search = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "meta" {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if r := search(c); r != nil {
				return r
			}
		}
		return nil
	}

	if doc = search(doc); doc == nil {
		return nil, errors.New("Invalid response data")
	}

	var ret []byte
	for i := range doc.Attr {
		if doc.Attr[i].Key == "CONTENT" || strings.Index(doc.Attr[i].Val, "sign") >= 0 {
			ret = *(*[]byte)(unsafe.Pointer(&doc.Attr[i].Val))
			break
		}
	}

	return ret, nil
}
