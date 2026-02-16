package APIs

import (
	"encoding/json"
	"strconv"
)

type UploadSHA1Body struct {
	ParentFileID int    `json:"parentFileID"`
	FileName     string `json:"fileName"`
	SHA1         string `json:"sha1"`
	Size         int    `json:"size"`
	Duplicate    int    `json:"duplicate"`
}

type UploadSHA1Data struct {
	FileID int  `json:"fileID"`
	Reuse  bool `json:"reuse"`
}
type UploadSHA1Response struct {
	Response
	Data UploadSHA1Data `json:"data"`
}

func (c *APIClient) UploadSHA1(uploadSHA1Body UploadSHA1Body) (UploadSHA1Response, error) {
	url := c.Config.Domain + c.Config.UploadSHA1API
	data := "parentFileID=" + strconv.Itoa(uploadSHA1Body.ParentFileID) +
		"&filename=" + uploadSHA1Body.FileName +
		"&sha1=" + uploadSHA1Body.SHA1 +
		"&size=" + strconv.Itoa(uploadSHA1Body.Size)
	if uploadSHA1Body.Duplicate != 0 {
		data += "&duplicate=" + strconv.Itoa(uploadSHA1Body.Duplicate)
	}

	body, err := c.PostData(url, data)
	if err != nil {
		return UploadSHA1Response{}, err
	}

	uploadSHA1Response := UploadSHA1Response{}
	err = json.Unmarshal(body, &uploadSHA1Response)
	return uploadSHA1Response, err
}
