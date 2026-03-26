package APIs

import (
	"bytes"
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

func EnableDirectLink(c *APIClient, enableDirectLinkBody EnableDirectLinkBody, config APIConfig) (EnableDirectLinkResponse, error) {
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
