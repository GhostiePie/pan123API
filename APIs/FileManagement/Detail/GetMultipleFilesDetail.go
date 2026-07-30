package Detail

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// GetMultipleFilesDetailBody 获取多文件详情请求体
type GetMultipleFilesDetailBody struct {
	FileIds []int `json:"fileIds"` // 文件ID列表
}

// FileDetailItem 单个文件详情条目
type FileDetailItem struct {
	FileId       int    `json:"fileId"`
	Filename     string `json:"filename"`
	ParentFileId int    `json:"parentFileId"`
	Type         int    `json:"type"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Category     int    `json:"category"`
	Status       int    `json:"status"`
	PunishFlag   int    `json:"punishFlag"`
	S3KeyFlag    string `json:"s3KeyFlag"`
	StorageNode  string `json:"storageNode"`
	Trashed      int    `json:"trashed"`
	CreateAt     string `json:"createAt"`
	UpdateAt     string `json:"updateAt"`
}

// GetMultipleFilesDetailData 多文件详情数据
type GetMultipleFilesDetailData struct {
	FileList []FileDetailItem `json:"fileList"` // 文件详情列表
}

// GetMultipleFilesDetailResponse 获取多文件详情响应
type GetMultipleFilesDetailResponse struct {
	Client.Response
	Data GetMultipleFilesDetailData `json:"data"`
}

// GetMultipleFilesDetail 批量获取多个文件的详细信息
func GetMultipleFilesDetail(c *Client.APIClient, getMultipleFilesDetailBody GetMultipleFilesDetailBody, config Client.APIConfig) (GetMultipleFilesDetailResponse, error) {
	url := config.Domain + config.GetMultipleFilesDetailAPI

	jsonData, err := json.Marshal(getMultipleFilesDetailBody)
	if err != nil {
		return GetMultipleFilesDetailResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return GetMultipleFilesDetailResponse{}, err
	}

	getMultipleFilesDetailResponse := GetMultipleFilesDetailResponse{}
	err = json.Unmarshal(body, &getMultipleFilesDetailResponse)
	return getMultipleFilesDetailResponse, err
}
