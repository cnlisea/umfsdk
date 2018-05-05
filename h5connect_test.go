package umfsdk

import (
	"testing"
	"time"
	"fmt"
	"unsafe"
)

func TestH5Connect(t *testing.T) {
	if err := Init("http://pay.soopay.net/spay/pay/payservice.do", "D:/key/60216202_.key.pem"); err != nil {
		t.Fatal(err)
	}
	cfg := &Config{
		MerId:     "60216202",
		RetUrl:    "http://test.yuanlaihuyu.com",
		NotifyUrl: "http://test.yuanlaihuyu.com/pay/notify",
		GoodsInfo: "元来棋牌测试",
		OrderId:   "22222222222222222222222222222222",
		Amount:    1,
		MerDate:   time.Now(),
		UserIp:    "127.0.0.1",
	}

	url, err := H5Connect(cfg)
	if err != nil {
		t.Fatal("err:", err)
	}

	fmt.Println("url:", *(*string)(unsafe.Pointer(&url)))
}
