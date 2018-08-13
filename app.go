package umfsdk

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/cnlisea/crypto"
)

func App(cfg *Config) (string, string, error) {
	param := map[string]string{
		"service":    "pay_req",                         // 请求类型
		"mer_id":     cfg.MerId,                         // 商户号
		"notify_url": cfg.NotifyUrl,                     // 异步通知地址
		"amt_type":   "RMB",                             // 付款币种
		"goods_id":   cfg.GoodsId,                       // 商品号
		"goods_inf":  cfg.GoodsInfo,                     // 商品描述信息
		"order_id":   cfg.OrderId,                       // 订单号
		"mer_date":   cfg.MerDate.Format("20060102"),    // 订单日期
		"amount":     strconv.FormatInt(cfg.Amount, 10), // 金额
		"mer_priv":   cfg.MerPriv,                       // 私有域
		"user_ip":    cfg.UserIp,                        // 用户IP地址
	}

	param = publicParams(param) // 公共参数
	param["sign"] = Sign(param) // 签名

	// http post request
	res, err := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(crypto.BuildQuery(param)))
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	parseData, err := ResponseParse(res.Body)
	if err != nil {
		return "", "", err
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

	if ret["ret_code"] != "0000" {
		return "", "", errors.New(ret["ret_msg"])
	}

	return ret["trade_no"], AppSign(cfg.MerId, cfg.OrderId, strconv.FormatInt(cfg.Amount, 10), cfg.MerDate), nil
}
