package DirectLink

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetDirectLinkOfflineLogsBody 获取直链离线日志的请求体
type GetDirectLinkOfflineLogsBody struct {
	StartHour string `json:"startHour"`
	EndHour   string `json:"endHour"`
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
}

// DirectLinkOfflineLog 直链离线日志条目
type DirectLinkOfflineLog struct {
	ID           string `json:"id"`
	FileName     string `json:"fileName"`
	FileSize     int    `json:"fileSize"`
	LogTimeRange string `json:"logTimeRange"`
	DownloadURL  string `json:"downloadURL"`
}

// GetDirectLinkOfflineLogsData 获取直链离线日志的返回数据
type GetDirectLinkOfflineLogsData struct {
	Total int                    `json:"total"`
	List  []DirectLinkOfflineLog `json:"list"`
}

// GetDirectLinkOfflineLogsResponse 获取直链离线日志的响应
type GetDirectLinkOfflineLogsResponse struct {
	Client.Response
	Data GetDirectLinkOfflineLogsData `json:"data"`
}

// GetDirectLinkOfflineLogs 获取直链离线日志
func GetDirectLinkOfflineLogs(c *Client.APIClient, getDirectLinkOfflineLogsBody GetDirectLinkOfflineLogsBody, config Client.APIConfig) (GetDirectLinkOfflineLogsResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkOfflineLogsAPI + "?startHour=" + url.QueryEscape(getDirectLinkOfflineLogsBody.StartHour) + "&endHour=" + url.QueryEscape(getDirectLinkOfflineLogsBody.EndHour) + "&pageNum=" + strconv.Itoa(getDirectLinkOfflineLogsBody.PageNum) + "&pageSize=" + strconv.Itoa(getDirectLinkOfflineLogsBody.PageSize)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkOfflineLogsResponse{}, err
	}

	getDirectLinkOfflineLogsResponse := GetDirectLinkOfflineLogsResponse{}
	err = json.Unmarshal(body, &getDirectLinkOfflineLogsResponse)
	return getDirectLinkOfflineLogsResponse, err
}
