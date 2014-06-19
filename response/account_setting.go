package response

import (
	"encoding/json"
	"net/http"
	)

type AccountSetting struct {
	Account
	ConnectedToFacebook int `json:"connected_to_facebook"`
	ConnectedToTwitter	int `json:"connected_to_twitter"`
}

func (resp AccountSetting) GetJSONResponseData() ([]byte, int) {
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusOK
}