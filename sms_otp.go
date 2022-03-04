package citcallgo

import (
	"context"
	"encoding/json"
	"net/http"
)

// citcall sms otp request body
type SMSOTPRequest struct {
	Msisdn   string `json:"msisdn"`
	SenderId string `json:"senderid"`
	Text     string `json:"text"`
}

// citcall sms otp response
type SMSOTPResponse struct {
	Rc       int    `json:"rc"`
	Info     string `json:"info"`
	SMSCount int    `json:"sms_count"`
	SenderId string `json:"senderid"`
	Msisdn   string `json:"msisdn"`
	Text     string `json:"text"`
	TrxID    string `json:"trxid"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

// citcall send sms otp
func (c *citcall) SendSMSOTP(ctx context.Context, requestBody *SMSOTPRequest) (*SMSOTPResponse, error) {
	res, err := c.request(ctx, http.MethodPost, c.citcallURL.smsOTP, requestBody)
	if err != nil {
		return nil, err
	}

	return handleSMSOTPResponse(res)
}

func handleSMSOTPResponse(res *http.Response) (*SMSOTPResponse, error) {
	var response SMSOTPResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
