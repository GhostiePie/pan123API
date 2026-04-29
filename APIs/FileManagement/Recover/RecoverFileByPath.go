package Recover

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type RecoverFileByPathBody struct {
	FileIDs      []int `json:"fileIDs"`
	ParentFileID int   `json:"parentFileID"`
}

type RecoverFileByPathData struct{}

type RecoverFileByPathResponse struct {
	ClientAndMethods.Response
	Data *RecoverFileByPathData `json:"data"`
}

func RecoverFileByPath(c *ClientAndMethods.APIClient, recoverFileByPathBody RecoverFileByPathBody) (RecoverFileByPathResponse, error) {
	url := c.Config.Domain + c.Config.RecoverFileByPathAPI

	jsonData, err := json.Marshal(recoverFileByPathBody)
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return RecoverFileByPathResponse{}, err
	}

	recoverFileByPathResponse := RecoverFileByPathResponse{}
	err = json.Unmarshal(body, &recoverFileByPathResponse)
	return recoverFileByPathResponse, err
}
