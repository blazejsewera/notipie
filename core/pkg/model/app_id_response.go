package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type AppIDResponse struct {
	AppID string `json:"appId"`
}

func AppIDResponseOf(appID string) AppIDResponse {
	return AppIDResponse{AppID: appID}
}

func (a AppIDResponse) ToJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("marshal AppIDResponse: %s", err)
	}
	return jsonBytes, nil
}

func AppIDResponseFromJSON(jsonBytes []byte) (AppIDResponse, error) {
	appIDResponse := AppIDResponse{}
	err := json.Unmarshal(jsonBytes, &appIDResponse)
	if err != nil {
		return AppIDResponse{}, err
	}
	return appIDResponse, nil
}

func AppIDResponseFromReader(r io.Reader) (AppIDResponse, error) {
	appIdResponse := AppIDResponse{}
	d := json.NewDecoder(r)
	err := d.Decode(&appIdResponse)
	if err != nil {
		return AppIDResponse{}, err
	}
	return appIdResponse, nil
}
