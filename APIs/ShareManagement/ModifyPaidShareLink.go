package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// ModifyPaidShareLinkBody 修改付费分享链接的请求体
type ModifyPaidShareLinkBody struct {
	ShareIdList        []uint64 `json:"shareIdList"`
	TrafficSwitch      int      `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int      `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64    `json:"trafficLimit,omitempty"`
}

// ModifyPaidShareLinkData 修改付费分享链接的返回数据
type ModifyPaidShareLinkData struct{}

// ModifyPaidShareLinkResponse 修改付费分享链接的响应
type ModifyPaidShareLinkResponse struct {
	Client.Response
	Data *ModifyPaidShareLinkData `json:"data"`
}

// ModifyPaidShareLink 修改付费分享链接
func ModifyPaidShareLink(c *Client.APIClient, modifyPaidShareLinkBody ModifyPaidShareLinkBody) (ModifyPaidShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.ModifyPaidShareLinkAPI

	jsonData, err := json.Marshal(modifyPaidShareLinkBody)
	if err != nil {
		return ModifyPaidShareLinkResponse{}, err
	}

	body, err := c.PutData(url, bytes.NewReader(jsonData))
	if err != nil {
		return ModifyPaidShareLinkResponse{}, err
	}

	modifyPaidShareLinkResponse := ModifyPaidShareLinkResponse{}
	err = json.Unmarshal(body, &modifyPaidShareLinkResponse)
	return modifyPaidShareLinkResponse, err
}
