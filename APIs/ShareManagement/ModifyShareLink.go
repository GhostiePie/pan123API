package ShareManagement

import (
	"bytes"
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

type ModifyShareLinkBody struct {
	ShareIdList        []uint64 `json:"shareIdList"`
	TrafficSwitch      int      `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int      `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64    `json:"trafficLimit,omitempty"`
}
type ModifyShareLinkData struct{}
type ModifyShareLinkResponse struct {
	Client.Response
	Data *ModifyShareLinkData `json:"data"`
}

func ModifyShareLink(c *Client.APIClient, modifyShareLinkBody ModifyShareLinkBody) (ModifyShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.ModifyShareLinkAPI

	jsonData, err := json.Marshal(modifyShareLinkBody)
	if err != nil {
		return ModifyShareLinkResponse{}, err
	}

	body, err := c.PutData(url, bytes.NewReader(jsonData))
	if err != nil {
		return ModifyShareLinkResponse{}, err
	}

	modifyShareLinkResponse := ModifyShareLinkResponse{}
	err = json.Unmarshal(body, &modifyShareLinkResponse)
	return modifyShareLinkResponse, err
}
