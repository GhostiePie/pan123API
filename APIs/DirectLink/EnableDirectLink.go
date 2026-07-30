package DirectLink

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// EnableDirectLinkBody 开启直链的请求体
type EnableDirectLinkBody struct {
	FileID int `json:"fileID"`
}

// EnableDirectLinkData 开启直链的返回数据
type EnableDirectLinkData struct {
	Filename string `json:"filename"`
}

// EnableDirectLinkResponse 开启直链的响应
type EnableDirectLinkResponse struct {
	Client.Response
	Data EnableDirectLinkData `json:"data"`
}

// EnableDirectLink 开启直链
func EnableDirectLink(c *Client.APIClient, enableDirectLinkBody EnableDirectLinkBody, config Client.APIConfig) (EnableDirectLinkResponse, error) {
	url := c.Config.Domain + c.Config.EnableDirectLinkAPI

	jsonData, err := json.Marshal(enableDirectLinkBody)
	if err != nil {
		return EnableDirectLinkResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return EnableDirectLinkResponse{}, err
	}

	enableDirectLinkResponse := EnableDirectLinkResponse{}
	err = json.Unmarshal(body, &enableDirectLinkResponse)
	return enableDirectLinkResponse, err
}
