package DirectLink

import (
	"encoding/json"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type DirectLinkCacheRefreshBody struct{}
type DirectLinkCacheRefreshData struct{}
type DirectLinkCacheRefreshResponse struct {
	ClientAndMethods.Response
	Data DirectLinkCacheRefreshData `json:"data"`
}

func DirectLinkCacheRefresh(c *ClientAndMethods.APIClient, directLinkCacheRefreshBody DirectLinkCacheRefreshBody, config ClientAndMethods.APIConfig) (DirectLinkCacheRefreshResponse, error) {
	url := c.Config.Domain + c.Config.DirectLinkCacheRefreshAPI

	body, err := c.PostQuery(url)
	if err != nil {
		return DirectLinkCacheRefreshResponse{}, err
	}

	directLinkCacheRefreshResponse := DirectLinkCacheRefreshResponse{}
	err = json.Unmarshal(body, &directLinkCacheRefreshResponse)
	return directLinkCacheRefreshResponse, err
}
