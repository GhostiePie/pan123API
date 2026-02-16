package APIs

import (
	"encoding/json"
)

type GetMultipleFilesDetailBody struct {
	FileIds []int `json:"fileIds"`
}
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
type GetMultipleFilesDetailData struct {
	FileList []FileDetailItem `json:"fileList"`
}
type GetMultipleFilesDetailResponse struct {
	Response
	Data GetMultipleFilesDetailData `json:"data"`
}

func (c *APIClient) GetMultipleFilesDetail(getMultipleFilesDetailBody GetMultipleFilesDetailBody, config APIConfig) (GetMultipleFilesDetailResponse, error) {
	url := config.Domain + config.GetMultipleFilesDetailAPI

	jsonData, err := json.Marshal(getMultipleFilesDetailBody)
	if err != nil {
		return GetMultipleFilesDetailResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return GetMultipleFilesDetailResponse{}, err
	}

	getMultipleFilesDetailResponse := GetMultipleFilesDetailResponse{}
	err = json.Unmarshal(body, &getMultipleFilesDetailResponse)
	return getMultipleFilesDetailResponse, err
}
