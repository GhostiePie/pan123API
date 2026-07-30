package IPBlacklistConfiguration

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// SwitchIPBlacklistListBody 切换IP黑名单开关的请求体
type SwitchIPBlacklistListBody struct {
	Status int `json:"Status"`
}

// SwitchIPBlacklistListData 切换IP黑名单开关的返回数据
type SwitchIPBlacklistListData struct {
	Done bool `json:"Done"`
}

// SwitchIPBlacklistListResponse 切换IP黑名单开关的响应
type SwitchIPBlacklistListResponse struct {
	Client.Response
	Data SwitchIPBlacklistListData `json:"data"`
}

// SwitchIPBlacklistList 切换IP黑名单开关
func SwitchIPBlacklistList(c *Client.APIClient, switchIPBlacklistListBody SwitchIPBlacklistListBody) (SwitchIPBlacklistListResponse, error) {
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
