package Upload

import (
	"encoding/json"
	"log"

	"github.com/GhostiePie/pan123API/APIs"
)

type GetUploadDomainResponse struct {
	APIs.Response
	Data []string `json:"data"`
}

func GetUploadDomain(c *APIs.APIClient) (GetUploadDomainResponse, error) {
	url := c.Config.Domain + c.Config.GetFileListAPI
	resp, err := c.GetQuery(url)
	if err != nil {
		log.Fatal(err)
	}
	getUploadDomainResponse := GetUploadDomainResponse{}
	err = json.Unmarshal(resp, &getUploadDomainResponse)
	if err != nil {
		log.Fatal(err)
	}
	return getUploadDomainResponse, nil
}
