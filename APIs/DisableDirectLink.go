package APIs

import (
	"bytes"
	"encoding/json"
)

type DisableDirectLinkBody struct {
	FileID int `json:"fileID"`
}
type DisableDirectLinkData struct {
	Filename string `json:"filename"`
}
type DisableDirectLinkResponse struct {
	Response
	Data DisableDirectLinkData `json:"data"`
}

func DisableDirectLink(c *APIClient, disableDirectLinkBody DisableDirectLinkBody, config APIConfig) (DisableDirectLinkResponse, error) {
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
