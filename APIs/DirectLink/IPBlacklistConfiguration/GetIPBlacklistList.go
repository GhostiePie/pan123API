package IPBlacklistConfiguration

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type GetIPBlacklistListData struct {
	IpList []string `json:"ipList"`
	Status int      `json:"status"`
}
type GetIPBlacklistListResponse struct {
	ClientAndMethods.Response
	Data GetIPBlacklistListData `json:"data"`
}

func GetIPBlacklistList(c *ClientAndMethods.APIClient) (GetIPBlacklistListResponse, error) {
	url := c.Config.Domain + c.Config.GetIPBlacklistListAPI
	resp, err := c.GetQuery(url)
	if err != nil {
		return GetIPBlacklistListResponse{}, err
	}
	getIPBlacklistListResponse := GetIPBlacklistListResponse{}
	err = json.Unmarshal(resp, &getIPBlacklistListResponse)
	if err != nil {
		return GetIPBlacklistListResponse{}, err
	}
	return getIPBlacklistListResponse, nil
}
