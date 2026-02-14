package pan123

import (
	"encoding/json"
	"strconv"
)

type CopyBatchFilesProgressBody struct {
	TaskId int `json:"taskId"`
}

type CopyBatchFilesProgressData struct {
	TaskId int `json:"taskId"`
	Status int `json:"status"`
}

type CopyBatchFilesProgressResponse struct {
	Response
	Data CopyBatchFilesProgressData `json:"data"`
}

func (c *APIClient) CopyBatchFilesProgressWithConfig(copyBatchFilesProgressBody CopyBatchFilesProgressBody, config Config) (CopyBatchFilesProgressResponse, error) {
	url := config.Domain + config.CopyBatchFilesProgressAPI

	data := "taskId=" + strconv.Itoa(copyBatchFilesProgressBody.TaskId)

	body, err := c.GetData(url, data)
	if err != nil {
		return CopyBatchFilesProgressResponse{}, err
	}

	copyBatchFilesProgressResponse := CopyBatchFilesProgressResponse{}
	err = json.Unmarshal(body, &copyBatchFilesProgressResponse)
	return copyBatchFilesProgressResponse, err
}

func (c *APIClient) CopyBatchFilesProgress(copyBatchFilesProgressBody CopyBatchFilesProgressBody) (CopyBatchFilesProgressResponse, error) {
	config := GetDefaultConfig()
	return c.CopyBatchFilesProgressWithConfig(copyBatchFilesProgressBody, config)
}
