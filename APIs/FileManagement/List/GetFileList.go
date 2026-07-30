package List

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetFileListBody 获取文件列表请求体
type GetFileListBody struct {
	ParentFileId int    `json:"parentFileId"`         // 父目录ID
	Limit        int    `json:"limit"`                // 每页数量
	SearchData   string `json:"searchData,omitempty"` // 搜索关键词
	SearchMode   int    `json:"searchMode,omitempty"` // 搜索模式
	LastFileId   int    `json:"lastFileId,omitempty"` // 上一页最后一个文件ID，用于翻页
}

// FileListItem 文件列表中的单个文件条目
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

// GetFileListData 文件列表数据
type GetFileListData struct {
	LastFileId int            `json:"lastFileId"` // 当前页最后一个文件ID
	FileList   []FileListItem `json:"fileList"`   // 文件列表
}

// GetFileListResponse 获取文件列表响应
type GetFileListResponse struct {
	Client.Response
	Data GetFileListData `json:"data"`
}

// GetFileList 获取指定目录下的文件列表
func GetFileList(c *Client.APIClient, getFileListBody GetFileListBody) (GetFileListResponse, error) {
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
