package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type CreatePaidShareLinkBody struct {
	ShareName          string `json:"shareName"`
	FileIDList         string `json:"fileIDList"`
	PayAmount          int    `json:"payAmount"`
	IsReward           int    `json:"isReward,omitempty"`
	ResourceDesc       string `json:"resourceDesc,omitempty"`
	TrafficSwitch      int    `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int    `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64  `json:"trafficLimit,omitempty"`
}

type CreatePaidShareLinkData struct {
	ShareID  int    `json:"shareID"`
	ShareKey string `json:"shareKey"`
}

type CreatePaidShareLinkResponse struct {
	ClientAndMethods.Response
	Data CreatePaidShareLinkData `json:"data"`
}

func CreatePaidShareLink(c *ClientAndMethods.APIClient, createPaidShareLinkBody CreatePaidShareLinkBody) (CreatePaidShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.CreatePaidShareLinkAPI

	jsonData, err := json.Marshal(createPaidShareLinkBody)
	if err != nil {
		return CreatePaidShareLinkResponse{}, err
	}

	body, err := c.PostData(url, bytes.NewReader(jsonData))
	if err != nil {
		return CreatePaidShareLinkResponse{}, err
	}

	createPaidShareLinkResponse := CreatePaidShareLinkResponse{}
	err = json.Unmarshal(body, &createPaidShareLinkResponse)
	return createPaidShareLinkResponse, err
}
