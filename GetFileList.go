package pan123

import (
	"encoding/json"
	"strconv"
)

type GetFileListBody struct {
	ParentFileId int    `json:"parentFileId"`
	Limit        int    `json:"limit"`
	SearchData   string `json:"searchData,omitempty"`
	SearchMode   int    `json:"searchMode,omitempty"`
	LastFileId   int    `json:"lastFileId,omitempty"`
}

type FileListItem struct {
	FileId       int    `json:"fileId"`
	Filename     string `json:"filename"`
	Type         int    `json:"type"`
	Size         int    `json:"size"`
	Etag         string `json:"etag"`
	Status       int    `json:"status"`
	ParentFileId int    `json:"parentFileId"`
	Category     int    `json:"category"`
	Trashed      int    `json:"trashed"`
	PunishFlag   int    `json:"punishFlag"`
	S3KeyFlag    string `json:"s3KeyFlag"`
	StorageNode  string `json:"storageNode"`
	CreateAt     string `json:"createAt"`
	UpdateAt     string `json:"updateAt"`
}

type GetFileListData struct {
	LastFileId int            `json:"lastFileId"`
	FileList   []FileListItem `json:"fileList"`
}

type GetFileListResponse struct {
	Response
	Data GetFileListData `json:"data"`
}

func (c *APIClient) GetFileListWithConfig(getFileListBody GetFileListBody, config Config) (GetFileListResponse, error) {
	url := config.Domain + config.GetFileListAPI

	data := "parentFileId=" + strconv.Itoa(getFileListBody.ParentFileId) +
		"&limit=" + strconv.Itoa(getFileListBody.Limit)
	if getFileListBody.SearchData != "" {
		data += "&searchData=" + getFileListBody.SearchData
	}
	if getFileListBody.SearchMode != 0 {
		data += "&searchMode=" + strconv.Itoa(getFileListBody.SearchMode)
	}
	if getFileListBody.LastFileId != 0 {
		data += "&lastFileId=" + strconv.Itoa(getFileListBody.LastFileId)
	}

	body, err := c.GetData(url, data)
	if err != nil {
		return GetFileListResponse{}, err
	}

	getFileListResponse := GetFileListResponse{}
	err = json.Unmarshal(body, &getFileListResponse)
	return getFileListResponse, err
}

func (c *APIClient) GetFileList(getFileListBody GetFileListBody) (GetFileListResponse, error) {
	config := GetDefaultConfig()
	return c.GetFileListWithConfig(getFileListBody, config)
}
