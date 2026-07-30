package OfflineDownloading

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CreateOfflineDownloadMissionBody 创建离线下载任务的请求体
type CreateOfflineDownloadMissionBody struct {
	URL         string `json:"url"`
	FileName    string `json:"fileName,omitempty"`
	DirID       int    `json:"dirID,omitempty"`
	CallBackUrl string `json:"callBackUrl,omitempty"`
}

// CreateOfflineDownloadMissionData 创建离线下载任务的返回数据
type CreateOfflineDownloadMissionData struct {
	TaskID int `json:"taskID"`
}

// CreateOfflineDownloadMissionResponse 创建离线下载任务的响应
type CreateOfflineDownloadMissionResponse struct {
	Client.Response
	Data CreateOfflineDownloadMissionData `json:"data"`
}

// CreateOfflineDownloadMission 创建离线下载任务
func CreateOfflineDownloadMission(c *Client.APIClient, createOfflineDownloadMissionBody CreateOfflineDownloadMissionBody) (CreateOfflineDownloadMissionResponse, error) {
	url := c.Config.Domain + c.Config.CreateOfflineDownloadMissionAPI

	jsonData, err := json.Marshal(createOfflineDownloadMissionBody)
	if err != nil {
		return CreateOfflineDownloadMissionResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CreateOfflineDownloadMissionResponse{}, err
	}

	createOfflineDownloadMissionResponse := CreateOfflineDownloadMissionResponse{}
	err = json.Unmarshal(body, &createOfflineDownloadMissionResponse)
	return createOfflineDownloadMissionResponse, err
}
