package APIs

import (
	"encoding/json"
)

type DirectLinkCacheRefreshBody struct{}
type DirectLinkCacheRefreshData struct{}
type DirectLinkCacheRefreshResponse struct {
	Response
	Data DirectLinkCacheRefreshData `json:"data"`
}

func (c *APIClient) DirectLinkCacheRefresh(directLinkCacheRefreshBody DirectLinkCacheRefreshBody, config APIConfig) (DirectLinkCacheRefreshResponse, error) {
	url := c.Config.Domain + c.Config.DirectLinkCacheRefreshAPI

	body, err := c.PostQuery(url)
	if err != nil {
		return DirectLinkCacheRefreshResponse{}, err
	}

	directLinkCacheRefreshResponse := DirectLinkCacheRefreshResponse{}
	err = json.Unmarshal(body, &directLinkCacheRefreshResponse)
	return directLinkCacheRefreshResponse, err
}
