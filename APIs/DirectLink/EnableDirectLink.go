package DirectLink

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type EnableDirectLinkBody struct {
	FileID int `json:"fileID"`
}
type EnableDirectLinkData struct {
	Filename string `json:"filename"`
}
type EnableDirectLinkResponse struct {
	ClientAndMethods.Response
	Data EnableDirectLinkData `json:"data"`
}

func EnableDirectLink(c *ClientAndMethods.APIClient, enableDirectLinkBody EnableDirectLinkBody, config ClientAndMethods.APIConfig) (EnableDirectLinkResponse, error) {
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
