package APIs

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

func (c *APIClient) GetFileList(getFileListBody GetFileListBody, config APIConfig) (GetFileListResponse, error) {
	url := c.Config.Domain + c.Config.GetFileListAPI +
		"?parentFileId=" + strconv.Itoa(getFileListBody.ParentFileId) +
		"&limit=" + strconv.Itoa(getFileListBody.Limit)
	if getFileListBody.SearchData != "" {
		url += "&searchData=" + getFileListBody.SearchData
	}
	if getFileListBody.SearchMode != 0 {
		url += "&searchMode=" + strconv.Itoa(getFileListBody.SearchMode)
	}
	if getFileListBody.LastFileId != 0 {
		url += "&lastFileId=" + strconv.Itoa(getFileListBody.LastFileId)
	}

	body, err := c.GetQuery(url)
	if err != nil {
		return GetFileListResponse{}, err
	}

	getFileListResponse := GetFileListResponse{}
	err = json.Unmarshal(body, &getFileListResponse)
	return getFileListResponse, err
}
