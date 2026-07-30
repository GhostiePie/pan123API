package DirectLink

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// DisableDirectLinkBody 关闭直链的请求体
type DisableDirectLinkBody struct {
	FileID int `json:"fileID"`
}

// DisableDirectLinkData 关闭直链的返回数据
type DisableDirectLinkData struct {
	Filename string `json:"filename"`
}

// DisableDirectLinkResponse 关闭直链的响应
type DisableDirectLinkResponse struct {
	Client.Response
	Data DisableDirectLinkData `json:"data"`
}

// DisableDirectLink 关闭直链
func DisableDirectLink(c *Client.APIClient, disableDirectLinkBody DisableDirectLinkBody, config Client.APIConfig) (DisableDirectLinkResponse, error) {
	url := c.Config.Domain + c.Config.DisableDirectLinkAPI

	jsonData, err := json.Marshal(disableDirectLinkBody)
	if err != nil {
		return DisableDirectLinkResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return DisableDirectLinkResponse{}, err
	}

	disableDirectLinkResponse := DisableDirectLinkResponse{}
	err = json.Unmarshal(body, &disableDirectLinkResponse)
	return disableDirectLinkResponse, err
}
