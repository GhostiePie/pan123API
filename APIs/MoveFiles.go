package APIs

import (
	"encoding/json"
)

type MoveFilesBody struct {
	FileIDs        []int `json:"fileIDs"`
	ToParentFileID int   `json:"toParentFileID"`
}
type MoveFilesData struct{}
type MoveFilesResponse struct {
	Response
	Data *MoveFilesData `json:"data"`
}

func (c *APIClient) MoveFiles(moveFilesBody MoveFilesBody) (MoveFilesResponse, error) {
	url := c.Config.Domain + c.Config.MoveFilesAPI

	jsonData, err := json.Marshal(moveFilesBody)
	if err != nil {
		return MoveFilesResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return MoveFilesResponse{}, err
	}

	moveFilesResponse := MoveFilesResponse{}
	err = json.Unmarshal(body, &moveFilesResponse)
	return moveFilesResponse, err
}
