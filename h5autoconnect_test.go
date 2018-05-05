package umfsdk

import (
	"log"
	"testing"
	"time"
)

func TestH5AutoConnect(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", ""); err != nil {
		log.Fatal(err)
	}
	cfg := &Config{
		MerId:     "",
		RetUrl:    "http://test.lisea.cn",
		NotifyUrl: "http://test.lisea.cn/pay/notify",
		GoodsInfo: "sea测试",
		OrderId:   "11111111111111111111111111111111",
		Amount:    1,
		MerDate:   time.Now(),
		UserIp:    "127.0.0.1",
	}

	t.Log(H5AutoConnect(cfg))
}
