package DirectLink

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/Client"
)

type DirectLinkCacheRefreshBody struct{}
type DirectLinkCacheRefreshData struct{}
type DirectLinkCacheRefreshResponse struct {
	Client.Response
	Data DirectLinkCacheRefreshData `json:"data"`
}

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
