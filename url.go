package citcallgo

type CitcallURL struct {
	DefaultApiURL string
	ApiVersion    string
	MisscallOtp   string
	Sms           string
	SmsOTP        string
}

func NewCitcallURL() *CitcallURL {
	var (
		defaultApiURL = "https://citcall.pub"
		apiVersion    = "v3"
		misscallOtp   = "motp"
		sms           = "sms"
		smsOtp        = "smsotp"
	)

	u := &CitcallURL{
		DefaultApiURL: defaultApiURL,
		ApiVersion:    apiVersion,
		MisscallOtp:   misscallOtp,
		Sms:           sms,
		SmsOTP:        smsOtp,
	}

	return u
}
