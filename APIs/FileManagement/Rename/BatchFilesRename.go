package Rename

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// BatchFilesRenameItem 批量重命名单条记录
type BatchFilesRenameItem struct {
	FileId   int    `json:"fileId"`   // 文件ID
	FileName string `json:"fileName"` // 新文件名
}

// BatchFilesRenameBody 批量重命名请求体
type BatchFilesRenameBody struct {
	RenameList []BatchFilesRenameItem `json:"renameList"` // 重命名列表
}

// BatchFilesRenameSuccessItem 批量重命名成功的条目
type BatchFilesRenameSuccessItem struct {
	FileID   int    `json:"fileID"`   // 文件ID
	UpdateAt string `json:"updateAt"` // 更新时间
}

// BatchFilesRenameFailItem 批量重命名失败的条目
type BatchFilesRenameFailItem struct {
	FileID  int    `json:"fileID"`  // 文件ID
	Message string `json:"message"` // 失败原因
}

// BatchFilesRenameData 批量重命名返回的数据
type BatchFilesRenameData struct {
	SuccessList []BatchFilesRenameSuccessItem `json:"successList"` // 成功列表
	FailList    []BatchFilesRenameFailItem    `json:"failList"`    // 失败列表
}

// BatchFilesRenameResponse 批量重命名响应
type BatchFilesRenameResponse struct {
	Client.Response
	Data BatchFilesRenameData `json:"data"`
}

// BatchFilesRename 批量重命名文件
func BatchFilesRename(c *Client.APIClient, batchFilesRenameBody BatchFilesRenameBody) (BatchFilesRenameResponse, error) {
	url := c.Config.Domain + c.Config.BatchFilesRenameAPI

	jsonData, err := json.Marshal(batchFilesRenameBody)
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	batchFilesRenameResponse := BatchFilesRenameResponse{}
	err = json.Unmarshal(body, &batchFilesRenameResponse)
	return batchFilesRenameResponse, err
}
