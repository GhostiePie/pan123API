package APIs

import (
	"encoding/json"
	"strconv"
)

type GetDirectLinkURLBody struct {
	FileID int `json:"fileID"`
}
type GetDirectLinkURLData struct {
	URL string `json:"url"`
}
type GetDirectLinkURLResponse struct {
	Response
	Data GetDirectLinkURLData `json:"data"`
}

func (c *APIClient) GetDirectLinkURL(getDirectLinkURLBody GetDirectLinkURLBody, config APIConfig) (GetDirectLinkURLResponse, error) {
	url := c.Config.Domain + c.Config.GetDirectLinkURLAPI + "?fileID=" + strconv.Itoa(getDirectLinkURLBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetDirectLinkURLResponse{}, err
	}

	getDirectLinkURLResponse := GetDirectLinkURLResponse{}
	err = json.Unmarshal(body, &getDirectLinkURLResponse)
	return getDirectLinkURLResponse, err
}
