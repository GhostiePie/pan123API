package APIs

import (
	"encoding/json"
)

type EnableDirectLinkBody struct {
	FileID int `json:"fileID"`
}
type EnableDirectLinkData struct {
	Filename string `json:"filename"`
}
type EnableDirectLinkResponse struct {
	Response
	Data EnableDirectLinkData `json:"data"`
}

func (c *APIClient) EnableDirectLink(enableDirectLinkBody EnableDirectLinkBody, config APIConfig) (EnableDirectLinkResponse, error) {
	url := c.Config.Domain + c.Config.EnableDirectLinkAPI

	jsonData, err := json.Marshal(enableDirectLinkBody)
	if err != nil {
		return EnableDirectLinkResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return EnableDirectLinkResponse{}, err
	}

	enableDirectLinkResponse := EnableDirectLinkResponse{}
	err = json.Unmarshal(body, &enableDirectLinkResponse)
	return enableDirectLinkResponse, err
}
