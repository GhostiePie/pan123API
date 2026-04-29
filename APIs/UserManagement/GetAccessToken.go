package UserManagement

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/GhostiePie/pan123API/ClientAndMethods"
)

type GetAccessTokenData struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}
type GetAccessTokenResponse struct {
	ClientAndMethods.Response
	Data GetAccessTokenData `json:"data"`
}

type GetAccessTokenBody struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

func GetAccessToken(c *ClientAndMethods.APIClient, getAccessTokenBody GetAccessTokenBody) (GetAccessTokenResponse, error) {
	url := c.Config.Domain + c.Config.AccessTokenAPI
	data, err := json.Marshal(getAccessTokenBody)
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	body, err := c.PostData(url, bytes.NewReader(data))
	if err != nil {
		return GetAccessTokenResponse{}, err
	}
	getAccessTokenResponse := GetAccessTokenResponse{}
	err = json.Unmarshal(body, &getAccessTokenResponse)

	return getAccessTokenResponse, err

}
