package IPBlacklistConfiguration

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type SwitchIPBlacklistListBody struct {
	Status int `json:"Status"`
}

type SwitchIPBlacklistListData struct {
	Done bool `json:"Done"`
}

type SwitchIPBlacklistListResponse struct {
	ClientAndMethods.Response
	Data SwitchIPBlacklistListData `json:"data"`
}

func SwitchIPBlacklistList(c *ClientAndMethods.APIClient, switchIPBlacklistListBody SwitchIPBlacklistListBody) (SwitchIPBlacklistListResponse, error) {
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
