package DirectLink

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetDirectLinkTrafficLogsBody 获取直链流量日志的请求体
type GetDirectLinkTrafficLogsBody struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// DirectLinkTrafficLog 直链流量日志条目
type DirectLinkTrafficLog struct {
	UniqueID      string `json:"uniqueID"`
	FileName      string `json:"fileName"`
	FileSize      int    `json:"fileSize"`
	FilePath      string `json:"filePath"`
	DirectLinkURL string `json:"directLinkURL"`
	FileSource    string `json:"fileSource"`
	TotalTraffic  int    `json:"totalTraffic"`
}

// GetDirectLinkTrafficLogsData 获取直链流量日志的返回数据
type GetDirectLinkTrafficLogsData struct {
	Total int                    `json:"total"`
	List  []DirectLinkTrafficLog `json:"list"`
}

// GetDirectLinkTrafficLogsResponse 获取直链流量日志的响应
type GetDirectLinkTrafficLogsResponse struct {
	Client.Response
	Data GetDirectLinkTrafficLogsData `json:"data"`
}

// GetDirectLinkTrafficLogs 获取直链流量日志
func GetDirectLinkTrafficLogs(c *Client.APIClient, getDirectLinkTrafficLogsBody GetDirectLinkTrafficLogsBody, config Client.APIConfig) (GetDirectLinkTrafficLogsResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkTrafficLogsAPI + "?pageNum=" + strconv.Itoa(getDirectLinkTrafficLogsBody.PageNum) + "&pageSize=" + strconv.Itoa(getDirectLinkTrafficLogsBody.PageSize) + "&startTime=" + url.QueryEscape(getDirectLinkTrafficLogsBody.StartTime) + "&endTime=" + url.QueryEscape(getDirectLinkTrafficLogsBody.EndTime)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkTrafficLogsResponse{}, err
	}

	getDirectLinkTrafficLogsResponse := GetDirectLinkTrafficLogsResponse{}
	err = json.Unmarshal(body, &getDirectLinkTrafficLogsResponse)
	return getDirectLinkTrafficLogsResponse, err
}
