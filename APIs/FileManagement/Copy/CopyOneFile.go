package Copy

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
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
	ClientAndMethods.Response
	Data CopyOneFileData `json:"data"`
}

func CopyOneFile(c *ClientAndMethods.APIClient, copyOneFileBody CopyOneFileBody) (CopyOneFileResponse, error) {
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
