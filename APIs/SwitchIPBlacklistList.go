package APIs

import (
	"bytes"
	"encoding/json"
)

type SwitchIPBlacklistListBody struct {
	Status int `json:"Status"`
}

type SwitchIPBlacklistListData struct {
	Done bool `json:"Done"`
}

type SwitchIPBlacklistListResponse struct {
	Response
	Data SwitchIPBlacklistListData `json:"data"`
}

func SwitchIPBlacklistList(c *APIClient, switchIPBlacklistListBody SwitchIPBlacklistListBody) (SwitchIPBlacklistListResponse, error) {
	url := c.Config.Domain + c.Config.SwitchIPBlacklistListAPI

	jsonData, err := json.Marshal(switchIPBlacklistListBody)
	if err != nil {
		return SwitchIPBlacklistListResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return SwitchIPBlacklistListResponse{}, err
	}

	switchIPBlacklistListResponse := SwitchIPBlacklistListResponse{}
	err = json.Unmarshal(body, &switchIPBlacklistListResponse)
	return switchIPBlacklistListResponse, err
}
