package IPBlacklistConfiguration

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// GetIPBlacklistListData 获取IP黑名单列表的返回数据
type GetIPBlacklistListData struct {
	IpList []string `json:"ipList"`
	Status int      `json:"status"`
}

// GetIPBlacklistListResponse 获取IP黑名单列表的响应
type GetIPBlacklistListResponse struct {
	Client.Response
	Data GetIPBlacklistListData `json:"data"`
}

// GetIPBlacklistList 获取IP黑名单列表
func GetIPBlacklistList(c *Client.APIClient) (GetIPBlacklistListResponse, error) {
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
