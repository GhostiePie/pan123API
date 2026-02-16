package APIs

import (
	"encoding/json"
	"strconv"
)

type GetOfflineDownloadProgressBody struct {
	TaskID int `json:"taskID"`
}
type GetOfflineDownloadProgressData struct {
	Process float64 `json:"process"`
	Status  int     `json:"status"`
}
type GetOfflineDownloadProgressResponse struct {
	Response
	Data GetOfflineDownloadProgressData `json:"data"`
}

func (c *APIClient) GetOfflineDownloadProgress(getOfflineDownloadProgressBody GetOfflineDownloadProgressBody, config APIConfig) (GetOfflineDownloadProgressResponse, error) {
	url := config.Domain + config.GetOfflineDownloadProgressAPI + "?taskID=" + strconv.Itoa(getOfflineDownloadProgressBody.TaskID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetOfflineDownloadProgressResponse{}, err
	}

	getOfflineDownloadProgressResponse := GetOfflineDownloadProgressResponse{}
	err = json.Unmarshal(body, &getOfflineDownloadProgressResponse)
	return getOfflineDownloadProgressResponse, err
}
