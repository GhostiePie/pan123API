package Copy

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CopyBatchFilesBody 批量拷贝文件请求体
type CopyBatchFilesBody struct {
	FileIds     []int `json:"fileIds"`     // 源文件ID列表
	TargetDirId int   `json:"targetDirId"` // 目标目录ID
}

// CopyBatchFilesData 批量拷贝文件返回的数据
type CopyBatchFilesData struct {
	TaskId int `json:"taskId"` // 异步任务ID
}

// CopyBatchFilesResponse 批量拷贝文件响应
type CopyBatchFilesResponse struct {
	Client.Response
	Data CopyBatchFilesData `json:"data"`
}

// CopyBatchFiles 批量拷贝文件到指定目录，返回异步任务ID
func CopyBatchFiles(c *Client.APIClient, copyBatchFilesBody CopyBatchFilesBody) (CopyBatchFilesResponse, error) {
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
