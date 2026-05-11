package Rename

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/Client"
)

type OneFileRenameBody struct {
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}

type OneFileRenameData struct{}

type OneFileRenameResponse struct {
	Client.Response
	Data *OneFileRenameData `json:"data"`
}

func OneFileRename(c *Client.APIClient, oneFileRenameBody OneFileRenameBody) (OneFileRenameResponse, error) {
	url := c.Config.Domain + c.Config.OneFileRenameAPI + "?fileId=" + strconv.Itoa(oneFileRenameBody.FileId) + "&fileName=" + oneFileRenameBody.FileName

	body, err := c.PutQuery(url)
	if err != nil {
		return OneFileRenameResponse{}, err
	}

	oneFileRenameResponse := OneFileRenameResponse{}
	err = json.Unmarshal(body, &oneFileRenameResponse)
	return oneFileRenameResponse, err
}
