package OfflineDownloading

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

type GetOfflineDownloadProgressBody struct {
	TaskID int `json:"taskID"`
}
type GetOfflineDownloadProgressData struct {
	Process float64 `json:"process"`
	Status  int     `json:"status"`
}
type GetOfflineDownloadProgressResponse struct {
	Client.Response
	Data GetOfflineDownloadProgressData `json:"data"`
}

func GetOfflineDownloadProgress(c *Client.APIClient, getOfflineDownloadProgressBody GetOfflineDownloadProgressBody, config Client.APIConfig) (GetOfflineDownloadProgressResponse, error) {
	url := config.Domain + config.GetOfflineDownloadProgressAPI + "?taskID=" + strconv.Itoa(getOfflineDownloadProgressBody.TaskID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetOfflineDownloadProgressResponse{}, err
	}

	getOfflineDownloadProgressResponse := GetOfflineDownloadProgressResponse{}
	err = json.Unmarshal(body, &getOfflineDownloadProgressResponse)
	return getOfflineDownloadProgressResponse, err
}
