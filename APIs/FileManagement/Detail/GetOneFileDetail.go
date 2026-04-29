package Detail

import (
	"encoding/json"
	"strconv"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
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
	ClientAndMethods.Response
	Data GetOneFileDetailData `json:"data"`
}

func GetOneFileDetail(c *ClientAndMethods.APIClient, getOneFileDetailBody GetOneFileDetailBody) (GetOneFileDetailResponse, error) {
	url := c.Config.Domain + c.Config.GetOneFileDetailAPI + "?fileID=" + strconv.Itoa(getOneFileDetailBody.FileID)

	body, err := c.GetQuery(url)
	if err != nil {
		return GetOneFileDetailResponse{}, err
	}

	getOneFileDetailResponse := GetOneFileDetailResponse{}
	err = json.Unmarshal(body, &getOneFileDetailResponse)
	return getOneFileDetailResponse, err
}
