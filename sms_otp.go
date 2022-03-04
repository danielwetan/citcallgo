package citcallgo

import (
	"context"
	"encoding/json"
	"net/http"
)

type SMSOTPRequest struct {
	Msisdn   string `json:"msisdn"`
	SenderId string `json:"senderid"`
	Text     string `json:"text"`
}

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
