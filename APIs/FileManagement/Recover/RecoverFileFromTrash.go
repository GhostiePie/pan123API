package Recover

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// RecoverFileFromTrashBody 从回收站恢复文件请求体
type RecoverFileFromTrashBody struct {
	FileIDs []int `json:"fileIDs"` // 待恢复的文件ID列表
}

// RecoverFileFromTrashData 从回收站恢复文件返回的数据
type RecoverFileFromTrashData struct {
	AbnormalFileIDs []int `json:"abnormalFileIDs"` // 恢复异常的文件ID列表
}

// RecoverFileFromTrashResponse 从回收站恢复文件响应
type RecoverFileFromTrashResponse struct {
	Client.Response
	Data RecoverFileFromTrashData `json:"data"`
}

// RecoverFileFromTrash 从回收站恢复文件到原位置
func RecoverFileFromTrash(c *Client.APIClient, recoverFileFromTrashBody RecoverFileFromTrashBody) (RecoverFileFromTrashResponse, error) {
	url := c.Config.Domain + c.Config.RecoverFileFromTrashAPI

	jsonData, err := json.Marshal(recoverFileFromTrashBody)
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	recoverFileFromTrashResponse := RecoverFileFromTrashResponse{}
	err = json.Unmarshal(body, &recoverFileFromTrashResponse)
	return recoverFileFromTrashResponse, err
}
