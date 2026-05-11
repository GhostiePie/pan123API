package Upload

import (
	"encoding/json"
	"strings"

	"github.com/GhostiePie/pan123API/Client"
)

type UploadCompleteBody struct {
	PreuploadID string `json:"preuploadID"`
}

type UploadCompleteData struct {
	Completed bool  `json:"completed"`
	FileID    int64 `json:"fileId"`
}

type UploadCompleteResponse struct {
	Client.Response
	Data UploadCompleteData `json:"data"`
}

func UploadComplete(c *Client.APIClient, uploadCompleteBody UploadCompleteBody) (UploadCompleteResponse, error) {
	url := c.Config.Domain + c.Config.UploadCompleteAPI
	data := "preuploadID=" + uploadCompleteBody.PreuploadID

	body, err := c.PostData(url, strings.NewReader(data))
	if err != nil {
		return UploadCompleteResponse{}, err
	}
	uploadCompleteResponse := UploadCompleteResponse{}
	err = json.Unmarshal(body, &uploadCompleteResponse)

	return uploadCompleteResponse, err
}
