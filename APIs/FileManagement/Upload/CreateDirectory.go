package Upload

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type CreateDirectoryBody struct {
	Name     string `json:"name"`
	ParentID string `json:"parentID"`
}

type CreateDirectoryData struct {
	DirID int `json:"dirID"`
}
type CreateDirectoryResponse struct {
	ClientAndMethods.Response
	Data CreateDirectoryData `json:"data"`
}

func CreateDirectory(c *ClientAndMethods.APIClient, createDirectoryBody CreateDirectoryBody) (CreateDirectoryResponse, error) {
	url := c.Config.Domain + c.Config.CreateDirectoryAPI
	data, err := json.Marshal(createDirectoryBody)
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	body, err := c.PostData(url, bytes.NewReader(data))
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	createDirectoryResponse := CreateDirectoryResponse{}
	err = json.Unmarshal(body, &createDirectoryResponse)

	return createDirectoryResponse, err
}
