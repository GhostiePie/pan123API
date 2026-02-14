package pan123

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

func (c *APIClient) MoveFilesWithConfig(moveFilesBody MoveFilesBody, config Config) (MoveFilesResponse, error) {
	url := config.Domain + config.MoveFilesAPI

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

func (c *APIClient) MoveFiles(moveFilesBody MoveFilesBody) (MoveFilesResponse, error) {
	config := GetDefaultConfig()
	return c.MoveFilesWithConfig(moveFilesBody, config)
}
