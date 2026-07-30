package DirectLink

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

// DirectLinkCacheRefreshBody 刷新直链缓存的请求体
type DirectLinkCacheRefreshBody struct{}

// DirectLinkCacheRefreshData 刷新直链缓存的返回数据
type DirectLinkCacheRefreshData struct{}

// DirectLinkCacheRefreshResponse 刷新直链缓存的响应
type DirectLinkCacheRefreshResponse struct {
	Client.Response
	Data DirectLinkCacheRefreshData `json:"data"`
}

// DirectLinkCacheRefresh 刷新直链缓存
func DirectLinkCacheRefresh(c *Client.APIClient, directLinkCacheRefreshBody DirectLinkCacheRefreshBody, config Client.APIConfig) (DirectLinkCacheRefreshResponse, error) {
	url := c.Config.Domain + c.Config.DirectLinkCacheRefreshAPI

	body, err := c.PostQuery(url)
	if err != nil {
		return DirectLinkCacheRefreshResponse{}, err
	}

	directLinkCacheRefreshResponse := DirectLinkCacheRefreshResponse{}
	err = json.Unmarshal(body, &directLinkCacheRefreshResponse)
	return directLinkCacheRefreshResponse, err
}
