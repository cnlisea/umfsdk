package umfsdk

func publicParams(param map[string]string) map[string]string {
	param["charset"] = "UTF-8"
	param["sign_type"] = "RSA"
	param["res_format"] = "HTML"
	param["version"] = "4.0"
	param["amt_type"] = "RMB"

	return param
}
