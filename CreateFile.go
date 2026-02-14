package pan123

import (
	"encoding/json"
	"strconv"
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
	Response
	Data CreateFileData `json:"data"`
}

func (c *APIClient) CreateFileWithConfig(createFileBody CreateFileBody, config Config) (CreateFileResponse, error) {
	url := config.Domain + config.CreateFileAPI
	data := "parentFileID=" + strconv.Itoa(createFileBody.ParentFileID) + "&filename=" + createFileBody.FileName + "&etag=" + createFileBody.Etag + "&size=" + strconv.Itoa(createFileBody.Size)
	if createFileBody.Duplicate != 0 {
		data += "&duplicate=" + strconv.Itoa(createFileBody.Duplicate)
	}
	if createFileBody.ContainDir {
		data += "&containDir=true"
	}
	body, err := c.PostData(url, data)
	if err != nil {
		return CreateFileResponse{}, err
	}
	createFileResponse := CreateFileResponse{}
	err = json.Unmarshal(body, &createFileResponse)

	return createFileResponse, err
}

func (c *APIClient) CreateFile(createFileBody CreateFileBody) (CreateFileResponse, error) {
	config := GetDefaultConfig()
	return c.CreateFileWithConfig(createFileBody, config)
}
