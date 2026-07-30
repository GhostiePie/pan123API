package DirectLink

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

// GetDirectLinkURLBody 获取直链URL的请求体
type GetDirectLinkURLBody struct {
	FileID int `json:"fileID"`
}

// GetDirectLinkURLData 获取直链URL的返回数据
type GetDirectLinkURLData struct {
	URL string `json:"url"`
}

// GetDirectLinkURLResponse 获取直链URL的响应
type GetDirectLinkURLResponse struct {
	Client.Response
	Data GetDirectLinkURLData `json:"data"`
}

// GetDirectLinkURL 获取直链URL
func GetDirectLinkURL(c *Client.APIClient, getDirectLinkURLBody GetDirectLinkURLBody, config Client.APIConfig) (GetDirectLinkURLResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkURLAPI + "?fileID=" + strconv.Itoa(getDirectLinkURLBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkURLResponse{}, err
	}

	getDirectLinkURLResponse := GetDirectLinkURLResponse{}
	err = json.Unmarshal(body, &getDirectLinkURLResponse)
	return getDirectLinkURLResponse, err
}
