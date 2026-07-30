package Upload

import (
	"encoding/json"
	"log"

	"github.com/GhostiePie/pan123API/Client"
)

// GetUploadURLResponse 获取上传URL响应
type GetUploadURLResponse struct {
	Client.Response
	Data []string `json:"data"` // 上传服务器URL列表
}

// GetUploadURL 获取可用的上传服务器URL列表
func GetUploadURL(c *Client.APIClient) (GetUploadURLResponse, error) {
	url := c.Config.Domain + c.Config.GetFileListAPI
	resp, err := c.GetQuery(url)
	if err != nil {
		log.Fatal(err)
	}
	getUploadURLResponse := GetUploadURLResponse{}
	err = json.Unmarshal(resp, &getUploadURLResponse)
	if err != nil {
		log.Fatal(err)
	}
	return getUploadURLResponse, nil
}
