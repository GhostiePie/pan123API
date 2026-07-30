package Upload

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CreateDirectoryBody 创建目录请求体
type CreateDirectoryBody struct {
	Name     string `json:"name"`     // 目录名称
	ParentID string `json:"parentID"` // 父目录ID
}

// CreateDirectoryData 创建目录返回的数据
type CreateDirectoryData struct {
	DirID int `json:"dirID"` // 新创建的目录ID
}

// CreateDirectoryResponse 创建目录响应
type CreateDirectoryResponse struct {
	Client.Response
	Data CreateDirectoryData `json:"data"`
}

// CreateDirectory 在指定父目录下创建新目录
func CreateDirectory(c *Client.APIClient, createDirectoryBody CreateDirectoryBody) (CreateDirectoryResponse, error) {
	url := c.Config.Domain + c.Config.CreateDirectoryAPI
	data, err := json.Marshal(createDirectoryBody)
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	body, err := c.PostData(url, bytes.NewReader(data))
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	createDirectoryResponse := CreateDirectoryResponse{}
	err = json.Unmarshal(body, &createDirectoryResponse)

	return createDirectoryResponse, err
}
