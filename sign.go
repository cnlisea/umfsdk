package umfsdk

import (
	"bytes"
	gcrypto "crypto"
	"github.com/cnlisea/crypto"
	"sort"
	"net/url"
)

func Sign(param map[string]string) string {
	// key sort
	keys := make([]string, 0, len(param))
	for k := range param {
		if k == "sign_type" || param[k] == "" {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)
	var b bytes.Buffer
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(param[k])
		b.WriteString("&")
	}

	return crypto.EncryptBase64(crypto.SignRSA(b.Bytes()[:b.Len()-1], gcrypto.SHA1, string(privateKey)))
}

func VerifySign(data string) bool {
	values, err := url.ParseQuery(data)
	if err != nil {
		return false
	}

	sign := values.Get("sign")
	if sign == "" {
		return false
	}
	values.Del("sign")
	values.Del("sign_type")

	return crypto.VerifySignature([]byte(values.Encode()), sign, gcrypto.SHA1, string(publicKey))
}
