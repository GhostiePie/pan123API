package Rename

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// OneFileRenameBody 单文件重命名请求体
type OneFileRenameBody struct {
	FileId   int    `json:"fileId"`   // 文件ID
	FileName string `json:"fileName"` // 新文件名
}

// OneFileRenameData 单文件重命名返回的数据（无实际字段）
type OneFileRenameData struct{}

// OneFileRenameResponse 单文件重命名响应
type OneFileRenameResponse struct {
	Client.Response
	Data *OneFileRenameData `json:"data"`
}

// OneFileRename 重命名单个文件
func OneFileRename(c *Client.APIClient, oneFileRenameBody OneFileRenameBody) (OneFileRenameResponse, error) {
	url := c.Config.Domain + c.Config.OneFileRenameAPI + "?fileId=" + strconv.Itoa(oneFileRenameBody.FileId) + "&fileName=" + oneFileRenameBody.FileName

	body, err := c.PutQuery(url)
	if err != nil {
		return OneFileRenameResponse{}, err
	}

	oneFileRenameResponse := OneFileRenameResponse{}
	err = json.Unmarshal(body, &oneFileRenameResponse)
	return oneFileRenameResponse, err
}
