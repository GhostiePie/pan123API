package Rename

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type BatchFilesRenameItem struct {
	FileId   int    `json:"fileId"`
	FileName string `json:"fileName"`
}

type BatchFilesRenameBody struct {
	RenameList []BatchFilesRenameItem `json:"renameList"`
}

type BatchFilesRenameSuccessItem struct {
	FileID   int    `json:"fileID"`
	UpdateAt string `json:"updateAt"`
}

type BatchFilesRenameFailItem struct {
	FileID  int    `json:"fileID"`
	Message string `json:"message"`
}

type BatchFilesRenameData struct {
	SuccessList []BatchFilesRenameSuccessItem `json:"successList"`
	FailList    []BatchFilesRenameFailItem    `json:"failList"`
}

type BatchFilesRenameResponse struct {
	ClientAndMethods.Response
	Data BatchFilesRenameData `json:"data"`
}

func BatchFilesRename(c *ClientAndMethods.APIClient, batchFilesRenameBody BatchFilesRenameBody) (BatchFilesRenameResponse, error) {
	url := c.Config.Domain + c.Config.BatchFilesRenameAPI

	jsonData, err := json.Marshal(batchFilesRenameBody)
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	batchFilesRenameResponse := BatchFilesRenameResponse{}
	err = json.Unmarshal(body, &batchFilesRenameResponse)
	return batchFilesRenameResponse, err
}
