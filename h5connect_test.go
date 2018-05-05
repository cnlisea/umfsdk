package umfsdk

import (
	"testing"
	"time"
	"unsafe"
)

func TestH5Connect(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", ""); err != nil {
		t.Fatal(err)
	}
	cfg := &Config{
		MerId:     "",
		RetUrl:    "http://test.lisea.cn",
		NotifyUrl: "http://test.lisea.cn/pay/notify",
		GoodsInfo: "lisea测试",
		OrderId:   "22222222222222222222222222222222",
		Amount:    1,
		MerDate:   time.Now(),
		UserIp:    "127.0.0.1",
	}

	url, err := H5Connect(cfg)
	if err != nil {
		t.Fatal("err:", err)
	}

	t.Log("url:", *(*string)(unsafe.Pointer(&url)))
}
