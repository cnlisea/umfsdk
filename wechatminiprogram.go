package umfsdk

import (
	"errors"
	"github.com/cnlisea/crypto"
	"net/http"
	"strconv"
	"strings"
)

type WeChatMiniProgramInfo struct {
	AppId     string // 小程序唯一编号
	TimeStamp string // 时间戳
	NonceStr  string // 随机字符串
	Package   string // 数据包
	SignType  string // 签名方式
	PaySign   string // 签名
}

// 微信小程序支付
func WeChatMiniProgram(cfg *Config, appId string, openId string) (*WeChatMiniProgramInfo, error) {
	param := map[string]string{
		"service":       "wechatminiprogram_order",         // 请求类型
		"mer_id":        cfg.MerId,                         // 商户号
		"notify_url":    cfg.NotifyUrl,                     // 异步通知地址
		"amt_type":      "RMB",                             // 付款币种
		"goods_id":      cfg.GoodsId,                       // 商品号
		"goods_inf":     cfg.GoodsInfo,                     // 商品描述信息
		"order_id":      cfg.OrderId,                       // 订单号
		"mer_date":      cfg.MerDate.Format("20060102"),    // 订单日期
		"amount":        strconv.FormatInt(cfg.Amount, 10), // 金额
		"mer_priv":      cfg.MerPriv,                       // 私有域
		"user_ip":       cfg.UserIp,                        // 用户IP地址
		"open_id":       openId,                            // 微信用户授权openId
		"app_id":        appId,                             // 小程序的唯一标识
		"scancode_type": "WECHAT",                          // 扫码类型
	}

	param = publicParams(param) // 公共参数
	param["sign"] = Sign(param) // 签名

	// http post request
	res, err := http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(crypto.BuildQuery(param)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	parseData, err := ResponseParse(res.Body)
	if err != nil {
		return nil, err
	}

	vals := strings.Split(string(parseData), "&")
	ret := make(map[string]string, len(vals))
	for i := range vals {
		s := strings.SplitN(vals[i], "=", 2)
		if len(s) < 2 {
			continue
		}
		ret[s[0]] = s[1]
	}

	if ret["ret_code"] != "0000" {
		return nil, errors.New(ret["ret_msg"])
	}

	return &WeChatMiniProgramInfo{
		AppId:     appId,
		TimeStamp: ret["time_stamp"],
		NonceStr:  ret["nonce_str"],
		Package:   ret["package"],
		SignType:  ret["weixin_sign_type"],
		PaySign:   ret["pay_sign"],
	}, nil
}
