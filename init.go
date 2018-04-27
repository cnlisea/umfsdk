package umfsdk

import "github.com/cnlisea/crypto"

var (
	requestUrl string                                                                                                                                                                                                                                                                                       // 请求url
	privateKey []byte                                                                                                                                                                                                                                                                                       // 私钥
	publicKey  = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDGWAu4p6aK1SiQqNKT1nTgYwA8
cz0Rde9fRtmLJAx1QxLqrerAUVl/uuXV7NQFSkTipouo3cwEEpae89267AeLJBzK
PbKnUID6JYGbwnq7CiRR4E244zcgqE8jo8DnkbH3KkiWonoUMD1uHy6TUFv5W7zr
haz/E59MVmbzrp1TwwIDAQAB
-----END PUBLIC KEY-----
`) // 公钥
)

func Init(url string, privateKeyPath string) error {
	var err error

	privateKey, err = ReadFileAll(privateKeyPath)
	if err != nil {
		return err
	}

	// private key pkcs1 to pkcs8 format
	privateKey, err = crypto.RsaPrivateKeyPkcs1ToPkcs8(privateKey)
	if err != nil {
		return err
	}

	requestUrl = url

	return err
}
