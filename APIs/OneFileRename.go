package APIs

import (
	"encoding/json"
	"strconv"
)

type OneFileRenameBody struct {
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}

type OneFileRenameData struct{}

type OneFileRenameResponse struct {
	Response
	Data *OneFileRenameData `json:"data"`
}

func (c *APIClient) OneFileRename(oneFileRenameBody OneFileRenameBody) (OneFileRenameResponse, error) {
	url := c.Config.Domain + c.Config.OneFileRenameAPI + "?fileId=" + strconv.Itoa(oneFileRenameBody.FileId) + "&fileName=" + oneFileRenameBody.FileName

	body, err := c.PutQuery(url)
	if err != nil {
		return OneFileRenameResponse{}, err
	}

	oneFileRenameResponse := OneFileRenameResponse{}
	err = json.Unmarshal(body, &oneFileRenameResponse)
	return oneFileRenameResponse, err
}
