package pan123

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

func (c *APIClient) DownloadFileWithConfig(downloadFileBody DownloadFileBody, config Config) (DownloadFileResponse, error) {
	url := config.Domain + config.DownloadFileAPI

	data := "fileId=" + strconv.Itoa(downloadFileBody.FileID)

	body, err := c.GetData(url, data)
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
			return downloadFileResponse, errInsufficientDownloadTraffic
		}
		if downloadFileResponse.Code == 5066 {
			return downloadFileResponse, errFileNotExists
		}
		return downloadFileResponse, errors.New(downloadFileResponse.Message)
	}

	return downloadFileResponse, nil
}

func (c *APIClient) DownloadFile(downloadFileBody DownloadFileBody) (DownloadFileResponse, error) {
	config := GetDefaultConfig()
	return c.DownloadFileWithConfig(downloadFileBody, config)
}
