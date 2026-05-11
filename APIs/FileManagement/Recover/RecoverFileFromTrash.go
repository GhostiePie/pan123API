package Recover

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

type RecoverFileFromTrashBody struct {
	FileIDs []int `json:"fileIDs"`
}

type RecoverFileFromTrashData struct {
	AbnormalFileIDs []int `json:"abnormalFileIDs"`
}

type RecoverFileFromTrashResponse struct {
	Client.Response
	Data RecoverFileFromTrashData `json:"data"`
}

func RecoverFileFromTrash(c *Client.APIClient, recoverFileFromTrashBody RecoverFileFromTrashBody) (RecoverFileFromTrashResponse, error) {
	url := c.Config.Domain + c.Config.RecoverFileFromTrashAPI

	jsonData, err := json.Marshal(recoverFileFromTrashBody)
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return RecoverFileFromTrashResponse{}, err
	}

	recoverFileFromTrashResponse := RecoverFileFromTrashResponse{}
	err = json.Unmarshal(body, &recoverFileFromTrashResponse)
	return recoverFileFromTrashResponse, err
}
