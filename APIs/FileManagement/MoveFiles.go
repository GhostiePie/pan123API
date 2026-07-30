package FileManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// MoveFilesBody 移动文件请求体
type MoveFilesBody struct {
	FileIDs        []int `json:"fileIDs"`        // 待移动的文件ID列表
	ToParentFileID int   `json:"toParentFileID"` // 目标目录ID
}

// MoveFilesData 移动文件返回的数据（无实际字段）
type MoveFilesData struct{}

// MoveFilesResponse 移动文件响应
type MoveFilesResponse struct {
	Client.Response
	Data *MoveFilesData `json:"data"`
}

// MoveFiles 批量移动文件到指定目录
func MoveFiles(c *Client.APIClient, moveFilesBody MoveFilesBody) (MoveFilesResponse, error) {
	url := c.Config.Domain + c.Config.MoveFilesAPI

	jsonData, err := json.Marshal(moveFilesBody)
	if err != nil {
		return MoveFilesResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return MoveFilesResponse{}, err
	}

	moveFilesResponse := MoveFilesResponse{}
	err = json.Unmarshal(body, &moveFilesResponse)
	return moveFilesResponse, err
}
