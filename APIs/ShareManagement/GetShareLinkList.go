package ShareManagement

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

type GetShareLinkListBody struct {
	Limit       int `json:"limit"`
	LastShareId int `json:"lastShareId,omitempty"`
}
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
type GetShareLinkListData struct {
	LastShareId int             `json:"lastShareId"`
	ShareList   []ShareListItem `json:"shareList"`
}
type GetShareLinkListResponse struct {
	Client.Response
	Data GetShareLinkListData `json:"data"`
}

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
