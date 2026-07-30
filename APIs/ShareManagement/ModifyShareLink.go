package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// ModifyShareLinkBody 修改分享链接的请求体
type ModifyShareLinkBody struct {
	ShareIdList        []uint64 `json:"shareIdList"`
	TrafficSwitch      int      `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int      `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64    `json:"trafficLimit,omitempty"`
}

// ModifyShareLinkData 修改分享链接的返回数据
type ModifyShareLinkData struct{}

// ModifyShareLinkResponse 修改分享链接的响应
type ModifyShareLinkResponse struct {
	Client.Response
	Data *ModifyShareLinkData `json:"data"`
}

// ModifyShareLink 修改分享链接
func ModifyShareLink(c *Client.APIClient, modifyShareLinkBody ModifyShareLinkBody) (ModifyShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.ModifyShareLinkAPI

	jsonData, err := json.Marshal(modifyShareLinkBody)
	if err != nil {
		return ModifyShareLinkResponse{}, err
	}

	body, err := c.PutData(url, bytes.NewReader(jsonData))
	if err != nil {
		return ModifyShareLinkResponse{}, err
	}

	modifyShareLinkResponse := ModifyShareLinkResponse{}
	err = json.Unmarshal(body, &modifyShareLinkResponse)
	return modifyShareLinkResponse, err
}
