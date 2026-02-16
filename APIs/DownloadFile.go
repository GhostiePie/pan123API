package APIs

import (
	"encoding/json"
	"errors"
	"strconv"
)

type DownloadFileBody struct {
	FileID int `json:"fileID"`
}

type DownloadFileData struct {
	DownloadUrl string `json:"downloadUrl"`
}

type DownloadFileResponse struct {
	Response
	Data DownloadFileData `json:"data"`
}

func (c *APIClient) DownloadFile(downloadFileBody DownloadFileBody) (DownloadFileResponse, error) {
	url := c.Config.Domain + c.Config.DownloadFileAPI + "?fileId=" + strconv.Itoa(downloadFileBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return DownloadFileResponse{}, err
	}

	downloadFileResponse := DownloadFileResponse{}
	err = json.Unmarshal(body, &downloadFileResponse)
	if err != nil {
		return DownloadFileResponse{}, err
	}

	if downloadFileResponse.Code != 0 {
		if downloadFileResponse.Code == 5113 {
			return downloadFileResponse, ErrInsufficientDownloadTraffic
		}
		if downloadFileResponse.Code == 5066 {
			return downloadFileResponse, ErrFileNotExists
		}
		return downloadFileResponse, errors.New(downloadFileResponse.Message)
	}

	return downloadFileResponse, nil
}
