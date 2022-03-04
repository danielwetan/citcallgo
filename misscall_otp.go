package citcallgo

import (
	"context"
	"encoding/json"
	"net/http"
)

// citcall misscall otp request body
type MisscallOtpRequest struct {
	Msisdn  string `json:"msisdn"`
	Gateway int    `json:"gateway"`
}

// citcall misscall otp response 
type MisscallOtpResponse struct {
	Rc      int    `json:"rc"`
	Trxid   string `json:"trxid"`
	Msisdn  string `json:"msisdn"`
	Token   string `json:"token"`
	Gateway int    `json:"gateway"`
}

// citcall send misscall
func (c *citcall) SendMisscall(ctx context.Context, requestBody *MisscallOtpRequest) (*MisscallOtpResponse, error) {
	res, err := c.request(ctx, http.MethodPost, c.citcallURL.misscallOtp, requestBody)
	if err != nil {
		return nil, err
	}

	return handleMisscallResponse(res)
}

func handleMisscallResponse(res *http.Response) (*MisscallOtpResponse, error) {
	var response MisscallOtpResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
