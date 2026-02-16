package APIs

import (
	"encoding/json"
)

type CreateOfflineDownloadMissionBody struct {
	URL         string `json:"url"`
	FileName    string `json:"fileName,omitempty"`
	DirID       int    `json:"dirID,omitempty"`
	CallBackUrl string `json:"callBackUrl,omitempty"`
}

type CreateOfflineDownloadMissionData struct {
	TaskID int `json:"taskID"`
}

type CreateOfflineDownloadMissionResponse struct {
	Response
	Data CreateOfflineDownloadMissionData `json:"data"`
}

func (c *APIClient) CreateOfflineDownloadMission(createOfflineDownloadMissionBody CreateOfflineDownloadMissionBody) (CreateOfflineDownloadMissionResponse, error) {
	url := c.Config.Domain + c.Config.CreateOfflineDownloadMissionAPI

	jsonData, err := json.Marshal(createOfflineDownloadMissionBody)
	if err != nil {
		return CreateOfflineDownloadMissionResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return CreateOfflineDownloadMissionResponse{}, err
	}

	createOfflineDownloadMissionResponse := CreateOfflineDownloadMissionResponse{}
	err = json.Unmarshal(body, &createOfflineDownloadMissionResponse)
	return createOfflineDownloadMissionResponse, err
}
