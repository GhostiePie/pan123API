package pan123

import "encoding/json"

type GetUploadDomainResponse struct {
	Response
	Data []string `json:"data"`
}

func (c *APIClient) GetUploadDomainWithConfig(config Config) (getUploadDomainResponse GetUploadDomainResponse, err error) {
	url := config.Domain + config.GetUploadDomainAPI
	resp, err := c.GetData(url, "")
	if err != nil {
		return GetUploadDomainResponse{}, err
	}

	getUploadDomainResponse = GetUploadDomainResponse{}
	err = json.Unmarshal(resp, &getUploadDomainResponse)
	if err != nil {
		return GetUploadDomainResponse{}, err
	}
	return getUploadDomainResponse, nil
}

func (c *APIClient) GetUploadDomain() (getUploadDomainResponse GetUploadDomainResponse, err error) {
	config := GetDefaultConfig()
	return c.GetUploadDomainWithConfig(config)
}
