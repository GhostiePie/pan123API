package Copy

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type CopyBatchFilesProgressBody struct {
	TaskId int `json:"taskId"`
}

type CopyBatchFilesProgressData struct {
	TaskId int `json:"taskId"`
	Status int `json:"status"`
}

type CopyBatchFilesProgressResponse struct {
	ClientAndMethods.Response
	Data CopyBatchFilesProgressData `json:"data"`
}

func CopyBatchFilesProgress(c *ClientAndMethods.APIClient, copyBatchFilesProgressBody CopyBatchFilesProgressBody) (CopyBatchFilesProgressResponse, error) {
	url := c.Config.Domain + c.Config.CopyBatchFilesProgressAPI + "?taskId=" + strconv.Itoa(copyBatchFilesProgressBody.TaskId)

	body, err := c.GetQuery(url)
	if err != nil {
		return CopyBatchFilesProgressResponse{}, err
	}

	copyBatchFilesProgressResponse := CopyBatchFilesProgressResponse{}
	err = json.Unmarshal(body, &copyBatchFilesProgressResponse)
	return copyBatchFilesProgressResponse, err
}
