package citcallgo

type citcallURL struct {
	DefaultApiURL string
	ApiVersion    string
	MisscallOtp   string
	Sms           string
	SmsOTP        string
}

func NewCitcallURL() *citcallURL {
	var (
		defaultApiURL = "https://citcall.pub"
		apiVersion    = "v3"
		misscallOtp   = "motp"
		sms           = "sms"
		smsOtp        = "smsotp"
	)

	u := &citcallURL{
		DefaultApiURL: defaultApiURL,
		ApiVersion:    apiVersion,
		MisscallOtp:   misscallOtp,
		Sms:           sms,
		SmsOTP:        smsOtp,
	}

	return u
}
