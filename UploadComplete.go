package pan123

import "encoding/json"

type UploadCompleteBody struct {
	PreuploadID string `json:"preuploadID"`
}

type UploadCompleteData struct {
	Completed bool   `json:"completed"`
	FileID    string `json:"fileId"`
}

type UploadCompleteResponse struct {
	Response
	Data UploadCompleteData `json:"data"`
}

func (c *APIClient) UploadCompleteWithConfig(uploadCompleteBody UploadCompleteBody, config Config) (UploadCompleteResponse, error) {
	url := config.Domain + config.UploadCompleteAPI
	data := "preuploadID=" + uploadCompleteBody.PreuploadID

	body, err := c.PostData(url, data)
	if err != nil {
		return UploadCompleteResponse{}, err
	}
	uploadCompleteResponse := UploadCompleteResponse{}
	err = json.Unmarshal(body, &uploadCompleteResponse)

	return uploadCompleteResponse, err
}

func (c *APIClient) UploadComplete(uploadCompleteBody UploadCompleteBody) (UploadCompleteResponse, error) {
	config := GetDefaultConfig()
	return c.UploadCompleteWithConfig(uploadCompleteBody, config)
}
