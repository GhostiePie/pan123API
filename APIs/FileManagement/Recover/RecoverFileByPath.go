package Recover

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// RecoverFileByPathBody 按路径恢复文件请求体
type RecoverFileByPathBody struct {
	FileIDs      []int `json:"fileIDs"`      // 待恢复的文件ID列表
	ParentFileID int   `json:"parentFileID"` // 目标父目录ID
}

// RecoverFileByPathData 按路径恢复文件返回的数据（无实际字段）
type RecoverFileByPathData struct{}

// RecoverFileByPathResponse 按路径恢复文件响应
type RecoverFileByPathResponse struct {
	Client.Response
	Data *RecoverFileByPathData `json:"data"`
}

// RecoverFileByPath 从回收站恢复文件到指定目录
func RecoverFileByPath(c *Client.APIClient, recoverFileByPathBody RecoverFileByPathBody) (RecoverFileByPathResponse, error) {
	url := c.Config.Domain + c.Config.RecoverFileByPathAPI

	jsonData, err := json.Marshal(recoverFileByPathBody)
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	recoverFileByPathResponse := RecoverFileByPathResponse{}
	err = json.Unmarshal(body, &recoverFileByPathResponse)
	return recoverFileByPathResponse, err
}
