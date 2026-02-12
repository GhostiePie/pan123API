package pan123

import (
	"encoding/json"
	"strconv"
)

type CreateFileBody struct {
	ParentFileID int    `json:"parentFileId"`
	FileName     string `json:"fileName"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	Duplicate    int    `json:"duplicate"`
	ContainDir   bool   `json:"containDir"`
}

type CreateFileData struct {
	FileID      int      `json:"fileId"`
	PreUploadID string   `json:"preuploadId"`
	Reuse       bool     `json:"reuse"`
	SliceSize   int      `json:"sliceSize"`
	Servers     []string `json:"servers"`
}
type CreateFileResponse struct {
	Response
	Data CreateFileData `json:"data"`
}

func (c APIClient) createFileWithConfig(createFileBody CreateFileBody, config Config) (CreateFileResponse, error) {
	url := config.Domain + config.CreateFileAPI
	data := "parentFileID=" + strconv.Itoa(createFileBody.ParentFileID) + "&filename=" + createFileBody.FileName + "&etag=" + createFileBody.Etag + "&size=" + strconv.Itoa(createFileBody.Size)
	if createFileBody.Duplicate != 0 {
		data += "&duplicate=" + strconv.Itoa(createFileBody.Duplicate)
	}
	if createFileBody.ContainDir {
		data += "&containDir=true"
	}
	body, _ := c.PostData(url, data)
	createFileResponse := CreateFileResponse{}
	err := json.Unmarshal(body, &createFileResponse)

	return createFileResponse, err
}

func (c APIClient) createFile(createFileBody CreateFileBody) (CreateFileResponse, error) {
	config := getDefaultConfig()
	return c.createFileWithConfig(createFileBody, config)
}
