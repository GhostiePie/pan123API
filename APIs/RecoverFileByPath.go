package APIs

import (
	"encoding/json"
)

type RecoverFileByPathBody struct {
	FileIDs      []int `json:"fileIDs"`
	ParentFileID int   `json:"parentFileID"`
}

type RecoverFileByPathData struct{}

type RecoverFileByPathResponse struct {
	Response
	Data *RecoverFileByPathData `json:"data"`
}

func (c *APIClient) RecoverFileByPath(recoverFileByPathBody RecoverFileByPathBody) (RecoverFileByPathResponse, error) {
	url := c.Config.Domain + c.Config.RecoverFileByPathAPI

	jsonData, err := json.Marshal(recoverFileByPathBody)
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	recoverFileByPathResponse := RecoverFileByPathResponse{}
	err = json.Unmarshal(body, &recoverFileByPathResponse)
	return recoverFileByPathResponse, err
}
