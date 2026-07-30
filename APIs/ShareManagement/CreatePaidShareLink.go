package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CreatePaidShareLinkBody 创建付费分享链接的请求体
type CreatePaidShareLinkBody struct {
	ShareName          string `json:"shareName"`
	FileIDList         string `json:"fileIDList"`
	PayAmount          int    `json:"payAmount"`
	IsReward           int    `json:"isReward,omitempty"`
	ResourceDesc       string `json:"resourceDesc,omitempty"`
	TrafficSwitch      int    `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64  `json:"trafficLimit,omitempty"`
}

// CreatePaidShareLinkData 创建付费分享链接的返回数据
type CreatePaidShareLinkData struct {
	ShareID  int    `json:"shareID"`
	ShareKey string `json:"shareKey"`
}

// CreatePaidShareLinkResponse 创建付费分享链接的响应
type CreatePaidShareLinkResponse struct {
	Client.Response
	Data CreatePaidShareLinkData `json:"data"`
}

// CreatePaidShareLink 创建付费分享链接
func CreatePaidShareLink(c *Client.APIClient, createPaidShareLinkBody CreatePaidShareLinkBody) (CreatePaidShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.CreatePaidShareLinkAPI

	jsonData, err := json.Marshal(createPaidShareLinkBody)
	if err != nil {
		return CreatePaidShareLinkResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CreatePaidShareLinkResponse{}, err
	}

	createPaidShareLinkResponse := CreatePaidShareLinkResponse{}
	err = json.Unmarshal(body, &createPaidShareLinkResponse)
	return createPaidShareLinkResponse, err
}
