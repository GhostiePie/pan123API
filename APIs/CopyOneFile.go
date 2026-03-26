package APIs

import (
	"bytes"
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

func CopyOneFile(c *APIClient, copyOneFileBody CopyOneFileBody) (CopyOneFileResponse, error) {
	url := c.Config.Domain + c.Config.CopyOneFileAPI

	jsonData, err := json.Marshal(copyOneFileBody)
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CopyOneFileResponse{}, err
	}

	copyOneFileResponse := CopyOneFileResponse{}
	err = json.Unmarshal(body, &copyOneFileResponse)
	return copyOneFileResponse, err
}
