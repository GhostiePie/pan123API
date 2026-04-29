package Rename

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type OneFileRenameBody struct {
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}

type OneFileRenameData struct{}

type OneFileRenameResponse struct {
	ClientAndMethods.Response
	Data *OneFileRenameData `json:"data"`
}

func OneFileRename(c *ClientAndMethods.APIClient, oneFileRenameBody OneFileRenameBody) (OneFileRenameResponse, error) {
	url := c.Config.Domain + c.Config.OneFileRenameAPI + "?fileId=" + strconv.Itoa(oneFileRenameBody.FileId) + "&fileName=" + oneFileRenameBody.FileName

	body, err := c.PutQuery(url)
	if err != nil {
		return OneFileRenameResponse{}, err
	}

	oneFileRenameResponse := OneFileRenameResponse{}
	err = json.Unmarshal(body, &oneFileRenameResponse)
	return oneFileRenameResponse, err
}
