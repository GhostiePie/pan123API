package Copy

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// CopyBatchFilesProgressBody 批量拷贝进度查询请求体
type CopyBatchFilesProgressBody struct {
	TaskId int `json:"taskId"` // 异步任务ID
}

// CopyBatchFilesProgressData 批量拷贝进度查询返回的数据
type CopyBatchFilesProgressData struct {
	TaskId int `json:"taskId"` // 任务ID
	Status int `json:"status"` // 任务状态
}

// CopyBatchFilesProgressResponse 批量拷贝进度查询响应
type CopyBatchFilesProgressResponse struct {
	Client.Response
	Data CopyBatchFilesProgressData `json:"data"`
}

// CopyBatchFilesProgress 查询批量拷贝任务的进度
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
