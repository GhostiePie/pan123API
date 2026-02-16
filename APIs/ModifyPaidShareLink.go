package APIs

import (
	"encoding/json"
)

type ModifyPaidShareLinkBody struct {
	ShareIdList        []uint64 `json:"shareIdList"`
	TrafficSwitch      int      `json:"trafficSwitch,omitempty"`
	TrafficLimitSwitch int      `json:"trafficLimitSwitch,omitempty"`
	TrafficLimit       int64    `json:"trafficLimit,omitempty"`
}
type ModifyPaidShareLinkData struct{}
type ModifyPaidShareLinkResponse struct {
	Response
	Data *ModifyPaidShareLinkData `json:"data"`
}

func (c *APIClient) ModifyPaidShareLink(modifyPaidShareLinkBody ModifyPaidShareLinkBody) (ModifyPaidShareLinkResponse, error) {
	url := c.Config.Domain + c.Config.ModifyPaidShareLinkAPI

	jsonData, err := json.Marshal(modifyPaidShareLinkBody)
	if err != nil {
		return ModifyPaidShareLinkResponse{}, err
	}

	body, err := c.PutData(url, string(jsonData))
	if err != nil {
		return ModifyPaidShareLinkResponse{}, err
	}

	modifyPaidShareLinkResponse := ModifyPaidShareLinkResponse{}
	err = json.Unmarshal(body, &modifyPaidShareLinkResponse)
	return modifyPaidShareLinkResponse, err
}
