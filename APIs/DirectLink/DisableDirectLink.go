package DirectLink

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type DisableDirectLinkBody struct {
	FileID int `json:"fileID"`
}
type DisableDirectLinkData struct {
	Filename string `json:"filename"`
}
type DisableDirectLinkResponse struct {
	ClientAndMethods.Response
	Data DisableDirectLinkData `json:"data"`
}

func DisableDirectLink(c *ClientAndMethods.APIClient, disableDirectLinkBody DisableDirectLinkBody, config ClientAndMethods.APIConfig) (DisableDirectLinkResponse, error) {
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
