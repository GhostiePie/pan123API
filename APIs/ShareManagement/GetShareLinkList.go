package ShareManagement

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetShareLinkListBody 获取分享链接列表的请求体
type GetShareLinkListBody struct {
	Limit       int `json:"limit"`
	LastShareId int `json:"lastShareId,omitempty"`
}

// ShareListItem 分享链接列表项
type ShareListItem struct {
	ShareId            int    `json:"shareId"`
	ShareKey           string `json:"shareKey"`
	ShareName          string `json:"shareName"`
	Expiration         string `json:"expiration"`
	Expired            int    `json:"expired"`
	SharePwd           string `json:"sharePwd"`
	TrafficSwitch      int    `json:"trafficSwitch"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch"`
	TrafficLimit       int64  `json:"trafficLimit"`
	BytesCharge        int64  `json:"bytesCharge"`
	PreviewCount       int    `json:"previewCount"`
	DownloadCount      int    `json:"downloadCount"`
	SaveCount          int    `json:"saveCount"`
}

// GetShareLinkListData 获取分享链接列表的返回数据
type GetShareLinkListData struct {
	LastShareId int             `json:"lastShareId"`
	ShareList   []ShareListItem `json:"shareList"`
}

// GetShareLinkListResponse 获取分享链接列表的响应
type GetShareLinkListResponse struct {
	Client.Response
	Data GetShareLinkListData `json:"data"`
}

// GetShareLinkList 获取分享链接列表
func GetShareLinkList(c *Client.APIClient, getShareLinkListBody GetShareLinkListBody, config Client.APIConfig) (GetShareLinkListResponse, error) {
	url := config.Domain + config.GetShareLinkListAPI +
		"?limit=" + strconv.Itoa(getShareLinkListBody.Limit)
	if getShareLinkListBody.LastShareId != 0 {
		url += "&lastShareId=" + strconv.Itoa(getShareLinkListBody.LastShareId)
	}

	body, err := c.GetQuery(url)
	if err != nil {
		return GetShareLinkListResponse{}, err
	}

	getShareLinkListResponse := GetShareLinkListResponse{}
	err = json.Unmarshal(body, &getShareLinkListResponse)
	return getShareLinkListResponse, err
}
