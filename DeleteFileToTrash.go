package pan123

import (
	"encoding/json"
)

type DeleteFileToTrashBody struct {
	FileIDs []int `json:"fileIDs"`
}

type DeleteFileToTrashData struct{}

type DeleteFileToTrashResponse struct {
	Response
	Data *DeleteFileToTrashData `json:"data"`
}

func (c *APIClient) DeleteFileToTrashWithConfig(deleteFileToTrashBody DeleteFileToTrashBody, config Config) (DeleteFileToTrashResponse, error) {
	url := config.Domain + config.DeleteFileToTrashAPI

	jsonData, err := json.Marshal(deleteFileToTrashBody)
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	deleteFileToTrashResponse := DeleteFileToTrashResponse{}
	err = json.Unmarshal(body, &deleteFileToTrashResponse)
	return deleteFileToTrashResponse, err
}

func (c *APIClient) DeleteFileToTrash(deleteFileToTrashBody DeleteFileToTrashBody) (DeleteFileToTrashResponse, error) {
	config := GetDefaultConfig()
	return c.DeleteFileToTrashWithConfig(deleteFileToTrashBody, config)
}
