package FileManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

type MoveFilesBody struct {
	FileIDs        []int `json:"fileIDs"`
	ToParentFileID int   `json:"toParentFileID"`
}
type MoveFilesData struct{}
type MoveFilesResponse struct {
	Client.Response
	Data *MoveFilesData `json:"data"`
}

func MoveFiles(c *Client.APIClient, moveFilesBody MoveFilesBody) (MoveFilesResponse, error) {
	url := c.Config.Domain + c.Config.MoveFilesAPI

	jsonData, err := json.Marshal(moveFilesBody)
	if err != nil {
		return MoveFilesResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return MoveFilesResponse{}, err
	}

	moveFilesResponse := MoveFilesResponse{}
	err = json.Unmarshal(body, &moveFilesResponse)
	return moveFilesResponse, err
}
