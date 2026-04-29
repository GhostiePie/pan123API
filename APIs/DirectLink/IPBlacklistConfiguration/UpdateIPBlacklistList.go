package IPBlacklistConfiguration

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type UpdateIPBlacklistListBody struct {
	IpList []string `json:"IpList"`
}

type UpdateIPBlacklistListData struct{}

type UpdateIPBlacklistListResponse struct {
	ClientAndMethods.Response
	Data *UpdateIPBlacklistListData `json:"data"`
}

func UpdateIPBlacklistList(c *ClientAndMethods.APIClient, updateIPBlacklistListBody UpdateIPBlacklistListBody) (UpdateIPBlacklistListResponse, error) {
	url := c.Config.Domain + c.Config.UpdateIPBlacklistListAPI

	jsonData, err := json.Marshal(updateIPBlacklistListBody)
	if err != nil {
		return UpdateIPBlacklistListResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return UpdateIPBlacklistListResponse{}, err
	}

	updateIPBlacklistListResponse := UpdateIPBlacklistListResponse{}
	err = json.Unmarshal(body, &updateIPBlacklistListResponse)
	return updateIPBlacklistListResponse, err
}
