package FileManagement

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// DownloadFileBody 下载文件请求体
type DownloadFileBody struct {
	FileID int `json:"fileID"` // 文件ID
}

// DownloadFileData 下载文件返回的数据
type DownloadFileData struct {
	DownloadUrl string `json:"downloadUrl"` // 下载链接
}

// DownloadFileResponse 下载文件响应
type DownloadFileResponse struct {
	Client.Response
	Data DownloadFileData `json:"data"`
}

// DownloadFile 获取文件的下载链接
func DownloadFile(c *Client.APIClient, downloadFileBody DownloadFileBody) (DownloadFileResponse, error) {
	url := c.Config.Domain + c.Config.DownloadFileAPI + "?fileId=" + strconv.Itoa(downloadFileBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return DownloadFileResponse{}, err
	}

	downloadFileResponse := DownloadFileResponse{}
	err = json.Unmarshal(body, &downloadFileResponse)
	if err != nil {
		return DownloadFileResponse{}, err
	}

	if downloadFileResponse.Code != 0 {
		if downloadFileResponse.Code == 5113 {
			return downloadFileResponse, Client.ErrInsufficientDownloadTraffic
		}
		if downloadFileResponse.Code == 5066 {
			return downloadFileResponse, Client.ErrFileNotExists
		}
		return downloadFileResponse, errors.New(downloadFileResponse.Message)
	}

	return downloadFileResponse, nil
}
