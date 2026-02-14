package pan123

import (
	"encoding/json"
	"strconv"
)

type GetOneFileDetailBody struct {
	FileID int `json:"fileID"`
}

type GetOneFileDetailData struct {
	FileID       int    `json:"fileID"`
	Filename     string `json:"filename"`
	Type         int    `json:"type"`
	Size         int    `json:"size"`
	Etag         string `json:"etag"`
	Status       int    `json:"status"`
	ParentFileID int    `json:"parentFileID"`
	CreateAt     string `json:"createAt"`
	Trashed      int    `json:"trashed"`
}

type GetOneFileDetailResponse struct {
	Response
	Data GetOneFileDetailData `json:"data"`
}

func (c *APIClient) GetOneFileDetailWithConfig(getOneFileDetailBody GetOneFileDetailBody, config Config) (GetOneFileDetailResponse, error) {
	url := config.Domain + config.GetOneFileDetailAPI

	data := "fileID=" + strconv.Itoa(getOneFileDetailBody.FileID)

	body, err := c.GetData(url, data)
	if err != nil {
		return GetOneFileDetailResponse{}, err
	}

	getOneFileDetailResponse := GetOneFileDetailResponse{}
	err = json.Unmarshal(body, &getOneFileDetailResponse)
	return getOneFileDetailResponse, err
}

func (c *APIClient) GetOneFileDetail(getOneFileDetailBody GetOneFileDetailBody) (GetOneFileDetailResponse, error) {
	config := GetDefaultConfig()
	return c.GetOneFileDetailWithConfig(getOneFileDetailBody, config)
}
