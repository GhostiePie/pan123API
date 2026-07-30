package Upload

import (
	"encoding/json"
	"strings"

	"github.com/GhostiePie/pan123API/Client"
)

// UploadCompleteBody 上传完成请求体
type UploadCompleteBody struct {
	PreuploadID string `json:"preuploadID"` // 预上传ID
}

// UploadCompleteData 上传完成返回的数据
type UploadCompleteData struct {
	Completed bool  `json:"completed"` // 是否完成
	FileID    int64 `json:"fileId"`    // 文件ID
}

// UploadCompleteResponse 上传完成响应
type UploadCompleteResponse struct {
	Client.Response
	Data UploadCompleteData `json:"data"`
}

// UploadComplete 通知服务端所有分片已上传完毕
func UploadComplete(c *Client.APIClient, uploadCompleteBody UploadCompleteBody) (UploadCompleteResponse, error) {
	url := c.Config.Domain + c.Config.UploadCompleteAPI
	data := "preuploadID=" + uploadCompleteBody.PreuploadID

	body, err := c.PostData(url, strings.NewReader(data))
	if err != nil {
		return UploadCompleteResponse{}, err
	}
	uploadCompleteResponse := UploadCompleteResponse{}
	err = json.Unmarshal(body, &uploadCompleteResponse)

	return uploadCompleteResponse, err
}
