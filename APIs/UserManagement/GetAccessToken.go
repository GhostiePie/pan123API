package UserManagement

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/GhostiePie/pan123API/Client"
)

// GetAccessTokenData 获取AccessToken的返回数据
type GetAccessTokenData struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

// GetAccessTokenResponse 获取AccessToken的响应
type GetAccessTokenResponse struct {
	Client.Response
	Data GetAccessTokenData `json:"data"`
}

// GetAccessTokenBody 获取AccessToken的请求体
type GetAccessTokenBody struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// GetAccessToken 获取AccessToken
func GetAccessToken(c *Client.APIClient, getAccessTokenBody GetAccessTokenBody) (GetAccessTokenResponse, error) {
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
