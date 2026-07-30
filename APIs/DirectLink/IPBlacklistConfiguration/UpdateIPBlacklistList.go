package IPBlacklistConfiguration

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// UpdateIPBlacklistListBody 更新IP黑名单列表的请求体
type UpdateIPBlacklistListBody struct {
	IpList []string `json:"IpList"`
}

// UpdateIPBlacklistListData 更新IP黑名单列表的返回数据
type UpdateIPBlacklistListData struct{}

// UpdateIPBlacklistListResponse 更新IP黑名单列表的响应
type UpdateIPBlacklistListResponse struct {
	Client.Response
	Data *UpdateIPBlacklistListData `json:"data"`
}

// UpdateIPBlacklistList 更新IP黑名单列表
func UpdateIPBlacklistList(c *Client.APIClient, updateIPBlacklistListBody UpdateIPBlacklistListBody) (UpdateIPBlacklistListResponse, error) {
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
