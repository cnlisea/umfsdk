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
		MerId:   "",
		OrderId: "",
		MerDate: time.Date(2018, 5, 3, 0, 0, 0, 0, time.Local),
	}

	res, err := QueryHistoryOrder(cfg)
	if err != nil {
		t.Fatal("err:", err)
		return
	}

	for k, v := range res {
		t.Log(k, " = ", v)
	}
}
