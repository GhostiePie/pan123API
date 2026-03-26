package Upload

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/APIs"
)

type CreateFileBody struct {
	ParentFileID int    `json:"parentFileID"`
	FileName     string `json:"fileName"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Duplicate    int    `json:"duplicate"`
	ContainDir   bool   `json:"containDir"`
}

type CreateFileData struct {
	FileID      int      `json:"fileID"`
	PreUploadID string   `json:"preuploadID"`
	Reuse       bool     `json:"reuse"`
	SliceSize   int      `json:"sliceSize"`
	Servers     []string `json:"servers"`
}
type CreateFileResponse struct {
	APIs.Response
	Data CreateFileData `json:"data"`
}

func CreateFile(c *APIs.APIClient, createFileBody CreateFileBody) (CreateFileResponse, error) {
	url := c.Config.Domain + c.Config.CreateFileAPI
	data, err := json.Marshal(createFileBody)
	if err != nil {
		return CreateFileResponse{}, err
	}
	body, err := c.PostData(url, string(data))
	if err != nil {
		return CreateFileResponse{}, err
	}
	createFileResponse := CreateFileResponse{}
	err = json.Unmarshal(body, &createFileResponse)

	return createFileResponse, err
}
