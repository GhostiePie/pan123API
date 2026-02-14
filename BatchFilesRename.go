package pan123

import (
	"encoding/json"
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
	Response
	Data BatchFilesRenameData `json:"data"`
}

func (c *APIClient) BatchFilesRenameWithConfig(batchFilesRenameBody BatchFilesRenameBody, config Config) (BatchFilesRenameResponse, error) {
	url := config.Domain + config.BatchFilesRenameAPI

	jsonData, err := json.Marshal(batchFilesRenameBody)
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return BatchFilesRenameResponse{}, err
	}

	batchFilesRenameResponse := BatchFilesRenameResponse{}
	err = json.Unmarshal(body, &batchFilesRenameResponse)
	return batchFilesRenameResponse, err
}

func (c *APIClient) BatchFilesRename(batchFilesRenameBody BatchFilesRenameBody) (BatchFilesRenameResponse, error) {
	config := GetDefaultConfig()
	return c.BatchFilesRenameWithConfig(batchFilesRenameBody, config)
}
