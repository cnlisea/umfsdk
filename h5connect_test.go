package umfsdk

import (
	"log"
	"testing"
	"time"
)

func TestH5Connect(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", "", "m"); err != nil {
		log.Fatal(err)
	}
	cfg := &Config{
		MerId:     "",
		RetUrl:    "http://test.lisea.cn",
		NotifyUrl: "http://test.lisea.cn/pay/notify",
		GoodsInfo: "sea测试",
		OrderId:   "11111111111111111111111111111111",
		Amount:    1,
		MerData:   time.Now(),
		UserIp:    "127.0.0.1",
	}

	t.Log(H5Connect(cfg, true))
}