package Delete

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type DeleteFileToTrashBody struct {
	FileIDs []int `json:"fileIDs"`
}

type DeleteFileToTrashData struct{}

type DeleteFileToTrashResponse struct {
	ClientAndMethods.Response
	Data *DeleteFileToTrashData `json:"data"`
}

func DeleteFileToTrash(c *ClientAndMethods.APIClient, deleteFileToTrashBody DeleteFileToTrashBody) (DeleteFileToTrashResponse, error) {
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
