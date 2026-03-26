package APIs

import (
	"bytes"
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

func DeleteFileToTrash(c *APIClient, deleteFileToTrashBody DeleteFileToTrashBody) (DeleteFileToTrashResponse, error) {
	url := c.Config.Domain + c.Config.DeleteFileToTrashAPI

	jsonData, err := json.Marshal(deleteFileToTrashBody)
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return DeleteFileToTrashResponse{}, err
	}

	deleteFileToTrashResponse := DeleteFileToTrashResponse{}
	err = json.Unmarshal(body, &deleteFileToTrashResponse)
	return deleteFileToTrashResponse, err
}
