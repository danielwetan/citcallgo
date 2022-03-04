package citcallgo

import (
	"context"
	"encoding/json"
	"net/http"
)

// citcall sms request body
type SMSRequest struct {
	Msisdn   string `json:"msisdn"`
	SenderId string `json:"senderid"`
	Text     string `json:"text"`
}

// citcall sms response
type SMSResponse struct {
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

// citcall send sms
func (c *citcall) SendSMS(ctx context.Context, requestBody *SMSRequest) (*SMSResponse, error) {
	res, err := c.request(ctx, http.MethodPost, c.citcallURL.sms, requestBody)
	if err != nil {
		return nil, err
	}

	return handleSMSResponse(res)
}

func handleSMSResponse(res *http.Response) (*SMSResponse, error) {
	var response SMSResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
