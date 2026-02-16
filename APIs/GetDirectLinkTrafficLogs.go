package APIs

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GetDirectLinkTrafficLogsBody struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type DirectLinkTrafficLog struct {
	UniqueID      string `json:"uniqueID"`
	FileName      string `json:"fileName"`
	FileSize      int    `json:"fileSize"`
	FilePath      string `json:"filePath"`
	DirectLinkURL string `json:"directLinkURL"`
	FileSource    string `json:"fileSource"`
	TotalTraffic  int    `json:"totalTraffic"`
}
type GetDirectLinkTrafficLogsData struct {
	Total int                    `json:"total"`
	List  []DirectLinkTrafficLog `json:"list"`
}
type GetDirectLinkTrafficLogsResponse struct {
	Response
	Data GetDirectLinkTrafficLogsData `json:"data"`
}

func (c *APIClient) GetDirectLinkTrafficLogs(getDirectLinkTrafficLogsBody GetDirectLinkTrafficLogsBody, config APIConfig) (GetDirectLinkTrafficLogsResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkTrafficLogsAPI + "?pageNum=" + strconv.Itoa(getDirectLinkTrafficLogsBody.PageNum) + "&pageSize=" + strconv.Itoa(getDirectLinkTrafficLogsBody.PageSize) + "&startTime=" + url.QueryEscape(getDirectLinkTrafficLogsBody.StartTime) + "&endTime=" + url.QueryEscape(getDirectLinkTrafficLogsBody.EndTime)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkTrafficLogsResponse{}, err
	}

	getDirectLinkTrafficLogsResponse := GetDirectLinkTrafficLogsResponse{}
	err = json.Unmarshal(body, &getDirectLinkTrafficLogsResponse)
	return getDirectLinkTrafficLogsResponse, err
}
