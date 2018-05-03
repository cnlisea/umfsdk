package umfsdk

import "time"

type Config struct {
	MerId     string    // 商户编号
	RetUrl    string    // 同步通知地址
	NotifyUrl string    // 异步通知地址
	GoodsId   string    // 商品号
	GoodsInfo string    // 商品描述信息
	OrderId   string    // 订单号
	Amount    int64     // 金额, 单位分
	MerDate   time.Time // 订单日期
	UserIp    string    // 用户IP
	MerPriv   string    // 私有域
	TradeNo   string    // 交易号
}
