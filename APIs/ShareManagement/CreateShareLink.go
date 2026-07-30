package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CreateShareLinkBody 创建分享链接的请求体
type CreateShareLinkBody struct {
	ShareName          string `json:"shareName"`
	ShareExpire        int    `json:"shareExpire"`
	FileIDList         string `json:"fileIDList"`
	SharePwd           string `json:"sharePwd,omitempty"`
	TrafficSwitch      int    `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64  `json:"trafficLimit,omitempty"`
}

// CreateShareLinkData 创建分享链接的返回数据
type CreateShareLinkData struct {
	ShareID  int    `json:"shareID"`
	ShareKey string `json:"shareKey"`
}

// CreateShareLinkResponse 创建分享链接的响应
type CreateShareLinkResponse struct {
	Client.Response
	Data CreateShareLinkData `json:"data"`
}

// CreateShareLink 创建分享链接
func CreateShareLink(c *Client.APIClient, createShareLinkBody CreateShareLinkBody) (CreateShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.CreateShareLinkAPI

	jsonData, err := json.Marshal(createShareLinkBody)
	if err != nil {
		return CreateShareLinkResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CreateShareLinkResponse{}, err
	}

	createShareLinkResponse := CreateShareLinkResponse{}
	err = json.Unmarshal(body, &createShareLinkResponse)
	return createShareLinkResponse, err
}
