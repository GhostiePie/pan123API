package APIs

import (
	"encoding/json"
)

type GetUserInfoBody struct{}
type VipInfoItem struct {
	VipLevel  int    `json:"vipLevel"`
	VipLabel  string `json:"vipLabel"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type DeveloperInfo struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
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
type GetUserInfoResponse struct {
	Response
	Data GetUserInfoData `json:"data"`
}
