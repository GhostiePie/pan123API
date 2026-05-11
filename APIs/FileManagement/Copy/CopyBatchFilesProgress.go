package Copy

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

type CopyBatchFilesProgressBody struct {
	TaskId int `json:"taskId"`
}

type CopyBatchFilesProgressData struct {
	TaskId int `json:"taskId"`
	Status int `json:"status"`
}

type CopyBatchFilesProgressResponse struct {
	Client.Response
	Data CopyBatchFilesProgressData `json:"data"`
}

func CopyBatchFilesProgress(c *Client.APIClient, copyBatchFilesProgressBody CopyBatchFilesProgressBody) (CopyBatchFilesProgressResponse, error) {
	url := c.Config.Domain + c.Config.CopyBatchFilesProgressAPI + "?taskId=" + strconv.Itoa(copyBatchFilesProgressBody.TaskId)

	body, err := c.GetQuery(url)
	if err != nil {
		return CopyBatchFilesProgressResponse{}, err
	}

	copyBatchFilesProgressResponse := CopyBatchFilesProgressResponse{}
	err = json.Unmarshal(body, &copyBatchFilesProgressResponse)
	return copyBatchFilesProgressResponse, err
}
