package DirectLink

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

type GetDirectLinkURLBody struct {
	FileID int `json:"fileID"`
}
type GetDirectLinkURLData struct {
	URL string `json:"url"`
}
type GetDirectLinkURLResponse struct {
	Client.Response
	Data GetDirectLinkURLData `json:"data"`
}

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
