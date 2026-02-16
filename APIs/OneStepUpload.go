package APIs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
)

type OneStepUploadBody struct {
	ParentFileID int    `json:"parentFileID"`
	FileName     string `json:"fileName"`
	Etag         string `json:"etag"`
	Size         int    `json:"size"`
	File         []byte `json:"file"`
	Duplicate    int    `json:"duplicate"`
	ContainDir   bool   `json:"containDir"`
}

type OneStepUploadData struct {
	FileID    int  `json:"fileID"`
	Completed bool `json:"completed"`
}
type OneStepUploadResponse struct {
	Response
	Data OneStepUploadData `json:"data"`
}

func (c *APIClient) OneStepUpload(oneStepUploadBody OneStepUploadBody) (OneStepUploadResponse, error) {
	url := c.Config.Domain + c.Config.OneStepUploadAPI
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("parentFileID", strconv.Itoa(oneStepUploadBody.ParentFileID))
	writer.WriteField("fileName", oneStepUploadBody.FileName)
	writer.WriteField("etag", oneStepUploadBody.Etag)
	writer.WriteField("size", strconv.Itoa(oneStepUploadBody.Size))
	if oneStepUploadBody.Duplicate != 0 {
		writer.WriteField("duplicate", strconv.Itoa(oneStepUploadBody.Duplicate))
	}
	if oneStepUploadBody.ContainDir {
		writer.WriteField("containDir", "true")
	}

	part, err := writer.CreateFormFile("file", oneStepUploadBody.FileName)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During CreateFormFile: " + err.Error())
	}
	part.Write(oneStepUploadBody.File)
	writer.Close()

	resp, err := c.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During c.Post: " + err.Error())
	}
	fmt.Println(string(resp))

	oneStepUploadResponse := OneStepUploadResponse{}
	err = json.Unmarshal(resp, &oneStepUploadResponse)
	if err != nil {
		return OneStepUploadResponse{}, errors.New("err During Unmarshal: " + err.Error())
	}

	return oneStepUploadResponse, nil
}
