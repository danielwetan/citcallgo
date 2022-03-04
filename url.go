package citcallgo

type citcallURL struct {
	defaultApiURL string
	apiVersion    string
	misscallOtp   string
	sms           string
	smsOTP        string
}

func newCitcallURL() *citcallURL {
	var (
		defaultApiURL = "https://citcall.pub"
		apiVersion    = "v3"
		misscallOtp   = "motp"
		sms           = "sms"
		smsOtp        = "smsotp"
	)

	u := &citcallURL{
		defaultApiURL: defaultApiURL,
		apiVersion:    apiVersion,
		misscallOtp:   misscallOtp,
		sms:           sms,
		smsOTP:        smsOtp,
	}

	return u
}
