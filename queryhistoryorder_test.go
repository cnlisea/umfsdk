package umfsdk

import (
	"testing"
	"time"
)

func TestQueryHistoryOrder(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", ""); err != nil {
		t.Fatal("err:", err)
	}

	cfg := &Config{
		MerId:   "60216202",
		OrderId: "20184298504380172294697261864686",
		TradeNo: "3804291651445650",
		MerData: time.Date(2018, 4, 29, 0, 0, 0, 0, time.Local),
	}

	res, err := QueryHistoryOrder(cfg)
	if err != nil {
		t.Fatal("err:", err)
		return
	}

	t.Log("res:", res)
}
