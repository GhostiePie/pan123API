package Detail

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
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
	Client.Response
	Data GetMultipleFilesDetailData `json:"data"`
}

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
