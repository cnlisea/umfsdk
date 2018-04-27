package umfsdk

import (
	"testing"
)

func TestVerifySign(t *testing.T) {
	if !VerifySign(`amount=1&amt_type=RMB&charset=UTF-8&error_code=0000&media_type=MOBILE&mer_date=20180427&mer_id=60216202&order_id=11111111111111111111111111111128&pay_date=20180427&pay_seq=WEBANK20180427173930041643371085&pay_type=0&service=pay_result_notify&settle_date=20180427&trade_no=3804271739600835&trade_state=TRADE_SUCCESS&version=4.0&sign=eFekonJ1RunxANQuHtGKoKLxhGG%2F1oK2hzTXj16aCdNzoozf%2BsqD8UDKRn2LBkKmERwWWBV94w0Xm99p9QeRzYVkjhP0bqvtiV7O%2FDgaAbnc6fnLYObfRhA6Ohmz3dzfz165fQnU2flkmt28IBD0El7le1Tv1Kjs%2Bwlhpdu8V6I%3D&sign_type=RSA`) {
		t.Fatal("failure..")
	}
	t.Log("success...")
}
