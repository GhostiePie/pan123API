package OfflineDownloading

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetOfflineDownloadProgressBody 获取离线下载进度的请求体
type GetOfflineDownloadProgressBody struct {
	TaskID int `json:"taskID"`
}

// GetOfflineDownloadProgressData 获取离线下载进度的返回数据
type GetOfflineDownloadProgressData struct {
	Process float64 `json:"process"`
	Status  int     `json:"status"`
}

// GetOfflineDownloadProgressResponse 获取离线下载进度的响应
type GetOfflineDownloadProgressResponse struct {
	Client.Response
	Data GetOfflineDownloadProgressData `json:"data"`
}

// GetOfflineDownloadProgress 获取离线下载进度
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
