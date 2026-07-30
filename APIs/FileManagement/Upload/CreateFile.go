package Upload

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CreateFileBody 创建文件上传任务请求体
type CreateFileBody struct {
	ParentFileID int    `json:"parentFileID"` // 父目录ID
	FileName     string `json:"fileName"`     // 文件名
	Etag         string `json:"etag"`         // 文件Etag值
	Size         int    `json:"size"`         // 文件大小(bytes)
	Duplicate    int    `json:"duplicate"`    // 重名策略：1重命名 2覆盖 3跳过
	ContainDir   bool   `json:"containDir"`   // 是否包含目录
}

// CreateFileData 创建文件上传任务返回的数据
type CreateFileData struct {
	FileID      int      `json:"fileID"`      // 文件ID
	PreUploadID string   `json:"preuploadID"` // 预上传ID，用于后续分片上传
	Reuse       bool     `json:"reuse"`       // 是否秒传命中
	SliceSize   int      `json:"sliceSize"`   // 分片大小(bytes)
	Servers     []string `json:"servers"`     // 上传服务器地址列表
}

// CreateFileResponse 创建文件上传任务响应
type CreateFileResponse struct {
	Client.Response
	Data CreateFileData `json:"data"`
}

// CreateFile 创建上传任务，返回预上传ID和分片信息
func CreateFile(c *Client.APIClient, createFileBody CreateFileBody) (CreateFileResponse, error) {
	url := c.Config.Domain + c.Config.CreateFileAPI
	data, err := json.Marshal(createFileBody)
	if err != nil {
		return CreateFileResponse{}, err
	}
	body, err := c.PostData(url, bytes.NewReader(data))
	if err != nil {
		return CreateFileResponse{}, err
	}
	createFileResponse := CreateFileResponse{}
	err = json.Unmarshal(body, &createFileResponse)

	return createFileResponse, err
}
