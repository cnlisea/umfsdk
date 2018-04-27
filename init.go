package umfsdk

import "github.com/cnlisea/crypto"

var (
	requestUrl string // 请求url
	privateKey []byte // 私钥
	publicKey  []byte // 公钥
)

func Init(url string, publicKeyPath string, privateKeyPath string) error {
	var err error

	publicKey, err = ReadFileAll(publicKeyPath)
	if err != nil {
		return err
	}

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
