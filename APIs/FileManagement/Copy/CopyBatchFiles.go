package Copy

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type CopyBatchFilesBody struct {
	FileIds     []int `json:"fileIds"`
	TargetDirId int   `json:"targetDirId"`
}

type CopyBatchFilesData struct {
	TaskId int `json:"taskId"`
}

type CopyBatchFilesResponse struct {
	ClientAndMethods.Response
	Data CopyBatchFilesData `json:"data"`
}

func CopyBatchFiles(c *ClientAndMethods.APIClient, copyBatchFilesBody CopyBatchFilesBody) (CopyBatchFilesResponse, error) {
	url := c.Config.Domain + c.Config.CopyBatchFilesAPI

	jsonData, err := json.Marshal(copyBatchFilesBody)
	if err != nil {
		return CopyBatchFilesResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CopyBatchFilesResponse{}, err
	}

	copyBatchFilesResponse := CopyBatchFilesResponse{}
	err = json.Unmarshal(body, &copyBatchFilesResponse)
	return copyBatchFilesResponse, err
}
