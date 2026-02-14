package pan123

import (
	"encoding/json"
)

type RecoverFileFromTrashBody struct {
	FileIDs []int `json:"fileIDs"`
}

type RecoverFileFromTrashData struct {
	AbnormalFileIDs []int `json:"abnormalFileIDs"`
}

type RecoverFileFromTrashResponse struct {
	Response
	Data RecoverFileFromTrashData `json:"data"`
}

func (c *APIClient) RecoverFileFromTrashWithConfig(recoverFileFromTrashBody RecoverFileFromTrashBody, config Config) (RecoverFileFromTrashResponse, error) {
	url := config.Domain + config.RecoverFileFromTrashAPI

	jsonData, err := json.Marshal(recoverFileFromTrashBody)
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	recoverFileFromTrashResponse := RecoverFileFromTrashResponse{}
	err = json.Unmarshal(body, &recoverFileFromTrashResponse)
	return recoverFileFromTrashResponse, err
}

func (c *APIClient) RecoverFileFromTrash(recoverFileFromTrashBody RecoverFileFromTrashBody) (RecoverFileFromTrashResponse, error) {
	config := GetDefaultConfig()
	return c.RecoverFileFromTrashWithConfig(recoverFileFromTrashBody, config)
}
