package APIs

import (
	"bytes"
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

func RecoverFileFromTrash(c *APIClient, recoverFileFromTrashBody RecoverFileFromTrashBody) (RecoverFileFromTrashResponse, error) {
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
