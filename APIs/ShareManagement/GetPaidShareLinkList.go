package ShareManagement

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetPaidShareLinkListBody 获取付费分享链接列表的请求体
type GetPaidShareLinkListBody struct {
	Limit       int `json:"limit"`
	LastShareId int `json:"lastShareId,omitempty"`
}

// PaidShareListItem 付费分享链接列表项
type PaidShareListItem struct {
	ShareId            int     `json:"shareId"`
	ShareKey           string  `json:"shareKey"`
	ShareName          string  `json:"shareName"`
	PayAmount          float64 `json:"payAmount"`
	Amount             float64 `json:"amount"`
	Expiration         string  `json:"expiration"`
	Expired            int     `json:"expired"`
	TrafficSwitch      int     `json:"trafficSwitch"`
	TrafficLimitSwitch int     `json:"trafficLimitSwitch"`
	TrafficLimit       int64   `json:"trafficLimit"`
	BytesCharge        int64   `json:"bytesCharge"`
	PreviewCount       int     `json:"previewCount"`
	DownloadCount      int     `json:"downloadCount"`
	SaveCount          int     `json:"saveCount"`
	AuditStatus        int     `json:"auditStatus,omitempty"`
	CreateAt           string  `json:"createAt,omitempty"`
	UpdateAt           string  `json:"updateAt,omitempty"`
}

// GetPaidShareLinkListData 获取付费分享链接列表的返回数据
type GetPaidShareLinkListData struct {
	LastShareId int                 `json:"lastShareId"`
	ShareList   []PaidShareListItem `json:"shareList"`
}

// GetPaidShareLinkListResponse 获取付费分享链接列表的响应
type GetPaidShareLinkListResponse struct {
	Client.Response
	Data GetPaidShareLinkListData `json:"data"`
}

// GetPaidShareLinkList 获取付费分享链接列表
func GetPaidShareLinkList(c *Client.APIClient, getPaidShareLinkListBody GetPaidShareLinkListBody, config Client.APIConfig) (GetPaidShareLinkListResponse, error) {
	url := config.Domain + config.GetPaidShareLinkListAPI +
		"?limit=" + strconv.Itoa(getPaidShareLinkListBody.Limit)
	if getPaidShareLinkListBody.LastShareId != 0 {
		url += "&lastShareId=" + strconv.Itoa(getPaidShareLinkListBody.LastShareId)
	}

	body, err := c.GetQuery(url)
	if err != nil {
		return GetPaidShareLinkListResponse{}, err
	}

	getPaidShareLinkListResponse := GetPaidShareLinkListResponse{}
	err = json.Unmarshal(body, &getPaidShareLinkListResponse)
	return getPaidShareLinkListResponse, err
}
