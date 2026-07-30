package Copy

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// CopyOneFileBody 单文件拷贝请求体
type CopyOneFileBody struct {
	FileId      int `json:"fileId"`      // 源文件ID
	TargetDirId int `json:"targetDirId"` // 目标目录ID
}

// CopyOneFileData 单文件拷贝返回的数据
type CopyOneFileData struct {
	SourceFileId int `json:"sourceFileId"` // 源文件ID
	TargetFileId int `json:"targetFileId"` // 目标文件ID
}

// CopyOneFileResponse 单文件拷贝响应
type CopyOneFileResponse struct {
	Client.Response
	Data CopyOneFileData `json:"data"`
}

// CopyOneFile 拷贝单个文件到指定目录
func CopyOneFile(c *Client.APIClient, copyOneFileBody CopyOneFileBody) (CopyOneFileResponse, error) {
	url := c.Config.Domain + c.Config.CopyOneFileAPI

	jsonData, err := json.Marshal(copyOneFileBody)
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	copyOneFileResponse := CopyOneFileResponse{}
	err = json.Unmarshal(body, &copyOneFileResponse)
	return copyOneFileResponse, err
}
