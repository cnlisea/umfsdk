package umfsdk

import (
	"testing"
)

func TestVerifySignNotify(t *testing.T) {
	if !VerifySignNotify(`amount=1&amt_type=RMB&charset=UTF-8&error_code=0000&media_type=MOBILE&mer_date=20180503&mer_id=60216202&order_id=20185363493698436878823217403188&pay_date=20180503&pay_seq=WEBANK20180503140338041666891384&pay_type=0&service=pay_result_notify&settle_date=20180503&trade_no=3805031403717574&trade_state=TRADE_SUCCESS&version=4.0&sign=Nx1w45vYVrdZ%2FUzVN5zE4eb8TB2dGiBTM%2F707VV4BQfn2oltm11oaOHuxRQPjoJFOfTT6mQLXB1UyzIiGRZ8duGei4ftHIHfB97zklfMlAEYZM6QHQUY3f%2BfyY56h0g4WNlMM2X9WXPFqd%2FdHGpZ0hfmSzKB92APqFJzIiJ5j8w%3D&sign_type=RSA`) {
		t.Fatal("failure..")
	}
	t.Log("success...")
}

func TestVerifySignQuery(t *testing.T) {
	if !VerifySignQuery(`amount=1&amt_type=RMB&media_id=00000000000&media_type=MOBILE&mer_date=20180429&mer_id=60216202&order_id=20184298504380172294697261864686&orig_retcode=&orig_retmsg=&pay_date=20180429&pay_seq=WEBANK20180429165115041652262980&pay_type=&product_id=P15Y0025&refund_amt=0&ret_code=0000&ret_msg=操作成功&settle_date=20180429&sign_type=RSA&trade_no=3804291651445650&trade_state=TRADE_SUCCESS&version=4.0&sign=fQOQW7k60x5VVX79zD4vigGL9VEKIslMKPZLNFFem44Z5L+uLHk3WbF+D2+QHudpWs5kjsf/4p9i1qjs97vYsfI7jcy1wCP1lm7qQM4D9KIKTmwSsXF78tM/E6F3SGWaPyrHuk3qWRXQnMdpndmwk+nprGWTeMZvfC1UWiVQCKI=`) {
		t.Fatal("failure")
	}
	t.Log("success...")
}
