package umfsdk

import (
	"testing"
	"time"
)

func TestNotifyResponse(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", ""); err != nil {
		t.Fatal(err)
	}

	cfg := &Config{
		MerId:   "",
		OrderId: "11111111111111111111111111111128",
		MerDate: time.Now(),
	}
	t.Log(NotifyResponse(cfg))
}
