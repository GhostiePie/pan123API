package APIs

import (
	"encoding/json"
)

type GetIPBlacklistListData struct {
	IpList []string `json:"ipList"`
	Status int      `json:"status"`
}
type GetIPBlacklistListResponse struct {
	Response
	Data GetIPBlacklistListData `json:"data"`
}
