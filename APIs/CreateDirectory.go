package APIs

import (
	"encoding/json"
)

type CreateDirectoryBody struct {
	Name     string `json:"name"`
	ParentID string `json:"parentID"`
}

type CreateDirectoryData struct {
	DirID int `json:"dirID"`
}
type CreateDirectoryResponse struct {
	Response
	Data CreateDirectoryData `json:"data"`
}

func (c *APIClient) CreateDirectory(createDirectoryBody CreateDirectoryBody) (CreateDirectoryResponse, error) {
	url := c.Config.Domain + c.Config.CreateDirectoryAPI
	data, err := json.Marshal(createDirectoryBody)
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	body, err := c.PostData(url, string(data))
	if err != nil {
		return CreateDirectoryResponse{}, err
	}
	createDirectoryResponse := CreateDirectoryResponse{}
	err = json.Unmarshal(body, &createDirectoryResponse)

	return createDirectoryResponse, err
}
