package APIs

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GetDirectLinkOfflineLogsBody struct {
	StartHour string `json:"startHour"`
	EndHour   string `json:"endHour"`
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
}
type DirectLinkOfflineLog struct {
	ID           string `json:"id"`
	FileName     string `json:"fileName"`
	FileSize     int    `json:"fileSize"`
	LogTimeRange string `json:"logTimeRange"`
	DownloadURL  string `json:"downloadURL"`
}
type GetDirectLinkOfflineLogsData struct {
	Total int                    `json:"total"`
	List  []DirectLinkOfflineLog `json:"list"`
}
type GetDirectLinkOfflineLogsResponse struct {
	Response
	Data GetDirectLinkOfflineLogsData `json:"data"`
}

func (c *APIClient) GetDirectLinkOfflineLogs(getDirectLinkOfflineLogsBody GetDirectLinkOfflineLogsBody, config APIConfig) (GetDirectLinkOfflineLogsResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkOfflineLogsAPI + "?startHour=" + url.QueryEscape(getDirectLinkOfflineLogsBody.StartHour) + "&endHour=" + url.QueryEscape(getDirectLinkOfflineLogsBody.EndHour) + "&pageNum=" + strconv.Itoa(getDirectLinkOfflineLogsBody.PageNum) + "&pageSize=" + strconv.Itoa(getDirectLinkOfflineLogsBody.PageSize)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkOfflineLogsResponse{}, err
	}

	getDirectLinkOfflineLogsResponse := GetDirectLinkOfflineLogsResponse{}
	err = json.Unmarshal(body, &getDirectLinkOfflineLogsResponse)
	return getDirectLinkOfflineLogsResponse, err
}
