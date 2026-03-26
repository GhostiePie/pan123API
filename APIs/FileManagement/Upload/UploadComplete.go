package Upload

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/APIs"
)

type UploadCompleteBody struct {
	PreuploadID string `json:"preuploadID"`
}

type UploadCompleteData struct {
	Completed bool   `json:"completed"`
	FileID    string `json:"fileId"`
}

type UploadCompleteResponse struct {
	APIs.Response
	Data UploadCompleteData `json:"data"`
}

func UploadComplete(c *APIs.APIClient, uploadCompleteBody UploadCompleteBody) (UploadCompleteResponse, error) {
	url := c.Config.Domain + c.Config.UploadCompleteAPI
	data := "preuploadID=" + uploadCompleteBody.PreuploadID

	body, err := c.PostData(url, data)
	if err != nil {
		return UploadCompleteResponse{}, err
	}
	uploadCompleteResponse := UploadCompleteResponse{}
	err = json.Unmarshal(body, &uploadCompleteResponse)

	return uploadCompleteResponse, err
}
