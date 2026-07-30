package UserManagement

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// GetUserInfoBody 获取用户信息的请求体
type GetUserInfoBody struct{}

// VipInfoItem VIP信息条目
type VipInfoItem struct {
	VipLevel  int    `json:"vipLevel"`
	VipLabel  string `json:"vipLabel"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// DeveloperInfo 开发者信息
type DeveloperInfo struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// GetUserInfoData 获取用户信息的返回数据
type GetUserInfoData struct {
	UID            int64         `json:"uid"`
	Nickname       string        `json:"nickname"`
	HeadImage      string        `json:"headImage"`
	Passport       string        `json:"passport"`
	Mail           string        `json:"mail"`
	SpaceUsed      int64         `json:"spaceUsed"`
	SpacePermanent int64         `json:"spacePermanent"`
	SpaceTemp      int64         `json:"spaceTemp"`
	SpaceTempExpr  int           `json:"spaceTempExpr"`
	Vip            bool          `json:"vip"`
	DirectTraffic  int64         `json:"directTraffic"`
	IsHideUID      bool          `json:"isHideUID"`
	HTTPSCount     int           `json:"httpsCount"`
	VipInfo        []VipInfoItem `json:"vipInfo"`
	DeveloperInfo  DeveloperInfo `json:"developerInfo"`
}

// GetUserInfoResponse 获取用户信息的响应
type GetUserInfoResponse struct {
	Client.Response
	Data GetUserInfoData `json:"data"`
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *Client.APIClient) (GetUserInfoResponse, error) {
	url := c.Config.Domain + c.Config.GetUserInfoAPI
	resp, err := c.GetQuery(url)
	if err != nil {
		return GetUserInfoResponse{}, err
	}
	getUserInfoResponse := GetUserInfoResponse{}
	err = json.Unmarshal(resp, &getUserInfoResponse)
	if err != nil {
		return GetUserInfoResponse{}, err
	}
	return getUserInfoResponse, nil
}
