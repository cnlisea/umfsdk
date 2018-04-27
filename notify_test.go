package umfsdk

import (
	"testing"
	"time"
)

func TestNotifyResponse(t *testing.T) {
	if err := Init("", "D:/key/60216202_.key.pem"); err != nil {
		t.Fatal(err)
	}

	cfg := &Config{
		MerId: "60216202",
		OrderId: "11111111111111111111111111111128",
		MerData: time.Now(),
	}
	t.Log(NotifyResponse(cfg))
}
