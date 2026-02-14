package pan123

import (
	"encoding/json"
)

type CopyBatchFilesBody struct {
	FileIds     []int `json:"fileIds"`
	TargetDirId int   `json:"targetDirId"`
}

type CopyBatchFilesData struct {
	TaskId int `json:"taskId"`
}

type CopyBatchFilesResponse struct {
	Response
	Data CopyBatchFilesData `json:"data"`
}

func (c *APIClient) CopyBatchFilesWithConfig(copyBatchFilesBody CopyBatchFilesBody, config Config) (CopyBatchFilesResponse, error) {
	url := config.Domain + config.CopyBatchFilesAPI

	jsonData, err := json.Marshal(copyBatchFilesBody)
	if err != nil {
		return CopyBatchFilesResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return CopyBatchFilesResponse{}, err
	}

	copyBatchFilesResponse := CopyBatchFilesResponse{}
	err = json.Unmarshal(body, &copyBatchFilesResponse)
	return copyBatchFilesResponse, err
}

func (c *APIClient) CopyBatchFiles(copyBatchFilesBody CopyBatchFilesBody) (CopyBatchFilesResponse, error) {
	config := GetDefaultConfig()
	return c.CopyBatchFilesWithConfig(copyBatchFilesBody, config)
}
