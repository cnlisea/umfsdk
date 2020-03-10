package umfsdk

import (
	"log"
	"testing"
	"time"
)

func TestWeChatMiniProgram(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", ""); err != nil {
		log.Fatal(err)
	}
	cfg := &Config{
		MerId:     "",
		NotifyUrl: "http://test.lisea.cn/pay/notify",
		GoodsInfo: "sea测试",
		OrderId:   "11111111111111111111111111119222",
		Amount:    1,
		MerDate:   time.Now(),
		UserIp:    "127.0.0.1",
	}

	t.Log(WeChatMiniProgram(cfg, "wx111222", "us11111"))
}
