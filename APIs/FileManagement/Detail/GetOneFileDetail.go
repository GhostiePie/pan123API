package Detail

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetOneFileDetailBody 获取单文件详情请求体
type GetOneFileDetailBody struct {
	FileID int `json:"fileID"` // 文件ID
}

// GetOneFileDetailData 单文件详情数据
type GetOneFileDetailData struct {
	FileID       int    `json:"fileID"`       // 文件ID
	Filename     string `json:"filename"`     // 文件名
	Type         int    `json:"type"`         // 文件类型
	Size         int    `json:"size"`         // 文件大小
	Etag         string `json:"etag"`         // 文件Etag
	Status       int    `json:"status"`       // 文件状态
	ParentFileID int    `json:"parentFileID"` // 父目录ID
	CreateAt     string `json:"createAt"`     // 创建时间
	Trashed      int    `json:"trashed"`      // 是否在回收站
}

// GetOneFileDetailResponse 获取单文件详情响应
type GetOneFileDetailResponse struct {
	Client.Response
	Data GetOneFileDetailData `json:"data"`
}

// GetOneFileDetail 获取单个文件的详细信息
func GetOneFileDetail(c *Client.APIClient, getOneFileDetailBody GetOneFileDetailBody) (GetOneFileDetailResponse, error) {
	url := c.Config.Domain + c.Config.GetOneFileDetailAPI + "?fileID=" + strconv.Itoa(getOneFileDetailBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetOneFileDetailResponse{}, err
	}

	getOneFileDetailResponse := GetOneFileDetailResponse{}
	err = json.Unmarshal(body, &getOneFileDetailResponse)
	return getOneFileDetailResponse, err
}
