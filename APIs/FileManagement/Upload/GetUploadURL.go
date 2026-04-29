package Upload

import (
	"encoding/json"
	"log"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type GetUploadURLResponse struct {
	ClientAndMethods.Response
	Data []string `json:"data"`
}

func GetUploadURL(c *ClientAndMethods.APIClient) (GetUploadURLResponse, error) {
	url := c.Config.Domain + c.Config.GetFileListAPI
	resp, err := c.GetQuery(url)
	if err != nil {
		log.Fatal(err)
	}
	getUploadURLResponse := GetUploadURLResponse{}
	err = json.Unmarshal(resp, &getUploadURLResponse)
	if err != nil {
		log.Fatal(err)
	}
	return getUploadURLResponse, nil
}
