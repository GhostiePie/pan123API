package pan123

import (
	"encoding/json"
	"strconv"
)

type OneFileRenameBody struct {
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}

type OneFileRenameData struct{}

type OneFileRenameResponse struct {
	Response
	Data *OneFileRenameData `json:"data"`
}

func (c *APIClient) OneFileRenameWithConfig(oneFileRenameBody OneFileRenameBody, config Config) (OneFileRenameResponse, error) {
	url := config.Domain + config.OneFileRenameAPI
	data := "fileId=" + strconv.Itoa(oneFileRenameBody.FileId) + "&fileName=" + oneFileRenameBody.FileName

	body, err := c.PutData(url, data)
	if err != nil {
		return OneFileRenameResponse{}, err
	}

	oneFileRenameResponse := OneFileRenameResponse{}
	err = json.Unmarshal(body, &oneFileRenameResponse)
	return oneFileRenameResponse, err
}

func (c *APIClient) OneFileRename(oneFileRenameBody OneFileRenameBody) (OneFileRenameResponse, error) {
	config := GetDefaultConfig()
	return c.OneFileRenameWithConfig(oneFileRenameBody, config)
}
