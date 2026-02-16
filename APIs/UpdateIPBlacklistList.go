package APIs

import (
	"encoding/json"
)

type UpdateIPBlacklistListBody struct {
	IpList []string `json:"IpList"`
}

type UpdateIPBlacklistListData struct{}

type UpdateIPBlacklistListResponse struct {
	Response
	Data *UpdateIPBlacklistListData `json:"data"`
}

func (c *APIClient) UpdateIPBlacklistList(updateIPBlacklistListBody UpdateIPBlacklistListBody) (UpdateIPBlacklistListResponse, error) {
	url := c.Config.Domain + c.Config.UpdateIPBlacklistListAPI

	jsonData, err := json.Marshal(updateIPBlacklistListBody)
	if err != nil {
		return UpdateIPBlacklistListResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return UpdateIPBlacklistListResponse{}, err
	}

	updateIPBlacklistListResponse := UpdateIPBlacklistListResponse{}
	err = json.Unmarshal(body, &updateIPBlacklistListResponse)
	return updateIPBlacklistListResponse, err
}
