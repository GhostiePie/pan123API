package APIs

import (
	"encoding/json"
	"time"
)

type GetAccessTokenData struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}
type GetAccessTokenResponse struct {
	Response
	Data GetAccessTokenData `json:"data"`
}

type GetAccessTokenBody struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

func (c *APIClient) getAccessToken(getAccessTokenBody GetAccessTokenBody) (GetAccessTokenResponse, error) {
	url := c.Config.Domain + c.Config.AccessTokenAPI
	data, err := json.Marshal(getAccessTokenBody)
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	body, err := c.PostData(url, string(data))
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	getAccessTokenResponse := GetAccessTokenResponse{}
	err = json.Unmarshal(body, &getAccessTokenResponse)

	return getAccessTokenResponse, err

}
