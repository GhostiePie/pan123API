package Delete

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// DeleteFileToTrashBody 删除文件到回收站请求体
type DeleteFileToTrashBody struct {
	FileIDs []int `json:"fileIDs"` // 待删除的文件ID列表
}

// DeleteFileToTrashData 删除文件到回收站返回的数据（无实际字段）
type DeleteFileToTrashData struct{}

// DeleteFileToTrashResponse 删除文件到回收站响应
type DeleteFileToTrashResponse struct {
	Client.Response
	Data *DeleteFileToTrashData `json:"data"`
}

// DeleteFileToTrash 将文件移入回收站
func DeleteFileToTrash(c *Client.APIClient, deleteFileToTrashBody DeleteFileToTrashBody) (DeleteFileToTrashResponse, error) {
	url := c.Config.Domain + c.Config.DeleteFileToTrashAPI

	jsonData, err := json.Marshal(deleteFileToTrashBody)
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	deleteFileToTrashResponse := DeleteFileToTrashResponse{}
	err = json.Unmarshal(body, &deleteFileToTrashResponse)
	return deleteFileToTrashResponse, err
}
