package pan123

import (
	"encoding/json"
)

type CopyOneFileBody struct {
	FileId      int `json:"fileId"`
	TargetDirId int `json:"targetDirId"`
}

type CopyOneFileData struct {
	SourceFileId int `json:"sourceFileId"`
	TargetFileId int `json:"targetFileId"`
}

type CopyOneFileResponse struct {
	Response
	Data CopyOneFileData `json:"data"`
}

func (c *APIClient) CopyOneFileWithConfig(copyOneFileBody CopyOneFileBody, config Config) (CopyOneFileResponse, error) {
	url := config.Domain + config.CopyOneFileAPI

	jsonData, err := json.Marshal(copyOneFileBody)
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	body, err := c.PostData(url, string(jsonData))
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	copyOneFileResponse := CopyOneFileResponse{}
	err = json.Unmarshal(body, &copyOneFileResponse)
	return copyOneFileResponse, err
}

func (c *APIClient) CopyOneFile(copyOneFileBody CopyOneFileBody) (CopyOneFileResponse, error) {
	config := GetDefaultConfig()
	return c.CopyOneFileWithConfig(copyOneFileBody, config)
}
